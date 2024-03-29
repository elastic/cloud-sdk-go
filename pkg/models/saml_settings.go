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

// SamlSettings The configuration for the Elasticsearch security SAML realm.
//
// swagger:model SamlSettings
type SamlSettings struct {

	// The SAML attribute mapping configuration
	// Required: true
	Attributes *SamlAttributeSettings `json:"attributes"`

	// When true, enables the security realm
	Enabled *bool `json:"enabled,omitempty"`

	// The SAML encryption certificate bundle URL. The bundle should be a zip file containing 'encryption.key' and 'encryption.pem' files in the directory '/saml/:id', where :id is the value of the [id] field.
	EncryptionCertificateURL string `json:"encryption_certificate_url,omitempty"`

	// The password to the encryption certificate bundle
	EncryptionCertificateURLPassword string `json:"encryption_certificate_url_password,omitempty"`

	// Specifies whether to set the ForceAuthn attribute when requesting that the IdP authenticate the current user. If set to true, the IdP is required to verify the user's identity, irrespective of any existing sessions they might have.
	ForceAuthn *bool `json:"force_authn,omitempty"`

	// The identifier for the security realm
	// Required: true
	ID *string `json:"id"`

	// The SAML Identity Provider configuration
	// Required: true
	Idp *SamlIdpSettings `json:"idp"`

	// The friendly name of the security realm
	// Required: true
	Name *string `json:"name"`

	// The NameID format. If not specified the IdP default is used. Example: 'urn:oasis:names:tc:SAML:2.0:nameid-format:persistent'
	NameidFormat string `json:"nameid_format,omitempty"`

	// The order that the security realm is evaluated
	Order int32 `json:"order,omitempty"`

	// Advanced configuration options in YAML format. Any settings defined here will override any configuration set via the API. Note that all keys should omit 'xpack.security.authc.realms.{realm_type}.{realm_id}'.
	OverrideYaml string `json:"override_yaml,omitempty"`

	// The role mapping rules associated with the security realm
	RoleMappings *SamlSecurityRealmRoleMappingRules `json:"role_mappings,omitempty"`

	// The SAML signing certificate bundle URL. The bundle should be a zip file containing 'signing.key' and 'signing.pem' files in the directory '/saml/:id', where :id is the value of the [id] field.
	SigningCertificateURL string `json:"signing_certificate_url,omitempty"`

	// The password to the signing certificate bundle
	SigningCertificateURLPassword string `json:"signing_certificate_url_password,omitempty"`

	// A list of SAML message types that should be signed. Each element in the list should be the local name of a SAML XML Element. Supported element types are AuthnRequest, LogoutRequest and LogoutResponse. Only valid if a signing certificate is also specified.
	SigningSamlMessages []string `json:"signing_saml_messages"`

	// The SAML Service Provider configuration
	// Required: true
	Sp *SamlSpSettings `json:"sp"`

	// The SSL trusted CA certificate bundle URL. The bundle should be a zip file containing a single keystore file 'keystore.ks' Note that all keys should omit the 'xpack.security.authc.realms.saml.{realm_id}' prefix. For example, when the realm ID is set to 'saml1', the advanced configuration 'xpack.security.authc.realms.saml.saml1.ssl.verification_mode: full' should be added as 'ssl.verification_mode: full'.
	SslCertificateURL string `json:"ssl_certificate_url,omitempty"`

	// The password to the SSL certificate bundle URL truststore
	SslCertificateURLTruststorePassword string `json:"ssl_certificate_url_truststore_password,omitempty"`

	// The format of the keystore file. Should be jks to use the Java Keystore format or PKCS12 to use PKCS#12 files. The default is jks.
	// Enum: [jks PKCS12]
	SslCertificateURLTruststoreType string `json:"ssl_certificate_url_truststore_type,omitempty"`
}

// Validate validates this saml settings
func (m *SamlSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCertificateURLTruststoreType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SamlSettings) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *SamlSettings) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SamlSettings) validateIdp(formats strfmt.Registry) error {

	if err := validate.Required("idp", "body", m.Idp); err != nil {
		return err
	}

	if m.Idp != nil {
		if err := m.Idp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("idp")
			}
			return err
		}
	}

	return nil
}

func (m *SamlSettings) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SamlSettings) validateRoleMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleMappings) { // not required
		return nil
	}

	if m.RoleMappings != nil {
		if err := m.RoleMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_mappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_mappings")
			}
			return err
		}
	}

	return nil
}

func (m *SamlSettings) validateSp(formats strfmt.Registry) error {

	if err := validate.Required("sp", "body", m.Sp); err != nil {
		return err
	}

	if m.Sp != nil {
		if err := m.Sp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sp")
			}
			return err
		}
	}

	return nil
}

var samlSettingsTypeSslCertificateURLTruststoreTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["jks","PKCS12"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		samlSettingsTypeSslCertificateURLTruststoreTypePropEnum = append(samlSettingsTypeSslCertificateURLTruststoreTypePropEnum, v)
	}
}

const (

	// SamlSettingsSslCertificateURLTruststoreTypeJks captures enum value "jks"
	SamlSettingsSslCertificateURLTruststoreTypeJks string = "jks"

	// SamlSettingsSslCertificateURLTruststoreTypePKCS12 captures enum value "PKCS12"
	SamlSettingsSslCertificateURLTruststoreTypePKCS12 string = "PKCS12"
)

// prop value enum
func (m *SamlSettings) validateSslCertificateURLTruststoreTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, samlSettingsTypeSslCertificateURLTruststoreTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SamlSettings) validateSslCertificateURLTruststoreType(formats strfmt.Registry) error {
	if swag.IsZero(m.SslCertificateURLTruststoreType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslCertificateURLTruststoreTypeEnum("ssl_certificate_url_truststore_type", "body", m.SslCertificateURLTruststoreType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this saml settings based on the context it is used
func (m *SamlSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SamlSettings) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {
		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *SamlSettings) contextValidateIdp(ctx context.Context, formats strfmt.Registry) error {

	if m.Idp != nil {
		if err := m.Idp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("idp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("idp")
			}
			return err
		}
	}

	return nil
}

func (m *SamlSettings) contextValidateRoleMappings(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleMappings != nil {
		if err := m.RoleMappings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_mappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_mappings")
			}
			return err
		}
	}

	return nil
}

func (m *SamlSettings) contextValidateSp(ctx context.Context, formats strfmt.Registry) error {

	if m.Sp != nil {
		if err := m.Sp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SamlSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SamlSettings) UnmarshalBinary(b []byte) error {
	var res SamlSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
