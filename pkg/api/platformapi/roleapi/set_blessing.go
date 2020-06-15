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

package roleapi

import (
	"context"
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// SetBlessingsParams is consumed by SetBlessings.
type SetBlessingsParams struct {
	*api.API

	Blessings *models.Blessings
	ID        string
	Region    string
}

// Validate ensures the parameters are usable.
func (params SetBlessingsParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid role set blessing params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Blessings == nil {
		merr = merr.Append(errors.New("blessing definitions not specified and is required for this operation"))
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("id not specified and is required for this operation"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// SetBlessings sets a role blessing definitions, the update uses PUT meaning
// it does not require a Role with the current + updated data, only requiring
// the changes which want to be updated.
func SetBlessings(params SetBlessingsParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.V1API.PlatformInfrastructure.SetBlueprinterBlessings(
			platform_infrastructure.NewSetBlueprinterBlessingsParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithBlueprinterRoleID(params.ID).
				WithBody(params.Blessings),
			params.AuthWriter,
		),
	)
}
