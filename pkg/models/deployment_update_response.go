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
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentUpdateResponse A response returned from the Deployment update endpoint
// swagger:model DeploymentUpdateResponse
type DeploymentUpdateResponse struct {

	// Backend diagnostics. Sent if validation is requested.
	Diagnostics *DeploymentDiagnostics `json:"diagnostics,omitempty"`

	// The id of the deployment
	// Required: true
	ID *string `json:"id"`

	// The name of the deployment
	// Required: true
	Name *string `json:"name"`

	// List of resources that are part of the deployment after the update operation.
	// Required: true
	Resources []*DeploymentResource `json:"resources"`

	// List of resources that have been shut down
	ShutdownResources *Orphaned `json:"shutdown_resources,omitempty"`
}

// Validate validates this deployment update response
func (m *DeploymentUpdateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiagnostics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShutdownResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentUpdateResponse) validateDiagnostics(formats strfmt.Registry) error {

	if swag.IsZero(m.Diagnostics) { // not required
		return nil
	}

	if m.Diagnostics != nil {
		if err := m.Diagnostics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("diagnostics")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentUpdateResponse) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentUpdateResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentUpdateResponse) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	for i := 0; i < len(m.Resources); i++ {
		if swag.IsZero(m.Resources[i]) { // not required
			continue
		}

		if m.Resources[i] != nil {
			if err := m.Resources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentUpdateResponse) validateShutdownResources(formats strfmt.Registry) error {

	if swag.IsZero(m.ShutdownResources) { // not required
		return nil
	}

	if m.ShutdownResources != nil {
		if err := m.ShutdownResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shutdown_resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentUpdateResponse) UnmarshalBinary(b []byte) error {
	var res DeploymentUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
