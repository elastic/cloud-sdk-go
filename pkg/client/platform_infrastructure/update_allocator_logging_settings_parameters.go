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

	strfmt "github.com/go-openapi/strfmt"
)

// NewUpdateAllocatorLoggingSettingsParams creates a new UpdateAllocatorLoggingSettingsParams object
// with the default values initialized.
func NewUpdateAllocatorLoggingSettingsParams() *UpdateAllocatorLoggingSettingsParams {
	var ()
	return &UpdateAllocatorLoggingSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAllocatorLoggingSettingsParamsWithTimeout creates a new UpdateAllocatorLoggingSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAllocatorLoggingSettingsParamsWithTimeout(timeout time.Duration) *UpdateAllocatorLoggingSettingsParams {
	var ()
	return &UpdateAllocatorLoggingSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateAllocatorLoggingSettingsParamsWithContext creates a new UpdateAllocatorLoggingSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAllocatorLoggingSettingsParamsWithContext(ctx context.Context) *UpdateAllocatorLoggingSettingsParams {
	var ()
	return &UpdateAllocatorLoggingSettingsParams{

		Context: ctx,
	}
}

// NewUpdateAllocatorLoggingSettingsParamsWithHTTPClient creates a new UpdateAllocatorLoggingSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAllocatorLoggingSettingsParamsWithHTTPClient(client *http.Client) *UpdateAllocatorLoggingSettingsParams {
	var ()
	return &UpdateAllocatorLoggingSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateAllocatorLoggingSettingsParams contains all the parameters to send to the API endpoint
for the update allocator logging settings operation typically these are written to a http.Request
*/
type UpdateAllocatorLoggingSettingsParams struct {

	/*AllocatorID
	  The allocator identifier.

	*/
	AllocatorID string
	/*Body
	  The logging settings to update

	*/
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) WithTimeout(timeout time.Duration) *UpdateAllocatorLoggingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) WithContext(ctx context.Context) *UpdateAllocatorLoggingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) WithHTTPClient(client *http.Client) *UpdateAllocatorLoggingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorID adds the allocatorID to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) WithAllocatorID(allocatorID string) *UpdateAllocatorLoggingSettingsParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WithBody adds the body to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) WithBody(body string) *UpdateAllocatorLoggingSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update allocator logging settings params
func (o *UpdateAllocatorLoggingSettingsParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAllocatorLoggingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocator_id
	if err := r.SetPathParam("allocator_id", o.AllocatorID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
