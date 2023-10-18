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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElasticsearchPlanControlConfiguration The configuration settings for the timeout and fallback parameters.
//
// swagger:model ElasticsearchPlanControlConfiguration
type ElasticsearchPlanControlConfiguration struct {

	// This timeout determines how long to give a cluster after it responds to API calls before performing actual operations on it. It defaults to 5s
	CalmWaitTime int64 `json:"calm_wait_time,omitempty"`

	// Set to 'forced' to force a reboot as part of the upgrade plan. NOTES: (ie taking an existing plan and leaving it alone except for setting 'transient.plan_configuration.cluster_reboot': 'forced' will reboot the cluster)
	// Enum: [forced]
	ClusterReboot string `json:"cluster_reboot,omitempty"`

	// If true (default false), does not clear the maintenance flag (which prevents its API from being accessed except by the constructor) on new instances added until after a snapshot has been restored, otherwise, the maintenance flag is cleared once the new instances successfully join the new cluster
	ExtendedMaintenance *bool `json:"extended_maintenance,omitempty"`

	// When you take a snapshot and 'skip_snapshots' is false, specifies the maximum age in seconds of the most recent snapshot before a new snapshot is created. Default is 300
	MaxSnapshotAge int64 `json:"max_snapshot_age,omitempty"`

	// If taking a snapshot (ie unless 'skip_snapshots': true) then will retry on failure at most this number of times (default: 5)
	MaxSnapshotAttempts int32 `json:"max_snapshot_attempts,omitempty"`

	// move allocators
	MoveAllocators []*AllocatorMoveRequest `json:"move_allocators"`

	// move instances
	MoveInstances []*InstanceMoveRequest `json:"move_instances"`

	// If true (default: false) only move_instances and move_allocators instructions will be executed, all other changes will be ignored
	MoveOnly *bool `json:"move_only,omitempty"`

	// If false (the default) then the plan will fail out if it believes the requested sequence of operations can result in data loss - this flag will override some of these restraints
	OverrideFailsafe *bool `json:"override_failsafe,omitempty"`

	// Map containing allocators tags in form of key value pairs, increasing the likelihood during move requests for allocators with matching tags, to be selected as target allocators
	PreferredAllocatorTags map[string]string `json:"preferred_allocator_tags,omitempty"`

	// List of allocators on which instances are placed if possible (if not possible/not specified then any available allocator with space is used)
	PreferredAllocators []string `json:"preferred_allocators"`

	// If true (default: false) does not allow re-using any existing instances currently in the cluster, ie even unchanged instances will be re-created
	ReallocateInstances *bool `json:"reallocate_instances,omitempty"`

	// If true (default: false) then the plan will not wait for data to be migrated from old instances to new instances before continuing the plan (potentially deleting the old instances and losing data)
	SkipDataMigration *bool `json:"skip_data_migration,omitempty"`

	// If false (the default), the cluster will run (currently) 2.x->5.x operations for any plan change ending with a 5.x cluster (eg apply a cluster license, ensure Monitoring is configured)
	SkipPostUpgradeSteps *bool `json:"skip_post_upgrade_steps,omitempty"`

	// If true (default: false), does not take (or require) a successful snapshot to be taken before performing any potentially destructive changes to this cluster
	SkipSnapshot *bool `json:"skip_snapshot,omitempty"`

	// If false (the default), the cluster will perform a snapshot after a major version upgrade takes place
	SkipSnapshotPostMajorUpgrade *bool `json:"skip_snapshot_post_major_upgrade,omitempty"`

	// If false, the cluster is checked for issues that should be resolved before migration (eg contains old Lucene segments), if true this is bypassed
	SkipUpgradeChecker *bool `json:"skip_upgrade_checker,omitempty"`

	// The total timeout in seconds after which the plan is cancelled even if it is not complete. Defaults to 4x the max memory capacity per node (in MB). NOTES: A 3 zone cluster with 2 nodes of 2048 each would have a timeout of 4*2048=8192 seconds. Timeout does not include time required to run rollback actions.
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this elasticsearch plan control configuration
func (m *ElasticsearchPlanControlConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterReboot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveAllocators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var elasticsearchPlanControlConfigurationTypeClusterRebootPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["forced"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		elasticsearchPlanControlConfigurationTypeClusterRebootPropEnum = append(elasticsearchPlanControlConfigurationTypeClusterRebootPropEnum, v)
	}
}

const (

	// ElasticsearchPlanControlConfigurationClusterRebootForced captures enum value "forced"
	ElasticsearchPlanControlConfigurationClusterRebootForced string = "forced"
)

// prop value enum
func (m *ElasticsearchPlanControlConfiguration) validateClusterRebootEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, elasticsearchPlanControlConfigurationTypeClusterRebootPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ElasticsearchPlanControlConfiguration) validateClusterReboot(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterReboot) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterRebootEnum("cluster_reboot", "body", m.ClusterReboot); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchPlanControlConfiguration) validateMoveAllocators(formats strfmt.Registry) error {
	if swag.IsZero(m.MoveAllocators) { // not required
		return nil
	}

	for i := 0; i < len(m.MoveAllocators); i++ {
		if swag.IsZero(m.MoveAllocators[i]) { // not required
			continue
		}

		if m.MoveAllocators[i] != nil {
			if err := m.MoveAllocators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("move_allocators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("move_allocators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchPlanControlConfiguration) validateMoveInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.MoveInstances) { // not required
		return nil
	}

	for i := 0; i < len(m.MoveInstances); i++ {
		if swag.IsZero(m.MoveInstances[i]) { // not required
			continue
		}

		if m.MoveInstances[i] != nil {
			if err := m.MoveInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("move_instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("move_instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this elasticsearch plan control configuration based on the context it is used
func (m *ElasticsearchPlanControlConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMoveAllocators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMoveInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchPlanControlConfiguration) contextValidateMoveAllocators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MoveAllocators); i++ {

		if m.MoveAllocators[i] != nil {
			if err := m.MoveAllocators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("move_allocators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("move_allocators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchPlanControlConfiguration) contextValidateMoveInstances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MoveInstances); i++ {

		if m.MoveInstances[i] != nil {
			if err := m.MoveInstances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("move_instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("move_instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchPlanControlConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchPlanControlConfiguration) UnmarshalBinary(b []byte) error {
	var res ElasticsearchPlanControlConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
