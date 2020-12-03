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
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deputil"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// ResyncParams is consumed by Resync
type ResyncParams struct {
	*api.API
	ID string
}

// Validate ensures the parameters are usable by the consuming function.
func (params ResyncParams) Validate() error {
	var merr = multierror.NewPrefixed("deployment resync",
		deputil.ValidateParams(&params),
	)

	return merr.ErrorOrNil()
}

// ResyncAllParams is consumed by ResyncAll
type ResyncAllParams struct {
	*api.API
}

// Validate ensures the parameters are usable by the consuming function.
func (params ResyncAllParams) Validate() error {
	if params.API == nil {
		return apierror.ErrMissingAPI
	}
	return nil
}

// Resync forces indexer to immediately resynchronize the search index
// and cache for a given deployment.
func Resync(params ResyncParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.API.V1API.Deployments.ResyncDeployment(
			deployments.NewResyncDeploymentParams().
				WithDeploymentID(params.ID),
			params.API.AuthWriter,
		),
	)
}

// ResyncAll asynchronously resynchronizes the search index for all deployments.
func ResyncAll(params ResyncAllParams) (*models.IndexSynchronizationResults, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.Deployments.ResyncDeployments(
		deployments.NewResyncDeploymentsParams(),
		params.API.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
