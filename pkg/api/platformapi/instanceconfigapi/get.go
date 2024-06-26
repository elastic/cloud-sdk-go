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

package instanceconfigapi

import (
	"context"
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_instances"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// GetParams is used to obtain an instance configuration from an ID.
type GetParams struct {
	*api.API
	ID            string
	Region        string
	ShowDeleted   bool
	ShowMaxZones  bool
	ConfigVersion *int64
}

// Validate ensures that the parameters are correct.
func (params GetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid instance config get params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("id not specified and is required for the operation"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Get obtains an instance configuration from an ID
func Get(params GetParams) (*models.InstanceConfiguration, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	requestParams := platform_configuration_instances.NewGetInstanceConfigurationParams().
		WithContext(api.WithRegion(context.Background(), params.Region)).
		WithID(params.ID).
		WithConfigVersion(params.ConfigVersion)

	if params.ShowDeleted {
		requestParams = requestParams.WithShowDeleted(ec.Bool(true))
	}
	if params.ShowMaxZones {
		requestParams = requestParams.WithShowMaxZones(ec.Bool(true))
	}

	res, err := params.API.V1API.PlatformConfigurationInstances.GetInstanceConfiguration(
		requestParams,
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
