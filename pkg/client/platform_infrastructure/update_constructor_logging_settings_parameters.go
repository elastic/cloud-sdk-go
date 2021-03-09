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

package platform_infrastructure

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

// NewUpdateConstructorLoggingSettingsParams creates a new UpdateConstructorLoggingSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConstructorLoggingSettingsParams() *UpdateConstructorLoggingSettingsParams {
	return &UpdateConstructorLoggingSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConstructorLoggingSettingsParamsWithTimeout creates a new UpdateConstructorLoggingSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateConstructorLoggingSettingsParamsWithTimeout(timeout time.Duration) *UpdateConstructorLoggingSettingsParams {
	return &UpdateConstructorLoggingSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateConstructorLoggingSettingsParamsWithContext creates a new UpdateConstructorLoggingSettingsParams object
// with the ability to set a context for a request.
func NewUpdateConstructorLoggingSettingsParamsWithContext(ctx context.Context) *UpdateConstructorLoggingSettingsParams {
	return &UpdateConstructorLoggingSettingsParams{
		Context: ctx,
	}
}

// NewUpdateConstructorLoggingSettingsParamsWithHTTPClient creates a new UpdateConstructorLoggingSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConstructorLoggingSettingsParamsWithHTTPClient(client *http.Client) *UpdateConstructorLoggingSettingsParams {
	return &UpdateConstructorLoggingSettingsParams{
		HTTPClient: client,
	}
}

/* UpdateConstructorLoggingSettingsParams contains all the parameters to send to the API endpoint
   for the update constructor logging settings operation.

   Typically these are written to a http.Request.
*/
type UpdateConstructorLoggingSettingsParams struct {

	/* Body.

	   The logging settings to update
	*/
	Body string

	/* ConstructorID.

	   Identifier for the constructor
	*/
	ConstructorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update constructor logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConstructorLoggingSettingsParams) WithDefaults() *UpdateConstructorLoggingSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update constructor logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConstructorLoggingSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) WithTimeout(timeout time.Duration) *UpdateConstructorLoggingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) WithContext(ctx context.Context) *UpdateConstructorLoggingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) WithHTTPClient(client *http.Client) *UpdateConstructorLoggingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) WithBody(body string) *UpdateConstructorLoggingSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) SetBody(body string) {
	o.Body = body
}

// WithConstructorID adds the constructorID to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) WithConstructorID(constructorID string) *UpdateConstructorLoggingSettingsParams {
	o.SetConstructorID(constructorID)
	return o
}

// SetConstructorID adds the constructorId to the update constructor logging settings params
func (o *UpdateConstructorLoggingSettingsParams) SetConstructorID(constructorID string) {
	o.ConstructorID = constructorID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConstructorLoggingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param constructor_id
	if err := r.SetPathParam("constructor_id", o.ConstructorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
