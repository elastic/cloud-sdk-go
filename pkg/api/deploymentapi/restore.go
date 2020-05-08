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
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// RestoreParams is consumed by Restore.
type RestoreParams struct {
	*api.API
	DeploymentID string

	RestoreSnapshot bool
}

// Validate ensures the parameters are usable by Restore.
func (params RestoreParams) Validate() error {
	var merr = multierror.NewPrefixed("deployment restore")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if len(params.DeploymentID) != 32 {
		merr = merr.Append(apierror.ErrDeploymentID)
	}

	return merr.ErrorOrNil()
}

// Restore restores a deployment which has been previously shut down.
func Restore(params RestoreParams) (*models.DeploymentRestoreResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.Deployments.RestoreDeployment(
		deployments.NewRestoreDeploymentParams().
			WithDeploymentID(params.DeploymentID).
			WithRestoreSnapshot(ec.Bool(params.RestoreSnapshot)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Unwrap(err)
	}

	return res.Payload, nil
}
