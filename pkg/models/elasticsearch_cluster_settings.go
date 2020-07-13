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
)

// ElasticsearchClusterSettings The settings for an Elasticsearch cluster.
// swagger:model ElasticsearchClusterSettings
type ElasticsearchClusterSettings struct {

	// ccs
	Ccs *CrossClusterSearchSettings `json:"ccs,omitempty"`

	// curation
	Curation *ClusterCurationSettings `json:"curation,omitempty"`

	// Threshold starting from which the number of instances in the cluster results in the introduction of dedicated masters. If the cluster is downscaled to a number of nodes below this one, dedicated masters will be removed. Limit is inclusive.
	DedicatedMastersThreshold int32 `json:"dedicated_masters_threshold,omitempty"`

	// DEPRECATED (Scheduled to be removed in the next major version): The set of rulesets to apply for all the resources in this cluster. When specified the same rulesets will be applied to Kibana and APM clusters as well
	IPFiltering *IPFilteringSettings `json:"ip_filtering,omitempty"`

	// metadata
	Metadata *ClusterMetadataSettings `json:"metadata,omitempty"`

	// monitoring
	Monitoring *ManagedMonitoringSettings `json:"monitoring,omitempty"`

	// snapshot
	Snapshot *ClusterSnapshotSettings `json:"snapshot,omitempty"`

	// The rulesets to apply to all resources in this cluster. When specified the same rulesets will be applied to Kibana and APM clusters as well
	TrafficFilter *TrafficFilterSettings `json:"traffic_filter,omitempty"`

	// Configuration of trust with other clusters
	Trust *TrustSettings `json:"trust,omitempty"`
}

// Validate validates this elasticsearch cluster settings
func (m *ElasticsearchClusterSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPFiltering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoring(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrafficFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrust(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterSettings) validateCcs(formats strfmt.Registry) error {

	if swag.IsZero(m.Ccs) { // not required
		return nil
	}

	if m.Ccs != nil {
		if err := m.Ccs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ccs")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateCuration(formats strfmt.Registry) error {

	if swag.IsZero(m.Curation) { // not required
		return nil
	}

	if m.Curation != nil {
		if err := m.Curation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("curation")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateIPFiltering(formats strfmt.Registry) error {

	if swag.IsZero(m.IPFiltering) { // not required
		return nil
	}

	if m.IPFiltering != nil {
		if err := m.IPFiltering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_filtering")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateMonitoring(formats strfmt.Registry) error {

	if swag.IsZero(m.Monitoring) { // not required
		return nil
	}

	if m.Monitoring != nil {
		if err := m.Monitoring.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoring")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateSnapshot(formats strfmt.Registry) error {

	if swag.IsZero(m.Snapshot) { // not required
		return nil
	}

	if m.Snapshot != nil {
		if err := m.Snapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateTrafficFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.TrafficFilter) { // not required
		return nil
	}

	if m.TrafficFilter != nil {
		if err := m.TrafficFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("traffic_filter")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterSettings) validateTrust(formats strfmt.Registry) error {

	if swag.IsZero(m.Trust) { // not required
		return nil
	}

	if m.Trust != nil {
		if err := m.Trust.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trust")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClusterSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClusterSettings) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClusterSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
