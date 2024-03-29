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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AllocatedInstanceStatus The status of the allocated Kibana instance or APM Server.
//
// swagger:model AllocatedInstanceStatus
type AllocatedInstanceStatus struct {

	// Indicates whether the cluster the instance belongs to is healthy
	ClusterHealthy *bool `json:"cluster_healthy,omitempty"`

	// Identifier for the cluster this instance belongs
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// Name of cluster this instance belongs, if available
	ClusterName string `json:"cluster_name,omitempty"`

	// Type of instance that is running
	// Required: true
	// Enum: [elasticsearch kibana apm integrations_server appsearch enterprise_search]
	ClusterType *string `json:"cluster_type"`

	// The id of the deployment this cluster belongs to.
	DeploymentID string `json:"deployment_id,omitempty"`

	// Indicates whether the instance is healthy
	Healthy *bool `json:"healthy,omitempty"`

	// The instance configuration id of this instance
	InstanceConfigurationID string `json:"instance_configuration_id,omitempty"`

	// Instance ID of the instance
	// Required: true
	InstanceName *string `json:"instance_name"`

	// Indicates whether the instance is vacating away from this allocator. Note that this is currently not populated when returned from the search endpoint.
	Moving *bool `json:"moving,omitempty"`

	// Memory assigned to this instance
	// Required: true
	NodeMemory *int32 `json:"node_memory"`

	// The plans associated with the current instance. Note that this is currently not populated when returned from the search endpoint.
	PlansInfo *AllocatedInstancePlansInfo `json:"plans_info,omitempty"`
}

// Validate validates this allocated instance status
func (m *AllocatedInstanceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlansInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatedInstanceStatus) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

var allocatedInstanceStatusTypeClusterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["elasticsearch","kibana","apm","integrations_server","appsearch","enterprise_search"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		allocatedInstanceStatusTypeClusterTypePropEnum = append(allocatedInstanceStatusTypeClusterTypePropEnum, v)
	}
}

const (

	// AllocatedInstanceStatusClusterTypeElasticsearch captures enum value "elasticsearch"
	AllocatedInstanceStatusClusterTypeElasticsearch string = "elasticsearch"

	// AllocatedInstanceStatusClusterTypeKibana captures enum value "kibana"
	AllocatedInstanceStatusClusterTypeKibana string = "kibana"

	// AllocatedInstanceStatusClusterTypeApm captures enum value "apm"
	AllocatedInstanceStatusClusterTypeApm string = "apm"

	// AllocatedInstanceStatusClusterTypeIntegrationsServer captures enum value "integrations_server"
	AllocatedInstanceStatusClusterTypeIntegrationsServer string = "integrations_server"

	// AllocatedInstanceStatusClusterTypeAppsearch captures enum value "appsearch"
	AllocatedInstanceStatusClusterTypeAppsearch string = "appsearch"

	// AllocatedInstanceStatusClusterTypeEnterpriseSearch captures enum value "enterprise_search"
	AllocatedInstanceStatusClusterTypeEnterpriseSearch string = "enterprise_search"
)

// prop value enum
func (m *AllocatedInstanceStatus) validateClusterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, allocatedInstanceStatusTypeClusterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AllocatedInstanceStatus) validateClusterType(formats strfmt.Registry) error {

	if err := validate.Required("cluster_type", "body", m.ClusterType); err != nil {
		return err
	}

	// value enum
	if err := m.validateClusterTypeEnum("cluster_type", "body", *m.ClusterType); err != nil {
		return err
	}

	return nil
}

func (m *AllocatedInstanceStatus) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instance_name", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *AllocatedInstanceStatus) validateNodeMemory(formats strfmt.Registry) error {

	if err := validate.Required("node_memory", "body", m.NodeMemory); err != nil {
		return err
	}

	return nil
}

func (m *AllocatedInstanceStatus) validatePlansInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.PlansInfo) { // not required
		return nil
	}

	if m.PlansInfo != nil {
		if err := m.PlansInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plans_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plans_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this allocated instance status based on the context it is used
func (m *AllocatedInstanceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlansInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatedInstanceStatus) contextValidatePlansInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.PlansInfo != nil {
		if err := m.PlansInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plans_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plans_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllocatedInstanceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocatedInstanceStatus) UnmarshalBinary(b []byte) error {
	var res AllocatedInstanceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
