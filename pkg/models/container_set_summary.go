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

// ContainerSetSummary Summarized information about a container set
//
// swagger:model ContainerSetSummary
type ContainerSetSummary struct {

	// The identifier for this container set
	// Required: true
	ContainerSetID *string `json:"container_set_id"`

	// Number of containers created in this container set
	// Required: true
	ContainersCreatedCount *int32 `json:"containers_created_count"`

	// Expected number of running containers in this container set
	// Required: true
	ContainersExpectedRunningCount *int32 `json:"containers_expected_running_count"`

	// Number of containers running in this container set
	// Required: true
	ContainersRunningCount *int32 `json:"containers_running_count"`

	// Number of containers started in this container set
	// Required: true
	ContainersStartedCount *int32 `json:"containers_started_count"`

	// Whether the container set is healthy
	// Required: true
	Healthy *bool `json:"healthy"`

	// True if the container set is hidden
	// Required: true
	Hidden *bool `json:"hidden"`
}

// Validate validates this container set summary
func (m *ContainerSetSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerSetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainersCreatedCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainersExpectedRunningCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainersRunningCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainersStartedCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHidden(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerSetSummary) validateContainerSetID(formats strfmt.Registry) error {

	if err := validate.Required("container_set_id", "body", m.ContainerSetID); err != nil {
		return err
	}

	return nil
}

func (m *ContainerSetSummary) validateContainersCreatedCount(formats strfmt.Registry) error {

	if err := validate.Required("containers_created_count", "body", m.ContainersCreatedCount); err != nil {
		return err
	}

	return nil
}

func (m *ContainerSetSummary) validateContainersExpectedRunningCount(formats strfmt.Registry) error {

	if err := validate.Required("containers_expected_running_count", "body", m.ContainersExpectedRunningCount); err != nil {
		return err
	}

	return nil
}

func (m *ContainerSetSummary) validateContainersRunningCount(formats strfmt.Registry) error {

	if err := validate.Required("containers_running_count", "body", m.ContainersRunningCount); err != nil {
		return err
	}

	return nil
}

func (m *ContainerSetSummary) validateContainersStartedCount(formats strfmt.Registry) error {

	if err := validate.Required("containers_started_count", "body", m.ContainersStartedCount); err != nil {
		return err
	}

	return nil
}

func (m *ContainerSetSummary) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ContainerSetSummary) validateHidden(formats strfmt.Registry) error {

	if err := validate.Required("hidden", "body", m.Hidden); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this container set summary based on context it is used
func (m *ContainerSetSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerSetSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerSetSummary) UnmarshalBinary(b []byte) error {
	var res ContainerSetSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
