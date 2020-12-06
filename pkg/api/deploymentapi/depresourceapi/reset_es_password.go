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

package depresourceapi

import (
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deputil"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util"
)

// ResetElasticsearchPasswordParams is consumed by ResetElasticsearchPassword.
type ResetElasticsearchPasswordParams struct {
	// Required API instance.
	*api.API

	// Required deployment ID.
	ID string

	// Optional RefID. If not specified, it is autodiscovered.
	RefID string
}

// Validate ensures the parameters are usable by ResetElasticsearchPassword.
func (params ResetElasticsearchPasswordParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment elasticsearch password reset params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if len(params.ID) != 32 {
		merr = merr.Append(deputil.NewInvalidDeploymentIDError(params.ID))
	}

	return merr.ErrorOrNil()
}

// ResetElasticsearchPassword resets an deployment's elasticsearch password and
// returns the Username and Password.
func ResetElasticsearchPassword(params ResetElasticsearchPasswordParams) (*models.ElasticsearchElasticUserPasswordResetResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	if params.RefID == "" {
		if err := deploymentapi.PopulateRefID(deploymentapi.PopulateRefIDParams{
			Kind:         util.Elasticsearch,
			API:          params.API,
			DeploymentID: params.ID,
			RefID:        &params.RefID,
		}); err != nil {
			return nil, err
		}
	}

	res, err := params.V1API.Deployments.ResetElasticsearchUserPassword(
		deployments.NewResetElasticsearchUserPasswordParams().
			WithDeploymentID(params.ID).
			WithRefID(params.RefID),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
