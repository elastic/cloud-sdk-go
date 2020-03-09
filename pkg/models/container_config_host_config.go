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

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContainerConfigHostConfig  Docker host config
// swagger:model ContainerConfigHostConfig
type ContainerConfigHostConfig struct {

	// Volume bindings for the container.
	// Required: true
	Binds []string `json:"binds"`

	// The length of a CPU period in microsecond
	CPUPeriod int32 `json:"cpu_period,omitempty"`

	// Ports that are exposed by the container.
	// Required: true
	ExtraHosts []string `json:"extra_hosts"`

	//  Sets the networking mode for the container.
	NetworkMode string `json:"network_mode,omitempty"`

	// Map of ports that should be exposed on the host.
	// Required: true
	PortBindings map[string][]PortBinding `json:"port_bindings"`

	// List of environment variables on the form KEY=value
	// Required: true
	Privileged *bool `json:"privileged"`

	// Docker behavior to apply when a container exits
	RestartPolicy *RestartPolicy `json:"restart_policy,omitempty"`
}

// Validate validates this container config host config
func (m *ContainerConfigHostConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBinds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtraHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortBindings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivileged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestartPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerConfigHostConfig) validateBinds(formats strfmt.Registry) error {

	if err := validate.Required("binds", "body", m.Binds); err != nil {
		return err
	}

	return nil
}

func (m *ContainerConfigHostConfig) validateExtraHosts(formats strfmt.Registry) error {

	if err := validate.Required("extra_hosts", "body", m.ExtraHosts); err != nil {
		return err
	}

	return nil
}

func (m *ContainerConfigHostConfig) validatePortBindings(formats strfmt.Registry) error {

	for k := range m.PortBindings {

		if err := validate.Required("port_bindings"+"."+k, "body", m.PortBindings[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.PortBindings[k]); i++ {

			if err := m.PortBindings[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("port_bindings" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *ContainerConfigHostConfig) validatePrivileged(formats strfmt.Registry) error {

	if err := validate.Required("privileged", "body", m.Privileged); err != nil {
		return err
	}

	return nil
}

func (m *ContainerConfigHostConfig) validateRestartPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.RestartPolicy) { // not required
		return nil
	}

	if m.RestartPolicy != nil {
		if err := m.RestartPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restart_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerConfigHostConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerConfigHostConfig) UnmarshalBinary(b []byte) error {
	var res ContainerConfigHostConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
