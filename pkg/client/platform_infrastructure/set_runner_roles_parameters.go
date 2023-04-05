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
	"github.com/go-openapi/swag"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewSetRunnerRolesParams creates a new SetRunnerRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetRunnerRolesParams() *SetRunnerRolesParams {
	return &SetRunnerRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetRunnerRolesParamsWithTimeout creates a new SetRunnerRolesParams object
// with the ability to set a timeout on a request.
func NewSetRunnerRolesParamsWithTimeout(timeout time.Duration) *SetRunnerRolesParams {
	return &SetRunnerRolesParams{
		timeout: timeout,
	}
}

// NewSetRunnerRolesParamsWithContext creates a new SetRunnerRolesParams object
// with the ability to set a context for a request.
func NewSetRunnerRolesParamsWithContext(ctx context.Context) *SetRunnerRolesParams {
	return &SetRunnerRolesParams{
		Context: ctx,
	}
}

// NewSetRunnerRolesParamsWithHTTPClient creates a new SetRunnerRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetRunnerRolesParamsWithHTTPClient(client *http.Client) *SetRunnerRolesParams {
	return &SetRunnerRolesParams{
		HTTPClient: client,
	}
}

/*
SetRunnerRolesParams contains all the parameters to send to the API endpoint

	for the set runner roles operation.

	Typically these are written to a http.Request.
*/
type SetRunnerRolesParams struct {

	/* Bless.

	   Assigns the runner to the roles.
	*/
	Bless *bool

	/* Body.

	   The roles for the runner that you want to apply.
	*/
	Body *models.RunnerRolesInfo

	/* RunnerID.

	   The identifier for the runner
	*/
	RunnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set runner roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetRunnerRolesParams) WithDefaults() *SetRunnerRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set runner roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetRunnerRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set runner roles params
func (o *SetRunnerRolesParams) WithTimeout(timeout time.Duration) *SetRunnerRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set runner roles params
func (o *SetRunnerRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set runner roles params
func (o *SetRunnerRolesParams) WithContext(ctx context.Context) *SetRunnerRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set runner roles params
func (o *SetRunnerRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set runner roles params
func (o *SetRunnerRolesParams) WithHTTPClient(client *http.Client) *SetRunnerRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set runner roles params
func (o *SetRunnerRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBless adds the bless to the set runner roles params
func (o *SetRunnerRolesParams) WithBless(bless *bool) *SetRunnerRolesParams {
	o.SetBless(bless)
	return o
}

// SetBless adds the bless to the set runner roles params
func (o *SetRunnerRolesParams) SetBless(bless *bool) {
	o.Bless = bless
}

// WithBody adds the body to the set runner roles params
func (o *SetRunnerRolesParams) WithBody(body *models.RunnerRolesInfo) *SetRunnerRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set runner roles params
func (o *SetRunnerRolesParams) SetBody(body *models.RunnerRolesInfo) {
	o.Body = body
}

// WithRunnerID adds the runnerID to the set runner roles params
func (o *SetRunnerRolesParams) WithRunnerID(runnerID string) *SetRunnerRolesParams {
	o.SetRunnerID(runnerID)
	return o
}

// SetRunnerID adds the runnerId to the set runner roles params
func (o *SetRunnerRolesParams) SetRunnerID(runnerID string) {
	o.RunnerID = runnerID
}

// WriteToRequest writes these params to a swagger request
func (o *SetRunnerRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Bless != nil {

		// query param bless
		var qrBless bool

		if o.Bless != nil {
			qrBless = *o.Bless
		}
		qBless := swag.FormatBool(qrBless)
		if qBless != "" {

			if err := r.SetQueryParam("bless", qBless); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param runner_id
	if err := r.SetPathParam("runner_id", o.RunnerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
