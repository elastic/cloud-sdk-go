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

// InstanceTypeResource instance type resource
//
// swagger:model InstanceTypeResource
type InstanceTypeResource struct {

	// Compatible versions
	// Required: true
	Compatibility []*CompatibleVersionResource `json:"compatibility"`

	// Id of the default instance configuration to use for this instance type.
	DefaultInstanceConfigurationID string `json:"default_instance_configuration_id,omitempty"`

	// Instance type description
	// Required: true
	Description *string `json:"description"`

	// Instance type, the key for this resource
	// Required: true
	InstanceType *string `json:"instance_type"`

	// Instance type name
	// Required: true
	Name *string `json:"name"`

	// Supported node types
	// Required: true
	NodeTypes []*NodeTypeResource `json:"node_types"`
}

// Validate validates this instance type resource
func (m *InstanceTypeResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompatibility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceTypeResource) validateCompatibility(formats strfmt.Registry) error {

	if err := validate.Required("compatibility", "body", m.Compatibility); err != nil {
		return err
	}

	for i := 0; i < len(m.Compatibility); i++ {
		if swag.IsZero(m.Compatibility[i]) { // not required
			continue
		}

		if m.Compatibility[i] != nil {
			if err := m.Compatibility[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("compatibility" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("compatibility" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceTypeResource) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *InstanceTypeResource) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instance_type", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *InstanceTypeResource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InstanceTypeResource) validateNodeTypes(formats strfmt.Registry) error {

	if err := validate.Required("node_types", "body", m.NodeTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.NodeTypes); i++ {
		if swag.IsZero(m.NodeTypes[i]) { // not required
			continue
		}

		if m.NodeTypes[i] != nil {
			if err := m.NodeTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this instance type resource based on the context it is used
func (m *InstanceTypeResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompatibility(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceTypeResource) contextValidateCompatibility(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Compatibility); i++ {

		if m.Compatibility[i] != nil {
			if err := m.Compatibility[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("compatibility" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("compatibility" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceTypeResource) contextValidateNodeTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeTypes); i++ {

		if m.NodeTypes[i] != nil {
			if err := m.NodeTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("node_types" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("node_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceTypeResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceTypeResource) UnmarshalBinary(b []byte) error {
	var res InstanceTypeResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
