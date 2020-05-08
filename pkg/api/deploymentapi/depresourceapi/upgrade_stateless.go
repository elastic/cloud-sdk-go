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
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// UpgradeStateless upgrades a stateless deployment resource like APM, Kibana
// and App Search.
func UpgradeStateless(params Params) (*models.DeploymentResourceUpgradeResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, multierror.NewPrefixed("deployment upgrade", err)
	}

	res, err := params.V1API.Deployments.UpgradeDeploymentStatelessResource(
		deployments.NewUpgradeDeploymentStatelessResourceParams().
			WithStatelessResourceKind(params.Kind).
			WithDeploymentID(params.DeploymentID).
			WithRefID(params.RefID),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Unwrap(err)
	}

	return res.Payload, nil
}
