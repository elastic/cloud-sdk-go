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

// InstanceMoveRequest The request that specifies the Elasticsearch or stateless (eg Kibana) instances to move to allocators as part of the upgrade plan. When used in conjunction with '\_\_all\_\_' (roll all instances as a single unit) strategy, these instances are not restarted, which can sometimes enable recovery plans when these instances are boot-looping.
//
// swagger:model InstanceMoveRequest
type InstanceMoveRequest struct {

	// The instance id that is going to be moved
	// Required: true
	From *string `json:"from"`

	// Tells the infrastructure that the instance should be considered as permanently down when deciding how to migrate data to new nodes. If left blank then the system will automatically decide (currently: will treat the instances as up)
	InstanceDown *bool `json:"instance_down,omitempty"`

	// An optional list of allocator ids to which the instance should be moved. If not specified then any available allocator can be used (including the current one if it is healthy)
	To []string `json:"to"`
}

// Validate validates this instance move request
func (m *InstanceMoveRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceMoveRequest) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this instance move request based on context it is used
func (m *InstanceMoveRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceMoveRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceMoveRequest) UnmarshalBinary(b []byte) error {
	var res InstanceMoveRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
