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

package trafficfilterapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments_traffic_filter"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// GetParams is consumed by Get.
type GetParams struct {
	// Required API instance.
	*api.API

	// Required rule identifier.
	ID string

	// Optionally include the rule associations.
	IncludeAssociations bool
}

// Validate ensures the parameters are usable by Get.
func (params GetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid traffic filter get params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("rule set id is not specified and is required for the operation"))
	}

	return merr.ErrorOrNil()
}

// Get obtains an existing traffic filter.
func Get(params GetParams) (*models.TrafficFilterRulesetInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.DeploymentsTrafficFilter.GetTrafficFilterRuleset(
		deployments_traffic_filter.NewGetTrafficFilterRulesetParams().
			WithRulesetID(params.ID).
			WithIncludeAssociations(&params.IncludeAssociations),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
