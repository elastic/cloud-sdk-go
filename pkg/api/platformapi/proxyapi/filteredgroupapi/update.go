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

package filteredgroupapi

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

var errVersionCannotBeLesserThanZero = errors.New("version cannot be empty")

// UpdateParams is the set of parameters required for updating proxies filtered group
type UpdateParams struct {
	*api.API

	ID                   string
	Region               string
	Filters              map[string]string
	ExpectedProxiesCount int32
	Version              string
}

// Validate parameters for Update function
func (params UpdateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid filtered group params")
	if params.ID == "" {
		merr = merr.Append(errIDCannotBeEmpty)
	}

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	if len(params.Filters) < 1 {
		merr = merr.Append(errFiltersCannotBeEmpty)
	}

	if params.ExpectedProxiesCount < 1 {
		merr = merr.Append(errExpectedProxiesCount)
	}

	if params.Version == "" {
		merr = merr.Append(errVersionCannotBeLesserThanZero)
	}

	return merr.ErrorOrNil()
}

// Update updates information for already existing proxies filtered group
func Update(params UpdateParams) (*models.ProxiesFilteredGroup, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	filters := createProxiesFilters(params.Filters)

	body := &models.ProxiesFilteredGroup{
		Filters:              filters,
		ID:                   params.ID,
		ExpectedProxiesCount: &params.ExpectedProxiesCount,
	}

	proxy, err := params.API.V1API.PlatformInfrastructure.UpdateProxiesFilteredGroup(
		platform_infrastructure.NewUpdateProxiesFilteredGroupParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithBody(body).
			WithVersion(&params.Version).
			WithProxiesFilteredGroupID(params.ID),
		params.AuthWriter,
	)
	if err != nil {
		return nil, api.UnwrapError(err)
	}

	return proxy.Payload, nil
}
