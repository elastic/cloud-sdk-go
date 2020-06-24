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
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeploymentSettings Additional configuration about the current deployment object.
// swagger:model DeploymentSettings
type DeploymentSettings struct {

	// The set of rulesets applies to this deployment.
	IPFilteringSettings *IPFilteringSettings `json:"ip_filtering_settings,omitempty"`
}

// Validate validates this deployment settings
func (m *DeploymentSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPFilteringSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentSettings) validateIPFilteringSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.IPFilteringSettings) { // not required
		return nil
	}

	if m.IPFilteringSettings != nil {
		if err := m.IPFilteringSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_filtering_settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentSettings) UnmarshalBinary(b []byte) error {
	var res DeploymentSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
