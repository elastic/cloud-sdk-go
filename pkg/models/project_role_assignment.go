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

// ProjectRoleAssignment Assignment for a role with project scope.
//
// swagger:model ProjectRoleAssignment
type ProjectRoleAssignment struct {

	// When true, the role applies to all projects in the organization, otherwise the role is scoped to the projects specified in `project_ids`.
	All *bool `json:"all,omitempty"`

	// If provided, the user assigned this role assignment will be granted this application role when signing in to the project(s) specified in the role assignment.
	ApplicationRoles []string `json:"application_roles"`

	// The ID of the organization the role is scoped to.
	// Required: true
	OrganizationID *string `json:"organization_id"`

	// The IDs of the projects the role is scoped to. Must be absent if `all` is true, and present if `all` is false.
	ProjectIds []string `json:"project_ids"`

	// The ID of the role that is assigned.
	// Required: true
	RoleID *string `json:"role_id"`
}

// Validate validates this project role assignment
func (m *ProjectRoleAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectRoleAssignment) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organization_id", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectRoleAssignment) validateRoleID(formats strfmt.Registry) error {

	if err := validate.Required("role_id", "body", m.RoleID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this project role assignment based on context it is used
func (m *ProjectRoleAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectRoleAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectRoleAssignment) UnmarshalBinary(b []byte) error {
	var res ProjectRoleAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
