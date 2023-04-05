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

// Blessings Describes whether certain runners are blessed to run a Role.
//
// swagger:model Blessings
type Blessings struct {

	// A mapping of Runner Ids to whether or not they are blessed to run the associated role
	// Required: true
	RunnerIdsToBlessing map[string]Blessing `json:"runner_ids_to_blessing"`
}

// Validate validates this blessings
func (m *Blessings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunnerIdsToBlessing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Blessings) validateRunnerIdsToBlessing(formats strfmt.Registry) error {

	if err := validate.Required("runner_ids_to_blessing", "body", m.RunnerIdsToBlessing); err != nil {
		return err
	}

	for k := range m.RunnerIdsToBlessing {

		if err := validate.Required("runner_ids_to_blessing"+"."+k, "body", m.RunnerIdsToBlessing[k]); err != nil {
			return err
		}
		if val, ok := m.RunnerIdsToBlessing[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runner_ids_to_blessing" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("runner_ids_to_blessing" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this blessings based on the context it is used
func (m *Blessings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRunnerIdsToBlessing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Blessings) contextValidateRunnerIdsToBlessing(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("runner_ids_to_blessing", "body", m.RunnerIdsToBlessing); err != nil {
		return err
	}

	for k := range m.RunnerIdsToBlessing {

		if val, ok := m.RunnerIdsToBlessing[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Blessings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Blessings) UnmarshalBinary(b []byte) error {
	var res Blessings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
