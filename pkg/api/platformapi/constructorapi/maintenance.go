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

package constructorapi

import (
	"context"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// MaintenanceParams is the set of parameters required for EnableMaintenace and
// DisableMaintenance
type MaintenanceParams struct {
	*api.API
	ID     string
	Region string
}

// Validate checks the parameters
func (params MaintenanceParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid constructor maintenance params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ID == "" {
		merr = merr.Append(errIDCannotBeEmpty)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// EnableMaintenace sets the constructor to operational mode
func EnableMaintenace(params MaintenanceParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.V1API.PlatformInfrastructure.StartConstructorMaintenanceMode(
			platform_infrastructure.NewStartConstructorMaintenanceModeParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithConstructorID(params.ID),
			params.AuthWriter,
		),
	)
}

// DisableMaintenance unsets the constructor to operational mode
func DisableMaintenance(params MaintenanceParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.API.V1API.PlatformInfrastructure.StopConstructorMaintenanceMode(
			platform_infrastructure.NewStopConstructorMaintenanceModeParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithConstructorID(params.ID),
			params.AuthWriter,
		),
	)
}
