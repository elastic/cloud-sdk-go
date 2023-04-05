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

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteUsersAPIKeysRequest The request payload that deletes the API keys.
//
// swagger:model DeleteUsersApiKeysRequest
type DeleteUsersAPIKeysRequest struct {

	// The list of API key IDs.
	// Required: true
	UserAPIKeys []*UserAPIKey `json:"user_api_keys"`
}

// Validate validates this delete users Api keys request
func (m *DeleteUsersAPIKeysRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserAPIKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteUsersAPIKeysRequest) validateUserAPIKeys(formats strfmt.Registry) error {

	if err := validate.Required("user_api_keys", "body", m.UserAPIKeys); err != nil {
		return err
	}

	for i := 0; i < len(m.UserAPIKeys); i++ {
		if swag.IsZero(m.UserAPIKeys[i]) { // not required
			continue
		}

		if m.UserAPIKeys[i] != nil {
			if err := m.UserAPIKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_api_keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("user_api_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete users Api keys request based on the context it is used
func (m *DeleteUsersAPIKeysRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserAPIKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteUsersAPIKeysRequest) contextValidateUserAPIKeys(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserAPIKeys); i++ {

		if m.UserAPIKeys[i] != nil {
			if err := m.UserAPIKeys[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_api_keys" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("user_api_keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteUsersAPIKeysRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteUsersAPIKeysRequest) UnmarshalBinary(b []byte) error {
	var res DeleteUsersAPIKeysRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
