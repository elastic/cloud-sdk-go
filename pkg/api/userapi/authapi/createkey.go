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

package userauthapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/authentication"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// CreateKeyParams is consumed by CreateKey
type CreateKeyParams struct {
	*api.API

	Description string
}

// Validate ensures the parameters are usable by the consuming function.
func (params CreateKeyParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid user auth params")

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}
	if params.Description == "" {
		merr = merr.Append(errors.New("key description is not specified and is required for this operation"))
	}

	return merr.ErrorOrNil()
}

// CreateKey creates a new API key for the current user.
func CreateKey(params CreateKeyParams) (*models.APIKeyResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.Authentication.CreateAPIKey(
		authentication.NewCreateAPIKeyParams().
			WithBody(&models.CreateAPIKeyRequest{
				Description: ec.String(params.Description),
			}),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
