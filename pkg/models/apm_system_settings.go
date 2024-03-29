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

// ApmSystemSettings A structure that defines a curated subset of the APM Server settings.
// TIP: To define the complete set of APM Server setting, use `ApmSystemSettings` with `user_settings_override_` and `user_settings_`.
//
// swagger:model ApmSystemSettings
type ApmSystemSettings struct {

	// Optionally enable debug mode for APM servers - defaults false
	DebugEnabled *bool `json:"debug_enabled,omitempty"`

	// Optionally override the account within APM - defaults to a system account that always exists (if specified, the username must also be specified). Note that this field is never returned from the API, it is write only.
	ElasticsearchPassword string `json:"elasticsearch_password,omitempty"`

	// DEPRECATED: Scheduled for removal in a future version of the API.
	//
	// Optionally override the URL to which to send data (for advanced users only, if unspecified the system selects an internal URL)
	ElasticsearchURL string `json:"elasticsearch_url,omitempty"`

	// Optionally override the account within APM - defaults to a system account that always exists (if specified, the password must also be specified). Note that this field is never returned from the API, it is write only.
	ElasticsearchUsername string `json:"elasticsearch_username,omitempty"`

	// DEPRECATED: Scheduled for removal in a future version of the API.
	//
	// Optionally override the URL to which to send data (for advanced users only, if unspecified the system selects an internal URL)
	KibanaURL string `json:"kibana_url,omitempty"`

	// Optionally override the secret token within APM - defaults to the previously existing secretToken
	SecretToken string `json:"secret_token,omitempty"`
}

// Validate validates this apm system settings
func (m *ApmSystemSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this apm system settings based on context it is used
func (m *ApmSystemSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApmSystemSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApmSystemSettings) UnmarshalBinary(b []byte) error {
	var res ApmSystemSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
