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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewSetAllocatorSettingsParams creates a new SetAllocatorSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetAllocatorSettingsParams() *SetAllocatorSettingsParams {
	return &SetAllocatorSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetAllocatorSettingsParamsWithTimeout creates a new SetAllocatorSettingsParams object
// with the ability to set a timeout on a request.
func NewSetAllocatorSettingsParamsWithTimeout(timeout time.Duration) *SetAllocatorSettingsParams {
	return &SetAllocatorSettingsParams{
		timeout: timeout,
	}
}

// NewSetAllocatorSettingsParamsWithContext creates a new SetAllocatorSettingsParams object
// with the ability to set a context for a request.
func NewSetAllocatorSettingsParamsWithContext(ctx context.Context) *SetAllocatorSettingsParams {
	return &SetAllocatorSettingsParams{
		Context: ctx,
	}
}

// NewSetAllocatorSettingsParamsWithHTTPClient creates a new SetAllocatorSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetAllocatorSettingsParamsWithHTTPClient(client *http.Client) *SetAllocatorSettingsParams {
	return &SetAllocatorSettingsParams{
		HTTPClient: client,
	}
}

/*
SetAllocatorSettingsParams contains all the parameters to send to the API endpoint

	for the set allocator settings operation.

	Typically these are written to a http.Request.
*/
type SetAllocatorSettingsParams struct {

	/* AllocatorID.

	   The allocator identifier.
	*/
	AllocatorID string

	/* Body.

	   The allocator settings to apply
	*/
	Body *models.AllocatorSettings

	/* Version.

	   Checks for conflicts against the metadata version, then returns the value in the `x-cloud-resource-version` header.
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set allocator settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAllocatorSettingsParams) WithDefaults() *SetAllocatorSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set allocator settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAllocatorSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set allocator settings params
func (o *SetAllocatorSettingsParams) WithTimeout(timeout time.Duration) *SetAllocatorSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set allocator settings params
func (o *SetAllocatorSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set allocator settings params
func (o *SetAllocatorSettingsParams) WithContext(ctx context.Context) *SetAllocatorSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set allocator settings params
func (o *SetAllocatorSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set allocator settings params
func (o *SetAllocatorSettingsParams) WithHTTPClient(client *http.Client) *SetAllocatorSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set allocator settings params
func (o *SetAllocatorSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorID adds the allocatorID to the set allocator settings params
func (o *SetAllocatorSettingsParams) WithAllocatorID(allocatorID string) *SetAllocatorSettingsParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the set allocator settings params
func (o *SetAllocatorSettingsParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WithBody adds the body to the set allocator settings params
func (o *SetAllocatorSettingsParams) WithBody(body *models.AllocatorSettings) *SetAllocatorSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set allocator settings params
func (o *SetAllocatorSettingsParams) SetBody(body *models.AllocatorSettings) {
	o.Body = body
}

// WithVersion adds the version to the set allocator settings params
func (o *SetAllocatorSettingsParams) WithVersion(version *string) *SetAllocatorSettingsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the set allocator settings params
func (o *SetAllocatorSettingsParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SetAllocatorSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocator_id
	if err := r.SetPathParam("allocator_id", o.AllocatorID); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
