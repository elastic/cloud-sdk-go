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

// TransientElasticsearchPlanConfiguration Defines the configuration parameters that control how the plan is applied. For example, the Elasticsearch cluster topology and Elasticsearch settings.
//
// swagger:model TransientElasticsearchPlanConfiguration
type TransientElasticsearchPlanConfiguration struct {

	// If specified, contains transient settings to be applied to an Elasticsearch cluster during changes, default values shown below applied.
	// These can be overridden by specifying them in the map (or null to unset). Additional settings can also be set. Settings will be cleared after the plan has finished. If not specified, no settings will be applied.
	// NOTE: These settings are only explicitly cleared for 5.x+ clusters, they must be hand-reset to their defaults in 2.x- (or a cluster reboot will clear them).
	// - indices.store.throttle.max_bytes_per_sec: 150Mb
	// - indices.recovery.max_bytes_per_sec: 150Mb
	// - cluster.routing.allocation.cluster_concurrent_rebalance: 10
	// - cluster.routing.allocation.node_initial_primaries_recoveries: 8
	// - cluster.routing.allocation.node_concurrent_incoming_recoveries: 8
	//
	ClusterSettingsJSON interface{} `json:"cluster_settings_json,omitempty"`

	// plan configuration
	PlanConfiguration *ElasticsearchPlanControlConfiguration `json:"plan_configuration,omitempty"`

	// The list of resources that will be configured as remote clusters
	RemoteClusters *RemoteResources `json:"remote_clusters,omitempty"`

	// restore snapshot
	RestoreSnapshot *RestoreSnapshotConfiguration `json:"restore_snapshot,omitempty"`

	// strategy
	Strategy *PlanStrategy `json:"strategy,omitempty"`
}

// Validate validates this transient elasticsearch plan configuration
func (m *TransientElasticsearchPlanConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlanConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestoreSnapshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransientElasticsearchPlanConfiguration) validatePlanConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.PlanConfiguration) { // not required
		return nil
	}

	if m.PlanConfiguration != nil {
		if err := m.PlanConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *TransientElasticsearchPlanConfiguration) validateRemoteClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteClusters) { // not required
		return nil
	}

	if m.RemoteClusters != nil {
		if err := m.RemoteClusters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote_clusters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remote_clusters")
			}
			return err
		}
	}

	return nil
}

func (m *TransientElasticsearchPlanConfiguration) validateRestoreSnapshot(formats strfmt.Registry) error {
	if swag.IsZero(m.RestoreSnapshot) { // not required
		return nil
	}

	if m.RestoreSnapshot != nil {
		if err := m.RestoreSnapshot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restore_snapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restore_snapshot")
			}
			return err
		}
	}

	return nil
}

func (m *TransientElasticsearchPlanConfiguration) validateStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.Strategy) { // not required
		return nil
	}

	if m.Strategy != nil {
		if err := m.Strategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this transient elasticsearch plan configuration based on the context it is used
func (m *TransientElasticsearchPlanConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlanConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestoreSnapshot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransientElasticsearchPlanConfiguration) contextValidatePlanConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.PlanConfiguration != nil {
		if err := m.PlanConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan_configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan_configuration")
			}
			return err
		}
	}

	return nil
}

func (m *TransientElasticsearchPlanConfiguration) contextValidateRemoteClusters(ctx context.Context, formats strfmt.Registry) error {

	if m.RemoteClusters != nil {
		if err := m.RemoteClusters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote_clusters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remote_clusters")
			}
			return err
		}
	}

	return nil
}

func (m *TransientElasticsearchPlanConfiguration) contextValidateRestoreSnapshot(ctx context.Context, formats strfmt.Registry) error {

	if m.RestoreSnapshot != nil {
		if err := m.RestoreSnapshot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restore_snapshot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("restore_snapshot")
			}
			return err
		}
	}

	return nil
}

func (m *TransientElasticsearchPlanConfiguration) contextValidateStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.Strategy != nil {
		if err := m.Strategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransientElasticsearchPlanConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransientElasticsearchPlanConfiguration) UnmarshalBinary(b []byte) error {
	var res TransientElasticsearchPlanConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
