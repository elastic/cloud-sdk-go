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
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// DeleteParams is used for the Delete call
type DeleteParams struct {
	*api.API
	Region string
	Name   string
}

// Validate ensures that parameters are correct
func (params DeleteParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid snapshot repository delete params")
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

// Delete removes a specified snapshot repository
func Delete(params DeleteParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	_, _, err := params.V1API.PlatformConfigurationSnapshots.DeleteSnapshotRepository(
		platform_configuration_snapshots.NewDeleteSnapshotRepositoryParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithRepositoryName(params.Name),
		params.AuthWriter,
	)

	return api.UnwrapError(err)
}
