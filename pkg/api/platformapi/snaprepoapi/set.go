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

package snaprepoapi

import (
	"context"
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_snapshots"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var (
	errNameCannotBeEmpty = errors.New("name not specified and is required for this operation")
	errConfigMustBeSet   = errors.New("config not specified and is required for this operation")
)

// SetParams is used for the Set Call, which will create or update a snapshot
// repository
type SetParams struct {
	*api.API
	Region string
	Name   string
	Type   string
	Config util.Validator
}

// Validate ensures that parameters are correct
func (params SetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid snapshot repository set params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Name == "" {
		merr = merr.Append(errNameCannotBeEmpty)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	if params.Config == nil {
		merr = merr.Append(errConfigMustBeSet)
	}

	if params.Config != nil {
		if err := params.Config.Validate(); err != nil {
			merr = merr.Append(err)
		}
	}

	return merr.ErrorOrNil()
}

// Set adds or updates a snapshot repository from a config
func Set(params SetParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.V1API.PlatformConfigurationSnapshots.SetSnapshotRepository(
			platform_configuration_snapshots.NewSetSnapshotRepositoryParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithRepositoryName(params.Name).
				WithBody(&models.SnapshotRepositoryConfiguration{
					Type:     ec.String(params.Type),
					Settings: params.Config,
				}),
			params.AuthWriter,
		),
	)
}
