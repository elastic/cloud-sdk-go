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

// PlanStrategy The options for performing a plan change. Specify only one property each time. The default is `grow_and_shrink`.
//
// swagger:model PlanStrategy
type PlanStrategy struct {

	// autodetect
	Autodetect AutodetectStrategyConfig `json:"autodetect,omitempty"`

	// grow and shrink
	GrowAndShrink GrowShrinkStrategyConfig `json:"grow_and_shrink,omitempty"`

	// rolling
	Rolling *RollingStrategyConfig `json:"rolling,omitempty"`

	// rolling grow and shrink
	RollingGrowAndShrink RollingGrowShrinkStrategyConfig `json:"rolling_grow_and_shrink,omitempty"`
}

// Validate validates this plan strategy
func (m *PlanStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRolling(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanStrategy) validateRolling(formats strfmt.Registry) error {
	if swag.IsZero(m.Rolling) { // not required
		return nil
	}

	if m.Rolling != nil {
		if err := m.Rolling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rolling")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plan strategy based on the context it is used
func (m *PlanStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRolling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanStrategy) contextValidateRolling(ctx context.Context, formats strfmt.Registry) error {

	if m.Rolling != nil {
		if err := m.Rolling.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rolling")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanStrategy) UnmarshalBinary(b []byte) error {
	var res PlanStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
