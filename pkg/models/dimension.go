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

// Dimension Billing dimension
//
// swagger:model Dimension
type Dimension struct {

	// The cost of the billing dimension
	// Required: true
	Cost *float64 `json:"cost"`

	// The type of the billing dimension
	// Required: true
	// Enum: [capacity data_in data_internode data_out storage_api storage_bytes]
	Type *string `json:"type"`
}

// Validate validates this dimension
func (m *Dimension) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dimension) validateCost(formats strfmt.Registry) error {

	if err := validate.Required("cost", "body", m.Cost); err != nil {
		return err
	}

	return nil
}

var dimensionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["capacity","data_in","data_internode","data_out","storage_api","storage_bytes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dimensionTypeTypePropEnum = append(dimensionTypeTypePropEnum, v)
	}
}

const (

	// DimensionTypeCapacity captures enum value "capacity"
	DimensionTypeCapacity string = "capacity"

	// DimensionTypeDataIn captures enum value "data_in"
	DimensionTypeDataIn string = "data_in"

	// DimensionTypeDataInternode captures enum value "data_internode"
	DimensionTypeDataInternode string = "data_internode"

	// DimensionTypeDataOut captures enum value "data_out"
	DimensionTypeDataOut string = "data_out"

	// DimensionTypeStorageAPI captures enum value "storage_api"
	DimensionTypeStorageAPI string = "storage_api"

	// DimensionTypeStorageBytes captures enum value "storage_bytes"
	DimensionTypeStorageBytes string = "storage_bytes"
)

// prop value enum
func (m *Dimension) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dimensionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Dimension) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dimension based on context it is used
func (m *Dimension) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Dimension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dimension) UnmarshalBinary(b []byte) error {
	var res Dimension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
