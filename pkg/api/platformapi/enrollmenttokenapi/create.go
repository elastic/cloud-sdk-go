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
	"fmt"
	"math"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_security"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// CreateParams is consumed by Create
type CreateParams struct {
	*api.API
	Roles    []string
	Duration time.Duration
	Region   string
}

// Validate ensures that there's no errors prior to performing the Create API
// call.
func (params CreateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid enrollment-token create params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if validity := int64(params.Duration.Seconds()); validity > math.MaxInt32 {
		merr = merr.Append(
			fmt.Errorf("validity value %d exceeds max allowed %d value in seconds", validity, math.MaxInt32),
		)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Create creates the token for the specific roles
func Create(params CreateParams) (*models.RequestEnrollmentTokenReply, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	var persistent = params.Duration.Seconds() <= 0
	var tokenConfig = models.EnrollmentTokenRequest{
		Persistent:        ec.Bool(persistent),
		Roles:             params.Roles,
		ValidityInSeconds: int32(params.Duration.Seconds()),
	}

	res, err := params.API.V1API.PlatformConfigurationSecurity.CreateEnrollmentToken(
		platform_configuration_security.NewCreateEnrollmentTokenParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithBody(&tokenConfig),
		params.AuthWriter,
	)

	if err != nil {
		return nil, api.UnwrapError(err)
	}

	return res.Payload, nil
}
