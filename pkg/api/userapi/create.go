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
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/elastic/cloud-sdk-go/pkg/util/email"
)

const minPasswordLength = 8

// CreateParams is consumed by Create
type CreateParams struct {
	*api.API

	Password                  []byte
	Roles                     []string
	UserName, FullName, Email string
}

// Validate ensures the parameters are usable by the consuming function.
func (params CreateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid user params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.UserName == "" {
		merr = merr.Append(errors.New("username is not specified and is required for this operation"))
	}

	if len(params.Password) < minPasswordLength {
		merr = merr.Append(errors.New("a password with a minimum of 8 characters is required for this operation"))
	}

	if len(params.Roles) == 0 {
		merr = merr.Append(errors.New("a minimum of 1 role is required for this operation"))
	}

	if params.Email != "" {
		if err := email.Validate(params.Email); err != nil {
			merr = merr.Append(err)
		}
	}

	merr = merr.Append(ValidateRoles(params.Roles))

	return merr.ErrorOrNil()
}

// Create creates a new user.
func Create(params CreateParams) (*models.User, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.Users.CreateUser(
		users.NewCreateUserParams().
			WithBody(&models.User{
				UserName: &params.UserName,
				FullName: params.FullName,
				Email:    params.Email,
				Security: &models.UserSecurity{
					Enabled:  ec.Bool(true),
					Password: string(params.Password),
					Roles:    params.Roles,
				},
			}),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
