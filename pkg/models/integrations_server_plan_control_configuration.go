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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntegrationsServerPlanControlConfiguration The plan control configuration options for the Integrations Server.
//
// swagger:model IntegrationsServerPlanControlConfiguration
type IntegrationsServerPlanControlConfiguration struct {

	// This timeout determines how long to give a cluster after it responds to API calls before performing actual operations on it. It defaults to 5s
	CalmWaitTime int64 `json:"calm_wait_time,omitempty"`

	// Set to 'forced' to force a reboot as part of the upgrade plan
	// Enum: [forced]
	ClusterReboot string `json:"cluster_reboot,omitempty"`

	// If true (default false), does not clear the maintenance flag (which prevents its API from being accessed except by the constructor) on new instances added until after a snapshot has been restored, otherwise, the maintenance flag is cleared once the new instances successfully join the new cluster
	ExtendedMaintenance *bool `json:"extended_maintenance,omitempty"`

	// move allocators
	MoveAllocators []*AllocatorMoveRequest `json:"move_allocators"`

	// move instances
	MoveInstances []*InstanceMoveRequest `json:"move_instances"`

	// If false (the default), the plan fails if it determines that the request can result in unsafe operations. Setting this flag overrides some of these check restraints.
	OverrideFailsafe *bool `json:"override_failsafe,omitempty"`

	// List of allocators on which instances are placed if possible (if not possible/not specified then any available allocator with space is used)
	PreferredAllocators []string `json:"preferred_allocators"`

	// If true (default: false) does not allow re-using any existing instances currently in the cluster, ie even unchanged instances will be re-created
	ReallocateInstances *bool `json:"reallocate_instances,omitempty"`

	// The total timeout in seconds after which the plan is cancelled even if it is not complete. Defaults to 4x the max memory capacity per node (in MB)
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this integrations server plan control configuration
func (m *IntegrationsServerPlanControlConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterReboot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveAllocators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var integrationsServerPlanControlConfigurationTypeClusterRebootPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["forced"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationsServerPlanControlConfigurationTypeClusterRebootPropEnum = append(integrationsServerPlanControlConfigurationTypeClusterRebootPropEnum, v)
	}
}

const (

	// IntegrationsServerPlanControlConfigurationClusterRebootForced captures enum value "forced"
	IntegrationsServerPlanControlConfigurationClusterRebootForced string = "forced"
)

// prop value enum
func (m *IntegrationsServerPlanControlConfiguration) validateClusterRebootEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, integrationsServerPlanControlConfigurationTypeClusterRebootPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntegrationsServerPlanControlConfiguration) validateClusterReboot(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterReboot) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterRebootEnum("cluster_reboot", "body", m.ClusterReboot); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationsServerPlanControlConfiguration) validateMoveAllocators(formats strfmt.Registry) error {
	if swag.IsZero(m.MoveAllocators) { // not required
		return nil
	}

	for i := 0; i < len(m.MoveAllocators); i++ {
		if swag.IsZero(m.MoveAllocators[i]) { // not required
			continue
		}

		if m.MoveAllocators[i] != nil {
			if err := m.MoveAllocators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("move_allocators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("move_allocators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IntegrationsServerPlanControlConfiguration) validateMoveInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.MoveInstances) { // not required
		return nil
	}

	for i := 0; i < len(m.MoveInstances); i++ {
		if swag.IsZero(m.MoveInstances[i]) { // not required
			continue
		}

		if m.MoveInstances[i] != nil {
			if err := m.MoveInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("move_instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("move_instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this integrations server plan control configuration based on the context it is used
func (m *IntegrationsServerPlanControlConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMoveAllocators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMoveInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsServerPlanControlConfiguration) contextValidateMoveAllocators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MoveAllocators); i++ {

		if m.MoveAllocators[i] != nil {
			if err := m.MoveAllocators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("move_allocators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("move_allocators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IntegrationsServerPlanControlConfiguration) contextValidateMoveInstances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MoveInstances); i++ {

		if m.MoveInstances[i] != nil {
			if err := m.MoveInstances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("move_instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("move_instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsServerPlanControlConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsServerPlanControlConfiguration) UnmarshalBinary(b []byte) error {
	var res IntegrationsServerPlanControlConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
