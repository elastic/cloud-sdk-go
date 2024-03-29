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

// StackVersionConfigPost The configuration for an Elastic Stack version.
//
// swagger:model StackVersionConfigPost
type StackVersionConfigPost struct {

	// apm
	Apm *StackVersionApmConfig `json:"apm,omitempty"`

	// appsearch
	Appsearch *StackVersionAppSearchConfig `json:"appsearch,omitempty"`

	// elasticsearch
	// Required: true
	Elasticsearch *StackVersionElasticsearchConfig `json:"elasticsearch"`

	// enterprise search
	EnterpriseSearch *StackVersionEnterpriseSearchConfig `json:"enterprise_search,omitempty"`

	// integrations server
	IntegrationsServer *StackVersionIntegrationsServerConfig `json:"integrations_server,omitempty"`

	// kibana
	// Required: true
	Kibana *StackVersionKibanaConfig `json:"kibana"`

	// metadata
	Metadata *StackVersionMetadata `json:"metadata,omitempty"`
}

// Validate validates this stack version config post
func (m *StackVersionConfigPost) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEnterpriseSearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationsServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKibana(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackVersionConfigPost) validateApm(formats strfmt.Registry) error {
	if swag.IsZero(m.Apm) { // not required
		return nil
	}

	if m.Apm != nil {
		if err := m.Apm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apm")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) validateAppsearch(formats strfmt.Registry) error {
	if swag.IsZero(m.Appsearch) { // not required
		return nil
	}

	if m.Appsearch != nil {
		if err := m.Appsearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appsearch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appsearch")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) validateElasticsearch(formats strfmt.Registry) error {

	if err := validate.Required("elasticsearch", "body", m.Elasticsearch); err != nil {
		return err
	}

	if m.Elasticsearch != nil {
		if err := m.Elasticsearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elasticsearch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elasticsearch")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) validateEnterpriseSearch(formats strfmt.Registry) error {
	if swag.IsZero(m.EnterpriseSearch) { // not required
		return nil
	}

	if m.EnterpriseSearch != nil {
		if err := m.EnterpriseSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enterprise_search")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) validateIntegrationsServer(formats strfmt.Registry) error {
	if swag.IsZero(m.IntegrationsServer) { // not required
		return nil
	}

	if m.IntegrationsServer != nil {
		if err := m.IntegrationsServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrations_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integrations_server")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) validateKibana(formats strfmt.Registry) error {

	if err := validate.Required("kibana", "body", m.Kibana); err != nil {
		return err
	}

	if m.Kibana != nil {
		if err := m.Kibana.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kibana")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kibana")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this stack version config post based on the context it is used
func (m *StackVersionConfigPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppsearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElasticsearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnterpriseSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIntegrationsServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKibana(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackVersionConfigPost) contextValidateApm(ctx context.Context, formats strfmt.Registry) error {

	if m.Apm != nil {
		if err := m.Apm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("apm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("apm")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) contextValidateAppsearch(ctx context.Context, formats strfmt.Registry) error {

	if m.Appsearch != nil {
		if err := m.Appsearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appsearch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appsearch")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) contextValidateElasticsearch(ctx context.Context, formats strfmt.Registry) error {

	if m.Elasticsearch != nil {
		if err := m.Elasticsearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elasticsearch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elasticsearch")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) contextValidateEnterpriseSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.EnterpriseSearch != nil {
		if err := m.EnterpriseSearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enterprise_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enterprise_search")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) contextValidateIntegrationsServer(ctx context.Context, formats strfmt.Registry) error {

	if m.IntegrationsServer != nil {
		if err := m.IntegrationsServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrations_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integrations_server")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) contextValidateKibana(ctx context.Context, formats strfmt.Registry) error {

	if m.Kibana != nil {
		if err := m.Kibana.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kibana")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kibana")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionConfigPost) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackVersionConfigPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackVersionConfigPost) UnmarshalBinary(b []byte) error {
	var res StackVersionConfigPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
