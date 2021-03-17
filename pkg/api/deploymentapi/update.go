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

package deploymentapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// UpdateParams is consumed by Update.
type UpdateParams struct {
	*api.API

	DeploymentID string
	Request      *models.DeploymentUpdateRequest

	// Optional values
	SkipSnapshot      bool
	HidePrunedOrphans bool

	// PayloadOverrides are used as a definition of values which want to
	// be overridden within the resources themselves.
	Overrides PayloadOverrides
}

// Validate ensures the parameters are usable by Update.
func (params UpdateParams) Validate() error {
	var merr = multierror.NewPrefixed("deployment update")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Request == nil {
		merr = merr.Append(errors.New("request payload cannot be empty"))
	}

	if len(params.DeploymentID) != 32 {
		merr = merr.Append(apierror.ErrDeploymentID)
	}

	return merr.ErrorOrNil()
}

// Update receives an update payload with an optional region override in case
// the region isn't specified in the update request payload. Additionally if
// Request.PruneOrphans is false then any omitted resources aren't shutdown.
// The opposite behavior can be expected when the flag is true since the update
// request is treated as the single source of truth and the complete desired
// deployment definition.
func Update(params UpdateParams) (*models.DeploymentUpdateResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := OverrideCreateOrUpdateRequest(
		params.Request, &params.Overrides,
	); err != nil {
		return nil, err
	}

	res, err := params.V1API.Deployments.UpdateDeployment(
		deployments.NewUpdateDeploymentParams().
			WithDeploymentID(params.DeploymentID).
			WithBody(params.Request).
			WithSkipSnapshot(&params.SkipSnapshot).
			WithHidePrunedOrphans(&params.HidePrunedOrphans),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
