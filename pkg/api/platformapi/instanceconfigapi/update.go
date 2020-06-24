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

// UpdateParams is used to overwrite an existing instance configuration.
type UpdateParams struct {
	*api.API
	ID     string
	Config *models.InstanceConfiguration
	Region string
}

// Validate ensures that the parameters are correct.
func (params UpdateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid instance config update params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Config == nil {
		merr = merr.Append(errors.New("request needs to have a config set"))
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("id must not be empty"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Update overwrites an already existing instance configuration.
func Update(params UpdateParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	_, _, err := params.API.V1API.PlatformConfigurationInstances.SetInstanceConfiguration(
		platform_configuration_instances.NewSetInstanceConfigurationParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithID(params.ID).
			WithInstance(params.Config),
		params.AuthWriter,
	)

	return api.UnwrapError(err)
}
