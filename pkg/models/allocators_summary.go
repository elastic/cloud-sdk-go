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

// AllocatorsSummary Summarized information about allocators.
//
// swagger:model AllocatorsSummary
type AllocatorsSummary struct {

	// Whether all allocators are healthy
	// Required: true
	Healthy *bool `json:"healthy"`

	// Summarized information on allocators in each zone
	// Required: true
	ZoneSummaries []*AllocatorsZoneSummary `json:"zone_summaries"`
}

// Validate validates this allocators summary
func (m *AllocatorsSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneSummaries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatorsSummary) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *AllocatorsSummary) validateZoneSummaries(formats strfmt.Registry) error {

	if err := validate.Required("zone_summaries", "body", m.ZoneSummaries); err != nil {
		return err
	}

	for i := 0; i < len(m.ZoneSummaries); i++ {
		if swag.IsZero(m.ZoneSummaries[i]) { // not required
			continue
		}

		if m.ZoneSummaries[i] != nil {
			if err := m.ZoneSummaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zone_summaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zone_summaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this allocators summary based on the context it is used
func (m *AllocatorsSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateZoneSummaries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatorsSummary) contextValidateZoneSummaries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ZoneSummaries); i++ {

		if m.ZoneSummaries[i] != nil {
			if err := m.ZoneSummaries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zone_summaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zone_summaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllocatorsSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocatorsSummary) UnmarshalBinary(b []byte) error {
	var res AllocatorsSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
