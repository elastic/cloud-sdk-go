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

package authentication

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

// NewGetUsersAPIKeysParams creates a new GetUsersAPIKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersAPIKeysParams() *GetUsersAPIKeysParams {
	return &GetUsersAPIKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersAPIKeysParamsWithTimeout creates a new GetUsersAPIKeysParams object
// with the ability to set a timeout on a request.
func NewGetUsersAPIKeysParamsWithTimeout(timeout time.Duration) *GetUsersAPIKeysParams {
	return &GetUsersAPIKeysParams{
		timeout: timeout,
	}
}

// NewGetUsersAPIKeysParamsWithContext creates a new GetUsersAPIKeysParams object
// with the ability to set a context for a request.
func NewGetUsersAPIKeysParamsWithContext(ctx context.Context) *GetUsersAPIKeysParams {
	return &GetUsersAPIKeysParams{
		Context: ctx,
	}
}

// NewGetUsersAPIKeysParamsWithHTTPClient creates a new GetUsersAPIKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersAPIKeysParamsWithHTTPClient(client *http.Client) *GetUsersAPIKeysParams {
	return &GetUsersAPIKeysParams{
		HTTPClient: client,
	}
}

/*
GetUsersAPIKeysParams contains all the parameters to send to the API endpoint

	for the get users api keys operation.

	Typically these are written to a http.Request.
*/
type GetUsersAPIKeysParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersAPIKeysParams) WithDefaults() *GetUsersAPIKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersAPIKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get users api keys params
func (o *GetUsersAPIKeysParams) WithTimeout(timeout time.Duration) *GetUsersAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users api keys params
func (o *GetUsersAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users api keys params
func (o *GetUsersAPIKeysParams) WithContext(ctx context.Context) *GetUsersAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users api keys params
func (o *GetUsersAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users api keys params
func (o *GetUsersAPIKeysParams) WithHTTPClient(client *http.Client) *GetUsersAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users api keys params
func (o *GetUsersAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
