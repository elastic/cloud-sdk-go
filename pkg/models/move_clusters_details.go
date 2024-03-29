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

// MoveClustersDetails Information about the Elasticsearch clusters, multiple Kibana instances, and multiple APM Servers that are moved off of the allocator.
//
// swagger:model MoveClustersDetails
type MoveClustersDetails struct {

	// Detailed information about the Apm clusters being moved.
	// Required: true
	ApmClusters []*MoveApmClusterDetails `json:"apm_clusters"`

	// Detailed information about the App Search clusters being moved.
	// Required: true
	AppsearchClusters []*MoveAppSearchDetails `json:"appsearch_clusters"`

	// Detailed information about the Elasticsearch clusters being moved.
	// Required: true
	ElasticsearchClusters []*MoveElasticsearchClusterDetails `json:"elasticsearch_clusters"`

	// Detailed information about the Elastic Enterprise Search clusters being moved.
	// Required: true
	EnterpriseSearchClusters []*MoveEnterpriseSearchDetails `json:"enterprise_search_clusters"`

	// Detailed information about the Kibana clusters being moved.
	// Required: true
	KibanaClusters []*MoveKibanaClusterDetails `json:"kibana_clusters"`
}

// Validate validates this move clusters details
func (m *MoveClustersDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApmClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppsearchClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElasticsearchClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnterpriseSearchClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKibanaClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveClustersDetails) validateApmClusters(formats strfmt.Registry) error {

	if err := validate.Required("apm_clusters", "body", m.ApmClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.ApmClusters); i++ {
		if swag.IsZero(m.ApmClusters[i]) { // not required
			continue
		}

		if m.ApmClusters[i] != nil {
			if err := m.ApmClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apm_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apm_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) validateAppsearchClusters(formats strfmt.Registry) error {

	if err := validate.Required("appsearch_clusters", "body", m.AppsearchClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.AppsearchClusters); i++ {
		if swag.IsZero(m.AppsearchClusters[i]) { // not required
			continue
		}

		if m.AppsearchClusters[i] != nil {
			if err := m.AppsearchClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appsearch_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appsearch_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) validateElasticsearchClusters(formats strfmt.Registry) error {

	if err := validate.Required("elasticsearch_clusters", "body", m.ElasticsearchClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.ElasticsearchClusters); i++ {
		if swag.IsZero(m.ElasticsearchClusters[i]) { // not required
			continue
		}

		if m.ElasticsearchClusters[i] != nil {
			if err := m.ElasticsearchClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elasticsearch_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elasticsearch_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) validateEnterpriseSearchClusters(formats strfmt.Registry) error {

	if err := validate.Required("enterprise_search_clusters", "body", m.EnterpriseSearchClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.EnterpriseSearchClusters); i++ {
		if swag.IsZero(m.EnterpriseSearchClusters[i]) { // not required
			continue
		}

		if m.EnterpriseSearchClusters[i] != nil {
			if err := m.EnterpriseSearchClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enterprise_search_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("enterprise_search_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) validateKibanaClusters(formats strfmt.Registry) error {

	if err := validate.Required("kibana_clusters", "body", m.KibanaClusters); err != nil {
		return err
	}

	for i := 0; i < len(m.KibanaClusters); i++ {
		if swag.IsZero(m.KibanaClusters[i]) { // not required
			continue
		}

		if m.KibanaClusters[i] != nil {
			if err := m.KibanaClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kibana_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kibana_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this move clusters details based on the context it is used
func (m *MoveClustersDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApmClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAppsearchClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElasticsearchClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnterpriseSearchClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKibanaClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveClustersDetails) contextValidateApmClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ApmClusters); i++ {

		if m.ApmClusters[i] != nil {
			if err := m.ApmClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apm_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apm_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) contextValidateAppsearchClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppsearchClusters); i++ {

		if m.AppsearchClusters[i] != nil {
			if err := m.AppsearchClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appsearch_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appsearch_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) contextValidateElasticsearchClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ElasticsearchClusters); i++ {

		if m.ElasticsearchClusters[i] != nil {
			if err := m.ElasticsearchClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elasticsearch_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("elasticsearch_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) contextValidateEnterpriseSearchClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EnterpriseSearchClusters); i++ {

		if m.EnterpriseSearchClusters[i] != nil {
			if err := m.EnterpriseSearchClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enterprise_search_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("enterprise_search_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) contextValidateKibanaClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KibanaClusters); i++ {

		if m.KibanaClusters[i] != nil {
			if err := m.KibanaClusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kibana_clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kibana_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MoveClustersDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MoveClustersDetails) UnmarshalBinary(b []byte) error {
	var res MoveClustersDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
