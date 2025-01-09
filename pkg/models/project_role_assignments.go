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
)

// ProjectRoleAssignments Assignments for roles with project scope.
//
// swagger:model ProjectRoleAssignments
type ProjectRoleAssignments struct {

	// The Elasticsearch project-scoped role assignments to set
	Elasticsearch []*ProjectRoleAssignment `json:"elasticsearch"`

	// The Observability project-scoped role assignments to set
	Observability []*ProjectRoleAssignment `json:"observability"`

	// The Security project-scoped role assignments to set
	Security []*ProjectRoleAssignment `json:"security"`
}

// Validate validates this project role assignments
func (m *ProjectRoleAssignments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElasticsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObservability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectRoleAssignments) validateElasticsearch(formats strfmt.Registry) error {
	if swag.IsZero(m.Elasticsearch) { // not required
		return nil
	}

	for i := 0; i < len(m.Elasticsearch); i++ {
		if swag.IsZero(m.Elasticsearch[i]) { // not required
			continue
		}

		if m.Elasticsearch[i] != nil {
			if err := m.Elasticsearch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elasticsearch" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elasticsearch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectRoleAssignments) validateObservability(formats strfmt.Registry) error {
	if swag.IsZero(m.Observability) { // not required
		return nil
	}

	for i := 0; i < len(m.Observability); i++ {
		if swag.IsZero(m.Observability[i]) { // not required
			continue
		}

		if m.Observability[i] != nil {
			if err := m.Observability[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("observability" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("observability" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectRoleAssignments) validateSecurity(formats strfmt.Registry) error {
	if swag.IsZero(m.Security) { // not required
		return nil
	}

	for i := 0; i < len(m.Security); i++ {
		if swag.IsZero(m.Security[i]) { // not required
			continue
		}

		if m.Security[i] != nil {
			if err := m.Security[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("security" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("security" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this project role assignments based on the context it is used
func (m *ProjectRoleAssignments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElasticsearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObservability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectRoleAssignments) contextValidateElasticsearch(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Elasticsearch); i++ {

		if m.Elasticsearch[i] != nil {
			if err := m.Elasticsearch[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elasticsearch" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elasticsearch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectRoleAssignments) contextValidateObservability(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Observability); i++ {

		if m.Observability[i] != nil {
			if err := m.Observability[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("observability" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("observability" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectRoleAssignments) contextValidateSecurity(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Security); i++ {

		if m.Security[i] != nil {
			if err := m.Security[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("security" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("security" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectRoleAssignments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectRoleAssignments) UnmarshalBinary(b []byte) error {
	var res ProjectRoleAssignments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
