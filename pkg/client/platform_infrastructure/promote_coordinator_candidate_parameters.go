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

// NewPromoteCoordinatorCandidateParams creates a new PromoteCoordinatorCandidateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPromoteCoordinatorCandidateParams() *PromoteCoordinatorCandidateParams {
	return &PromoteCoordinatorCandidateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPromoteCoordinatorCandidateParamsWithTimeout creates a new PromoteCoordinatorCandidateParams object
// with the ability to set a timeout on a request.
func NewPromoteCoordinatorCandidateParamsWithTimeout(timeout time.Duration) *PromoteCoordinatorCandidateParams {
	return &PromoteCoordinatorCandidateParams{
		timeout: timeout,
	}
}

// NewPromoteCoordinatorCandidateParamsWithContext creates a new PromoteCoordinatorCandidateParams object
// with the ability to set a context for a request.
func NewPromoteCoordinatorCandidateParamsWithContext(ctx context.Context) *PromoteCoordinatorCandidateParams {
	return &PromoteCoordinatorCandidateParams{
		Context: ctx,
	}
}

// NewPromoteCoordinatorCandidateParamsWithHTTPClient creates a new PromoteCoordinatorCandidateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPromoteCoordinatorCandidateParamsWithHTTPClient(client *http.Client) *PromoteCoordinatorCandidateParams {
	return &PromoteCoordinatorCandidateParams{
		HTTPClient: client,
	}
}

/* PromoteCoordinatorCandidateParams contains all the parameters to send to the API endpoint
   for the promote coordinator candidate operation.

   Typically these are written to a http.Request.
*/
type PromoteCoordinatorCandidateParams struct {

	/* CoordinatorCandidateID.

	   The identifier for the coordinator candidate
	*/
	CoordinatorCandidateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the promote coordinator candidate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PromoteCoordinatorCandidateParams) WithDefaults() *PromoteCoordinatorCandidateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the promote coordinator candidate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PromoteCoordinatorCandidateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the promote coordinator candidate params
func (o *PromoteCoordinatorCandidateParams) WithTimeout(timeout time.Duration) *PromoteCoordinatorCandidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the promote coordinator candidate params
func (o *PromoteCoordinatorCandidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the promote coordinator candidate params
func (o *PromoteCoordinatorCandidateParams) WithContext(ctx context.Context) *PromoteCoordinatorCandidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the promote coordinator candidate params
func (o *PromoteCoordinatorCandidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the promote coordinator candidate params
func (o *PromoteCoordinatorCandidateParams) WithHTTPClient(client *http.Client) *PromoteCoordinatorCandidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the promote coordinator candidate params
func (o *PromoteCoordinatorCandidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCoordinatorCandidateID adds the coordinatorCandidateID to the promote coordinator candidate params
func (o *PromoteCoordinatorCandidateParams) WithCoordinatorCandidateID(coordinatorCandidateID string) *PromoteCoordinatorCandidateParams {
	o.SetCoordinatorCandidateID(coordinatorCandidateID)
	return o
}

// SetCoordinatorCandidateID adds the coordinatorCandidateId to the promote coordinator candidate params
func (o *PromoteCoordinatorCandidateParams) SetCoordinatorCandidateID(coordinatorCandidateID string) {
	o.CoordinatorCandidateID = coordinatorCandidateID
}

// WriteToRequest writes these params to a swagger request
func (o *PromoteCoordinatorCandidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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