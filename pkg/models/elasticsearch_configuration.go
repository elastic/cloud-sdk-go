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
)

// ElasticsearchConfiguration The Elasticsearch cluster settings. When specified at the top level, provides a field-by-field default. When specified at the topology level, provides the override settings.
//
// swagger:model ElasticsearchConfiguration
type ElasticsearchConfiguration struct {

	// Defines the index curation routing for the cluster
	Curation *ElasticsearchCuration `json:"curation,omitempty"`

	// A docker URI that allows overriding of the default docker image specified for this version
	DockerImage string `json:"docker_image,omitempty"`

	// A list of plugin names from the Elastic-supported subset that are bundled with the version images. NOTES: (Users should consult the Elastic stack objects to see what plugins are available, this is currently only available from the UI)
	EnabledBuiltInPlugins []string `json:"enabled_built_in_plugins,omitempty"`

	// Defines the Elasticsearch node attributes for the instances in the topology
	NodeAttributes map[string]string `json:"node_attributes,omitempty"`

	// system settings
	SystemSettings *ElasticsearchSystemSettings `json:"system_settings,omitempty"`

	// A list of admin-uploaded bundle objects (eg scripts, synonym files) that are available for this user.
	UserBundles []*ElasticsearchUserBundle `json:"user_bundles,omitempty"`

	// A list of admin-uploaded plugin objects that are available for this user.
	UserPlugins []*ElasticsearchUserPlugin `json:"user_plugins,omitempty"`

	// An arbitrary JSON object allowing cluster owners to set their parameters (only one of this and 'user_settings_yaml' is allowed), provided the parameters arey are on the allowlist and not on the denylist. NOTES: (This field together with 'user_settings_override*' and 'system_settings' defines the total set of Elasticsearch settings)
	UserSettingsJSON interface{} `json:"user_settings_json,omitempty"`

	// An arbitrary JSON object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_yaml' is allowed), ie in addition to the documented 'system_settings'. NOTES: (This field together with 'system_settings' and 'user_settings*' defines the total set of Elasticsearch settings)
	UserSettingsOverrideJSON interface{} `json:"user_settings_override_json,omitempty"`

	// An arbitrary YAML object allowing ECE admins owners to set clusters' parameters (only one of this and 'user_settings_override_json' is allowed), ie in addition to the documented 'system_settings'. NOTES: (This field together with 'system_settings' and 'user_settings*' defines the total set of Elasticsearch settings)
	UserSettingsOverrideYaml string `json:"user_settings_override_yaml,omitempty"`

	// An arbitrary YAML object allowing cluster owners to set their parameters (only one of this and 'user_settings_json' is allowed), provided the parameters arey are on the allowlist and not on the denylist. NOTES: (This field together with 'user_settings_override*' and 'system_settings' defines the total set of Elasticsearch settings)
	UserSettingsYaml string `json:"user_settings_yaml,omitempty"`

	// The version of the Elasticsearch cluster (must be one of the ECE supported versions). Currently cannot be different across the topology (and is generally specified in the globals). Defaults to the latest version if not specified.
	Version string `json:"version,omitempty"`
}

// Validate validates this elasticsearch configuration
func (m *ElasticsearchConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserBundles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserPlugins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchConfiguration) validateCuration(formats strfmt.Registry) error {
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

func (m *ElasticsearchConfiguration) validateSystemSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemSettings) { // not required
		return nil
	}

	if m.SystemSettings != nil {
		if err := m.SystemSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system_settings")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchConfiguration) validateUserBundles(formats strfmt.Registry) error {
	if swag.IsZero(m.UserBundles) { // not required
		return nil
	}

	for i := 0; i < len(m.UserBundles); i++ {
		if swag.IsZero(m.UserBundles[i]) { // not required
			continue
		}

		if m.UserBundles[i] != nil {
			if err := m.UserBundles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchConfiguration) validateUserPlugins(formats strfmt.Registry) error {
	if swag.IsZero(m.UserPlugins) { // not required
		return nil
	}

	for i := 0; i < len(m.UserPlugins); i++ {
		if swag.IsZero(m.UserPlugins[i]) { // not required
			continue
		}

		if m.UserPlugins[i] != nil {
			if err := m.UserPlugins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this elasticsearch configuration based on the context it is used
func (m *ElasticsearchConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserBundles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserPlugins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchConfiguration) contextValidateCuration(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ElasticsearchConfiguration) contextValidateSystemSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemSettings != nil {
		if err := m.SystemSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system_settings")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchConfiguration) contextValidateUserBundles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserBundles); i++ {

		if m.UserBundles[i] != nil {
			if err := m.UserBundles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_bundles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchConfiguration) contextValidateUserPlugins(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserPlugins); i++ {

		if m.UserPlugins[i] != nil {
			if err := m.UserPlugins[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("user_plugins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchConfiguration) UnmarshalBinary(b []byte) error {
	var res ElasticsearchConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
