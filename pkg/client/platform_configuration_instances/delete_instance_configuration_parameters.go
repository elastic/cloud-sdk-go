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

package platform_configuration_instances

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
	"github.com/go-openapi/swag"
)

// NewDeleteInstanceConfigurationParams creates a new DeleteInstanceConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteInstanceConfigurationParams() *DeleteInstanceConfigurationParams {
	return &DeleteInstanceConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceConfigurationParamsWithTimeout creates a new DeleteInstanceConfigurationParams object
// with the ability to set a timeout on a request.
func NewDeleteInstanceConfigurationParamsWithTimeout(timeout time.Duration) *DeleteInstanceConfigurationParams {
	return &DeleteInstanceConfigurationParams{
		timeout: timeout,
	}
}

// NewDeleteInstanceConfigurationParamsWithContext creates a new DeleteInstanceConfigurationParams object
// with the ability to set a context for a request.
func NewDeleteInstanceConfigurationParamsWithContext(ctx context.Context) *DeleteInstanceConfigurationParams {
	return &DeleteInstanceConfigurationParams{
		Context: ctx,
	}
}

// NewDeleteInstanceConfigurationParamsWithHTTPClient creates a new DeleteInstanceConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteInstanceConfigurationParamsWithHTTPClient(client *http.Client) *DeleteInstanceConfigurationParams {
	return &DeleteInstanceConfigurationParams{
		HTTPClient: client,
	}
}

/*
DeleteInstanceConfigurationParams contains all the parameters to send to the API endpoint

	for the delete instance configuration operation.

	Typically these are written to a http.Request.
*/
type DeleteInstanceConfigurationParams struct {

	/* ID.

	   ID of the instance configuration
	*/
	ID string

	/* OnlyTestVersion.

	   If true, the testing version (version -1) will be permanently deleted but the latest IC version will remain untouched. Defaults to false
	*/
	OnlyTestVersion *bool

	/* Version.

	   This is a database-level field, not related to the application-level 'config_version', except as described in the following docs. If specified, checks for conflicts against 'x-cloud-resource-version' from the GET request (the GET's 'config_version' should be left blank if the IC is configuration controlled, ie to get the latest configuration)
	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete instance configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstanceConfigurationParams) WithDefaults() *DeleteInstanceConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete instance configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteInstanceConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithTimeout(timeout time.Duration) *DeleteInstanceConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithContext(ctx context.Context) *DeleteInstanceConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithHTTPClient(client *http.Client) *DeleteInstanceConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithID(id string) *DeleteInstanceConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetID(id string) {
	o.ID = id
}

// WithOnlyTestVersion adds the onlyTestVersion to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithOnlyTestVersion(onlyTestVersion *bool) *DeleteInstanceConfigurationParams {
	o.SetOnlyTestVersion(onlyTestVersion)
	return o
}

// SetOnlyTestVersion adds the onlyTestVersion to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetOnlyTestVersion(onlyTestVersion *bool) {
	o.OnlyTestVersion = onlyTestVersion
}

// WithVersion adds the version to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) WithVersion(version *int64) *DeleteInstanceConfigurationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete instance configuration params
func (o *DeleteInstanceConfigurationParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.OnlyTestVersion != nil {

		// query param only_test_version
		var qrOnlyTestVersion bool

		if o.OnlyTestVersion != nil {
			qrOnlyTestVersion = *o.OnlyTestVersion
		}
		qOnlyTestVersion := swag.FormatBool(qrOnlyTestVersion)
		if qOnlyTestVersion != "" {

			if err := r.SetQueryParam("only_test_version", qOnlyTestVersion); err != nil {
				return err
			}
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
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
