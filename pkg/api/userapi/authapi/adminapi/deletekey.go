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

package userauthadminapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/authentication"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// DeleteKeyParams is consumed by DeleteKey
type DeleteKeyParams struct {
	*api.API

	ID     string
	UserID string
}

// Validate ensures the parameters are usable by the consuming function.
func (params DeleteKeyParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid user auth admin params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ID == "" {
		merr = merr.Append(errors.New("key id is not specified and is required for this operation"))
	}

	if params.UserID == "" {
		merr = merr.Append(errors.New("user id is not specified and is required for this operation"))
	}

	return merr.ErrorOrNil()
}

// DeleteKey deletes a user's API Key.
func DeleteKey(params DeleteKeyParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(params.V1API.Authentication.DeleteUserAPIKey(
		authentication.NewDeleteUserAPIKeyParams().
			WithAPIKeyID(params.ID).
			WithUserID(params.UserID),
		params.AuthWriter,
	))
}
