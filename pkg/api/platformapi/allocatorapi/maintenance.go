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

package allocatorapi

import (
	"context"
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// MaintenanceParams is used to set / unset maintenance mode
type MaintenanceParams struct {
	*api.API
	ID     string
	Region string
}

// Validate ensures that the parameters are correct
func (params MaintenanceParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid allocator maintenance params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}
	if params.ID == "" {
		merr = merr.Append(errors.New("id cannot be empty"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// StartMaintenance sets an allocator to maintenance mode
func StartMaintenance(params MaintenanceParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.API.V1API.PlatformInfrastructure.StartAllocatorMaintenanceMode(
			platform_infrastructure.NewStartAllocatorMaintenanceModeParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithAllocatorID(params.ID),
			params.AuthWriter,
		),
	)
}

// StopMaintenance unsets an allocator to maintenance mode
func StopMaintenance(params MaintenanceParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.API.V1API.PlatformInfrastructure.StopAllocatorMaintenanceMode(
			platform_infrastructure.NewStopAllocatorMaintenanceModeParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithAllocatorID(params.ID),
			params.AuthWriter,
		),
	)
}
