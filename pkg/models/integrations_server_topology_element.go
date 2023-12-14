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

// IntegrationsServerTopologyElement Defines the topology of the Integrations Server nodes. For example, the number or capacity of the nodes, and where you can allocate the nodes.
//
// swagger:model IntegrationsServerTopologyElement
type IntegrationsServerTopologyElement struct {

	// Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the id of an existing instance configuration.
	InstanceConfigurationID string `json:"instance_configuration_id,omitempty"`

	// The version of the Instance Configuration Id. If it is unset, the meaning depends on read vs writes. For deployment reads, it is equivalent to version 0 (or the IC is unversioned); for deployment creates and deployment template use, it is equivalent to 'the latest version'; and for deployment updates, it is equivalent to 'retain the current version'.
	InstanceConfigurationVersion *int32 `json:"instance_configuration_version,omitempty"`

	// integrations server
	IntegrationsServer *IntegrationsServerConfiguration `json:"integrations_server,omitempty"`

	// size
	Size *TopologySize `json:"size,omitempty"`

	// number of zones in which nodes will be placed
	ZoneCount int32 `json:"zone_count,omitempty"`
}

// Validate validates this integrations server topology element
func (m *IntegrationsServerTopologyElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrationsServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsServerTopologyElement) validateIntegrationsServer(formats strfmt.Registry) error {
	if swag.IsZero(m.IntegrationsServer) { // not required
		return nil
	}

	if m.IntegrationsServer != nil {
		if err := m.IntegrationsServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrations_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integrations_server")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationsServerTopologyElement) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if m.Size != nil {
		if err := m.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this integrations server topology element based on the context it is used
func (m *IntegrationsServerTopologyElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegrationsServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsServerTopologyElement) contextValidateIntegrationsServer(ctx context.Context, formats strfmt.Registry) error {

	if m.IntegrationsServer != nil {
		if err := m.IntegrationsServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrations_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integrations_server")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationsServerTopologyElement) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if m.Size != nil {
		if err := m.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsServerTopologyElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsServerTopologyElement) UnmarshalBinary(b []byte) error {
	var res IntegrationsServerTopologyElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
