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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AutoOpsSettings AutoOps settings for this deployment.
//
// swagger:model AutoOpsSettings
type AutoOpsSettings struct {

	// See AutoOps integration status for this deployment.
	// Required: true
	// Enum: [connected not_connected excluded]
	Status *string `json:"status"`
}

// Validate validates this auto ops settings
func (m *AutoOpsSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var autoOpsSettingsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","not_connected","excluded"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoOpsSettingsTypeStatusPropEnum = append(autoOpsSettingsTypeStatusPropEnum, v)
	}
}

const (

	// AutoOpsSettingsStatusConnected captures enum value "connected"
	AutoOpsSettingsStatusConnected string = "connected"

	// AutoOpsSettingsStatusNotConnected captures enum value "not_connected"
	AutoOpsSettingsStatusNotConnected string = "not_connected"

	// AutoOpsSettingsStatusExcluded captures enum value "excluded"
	AutoOpsSettingsStatusExcluded string = "excluded"
)

// prop value enum
func (m *AutoOpsSettings) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, autoOpsSettingsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AutoOpsSettings) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this auto ops settings based on context it is used
func (m *AutoOpsSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AutoOpsSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoOpsSettings) UnmarshalBinary(b []byte) error {
	var res AutoOpsSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}