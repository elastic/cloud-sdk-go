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
	"encoding/json"
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/users"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// EnableParams is consumed by Enable
type EnableParams struct {
	*api.API

	Enabled  bool
	UserName string
}

// Validate ensures the parameters are usable by the consuming function.
func (params EnableParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid user params")
	if params.UserName == "" {
		merr = merr.Append(errors.New("username is not specified and is required for this operation"))
	}

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	return merr.ErrorOrNil()
}

// Enable enables or disables an existing user.
func Enable(params EnableParams) (*models.User, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	user := &models.User{
		UserName: &params.UserName,
		Security: &models.UserSecurity{
			Enabled: &params.Enabled,
		},
	}

	b, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	res, err := params.V1API.Users.UpdateUser(
		users.NewUpdateUserParams().
			WithUserName(params.UserName).
			WithBody(string(b)),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
