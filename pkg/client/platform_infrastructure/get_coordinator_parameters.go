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

// NewGetCoordinatorParams creates a new GetCoordinatorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCoordinatorParams() *GetCoordinatorParams {
	return &GetCoordinatorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCoordinatorParamsWithTimeout creates a new GetCoordinatorParams object
// with the ability to set a timeout on a request.
func NewGetCoordinatorParamsWithTimeout(timeout time.Duration) *GetCoordinatorParams {
	return &GetCoordinatorParams{
		timeout: timeout,
	}
}

// NewGetCoordinatorParamsWithContext creates a new GetCoordinatorParams object
// with the ability to set a context for a request.
func NewGetCoordinatorParamsWithContext(ctx context.Context) *GetCoordinatorParams {
	return &GetCoordinatorParams{
		Context: ctx,
	}
}

// NewGetCoordinatorParamsWithHTTPClient creates a new GetCoordinatorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCoordinatorParamsWithHTTPClient(client *http.Client) *GetCoordinatorParams {
	return &GetCoordinatorParams{
		HTTPClient: client,
	}
}

/* GetCoordinatorParams contains all the parameters to send to the API endpoint
   for the get coordinator operation.

   Typically these are written to a http.Request.
*/
type GetCoordinatorParams struct {

	/* CoordinatorID.

	   The identifier for the coordinator
	*/
	CoordinatorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get coordinator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoordinatorParams) WithDefaults() *GetCoordinatorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get coordinator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoordinatorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get coordinator params
func (o *GetCoordinatorParams) WithTimeout(timeout time.Duration) *GetCoordinatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get coordinator params
func (o *GetCoordinatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get coordinator params
func (o *GetCoordinatorParams) WithContext(ctx context.Context) *GetCoordinatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get coordinator params
func (o *GetCoordinatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get coordinator params
func (o *GetCoordinatorParams) WithHTTPClient(client *http.Client) *GetCoordinatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get coordinator params
func (o *GetCoordinatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCoordinatorID adds the coordinatorID to the get coordinator params
func (o *GetCoordinatorParams) WithCoordinatorID(coordinatorID string) *GetCoordinatorParams {
	o.SetCoordinatorID(coordinatorID)
	return o
}

// SetCoordinatorID adds the coordinatorId to the get coordinator params
func (o *GetCoordinatorParams) SetCoordinatorID(coordinatorID string) {
	o.CoordinatorID = coordinatorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCoordinatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param coordinator_id
	if err := r.SetPathParam("coordinator_id", o.CoordinatorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
