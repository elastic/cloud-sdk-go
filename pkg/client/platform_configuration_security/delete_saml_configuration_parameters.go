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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteSamlConfigurationParams creates a new DeleteSamlConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSamlConfigurationParams() *DeleteSamlConfigurationParams {
	return &DeleteSamlConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSamlConfigurationParamsWithTimeout creates a new DeleteSamlConfigurationParams object
// with the ability to set a timeout on a request.
func NewDeleteSamlConfigurationParamsWithTimeout(timeout time.Duration) *DeleteSamlConfigurationParams {
	return &DeleteSamlConfigurationParams{
		timeout: timeout,
	}
}

// NewDeleteSamlConfigurationParamsWithContext creates a new DeleteSamlConfigurationParams object
// with the ability to set a context for a request.
func NewDeleteSamlConfigurationParamsWithContext(ctx context.Context) *DeleteSamlConfigurationParams {
	return &DeleteSamlConfigurationParams{
		Context: ctx,
	}
}

// NewDeleteSamlConfigurationParamsWithHTTPClient creates a new DeleteSamlConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSamlConfigurationParamsWithHTTPClient(client *http.Client) *DeleteSamlConfigurationParams {
	return &DeleteSamlConfigurationParams{
		HTTPClient: client,
	}
}

/*
DeleteSamlConfigurationParams contains all the parameters to send to the API endpoint

	for the delete saml configuration operation.

	Typically these are written to a http.Request.
*/
type DeleteSamlConfigurationParams struct {

	/* RealmID.

	   The Elasticsearch Security realm identifier.
	*/
	RealmID string

	/* Version.

	   When specified, checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete saml configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSamlConfigurationParams) WithDefaults() *DeleteSamlConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete saml configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSamlConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) WithTimeout(timeout time.Duration) *DeleteSamlConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) WithContext(ctx context.Context) *DeleteSamlConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) WithHTTPClient(client *http.Client) *DeleteSamlConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRealmID adds the realmID to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) WithRealmID(realmID string) *DeleteSamlConfigurationParams {
	o.SetRealmID(realmID)
	return o
}

// SetRealmID adds the realmId to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) SetRealmID(realmID string) {
	o.RealmID = realmID
}

// WithVersion adds the version to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) WithVersion(version *string) *DeleteSamlConfigurationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete saml configuration params
func (o *DeleteSamlConfigurationParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSamlConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param realm_id
	if err := r.SetPathParam("realm_id", o.RealmID); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion string

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
