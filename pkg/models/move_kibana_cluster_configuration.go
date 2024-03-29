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

// MoveKibanaClusterConfiguration The configuration object for moving Kibana instances.
//
// swagger:model MoveKibanaClusterConfiguration
type MoveKibanaClusterConfiguration struct {

	// Identifiers for the Kibana clusters.
	// Required: true
	ClusterIds []string `json:"cluster_ids"`

	// Plan override to apply to the Kibana clusters being moved.
	PlanOverride *TransientKibanaPlanConfiguration `json:"plan_override,omitempty"`
}

// Validate validates this move kibana cluster configuration
func (m *MoveKibanaClusterConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanOverride(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveKibanaClusterConfiguration) validateClusterIds(formats strfmt.Registry) error {

	if err := validate.Required("cluster_ids", "body", m.ClusterIds); err != nil {
		return err
	}

	return nil
}

func (m *MoveKibanaClusterConfiguration) validatePlanOverride(formats strfmt.Registry) error {
	if swag.IsZero(m.PlanOverride) { // not required
		return nil
	}

	if m.PlanOverride != nil {
		if err := m.PlanOverride.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_override")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan_override")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this move kibana cluster configuration based on the context it is used
func (m *MoveKibanaClusterConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlanOverride(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveKibanaClusterConfiguration) contextValidatePlanOverride(ctx context.Context, formats strfmt.Registry) error {

	if m.PlanOverride != nil {
		if err := m.PlanOverride.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_override")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan_override")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveKibanaClusterConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveKibanaClusterConfiguration) UnmarshalBinary(b []byte) error {
	var res MoveKibanaClusterConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
