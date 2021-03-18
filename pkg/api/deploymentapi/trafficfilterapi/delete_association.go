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

// DeleteAssociationParams is consumed by DeleteAssociation.
type DeleteAssociationParams struct {
	// Required API instance.
	*api.API

	// Required ruleset identifier.
	ID string

	// Required entity identifier.
	EntityID string

	// Required ruleset type ("deployment" or "cluster").
	EntityType string
}

// Validate ensures the parameters are usable by DeleteAssociation.
func (params DeleteAssociationParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid traffic filter association delete params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("rule set id is not specified and is required for the operation"))
	}

	if params.EntityID == "" {
		merr = merr.Append(errors.New("entity id is not specified and is required for the operation"))
	}

	if params.EntityType == "" {
		merr = merr.Append(errors.New("entity type is not specified and is required for the operation"))
	}

	return merr.ErrorOrNil()
}

// DeleteAssociation deletes a new traffic filter association to the specified
// entity.
func DeleteAssociation(params DeleteAssociationParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.V1API.DeploymentsTrafficFilter.DeleteTrafficFilterRulesetAssociation(
			deployments_traffic_filter.NewDeleteTrafficFilterRulesetAssociationParams().
				WithRulesetID(params.ID).
				WithAssociatedEntityID(params.EntityID).
				WithAssociationType(params.EntityType),
			params.AuthWriter,
		),
	)
}
