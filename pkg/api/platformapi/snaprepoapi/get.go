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

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_snapshots"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// GetParams is used for the Get call
type GetParams struct {
	*api.API
	Region string
	Name   string
}

// Validate ensures that parameters are correct
func (params GetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid snapshot repository get params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Name == "" {
		merr = merr.Append(errNameCannotBeEmpty)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Get obtains the specified snapshot repository configuration
func Get(params GetParams) (*models.RepositoryConfig, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	repo, err := params.V1API.PlatformConfigurationSnapshots.GetSnapshotRepository(
		platform_configuration_snapshots.NewGetSnapshotRepositoryParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithRepositoryName(params.Name),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return repo.Payload, nil
}
