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
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentTemplateDefinitionRequest The request body for a deployment templatee.
// swagger:model DeploymentTemplateDefinitionRequest
type DeploymentTemplateDefinitionRequest struct {

	// apm
	Apm *CreateApmInCreateElasticsearchRequest `json:"apm,omitempty"`

	// appsearch
	Appsearch *CreateAppSearchRequest `json:"appsearch,omitempty"`

	// The human readable name for the cluster (defaults to the generated cluster id if not specified)
	ClusterName string `json:"cluster_name,omitempty"`

	// enterprise search
	EnterpriseSearch *CreateEnterpriseSearchRequest `json:"enterprise_search,omitempty"`

	// kibana
	Kibana *CreateKibanaInCreateElasticsearchRequest `json:"kibana,omitempty"`

	// plan
	// Required: true
	Plan *ElasticsearchClusterPlan `json:"plan"`

	// settings
	Settings *ElasticsearchClusterSettings `json:"settings,omitempty"`
}

// Validate validates this deployment template definition request
func (m *DeploymentTemplateDefinitionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnterpriseSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKibana(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentTemplateDefinitionRequest) validateApm(formats strfmt.Registry) error {

	if swag.IsZero(m.Apm) { // not required
		return nil
	}

	if m.Apm != nil {
		if err := m.Apm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apm")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentTemplateDefinitionRequest) validateAppsearch(formats strfmt.Registry) error {

	if swag.IsZero(m.Appsearch) { // not required
		return nil
	}

	if m.Appsearch != nil {
		if err := m.Appsearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appsearch")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentTemplateDefinitionRequest) validateEnterpriseSearch(formats strfmt.Registry) error {

	if swag.IsZero(m.EnterpriseSearch) { // not required
		return nil
	}

	if m.EnterpriseSearch != nil {
		if err := m.EnterpriseSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise_search")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentTemplateDefinitionRequest) validateKibana(formats strfmt.Registry) error {

	if swag.IsZero(m.Kibana) { // not required
		return nil
	}

	if m.Kibana != nil {
		if err := m.Kibana.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kibana")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentTemplateDefinitionRequest) validatePlan(formats strfmt.Registry) error {

	if err := validate.Required("plan", "body", m.Plan); err != nil {
		return err
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentTemplateDefinitionRequest) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentTemplateDefinitionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentTemplateDefinitionRequest) UnmarshalBinary(b []byte) error {
	var res DeploymentTemplateDefinitionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
