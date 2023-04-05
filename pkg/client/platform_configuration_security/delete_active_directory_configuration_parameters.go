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

// NewDeleteActiveDirectoryConfigurationParams creates a new DeleteActiveDirectoryConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteActiveDirectoryConfigurationParams() *DeleteActiveDirectoryConfigurationParams {
	return &DeleteActiveDirectoryConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteActiveDirectoryConfigurationParamsWithTimeout creates a new DeleteActiveDirectoryConfigurationParams object
// with the ability to set a timeout on a request.
func NewDeleteActiveDirectoryConfigurationParamsWithTimeout(timeout time.Duration) *DeleteActiveDirectoryConfigurationParams {
	return &DeleteActiveDirectoryConfigurationParams{
		timeout: timeout,
	}
}

// NewDeleteActiveDirectoryConfigurationParamsWithContext creates a new DeleteActiveDirectoryConfigurationParams object
// with the ability to set a context for a request.
func NewDeleteActiveDirectoryConfigurationParamsWithContext(ctx context.Context) *DeleteActiveDirectoryConfigurationParams {
	return &DeleteActiveDirectoryConfigurationParams{
		Context: ctx,
	}
}

// NewDeleteActiveDirectoryConfigurationParamsWithHTTPClient creates a new DeleteActiveDirectoryConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteActiveDirectoryConfigurationParamsWithHTTPClient(client *http.Client) *DeleteActiveDirectoryConfigurationParams {
	return &DeleteActiveDirectoryConfigurationParams{
		HTTPClient: client,
	}
}

/*
DeleteActiveDirectoryConfigurationParams contains all the parameters to send to the API endpoint

	for the delete active directory configuration operation.

	Typically these are written to a http.Request.
*/
type DeleteActiveDirectoryConfigurationParams struct {

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

// WithDefaults hydrates default values in the delete active directory configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteActiveDirectoryConfigurationParams) WithDefaults() *DeleteActiveDirectoryConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete active directory configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteActiveDirectoryConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) WithTimeout(timeout time.Duration) *DeleteActiveDirectoryConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) WithContext(ctx context.Context) *DeleteActiveDirectoryConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) WithHTTPClient(client *http.Client) *DeleteActiveDirectoryConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRealmID adds the realmID to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) WithRealmID(realmID string) *DeleteActiveDirectoryConfigurationParams {
	o.SetRealmID(realmID)
	return o
}

// SetRealmID adds the realmId to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) SetRealmID(realmID string) {
	o.RealmID = realmID
}

// WithVersion adds the version to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) WithVersion(version *string) *DeleteActiveDirectoryConfigurationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete active directory configuration params
func (o *DeleteActiveDirectoryConfigurationParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteActiveDirectoryConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
