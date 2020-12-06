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

package runnerapi

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

var errIDCannotBeEmpty = errors.New("id not specified and is required for the operation")

// ResyncParams is consumed by Resync
type ResyncParams struct {
	*api.API
	Region string
	ID     string
}

// Validate ensures the parameters are usable by the consuming function.
func (params ResyncParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid runner resync params")
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

// Resync forces indexer to immediately resynchronize the search index
// and cache for a given runner.
func Resync(params ResyncParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.API.V1API.PlatformInfrastructure.ResyncRunner(
			platform_infrastructure.NewResyncRunnerParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithRunnerID(params.ID),
			params.API.AuthWriter,
		),
	)
}

// ResyncAllParams is consumed by Resync
type ResyncAllParams struct {
	*api.API
	Region string
}

// Validate ensures the parameters are usable by the consuming function.
func (params ResyncAllParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid runner resync all params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// ResyncAll asynchronously resynchronizes the search index for all runners.
func ResyncAll(params ResyncAllParams) (*models.ModelVersionIndexSynchronizationResults, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.PlatformInfrastructure.ResyncRunners(
		platform_infrastructure.NewResyncRunnersParams().
			WithContext(api.WithRegion(context.Background(), params.Region)),
		params.API.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
