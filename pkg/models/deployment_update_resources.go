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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeploymentUpdateResources Describes the resources that will have the Deployment
// swagger:model DeploymentUpdateResources
type DeploymentUpdateResources struct {

	// A list of payloads for APM updates
	Apm []*ApmPayload `json:"apm"`

	// A list of payloads for AppSearch updates
	Appsearch []*AppSearchPayload `json:"appsearch"`

	// A list of payloads for Elasticsearch cluster updates
	Elasticsearch []*ElasticsearchPayload `json:"elasticsearch"`

	// A list of payloads for Kibana updates
	Kibana []*KibanaPayload `json:"kibana"`
}

// Validate validates this deployment update resources
func (m *DeploymentUpdateResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElasticsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKibana(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentUpdateResources) validateApm(formats strfmt.Registry) error {

	if swag.IsZero(m.Apm) { // not required
		return nil
	}

	for i := 0; i < len(m.Apm); i++ {
		if swag.IsZero(m.Apm[i]) { // not required
			continue
		}

		if m.Apm[i] != nil {
			if err := m.Apm[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apm" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentUpdateResources) validateAppsearch(formats strfmt.Registry) error {

	if swag.IsZero(m.Appsearch) { // not required
		return nil
	}

	for i := 0; i < len(m.Appsearch); i++ {
		if swag.IsZero(m.Appsearch[i]) { // not required
			continue
		}

		if m.Appsearch[i] != nil {
			if err := m.Appsearch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appsearch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentUpdateResources) validateElasticsearch(formats strfmt.Registry) error {

	if swag.IsZero(m.Elasticsearch) { // not required
		return nil
	}

	for i := 0; i < len(m.Elasticsearch); i++ {
		if swag.IsZero(m.Elasticsearch[i]) { // not required
			continue
		}

		if m.Elasticsearch[i] != nil {
			if err := m.Elasticsearch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elasticsearch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentUpdateResources) validateKibana(formats strfmt.Registry) error {

	if swag.IsZero(m.Kibana) { // not required
		return nil
	}

	for i := 0; i < len(m.Kibana); i++ {
		if swag.IsZero(m.Kibana[i]) { // not required
			continue
		}

		if m.Kibana[i] != nil {
			if err := m.Kibana[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kibana" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentUpdateResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentUpdateResources) UnmarshalBinary(b []byte) error {
	var res DeploymentUpdateResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}