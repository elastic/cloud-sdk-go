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

package userapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/users"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// GetParams is consumed by Get
type GetParams struct {
	*api.API

	UserName string
}

// Validate ensures the parameters are usable by the consuming function.
func (params GetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid user params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.UserName == "" {
		merr = merr.Append(errors.New("username is not specified and is required for this operation"))
	}

	return merr.ErrorOrNil()
}

// Get returns information about a specified user.
func Get(params GetParams) (*models.User, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.Users.GetUser(
		users.NewGetUserParams().
			WithUserName(params.UserName),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}

// GetCurrentParams is consumed by GetCurrent
type GetCurrentParams struct {
	*api.API
}

// Validate ensures the parameters are usable by the consuming function.
func (params GetCurrentParams) Validate() error {
	if params.API == nil {
		return apierror.ErrMissingAPI
	}

	return nil
}

// GetCurrent returns information about the current user.
func GetCurrent(params GetCurrentParams) (*models.User, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.Users.GetCurrentUser(
		users.NewGetCurrentUserParams(),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
