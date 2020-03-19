// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package plan

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	multierror "github.com/hashicorp/go-multierror"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_apm"
	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_elasticsearch"
	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_kibana"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/elastic/cloud-sdk-go/pkg/util/slice"
)

// GetParams (DEPRECATED) is used to obtain the plan activity logs of a cluster. If pending
// is true, the pending plan logs are obtained, otherwise the current plan logs
// are.
type GetParams struct {
	ID   string
	Kind string
	*api.API
	// Only used when MaxRetries is set.
	Cooldown time.Duration
	Pending  bool
	// If not set, disables retrying the request upon receiving an error.
	MaxRetries uint8
	retries    uint8
}

// Validate (DEPRECATED) verifies that the parameters being sent are usable by the consuming
// function.
func (params GetParams) Validate() error {
	var err = new(multierror.Error)
	if params.API == nil {
		err = multierror.Append(err, errors.New("plan get: API cannot be nil"))
	}

	if len(params.ID) != 32 {
		err = multierror.Append(err, errors.New("plan get: invalid cluster id"))
	}

	if !slice.HasString(allowedKinds, params.Kind) {
		err = multierror.Append(err, fmt.Errorf("plan get: invalid kind %s", params.Kind))
	}

	if params.MaxRetries > 0 && params.Cooldown.Nanoseconds() == 0 {
		err = multierror.Append(err, fmt.Errorf("plan get: both MaxRetries and Cooldown must be set"))
	}

	return err.ErrorOrNil()
}

func (params GetParams) needsRetry() bool {
	return params.retries < params.MaxRetries
}

// Get (DEPRECATED) obtains a cluster's plan logs from a set of parameters.
func Get(params GetParams) ([]*models.ClusterPlanStepInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	var res interface{}
	var err error
	switch params.Kind {
	case "elasticsearch":
		res, err = params.API.V1API.ClustersElasticsearch.GetEsClusterPlanActivity(
			clusters_elasticsearch.NewGetEsClusterPlanActivityParams().
				WithClusterID(params.ID).
				WithShowPlanLogs(ec.Bool(true)),
			params.API.AuthWriter,
		)
	case "kibana":
		res, err = params.V1API.ClustersKibana.GetKibanaClusterPlanActivity(
			clusters_kibana.NewGetKibanaClusterPlanActivityParams().
				WithClusterID(params.ID).
				WithShowPlanLogs(ec.Bool(true)),
			params.API.AuthWriter,
		)
	case "apm":
		res, err = params.V1API.ClustersApm.GetApmClusterPlanActivity(
			clusters_apm.NewGetApmClusterPlanActivityParams().
				WithClusterID(params.ID).
				WithShowPlanLogs(ec.Bool(true)),
			params.API.AuthWriter,
		)
	default:
		return nil, fmt.Errorf("poller: invalid kind %s", params.Kind)
	}

	if err != nil {
		if !params.needsRetry() {
			return nil, err
		}
		params.retries++
		// Cooldown for 1s + the retry number (Exponential backoff).
		<-time.After(params.Cooldown * time.Duration(params.retries))
		return Get(params)
	}
	return GetLogsFromPlanActivityResponse(res, params.Pending)
}

// GetLogsFromPlanActivityResponse (DEPRECATED) takes in a response from `ClusterPlanActivity`
// and returns the plan steps, either from the current or the pending plan.
func GetLogsFromPlanActivityResponse(res interface{}, pending bool) ([]*models.ClusterPlanStepInfo, error) {
	payload := reflect.ValueOf(res).Elem().FieldByName("Payload")
	if payload.IsNil() || !payload.IsValid() {
		return nil, errors.New("plan: failed to obtain Payload field from Response")
	}

	var prop = "Current"
	if pending {
		prop = "Pending"
	}

	var plan = payload.Elem().FieldByName(prop)
	if plan.IsNil() || !payload.IsValid() {
		return nil, fmt.Errorf("plan: failed to obtain %s from Payload", prop)
	}

	planAttemptLog := plan.Elem().FieldByName("PlanAttemptLog")
	if planAttemptLog.IsNil() || !planAttemptLog.IsValid() {
		return nil, errors.New("plan: failed to obtain PlanAttemptLog field from Plan")
	}

	if log, ok := planAttemptLog.Interface().([]*models.ClusterPlanStepInfo); ok {
		return log, nil
	}

	return nil, errors.New("plan: failed casting PlanAttemptLog to []*models.ClusterPlanStepInfo")
}
