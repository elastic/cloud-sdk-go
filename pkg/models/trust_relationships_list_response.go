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

// TrustRelationshipsListResponse Contains a list of trust relationships
// swagger:model TrustRelationshipsListResponse
type TrustRelationshipsListResponse struct {

	// The trust relationships
	// Required: true
	TrustRelationships []*TrustRelationshipGetResponse `json:"trust_relationships"`
}

// Validate validates this trust relationships list response
func (m *TrustRelationshipsListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrustRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrustRelationshipsListResponse) validateTrustRelationships(formats strfmt.Registry) error {

	if err := validate.Required("trust_relationships", "body", m.TrustRelationships); err != nil {
		return err
	}

	for i := 0; i < len(m.TrustRelationships); i++ {
		if swag.IsZero(m.TrustRelationships[i]) { // not required
			continue
		}

		if m.TrustRelationships[i] != nil {
			if err := m.TrustRelationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trust_relationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrustRelationshipsListResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrustRelationshipsListResponse) UnmarshalBinary(b []byte) error {
	var res TrustRelationshipsListResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
