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

package settingsapi

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

// UpdateParams is the set of parameters required for setting the settings for all proxies
type UpdateParams struct {
	*api.API

	Region, Version string
	*models.ProxiesSettings
}

// Validate parameters for set function
func (params UpdateParams) Validate() error {
	var errs = multierror.NewPrefixed("invalid set proxy settings params")

	if params.API == nil {
		errs = errs.Append(apierror.ErrMissingAPI)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		errs = errs.Append(err)
	}

	if params.ProxiesSettings == nil {
		errs = errs.Append(errors.New("proxy settings object is required for this operation"))
	}

	return errs.ErrorOrNil()
}

// Set sets the settings for all proxies
func Set(params UpdateParams) (*models.ProxiesSettings, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	settings, err := params.API.V1API.PlatformInfrastructure.SetProxiesSettings(
		platform_infrastructure.NewSetProxiesSettingsParams().
			WithVersion(&params.Version).
			WithBody(params.ProxiesSettings).
			WithContext(api.WithRegion(context.Background(), params.Region)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return settings.Payload, nil
}
