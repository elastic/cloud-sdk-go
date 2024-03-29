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

// RulesetAssociations The configuration settings for the traffic filter.
//
// swagger:model RulesetAssociations
type RulesetAssociations struct {

	// List of associations
	// Required: true
	Associations []*FilterAssociation `json:"associations"`

	// Total number of associations. This includes associations the user does not have permission to view.
	// Required: true
	TotalAssociations *int32 `json:"total_associations"`
}

// Validate validates this ruleset associations
func (m *RulesetAssociations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalAssociations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RulesetAssociations) validateAssociations(formats strfmt.Registry) error {

	if err := validate.Required("associations", "body", m.Associations); err != nil {
		return err
	}

	for i := 0; i < len(m.Associations); i++ {
		if swag.IsZero(m.Associations[i]) { // not required
			continue
		}

		if m.Associations[i] != nil {
			if err := m.Associations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("associations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RulesetAssociations) validateTotalAssociations(formats strfmt.Registry) error {

	if err := validate.Required("total_associations", "body", m.TotalAssociations); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ruleset associations based on the context it is used
func (m *RulesetAssociations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssociations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RulesetAssociations) contextValidateAssociations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Associations); i++ {

		if m.Associations[i] != nil {
			if err := m.Associations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("associations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RulesetAssociations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RulesetAssociations) UnmarshalBinary(b []byte) error {
	var res RulesetAssociations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
