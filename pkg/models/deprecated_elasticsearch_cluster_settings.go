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

// DeprecatedElasticsearchClusterSettings The settings for an Elasticsearch cluster.
//
// swagger:model DeprecatedElasticsearchClusterSettings
type DeprecatedElasticsearchClusterSettings struct {

	// > WARNING
	// > This endpoint is deprecated and scheduled to be removed in the next major version. Use the /remote-clusters endpoints instead.
	//
	// Configuration of remote clusters.
	Ccs *CrossClusterSearchSettings `json:"ccs,omitempty"`

	// curation
	Curation *ClusterCurationSettings `json:"curation,omitempty"`

	// Threshold starting from which the number of instances in the cluster results in the introduction of dedicated masters. If the cluster is downscaled to a number of nodes below this one, dedicated masters will be removed. Limit is inclusive.
	DedicatedMastersThreshold int32 `json:"dedicated_masters_threshold,omitempty"`

	// > WARNING
	// > This endpoint is deprecated and scheduled to be removed in the next major version. Use traffic filter settings instead.
	//
	// The set of rulesets to apply for all the resources in this cluster. When specified the same rulesets will be applied to Kibana and APM clusters as well
	IPFiltering *IPFilteringSettings `json:"ip_filtering,omitempty"`

	// metadata
	Metadata *ClusterMetadataSettings `json:"metadata,omitempty"`

	// monitoring
	Monitoring *ManagedMonitoringSettings `json:"monitoring,omitempty"`

	// snapshot
	Snapshot *ClusterSnapshotSettings `json:"snapshot,omitempty"`

	// The rulesets to apply to all resources in this cluster. When specified the same rulesets will be applied to Kibana and APM clusters as well
	TrafficFilter *TrafficFilterSettings `json:"traffic_filter,omitempty"`
}

// Validate validates this deprecated elasticsearch cluster settings
func (m *DeprecatedElasticsearchClusterSettings) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeprecatedElasticsearchClusterSettings) validateCcs(formats strfmt.Registry) error {
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

func (m *DeprecatedElasticsearchClusterSettings) validateCuration(formats strfmt.Registry) error {
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

func (m *DeprecatedElasticsearchClusterSettings) validateIPFiltering(formats strfmt.Registry) error {
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

func (m *DeprecatedElasticsearchClusterSettings) validateMetadata(formats strfmt.Registry) error {
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

func (m *DeprecatedElasticsearchClusterSettings) validateMonitoring(formats strfmt.Registry) error {
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

func (m *DeprecatedElasticsearchClusterSettings) validateSnapshot(formats strfmt.Registry) error {
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

func (m *DeprecatedElasticsearchClusterSettings) validateTrafficFilter(formats strfmt.Registry) error {
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

// ContextValidate validate this deprecated elasticsearch cluster settings based on the context it is used
func (m *DeprecatedElasticsearchClusterSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCcs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPFiltering(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonitoring(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrafficFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeprecatedElasticsearchClusterSettings) contextValidateCcs(ctx context.Context, formats strfmt.Registry) error {

	if m.Ccs != nil {
		if err := m.Ccs.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ccs")
			}
			return err
		}
	}

	return nil
}

func (m *DeprecatedElasticsearchClusterSettings) contextValidateCuration(ctx context.Context, formats strfmt.Registry) error {

	if m.Curation != nil {
		if err := m.Curation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("curation")
			}
			return err
		}
	}

	return nil
}

func (m *DeprecatedElasticsearchClusterSettings) contextValidateIPFiltering(ctx context.Context, formats strfmt.Registry) error {

	if m.IPFiltering != nil {
		if err := m.IPFiltering.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_filtering")
			}
			return err
		}
	}

	return nil
}

func (m *DeprecatedElasticsearchClusterSettings) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *DeprecatedElasticsearchClusterSettings) contextValidateMonitoring(ctx context.Context, formats strfmt.Registry) error {

	if m.Monitoring != nil {
		if err := m.Monitoring.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoring")
			}
			return err
		}
	}

	return nil
}

func (m *DeprecatedElasticsearchClusterSettings) contextValidateSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.Snapshot != nil {
		if err := m.Snapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot")
			}
			return err
		}
	}

	return nil
}

func (m *DeprecatedElasticsearchClusterSettings) contextValidateTrafficFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.TrafficFilter != nil {
		if err := m.TrafficFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("traffic_filter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeprecatedElasticsearchClusterSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeprecatedElasticsearchClusterSettings) UnmarshalBinary(b []byte) error {
	var res DeprecatedElasticsearchClusterSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
