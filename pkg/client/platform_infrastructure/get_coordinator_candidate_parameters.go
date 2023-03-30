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

// NewGetCoordinatorCandidateParams creates a new GetCoordinatorCandidateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCoordinatorCandidateParams() *GetCoordinatorCandidateParams {
	return &GetCoordinatorCandidateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCoordinatorCandidateParamsWithTimeout creates a new GetCoordinatorCandidateParams object
// with the ability to set a timeout on a request.
func NewGetCoordinatorCandidateParamsWithTimeout(timeout time.Duration) *GetCoordinatorCandidateParams {
	return &GetCoordinatorCandidateParams{
		timeout: timeout,
	}
}

// NewGetCoordinatorCandidateParamsWithContext creates a new GetCoordinatorCandidateParams object
// with the ability to set a context for a request.
func NewGetCoordinatorCandidateParamsWithContext(ctx context.Context) *GetCoordinatorCandidateParams {
	return &GetCoordinatorCandidateParams{
		Context: ctx,
	}
}

// NewGetCoordinatorCandidateParamsWithHTTPClient creates a new GetCoordinatorCandidateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCoordinatorCandidateParamsWithHTTPClient(client *http.Client) *GetCoordinatorCandidateParams {
	return &GetCoordinatorCandidateParams{
		HTTPClient: client,
	}
}

/* GetCoordinatorCandidateParams contains all the parameters to send to the API endpoint
   for the get coordinator candidate operation.

   Typically these are written to a http.Request.
*/
type GetCoordinatorCandidateParams struct {

	/* CoordinatorCandidateID.

	   The identifier for the coordinator candidate
	*/
	CoordinatorCandidateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get coordinator candidate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoordinatorCandidateParams) WithDefaults() *GetCoordinatorCandidateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get coordinator candidate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoordinatorCandidateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get coordinator candidate params
func (o *GetCoordinatorCandidateParams) WithTimeout(timeout time.Duration) *GetCoordinatorCandidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get coordinator candidate params
func (o *GetCoordinatorCandidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get coordinator candidate params
func (o *GetCoordinatorCandidateParams) WithContext(ctx context.Context) *GetCoordinatorCandidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get coordinator candidate params
func (o *GetCoordinatorCandidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get coordinator candidate params
func (o *GetCoordinatorCandidateParams) WithHTTPClient(client *http.Client) *GetCoordinatorCandidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get coordinator candidate params
func (o *GetCoordinatorCandidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCoordinatorCandidateID adds the coordinatorCandidateID to the get coordinator candidate params
func (o *GetCoordinatorCandidateParams) WithCoordinatorCandidateID(coordinatorCandidateID string) *GetCoordinatorCandidateParams {
	o.SetCoordinatorCandidateID(coordinatorCandidateID)
	return o
}

// SetCoordinatorCandidateID adds the coordinatorCandidateId to the get coordinator candidate params
func (o *GetCoordinatorCandidateParams) SetCoordinatorCandidateID(coordinatorCandidateID string) {
	o.CoordinatorCandidateID = coordinatorCandidateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCoordinatorCandidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param coordinator_candidate_id
	if err := r.SetPathParam("coordinator_candidate_id", o.CoordinatorCandidateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
