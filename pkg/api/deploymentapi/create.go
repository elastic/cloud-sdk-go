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

// CreateParams is consumed by Create.
type CreateParams struct {
	*api.API

	// Request from which to create the deployment, by cominging this with
	// the Overrides fields, certain fields can be overridden centrally through
	// that struct.
	//
	// Additionally, a Request can be generated with depresourceapi.NewPayload()
	// for more information on that struct, please refer to its godoc.
	Request *models.DeploymentCreateRequest

	// Optional request_id to be sent in the Create which acts as an idempotency
	// token.
	RequestID string

	// PayloadOverrides are used as a definition of values which want to
	// be overridden within the resources themselves.
	Overrides *PayloadOverrides
}

// Validate ensures the parameters are usable by Create.
func (params CreateParams) Validate() error {
	var merr = multierror.NewPrefixed("deployment create")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Request == nil {
		merr = merr.Append(errors.New("request payload cannot be empty"))
	}

	return merr.ErrorOrNil()
}

// Create performs a Create using the specified Request against the API. Also
// overrides the passed request with the PayloadOverrides set in the wrapping
// CreateParams.
func Create(params CreateParams) (*models.DeploymentCreateResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	setOverrides(params.Request, params.Overrides)

	var id *string
	if params.RequestID != "" {
		id = &params.RequestID
	}

	_, res, res2, err := params.V1API.Deployments.CreateDeployment(
		deployments.NewCreateDeploymentParams().
			WithRequestID(id).
			WithBody(params.Request),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Unwrap(err)
	}

	if res == nil {
		return res2.Payload, nil
	}

	return res.Payload, nil
}
