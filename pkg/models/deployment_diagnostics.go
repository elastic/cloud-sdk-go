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
)

// DeploymentDiagnostics Describes the diagnostics for a given Deployment-modifying payload
//
// swagger:model DeploymentDiagnostics
type DeploymentDiagnostics struct {

	// Diagnostics for resources to be created
	Creates *Creates `json:"creates,omitempty"`

	// Diagnostics for existing resources that may be updated
	Updates *Updates `json:"updates,omitempty"`
}

// Validate validates this deployment diagnostics
func (m *DeploymentDiagnostics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentDiagnostics) validateCreates(formats strfmt.Registry) error {
	if swag.IsZero(m.Creates) { // not required
		return nil
	}

	if m.Creates != nil {
		if err := m.Creates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creates")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentDiagnostics) validateUpdates(formats strfmt.Registry) error {
	if swag.IsZero(m.Updates) { // not required
		return nil
	}

	if m.Updates != nil {
		if err := m.Updates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this deployment diagnostics based on the context it is used
func (m *DeploymentDiagnostics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentDiagnostics) contextValidateCreates(ctx context.Context, formats strfmt.Registry) error {

	if m.Creates != nil {
		if err := m.Creates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creates")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentDiagnostics) contextValidateUpdates(ctx context.Context, formats strfmt.Registry) error {

	if m.Updates != nil {
		if err := m.Updates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updates")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentDiagnostics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentDiagnostics) UnmarshalBinary(b []byte) error {
	var res DeploymentDiagnostics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
