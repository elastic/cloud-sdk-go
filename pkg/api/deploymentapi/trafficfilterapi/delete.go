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
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// DeleteParams is consumed by Delete.
type DeleteParams struct {
	// Required API instance.
	*api.API

	// Required rule identifier.
	ID string

	// Optionally ignore the existing rule associations.
	IgnoreAssociations bool
}

// Validate ensures the parameters are usable by Delete.
func (params DeleteParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid traffic filter delete params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("rule set id cannot be empty"))
	}

	return merr.ErrorOrNil()
}

// Delete removes an existing traffic filter rule.
func Delete(params DeleteParams) error {
	if err := params.Validate(); err != nil {
		return err
	}
	return api.ReturnErrOnly(
		params.V1API.DeploymentsTrafficFilter.DeleteTrafficFilterRuleset(
			deployments_traffic_filter.NewDeleteTrafficFilterRulesetParams().
				WithIgnoreAssociations(&params.IgnoreAssociations).
				WithRulesetID(params.ID),
			params.AuthWriter,
		),
	)
}
