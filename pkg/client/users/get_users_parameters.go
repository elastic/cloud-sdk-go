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

package users

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

// NewGetUsersParams creates a new GetUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersParams() *GetUsersParams {
	return &GetUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersParamsWithTimeout creates a new GetUsersParams object
// with the ability to set a timeout on a request.
func NewGetUsersParamsWithTimeout(timeout time.Duration) *GetUsersParams {
	return &GetUsersParams{
		timeout: timeout,
	}
}

// NewGetUsersParamsWithContext creates a new GetUsersParams object
// with the ability to set a context for a request.
func NewGetUsersParamsWithContext(ctx context.Context) *GetUsersParams {
	return &GetUsersParams{
		Context: ctx,
	}
}

// NewGetUsersParamsWithHTTPClient creates a new GetUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersParamsWithHTTPClient(client *http.Client) *GetUsersParams {
	return &GetUsersParams{
		HTTPClient: client,
	}
}

/*
GetUsersParams contains all the parameters to send to the API endpoint

	for the get users operation.

	Typically these are written to a http.Request.
*/
type GetUsersParams struct {

	/* IncludeDisabled.

	   True if disabled users should be included in the response
	*/
	IncludeDisabled *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersParams) WithDefaults() *GetUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersParams) SetDefaults() {
	var (
		includeDisabledDefault = bool(false)
	)

	val := GetUsersParams{
		IncludeDisabled: &includeDisabledDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get users params
func (o *GetUsersParams) WithTimeout(timeout time.Duration) *GetUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users params
func (o *GetUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users params
func (o *GetUsersParams) WithContext(ctx context.Context) *GetUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users params
func (o *GetUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) WithHTTPClient(client *http.Client) *GetUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users params
func (o *GetUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeDisabled adds the includeDisabled to the get users params
func (o *GetUsersParams) WithIncludeDisabled(includeDisabled *bool) *GetUsersParams {
	o.SetIncludeDisabled(includeDisabled)
	return o
}

// SetIncludeDisabled adds the includeDisabled to the get users params
func (o *GetUsersParams) SetIncludeDisabled(includeDisabled *bool) {
	o.IncludeDisabled = includeDisabled
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeDisabled != nil {

		// query param include_disabled
		var qrIncludeDisabled bool

		if o.IncludeDisabled != nil {
			qrIncludeDisabled = *o.IncludeDisabled
		}
		qIncludeDisabled := swag.FormatBool(qrIncludeDisabled)
		if qIncludeDisabled != "" {

			if err := r.SetQueryParam("include_disabled", qIncludeDisabled); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
