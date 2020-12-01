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

package eskeystoreapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util"
)

// UpdateParams is consumed by the Update function.
type UpdateParams struct {
	*api.API

	DeploymentID string
	Contents     *models.KeystoreContents

	// Optional RefID, whne not specified, an API call will be issued to auto-
	// discover the resource's RefID.
	RefID string
}

// Validate ensures the parameters are usable by Update.
func (params UpdateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid elasticsearch keystore get params")

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if len(params.DeploymentID) != 32 {
		merr = merr.Append(apierror.ErrDeploymentID)
	}

	if params.Contents == nil {
		merr = merr.Append(errors.New("required keystore contents not provided"))
	}

	return merr.ErrorOrNil()
}

// Update changes the contents of the specified deployment's elasticsearch
// keystore by using the PATCH method, the payload is a partial payload where
// any ignored current keystore items, are not removed, unless the secrets are
// set to "null": {"secrets": {"my-secret": null}}.
func Update(params UpdateParams) (*models.KeystoreContents, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := deploymentapi.PopulateRefID(deploymentapi.PopulateRefIDParams{
		API:          params.API,
		DeploymentID: params.DeploymentID,
		RefID:        &params.RefID,
		Kind:         util.Elasticsearch,
	}); err != nil {
		return nil, err
	}

	res, err := params.V1API.Deployments.SetDeploymentEsResourceKeystore(
		deployments.NewSetDeploymentEsResourceKeystoreParams().
			WithDeploymentID(params.DeploymentID).
			WithRefID(params.RefID).
			WithBody(params.Contents),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
