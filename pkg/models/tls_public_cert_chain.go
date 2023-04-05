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

// TLSPublicCertChain The public portion of the certificate chain that contains the PEM encoded server certificate, intermediate certificates, and the CA certificate. NOTE: The private key, normally included in certificate chains, is omitted.
//
// swagger:model TlsPublicCertChain
type TLSPublicCertChain struct {

	// Whether or not this certificate can be updated using the API
	APIManaged *bool `json:"api_managed,omitempty"`

	// The list of PEM encoded X509 certificates that make up the certificate chain
	// Required: true
	Chain []string `json:"chain"`

	// Details on the validity and lifetime of the certification chain
	ChainStatus *ChainStatus `json:"chain_status,omitempty"`

	// Was this certificate chain user supplied or automatically generated?
	// Required: true
	UserSupplied *bool `json:"user_supplied"`
}

// Validate validates this Tls public cert chain
func (m *TLSPublicCertChain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChainStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserSupplied(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TLSPublicCertChain) validateChain(formats strfmt.Registry) error {

	if err := validate.Required("chain", "body", m.Chain); err != nil {
		return err
	}

	return nil
}

func (m *TLSPublicCertChain) validateChainStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ChainStatus) { // not required
		return nil
	}

	if m.ChainStatus != nil {
		if err := m.ChainStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chain_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chain_status")
			}
			return err
		}
	}

	return nil
}

func (m *TLSPublicCertChain) validateUserSupplied(formats strfmt.Registry) error {

	if err := validate.Required("user_supplied", "body", m.UserSupplied); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Tls public cert chain based on the context it is used
func (m *TLSPublicCertChain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChainStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TLSPublicCertChain) contextValidateChainStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.ChainStatus != nil {
		if err := m.ChainStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chain_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chain_status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TLSPublicCertChain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TLSPublicCertChain) UnmarshalBinary(b []byte) error {
	var res TLSPublicCertChain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
