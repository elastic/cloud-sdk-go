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

// EnterpriseSearchPlan The plan for the Enterprise Search cluster.
//
// swagger:model EnterpriseSearchPlan
type EnterpriseSearchPlan struct {

	// cluster topology
	ClusterTopology []*EnterpriseSearchTopologyElement `json:"cluster_topology"`

	// enterprise search
	// Required: true
	EnterpriseSearch *EnterpriseSearchConfiguration `json:"enterprise_search"`

	// transient
	Transient *TransientEnterpriseSearchPlanConfiguration `json:"transient,omitempty"`
}

// Validate validates this enterprise search plan
func (m *EnterpriseSearchPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterTopology(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnterpriseSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnterpriseSearchPlan) validateClusterTopology(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterTopology) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterTopology); i++ {
		if swag.IsZero(m.ClusterTopology[i]) { // not required
			continue
		}

		if m.ClusterTopology[i] != nil {
			if err := m.ClusterTopology[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnterpriseSearchPlan) validateEnterpriseSearch(formats strfmt.Registry) error {

	if err := validate.Required("enterprise_search", "body", m.EnterpriseSearch); err != nil {
		return err
	}

	if m.EnterpriseSearch != nil {
		if err := m.EnterpriseSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enterprise_search")
			}
			return err
		}
	}

	return nil
}

func (m *EnterpriseSearchPlan) validateTransient(formats strfmt.Registry) error {
	if swag.IsZero(m.Transient) { // not required
		return nil
	}

	if m.Transient != nil {
		if err := m.Transient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transient")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this enterprise search plan based on the context it is used
func (m *EnterpriseSearchPlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterTopology(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnterpriseSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnterpriseSearchPlan) contextValidateClusterTopology(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterTopology); i++ {

		if m.ClusterTopology[i] != nil {
			if err := m.ClusterTopology[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnterpriseSearchPlan) contextValidateEnterpriseSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.EnterpriseSearch != nil {
		if err := m.EnterpriseSearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enterprise_search")
			}
			return err
		}
	}

	return nil
}

func (m *EnterpriseSearchPlan) contextValidateTransient(ctx context.Context, formats strfmt.Registry) error {

	if m.Transient != nil {
		if err := m.Transient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transient")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnterpriseSearchPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnterpriseSearchPlan) UnmarshalBinary(b []byte) error {
	var res EnterpriseSearchPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
