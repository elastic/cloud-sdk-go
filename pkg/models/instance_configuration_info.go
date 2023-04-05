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

// InstanceConfigurationInfo The configuration template for Elasticsearch instances, Kibana instances, and APM Servers.
//
// swagger:model InstanceConfigurationInfo
type InstanceConfigurationInfo struct {

	// Settings for the instance CPU multiplier
	CPUMultiplier float64 `json:"cpu_multiplier,omitempty"`

	// Optional description for the instance configuration
	Description string `json:"description,omitempty"`

	// Numerics representing possible instance sizes that the instance configuration supports.
	// Required: true
	DiscreteSizes *DiscreteSizes `json:"discrete_sizes"`

	// Unique identifier for the instance configuration
	ID string `json:"id,omitempty"`

	// The type of instance (elasticsearch, kibana)
	// Required: true
	InstanceType *string `json:"instance_type"`

	// The maximum number of availability zones in which this instance configuration has allocators. This field will be missing unless explicitly requested with the show_max_zones parameter.
	// Read Only: true
	MaxZones int32 `json:"max_zones,omitempty"`

	// Display name for the instance configuration.
	// Required: true
	Name *string `json:"name"`

	// Node types (master, data) for the instance
	NodeTypes []string `json:"node_types"`

	// Settings for the instance storage multiplier
	StorageMultiplier float64 `json:"storage_multiplier,omitempty"`
}

// Validate validates this instance configuration info
func (m *InstanceConfigurationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscreteSizes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceConfigurationInfo) validateDiscreteSizes(formats strfmt.Registry) error {

	if err := validate.Required("discrete_sizes", "body", m.DiscreteSizes); err != nil {
		return err
	}

	if m.DiscreteSizes != nil {
		if err := m.DiscreteSizes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discrete_sizes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discrete_sizes")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceConfigurationInfo) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instance_type", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *InstanceConfigurationInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this instance configuration info based on the context it is used
func (m *InstanceConfigurationInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDiscreteSizes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceConfigurationInfo) contextValidateDiscreteSizes(ctx context.Context, formats strfmt.Registry) error {

	if m.DiscreteSizes != nil {
		if err := m.DiscreteSizes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("discrete_sizes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("discrete_sizes")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceConfigurationInfo) contextValidateMaxZones(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "max_zones", "body", int32(m.MaxZones)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceConfigurationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceConfigurationInfo) UnmarshalBinary(b []byte) error {
	var res InstanceConfigurationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
