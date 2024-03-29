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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuthenticationInfo A user's authentication info
//
// swagger:model AuthenticationInfo
type AuthenticationInfo struct {

	// Deprecated: The UTC time when elevated permissions will expire, if the user has elevated permissions
	// Format: date-time
	ElevatedPermissionsExpireAt strfmt.DateTime `json:"elevated_permissions_expire_at,omitempty"`

	// The UTC time when current authentication will expire. Applies to only token based authentication
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// Deprecated: True if the user has elevated permissions
	// Required: true
	HasElevatedPermissions *bool `json:"has_elevated_permissions"`

	// Deprecated: True if the user has an available TOTP device
	HasTotpDevice *bool `json:"has_totp_device,omitempty"`

	// The API to be used when refreshing the current user's JWT
	// Required: true
	RefreshTokenURL *string `json:"refresh_token_url"`

	// Deprecated: The TOTP device source
	TotpDeviceSource string `json:"totp_device_source,omitempty"`

	// Deprecated: URL for configuring an MFA TOTP device.  Does not apply when totp_device_source is 'native'.
	TotpDeviceSourceEnableMfaHref string `json:"totp_device_source_enable_mfa_href,omitempty"`
}

// Validate validates this authentication info
func (m *AuthenticationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElevatedPermissionsExpireAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasElevatedPermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefreshTokenURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthenticationInfo) validateElevatedPermissionsExpireAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ElevatedPermissionsExpireAt) { // not required
		return nil
	}

	if err := validate.FormatOf("elevated_permissions_expire_at", "body", "date-time", m.ElevatedPermissionsExpireAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationInfo) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationInfo) validateHasElevatedPermissions(formats strfmt.Registry) error {

	if err := validate.Required("has_elevated_permissions", "body", m.HasElevatedPermissions); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationInfo) validateRefreshTokenURL(formats strfmt.Registry) error {

	if err := validate.Required("refresh_token_url", "body", m.RefreshTokenURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this authentication info based on context it is used
func (m *AuthenticationInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthenticationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthenticationInfo) UnmarshalBinary(b []byte) error {
	var res AuthenticationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
