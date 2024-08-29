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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TopologySize Measured by the amount of a resource. The final cluster size is calculated using multipliers from the topology instance configuration.
//
// swagger:model TopologySize
type TopologySize struct {

	// Type of resource. In ESS the resource used should always be `memory`.
	// Required: true
	// Enum: [memory storage]
	Resource *string `json:"resource"`

	// Amount of resource
	// Required: true
	Value *int32 `json:"value"`
}

// Validate validates this topology size
func (m *TopologySize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var topologySizeTypeResourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["memory","storage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		topologySizeTypeResourcePropEnum = append(topologySizeTypeResourcePropEnum, v)
	}
}

const (

	// TopologySizeResourceMemory captures enum value "memory"
	TopologySizeResourceMemory string = "memory"

	// TopologySizeResourceStorage captures enum value "storage"
	TopologySizeResourceStorage string = "storage"
)

// prop value enum
func (m *TopologySize) validateResourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, topologySizeTypeResourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TopologySize) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", m.Resource); err != nil {
		return err
	}

	// value enum
	if err := m.validateResourceEnum("resource", "body", *m.Resource); err != nil {
		return err
	}

	return nil
}

func (m *TopologySize) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this topology size based on context it is used
func (m *TopologySize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TopologySize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TopologySize) UnmarshalBinary(b []byte) error {
	var res TopologySize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
