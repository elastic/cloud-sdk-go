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

package extensions

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

// NewDeleteExtensionParams creates a new DeleteExtensionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteExtensionParams() *DeleteExtensionParams {
	return &DeleteExtensionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExtensionParamsWithTimeout creates a new DeleteExtensionParams object
// with the ability to set a timeout on a request.
func NewDeleteExtensionParamsWithTimeout(timeout time.Duration) *DeleteExtensionParams {
	return &DeleteExtensionParams{
		timeout: timeout,
	}
}

// NewDeleteExtensionParamsWithContext creates a new DeleteExtensionParams object
// with the ability to set a context for a request.
func NewDeleteExtensionParamsWithContext(ctx context.Context) *DeleteExtensionParams {
	return &DeleteExtensionParams{
		Context: ctx,
	}
}

// NewDeleteExtensionParamsWithHTTPClient creates a new DeleteExtensionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteExtensionParamsWithHTTPClient(client *http.Client) *DeleteExtensionParams {
	return &DeleteExtensionParams{
		HTTPClient: client,
	}
}

/*
DeleteExtensionParams contains all the parameters to send to the API endpoint

	for the delete extension operation.

	Typically these are written to a http.Request.
*/
type DeleteExtensionParams struct {

	/* ExtensionID.

	   Id of an extension
	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete extension params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExtensionParams) WithDefaults() *DeleteExtensionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete extension params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteExtensionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete extension params
func (o *DeleteExtensionParams) WithTimeout(timeout time.Duration) *DeleteExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete extension params
func (o *DeleteExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete extension params
func (o *DeleteExtensionParams) WithContext(ctx context.Context) *DeleteExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete extension params
func (o *DeleteExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete extension params
func (o *DeleteExtensionParams) WithHTTPClient(client *http.Client) *DeleteExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete extension params
func (o *DeleteExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtensionID adds the extensionID to the delete extension params
func (o *DeleteExtensionParams) WithExtensionID(extensionID string) *DeleteExtensionParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the delete extension params
func (o *DeleteExtensionParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extension_id
	if err := r.SetPathParam("extension_id", o.ExtensionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
