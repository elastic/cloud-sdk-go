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

// MoveClustersDetails Information about the Elasticsearch clusters, multiple Kibana instances, and multiple APM Servers that are moved off of the allocator.
// swagger:model MoveClustersDetails
type MoveClustersDetails struct {

	// Detailed information about the Apm clusters being moved.
	ApmClusters []*MoveApmClusterDetails `json:"apm_clusters"`

	// Detailed information about the App Search clusters being moved.
	AppsearchClusters []*MoveAppSearchDetails `json:"appsearch_clusters"`

	// Detailed information about the Elasticsearch clusters being moved.
	ElasticsearchClusters []*MoveElasticsearchClusterDetails `json:"elasticsearch_clusters"`

	// Detailed information about the Kibana clusters being moved.
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

	if err := m.validateKibanaClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoveClustersDetails) validateApmClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.ApmClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.ApmClusters); i++ {
		if swag.IsZero(m.ApmClusters[i]) { // not required
			continue
		}

		if m.ApmClusters[i] != nil {
			if err := m.ApmClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apm_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) validateAppsearchClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.AppsearchClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.AppsearchClusters); i++ {
		if swag.IsZero(m.AppsearchClusters[i]) { // not required
			continue
		}

		if m.AppsearchClusters[i] != nil {
			if err := m.AppsearchClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appsearch_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) validateElasticsearchClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.ElasticsearchClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.ElasticsearchClusters); i++ {
		if swag.IsZero(m.ElasticsearchClusters[i]) { // not required
			continue
		}

		if m.ElasticsearchClusters[i] != nil {
			if err := m.ElasticsearchClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elasticsearch_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MoveClustersDetails) validateKibanaClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.KibanaClusters) { // not required
		return nil
	}

	for i := 0; i < len(m.KibanaClusters); i++ {
		if swag.IsZero(m.KibanaClusters[i]) { // not required
			continue
		}

		if m.KibanaClusters[i] != nil {
			if err := m.KibanaClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kibana_clusters" + "." + strconv.Itoa(i))
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
