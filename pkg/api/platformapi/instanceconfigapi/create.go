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

// CreateParams is used to create a new instance configuration.
type CreateParams struct {
	*api.API
	Config *models.InstanceConfiguration
	Region string
}

// Validate ensures that the parameters are correct.
func (params CreateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid instance config create params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Config == nil {
		merr = merr.Append(errors.New("config not specified and is required for the operation"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Create creates a new instance configuration.
func Create(params CreateParams) (*models.IDResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	if params.Config.ID != "" {
		if err := Update(UpdateParams{
			API:    params.API,
			ID:     params.Config.ID,
			Config: params.Config,
			Region: params.Region,
		}); err != nil {
			return nil, apierror.Wrap(err)
		}
		return &models.IDResponse{ID: ec.String(params.Config.ID)}, nil
	}

	res, err := params.API.V1API.PlatformConfigurationInstances.CreateInstanceConfiguration(
		platform_configuration_instances.NewCreateInstanceConfigurationParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithInstance(params.Config),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
