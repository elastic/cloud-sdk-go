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

package stack

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

// NewDeleteVersionStackParams creates a new DeleteVersionStackParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVersionStackParams() *DeleteVersionStackParams {
	return &DeleteVersionStackParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVersionStackParamsWithTimeout creates a new DeleteVersionStackParams object
// with the ability to set a timeout on a request.
func NewDeleteVersionStackParamsWithTimeout(timeout time.Duration) *DeleteVersionStackParams {
	return &DeleteVersionStackParams{
		timeout: timeout,
	}
}

// NewDeleteVersionStackParamsWithContext creates a new DeleteVersionStackParams object
// with the ability to set a context for a request.
func NewDeleteVersionStackParamsWithContext(ctx context.Context) *DeleteVersionStackParams {
	return &DeleteVersionStackParams{
		Context: ctx,
	}
}

// NewDeleteVersionStackParamsWithHTTPClient creates a new DeleteVersionStackParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVersionStackParamsWithHTTPClient(client *http.Client) *DeleteVersionStackParams {
	return &DeleteVersionStackParams{
		HTTPClient: client,
	}
}

/*
DeleteVersionStackParams contains all the parameters to send to the API endpoint

	for the delete version stack operation.

	Typically these are written to a http.Request.
*/
type DeleteVersionStackParams struct {

	/* Version.

	   The Elastic Stack version. For example, `5.3.1` or `5.0.0-RC4`.
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete version stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVersionStackParams) WithDefaults() *DeleteVersionStackParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete version stack params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVersionStackParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete version stack params
func (o *DeleteVersionStackParams) WithTimeout(timeout time.Duration) *DeleteVersionStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete version stack params
func (o *DeleteVersionStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete version stack params
func (o *DeleteVersionStackParams) WithContext(ctx context.Context) *DeleteVersionStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete version stack params
func (o *DeleteVersionStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete version stack params
func (o *DeleteVersionStackParams) WithHTTPClient(client *http.Client) *DeleteVersionStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete version stack params
func (o *DeleteVersionStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersion adds the version to the delete version stack params
func (o *DeleteVersionStackParams) WithVersion(version string) *DeleteVersionStackParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete version stack params
func (o *DeleteVersionStackParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVersionStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
