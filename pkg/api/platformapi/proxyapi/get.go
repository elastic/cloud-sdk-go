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

package proxyapi

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

var (
	errIDCannotBeEmpty = errors.New("proxy id is not specified and is required for the operation")
)

// GetParams is the set of parameters required for retrieving a proxy
type GetParams struct {
	*api.API
	ID     string
	Region string
}

// Validate checks the parameters
func (params GetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid proxy params")
	if params.ID == "" {
		merr = merr.Append(errIDCannotBeEmpty)
	}

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Get returns information about a specific proxy
func Get(params GetParams) (*models.ProxyInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	proxy, err := params.API.V1API.PlatformInfrastructure.GetProxy(
		platform_infrastructure.NewGetProxyParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithProxyID(params.ID),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return proxy.Payload, nil
}
