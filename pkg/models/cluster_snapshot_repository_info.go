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
)

// ClusterSnapshotRepositoryInfo Information about the Elasticsearch cluster snapshot repository.
//
// swagger:model ClusterSnapshotRepositoryInfo
type ClusterSnapshotRepositoryInfo struct {

	// Cluster snapshot default repository settings
	Default ClusterSnapshotRepositoryDefault `json:"default,omitempty"`

	// Cluster snapshot reference repository settings, containing the repository name in ECE fashion
	Reference *ClusterSnapshotRepositoryReference `json:"reference,omitempty"`

	// Cluster snapshot static repository settings, containing repository type and settings
	Static *ClusterSnapshotRepositoryStatic `json:"static,omitempty"`
}

// Validate validates this cluster snapshot repository info
func (m *ClusterSnapshotRepositoryInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSnapshotRepositoryInfo) validateReference(formats strfmt.Registry) error {
	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if m.Reference != nil {
		if err := m.Reference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reference")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSnapshotRepositoryInfo) validateStatic(formats strfmt.Registry) error {
	if swag.IsZero(m.Static) { // not required
		return nil
	}

	if m.Static != nil {
		if err := m.Static.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster snapshot repository info based on the context it is used
func (m *ClusterSnapshotRepositoryInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSnapshotRepositoryInfo) contextValidateReference(ctx context.Context, formats strfmt.Registry) error {

	if m.Reference != nil {
		if err := m.Reference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reference")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSnapshotRepositoryInfo) contextValidateStatic(ctx context.Context, formats strfmt.Registry) error {

	if m.Static != nil {
		if err := m.Static.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSnapshotRepositoryInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSnapshotRepositoryInfo) UnmarshalBinary(b []byte) error {
	var res ClusterSnapshotRepositoryInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
