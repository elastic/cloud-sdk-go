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

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElasticsearchScriptTypeSettings Enables scripting for the specified type and controls other parameters. Store scripts in indexes (`stored`), upload in file bundles (`file`), or use in API requests (`inline`).
//
// swagger:model ElasticsearchScriptTypeSettings
type ElasticsearchScriptTypeSettings struct {

	// If enabled (default: true) then scripts are enabled, either for sandboxing languages (by default), or for all installed languages if 'sandbox_mode' is disabled (or for 6.x). NOTES: (Corresponds to the parameter 'script.file|stored/indexed|inline')
	Enabled *bool `json:"enabled,omitempty"`

	// If enabled (default: true) and this script type is enabled, then only the sandbox languages are allowed. By default the sandbox languages are painless, expressions and mustache, but this can be restricted via the 'painless_enabled', 'mustache_enabled' 'expression_enabled' settings.NOTES: Not supported in 6.x. (Corresponds to the parameters 'script.engine.[painless|mustache|expressions].[file|stored|inline]')
	SandboxMode *bool `json:"sandbox_mode,omitempty"`
}

// Validate validates this elasticsearch script type settings
func (m *ElasticsearchScriptTypeSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this elasticsearch script type settings based on context it is used
func (m *ElasticsearchScriptTypeSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchScriptTypeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchScriptTypeSettings) UnmarshalBinary(b []byte) error {
	var res ElasticsearchScriptTypeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
