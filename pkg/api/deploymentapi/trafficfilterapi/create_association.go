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

// CreateAssociationParams is consumed by CreateAssociation.
type CreateAssociationParams struct {
	// Required API instance.
	*api.API

	// Required ruleset identifier.
	ID string

	// Required entity identifier.
	EntityID string

	// Required ruleset type ("deployment" or "cluster").
	EntityType string
}

// Validate ensures the parameters are usable by CreateAssociation.
func (params CreateAssociationParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid traffic filter association create params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("rule set id is not specified and is required for the operation"))
	}

	if params.EntityID == "" {
		merr = merr.Append(errors.New("entity id is not specified and is required for the operation"))
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("entity type is not specified and is required for the operation"))
	}

	return merr.ErrorOrNil()
}

// CreateAssociation creates a new traffic filter association to the specified
// entity.
func CreateAssociation(params CreateAssociationParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	_, _, err := params.V1API.DeploymentsTrafficFilter.CreateTrafficFilterRulesetAssociation(
		deployments_traffic_filter.NewCreateTrafficFilterRulesetAssociationParams().
			WithRulesetID(params.ID).
			WithBody(&models.FilterAssociation{
				ID: &params.EntityID, EntityType: &params.EntityType,
			}),
		params.AuthWriter,
	)
	return apierror.Wrap(err)
}
