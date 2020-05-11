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
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deputil"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// Params can be embedded in any structure which makes use of the
// deployment/resource API which always require the fields. Also provides
// RefID auto-discovery for the deployment resource kind if not specified.
type Params struct {
	*api.API

	DeploymentID string
	Kind         string
	RefID        string
}

// Validate ensures the parameters are usable by the consuming function.
func (params *Params) Validate() error {
	var merr = multierror.NewPrefixed("deployment resource")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if len(params.DeploymentID) != 32 {
		merr = merr.Append(deputil.NewInvalidDeploymentIDError(params.DeploymentID))
	}

	if params.Kind == "" {
		merr = merr.Append(errors.New("resource kind cannot be empty"))
	}

	// Ensures that RefID is either populated when the RefID isn't specified or
	// returns an error when it fails obtaining the ref ID.
	if err := params.fillDefaults(); err != nil {
		merr = merr.Append(multierror.NewPrefixed(
			"failed auto-discovering the resource ref id", err,
		))
	}

	return merr.ErrorOrNil()
}

// fillDefaults populates the RefID through an API call.
func (params *Params) fillDefaults() error {
	if params.RefID != "" {
		return nil
	}

	refID, err := deploymentapi.GetKindRefID(deploymentapi.GetResourceParams{
		Kind: params.Kind,
		GetParams: deploymentapi.GetParams{
			API:          params.API,
			DeploymentID: params.DeploymentID,
		},
	})

	params.RefID = refID
	return err
}
