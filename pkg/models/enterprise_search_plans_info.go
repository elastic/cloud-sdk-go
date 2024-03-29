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

// EnterpriseSearchPlansInfo Information about current, pending, and past Enterprise Search Server plans.
//
// swagger:model EnterpriseSearchPlansInfo
type EnterpriseSearchPlansInfo struct {

	// current
	Current *EnterpriseSearchPlanInfo `json:"current,omitempty"`

	// Whether the plan situation is healthy (if unhealthy, means the last plan attempt failed)
	// Required: true
	Healthy *bool `json:"healthy"`

	// history
	// Required: true
	History []*EnterpriseSearchPlanInfo `json:"history"`

	// pending
	Pending *EnterpriseSearchPlanInfo `json:"pending,omitempty"`
}

// Validate validates this enterprise search plans info
func (m *EnterpriseSearchPlansInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePending(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnterpriseSearchPlansInfo) validateCurrent(formats strfmt.Registry) error {
	if swag.IsZero(m.Current) { // not required
		return nil
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *EnterpriseSearchPlansInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *EnterpriseSearchPlansInfo) validateHistory(formats strfmt.Registry) error {

	if err := validate.Required("history", "body", m.History); err != nil {
		return err
	}

	for i := 0; i < len(m.History); i++ {
		if swag.IsZero(m.History[i]) { // not required
			continue
		}

		if m.History[i] != nil {
			if err := m.History[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("history" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnterpriseSearchPlansInfo) validatePending(formats strfmt.Registry) error {
	if swag.IsZero(m.Pending) { // not required
		return nil
	}

	if m.Pending != nil {
		if err := m.Pending.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pending")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pending")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this enterprise search plans info based on the context it is used
func (m *EnterpriseSearchPlansInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePending(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnterpriseSearchPlansInfo) contextValidateCurrent(ctx context.Context, formats strfmt.Registry) error {

	if m.Current != nil {
		if err := m.Current.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *EnterpriseSearchPlansInfo) contextValidateHistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.History); i++ {

		if m.History[i] != nil {
			if err := m.History[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("history" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnterpriseSearchPlansInfo) contextValidatePending(ctx context.Context, formats strfmt.Registry) error {

	if m.Pending != nil {
		if err := m.Pending.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pending")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pending")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnterpriseSearchPlansInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnterpriseSearchPlansInfo) UnmarshalBinary(b []byte) error {
	var res EnterpriseSearchPlansInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
