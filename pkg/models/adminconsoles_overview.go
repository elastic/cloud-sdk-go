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

// AdminconsolesOverview Information about all of the adminconsoles.
//
// swagger:model AdminconsolesOverview
type AdminconsolesOverview struct {

	// List of adminconsoles
	// Required: true
	Adminconsoles []*AdminconsoleInfo `json:"adminconsoles"`
}

// Validate validates this adminconsoles overview
func (m *AdminconsolesOverview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminconsoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminconsolesOverview) validateAdminconsoles(formats strfmt.Registry) error {

	if err := validate.Required("adminconsoles", "body", m.Adminconsoles); err != nil {
		return err
	}

	for i := 0; i < len(m.Adminconsoles); i++ {
		if swag.IsZero(m.Adminconsoles[i]) { // not required
			continue
		}

		if m.Adminconsoles[i] != nil {
			if err := m.Adminconsoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adminconsoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adminconsoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this adminconsoles overview based on the context it is used
func (m *AdminconsolesOverview) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminconsoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminconsolesOverview) contextValidateAdminconsoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Adminconsoles); i++ {

		if m.Adminconsoles[i] != nil {
			if err := m.Adminconsoles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adminconsoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adminconsoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdminconsolesOverview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminconsolesOverview) UnmarshalBinary(b []byte) error {
	var res AdminconsolesOverview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
