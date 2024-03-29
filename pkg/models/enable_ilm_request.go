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

// EnableIlmRequest Request sent to enable ILM on a deployment.
//
// swagger:model EnableIlmRequest
type EnableIlmRequest struct {

	// A locally-unique user-specified id for Kibana
	// Required: true
	IndexPatterns []*IndexPattern `json:"index_patterns"`
}

// Validate validates this enable ilm request
func (m *EnableIlmRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndexPatterns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnableIlmRequest) validateIndexPatterns(formats strfmt.Registry) error {

	if err := validate.Required("index_patterns", "body", m.IndexPatterns); err != nil {
		return err
	}

	for i := 0; i < len(m.IndexPatterns); i++ {
		if swag.IsZero(m.IndexPatterns[i]) { // not required
			continue
		}

		if m.IndexPatterns[i] != nil {
			if err := m.IndexPatterns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("index_patterns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("index_patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this enable ilm request based on the context it is used
func (m *EnableIlmRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIndexPatterns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnableIlmRequest) contextValidateIndexPatterns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IndexPatterns); i++ {

		if m.IndexPatterns[i] != nil {
			if err := m.IndexPatterns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("index_patterns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("index_patterns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnableIlmRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnableIlmRequest) UnmarshalBinary(b []byte) error {
	var res EnableIlmRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
