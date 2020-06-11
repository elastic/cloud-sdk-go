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

package enrollmenttokenapi

import (
	"context"
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_security"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// DeleteParams is consumed by Delete
type DeleteParams struct {
	*api.API
	Token  string
	Region string
}

// Validate ensures that there's no errors prior to performing the Delete API
// call.
func (params DeleteParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid enrollment-token delete params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Token == "" {
		merr = merr.Append(errors.New("token cannot be empty"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Delete deletes a persistent token
func Delete(params DeleteParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.API.V1API.PlatformConfigurationSecurity.DeleteEnrollmentToken(
			platform_configuration_security.NewDeleteEnrollmentTokenParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithToken(params.Token),
			params.AuthWriter,
		),
	)
}
