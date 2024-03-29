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

// RegionInfo Information about a region.
//
// swagger:model RegionInfo
type RegionInfo struct {

	// Information about allocators
	// Required: true
	Allocators *AllocatorsSummary `json:"allocators"`

	// Information about constructors
	// Required: true
	Constructors *ConstructorOverview `json:"constructors"`

	// Information about container sets
	// Required: true
	ContainerSetsStatus *ContainerSetsSummary `json:"container_sets_status"`

	// Information about coordinators
	// Required: true
	Coordinators *CoordinatorsSummary `json:"coordinators"`

	// Information about proxies
	// Required: true
	Proxies *ProxiesSummary `json:"proxies"`

	// Identifier of this region
	// Required: true
	RegionID *string `json:"region_id"`

	// Information about resources
	// Required: true
	Resources *PlatformResourcesSummary `json:"resources"`

	// Information about runners
	// Required: true
	Runners *RunnersSummary `json:"runners"`

	// Information about the Zookeeper state
	// Required: true
	ZookeeperStates *ZookeeperSummary `json:"zookeeper_states"`
}

// Validate validates this region info
func (m *RegionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstructors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerSetsStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoordinators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZookeeperStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegionInfo) validateAllocators(formats strfmt.Registry) error {

	if err := validate.Required("allocators", "body", m.Allocators); err != nil {
		return err
	}

	if m.Allocators != nil {
		if err := m.Allocators.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocators")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocators")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) validateConstructors(formats strfmt.Registry) error {

	if err := validate.Required("constructors", "body", m.Constructors); err != nil {
		return err
	}

	if m.Constructors != nil {
		if err := m.Constructors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constructors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constructors")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) validateContainerSetsStatus(formats strfmt.Registry) error {

	if err := validate.Required("container_sets_status", "body", m.ContainerSetsStatus); err != nil {
		return err
	}

	if m.ContainerSetsStatus != nil {
		if err := m.ContainerSetsStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container_sets_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("container_sets_status")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) validateCoordinators(formats strfmt.Registry) error {

	if err := validate.Required("coordinators", "body", m.Coordinators); err != nil {
		return err
	}

	if m.Coordinators != nil {
		if err := m.Coordinators.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coordinators")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("coordinators")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) validateProxies(formats strfmt.Registry) error {

	if err := validate.Required("proxies", "body", m.Proxies); err != nil {
		return err
	}

	if m.Proxies != nil {
		if err := m.Proxies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxies")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *RegionInfo) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) validateRunners(formats strfmt.Registry) error {

	if err := validate.Required("runners", "body", m.Runners); err != nil {
		return err
	}

	if m.Runners != nil {
		if err := m.Runners.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runners")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runners")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) validateZookeeperStates(formats strfmt.Registry) error {

	if err := validate.Required("zookeeper_states", "body", m.ZookeeperStates); err != nil {
		return err
	}

	if m.ZookeeperStates != nil {
		if err := m.ZookeeperStates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zookeeper_states")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zookeeper_states")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this region info based on the context it is used
func (m *RegionInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstructors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContainerSetsStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCoordinators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProxies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunners(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZookeeperStates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegionInfo) contextValidateAllocators(ctx context.Context, formats strfmt.Registry) error {

	if m.Allocators != nil {
		if err := m.Allocators.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocators")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocators")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) contextValidateConstructors(ctx context.Context, formats strfmt.Registry) error {

	if m.Constructors != nil {
		if err := m.Constructors.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constructors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constructors")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) contextValidateContainerSetsStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerSetsStatus != nil {
		if err := m.ContainerSetsStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container_sets_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("container_sets_status")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) contextValidateCoordinators(ctx context.Context, formats strfmt.Registry) error {

	if m.Coordinators != nil {
		if err := m.Coordinators.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coordinators")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("coordinators")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) contextValidateProxies(ctx context.Context, formats strfmt.Registry) error {

	if m.Proxies != nil {
		if err := m.Proxies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proxies")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {
		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) contextValidateRunners(ctx context.Context, formats strfmt.Registry) error {

	if m.Runners != nil {
		if err := m.Runners.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runners")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runners")
			}
			return err
		}
	}

	return nil
}

func (m *RegionInfo) contextValidateZookeeperStates(ctx context.Context, formats strfmt.Registry) error {

	if m.ZookeeperStates != nil {
		if err := m.ZookeeperStates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zookeeper_states")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zookeeper_states")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegionInfo) UnmarshalBinary(b []byte) error {
	var res RegionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
