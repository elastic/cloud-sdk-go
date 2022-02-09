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
	"reflect"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util"
)

const (
	pendingPlanPath = "Info.PlanInfo.Pending"
	currentPlanPath = "Info.PlanInfo.Current"
	historyPlanPath = "Info.PlanInfo.History"
	planAttemptLog  = "PlanAttemptLog"
)

var (
	// ErrPlanFinished is returned when a cluster has no plan PlanStepInfo
	ErrPlanFinished = errors.New("finished all the plan steps")

	errNoPendingPlan = errors.New("no pending plan")
)

// buildTrackResponse takes a models.DeploymentResources and iterates over the
// deployment's resources to obtain a TrackResponse from either the Pending,
// Current or History fields from the resource's plan. It returns []TrackResponse
// each of them being an individual update about a resource's plan. When
// getCurrentPlan is set to true, the plan which is looked up is either the current
// or the last plan in the resource's plan history in the case that the resource
// does not have any current plan, which is a common case for resources which failed
// to create properly.
func buildTrackResponse(res *models.DeploymentResources, getCurrentPlan bool) []TrackResponse {
	var pending = make([]TrackResponse, 0)
	for _, info := range res.Elasticsearch {
		p, err := parseResourceInfo(info, util.Elasticsearch, getCurrentPlan)
		if err != nil {
			continue
		}
		pending = append(pending, p)
	}

	for _, info := range res.Kibana {
		p, err := parseResourceInfo(info, util.Kibana, getCurrentPlan)
		if err != nil {
			continue
		}
		pending = append(pending, p)
	}

	for _, info := range res.Apm {
		p, err := parseResourceInfo(info, util.Apm, getCurrentPlan)
		if err != nil {
			continue
		}
		pending = append(pending, p)
	}

	for _, info := range res.IntegrationsServer {
		p, err := parseResourceInfo(info, util.IntegrationsServer, getCurrentPlan)
		if err != nil {
			continue
		}
		pending = append(pending, p)
	}

	for _, info := range res.Appsearch {
		p, err := parseResourceInfo(info, util.Appsearch, getCurrentPlan)
		if err != nil {
			continue
		}
		pending = append(pending, p)
	}

	for _, info := range res.EnterpriseSearch {
		p, err := parseResourceInfo(info, util.EnterpriseSearch, getCurrentPlan)
		if err != nil {
			continue
		}
		pending = append(pending, p)
	}

	return pending
}

// parseResourceInfo takes in a <kind>ResourceInfo type along with the Kind to
// be able to obtain the resource's plan using reflection, which is deferred to
// getPlanStepInfo. This function builds the TrackResponse structure.
func parseResourceInfo(info interface{}, kind string, getCurrentPlan bool) (TrackResponse, error) {
	stepLog, err := getPlanStepInfo(info, getCurrentPlan)
	if err != nil {
		return TrackResponse{}, err
	}

	var id, refID string
	if v := reflect.ValueOf(info); v.IsValid() {
		id, refID = stringPFieldValue(v, "ID"), stringPFieldValue(v, "RefID")
	}

	step, err := GetStepName(stepLog)
	if step == "" {
		return TrackResponse{}, ErrPlanFinished
	}

	return TrackResponse{
		Kind:           kind,
		ID:             id,
		RefID:          refID,
		Step:           step,
		Err:            err,
		FailureDetails: stepDetails(stepLog),
		Finished:       step == planCompleted,
		Duration:       getPlanDuration(stepLog),
	}, nil
}

// getPlanStepInfo takes in a resource's info and returns a slice of
// ClusterPlanStepInfo, which is obtained by accessing the named fields via
// reflection. The flow might be confusing but in a nutshell it tries to:
//   1. Obtain the Pending plan step log.
//   2. Obtain the Current plan step log when getCurrentPlan is true.
//   3. (if getCurrentPlan == true and the Current plan is empty) obtains the
//       "Current" plan accessing the last item in the plan history slice.
func getPlanStepInfo(workload interface{}, getCurrentPlan bool) ([]*models.ClusterPlanStepInfo, error) {
	var planName = pendingPlanPath
	if getCurrentPlan {
		planName = currentPlanPath
	}

	payloadValue := reflect.ValueOf(workload)
	if !payloadValue.IsValid() {
		return nil, errNoPendingPlan
	}

	// Get either the "Pending" or "Current" plan.
	plan := reflectElemFieldPath(payloadValue, planName)
	if !plan.IsValid() {
		return nil, errNoPendingPlan
	}

	// If the pending plan is nil and getCurrentPlan == false, return an error.
	var noPendingPlanAndWantPendingPlan = plan.IsNil() && !getCurrentPlan
	if noPendingPlanAndWantPendingPlan {
		return nil, errNoPendingPlan
	}

	// When either "Pending" or "Current" aren't nil, obtain the plan log.
	var planLog reflect.Value
	if !plan.IsNil() {
		planLog = plan.Elem().FieldByName(planAttemptLog)
	}

	// When either "Pending" or "Current" are nil, and the current plan needs
	// to be obtained as set by the "getCurrentPlan" bool. Another case is when
	// the planLog is empty, for whichever case, the latest plan in the plan
	// history trail is obtained.
	var currentPlanIsNilAndPlanLogIsNil = plan.IsNil() && getCurrentPlan || planLog.IsNil()
	if currentPlanIsNilAndPlanLogIsNil {
		if history := reflectElemFieldPath(payloadValue, historyPlanPath); history.Len() > 0 {
			var lastPlan = history.Index(history.Len() - 1)
			planLog = lastPlan.Elem().FieldByName(planAttemptLog)
		}
	}

	return getPlanLog(planLog)
}

func getPlanDuration(log []*models.ClusterPlanStepInfo) strfmt.Duration {
	for i := range log {
		return strfmt.Duration(time.Since(time.Time(*log[i].Started)))
	}

	return 0
}

// stringPFieldValue obtains the value of a string pointer field.
func stringPFieldValue(v reflect.Value, field string) string {
	if value := v.Elem().FieldByName(field); value.IsValid() {
		valueStrP := value.Interface().(*string)
		return *valueStrP
	}
	return ""
}

// reflectElemFieldPath obtains the reflect.Value at the end of the specified
// path. Path is in the format of <Property>.<Property> as many times as
// required to obtain the end field.
func reflectElemFieldPath(v reflect.Value, p string) reflect.Value {
	properties := strings.Split(p, ".")
	if len(properties) == 1 {
		return v.Elem().FieldByName(p)
	}

	var fValue reflect.Value
	for i := range properties {
		field := properties[i]
		if fValue = v.Elem().FieldByName(field); fValue.IsValid() {
			return reflectElemFieldPath(fValue, strings.Join(properties[i+1:], "."))
		}
	}

	return fValue
}

// getPlanLog gets the PlanStepInfo slice from PlanAttemptLog as reflect.Value.
func getPlanLog(planLog reflect.Value) ([]*models.ClusterPlanStepInfo, error) {
	if !planLog.IsValid() {
		return nil, errNoPendingPlan
	}

	if log, ok := planLog.Interface().([]*models.ClusterPlanStepInfo); ok {
		return log, nil
	}

	return nil, errors.New("plan: failed casting PlanAttemptLog to []*models.ClusterPlanStepInfo")
}
