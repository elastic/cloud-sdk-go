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

// NewGetRunnerLoggingSettingsParams creates a new GetRunnerLoggingSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunnerLoggingSettingsParams() *GetRunnerLoggingSettingsParams {
	return &GetRunnerLoggingSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunnerLoggingSettingsParamsWithTimeout creates a new GetRunnerLoggingSettingsParams object
// with the ability to set a timeout on a request.
func NewGetRunnerLoggingSettingsParamsWithTimeout(timeout time.Duration) *GetRunnerLoggingSettingsParams {
	return &GetRunnerLoggingSettingsParams{
		timeout: timeout,
	}
}

// NewGetRunnerLoggingSettingsParamsWithContext creates a new GetRunnerLoggingSettingsParams object
// with the ability to set a context for a request.
func NewGetRunnerLoggingSettingsParamsWithContext(ctx context.Context) *GetRunnerLoggingSettingsParams {
	return &GetRunnerLoggingSettingsParams{
		Context: ctx,
	}
}

// NewGetRunnerLoggingSettingsParamsWithHTTPClient creates a new GetRunnerLoggingSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunnerLoggingSettingsParamsWithHTTPClient(client *http.Client) *GetRunnerLoggingSettingsParams {
	return &GetRunnerLoggingSettingsParams{
		HTTPClient: client,
	}
}

/*
GetRunnerLoggingSettingsParams contains all the parameters to send to the API endpoint

	for the get runner logging settings operation.

	Typically these are written to a http.Request.
*/
type GetRunnerLoggingSettingsParams struct {

	/* RunnerID.

	   The identifier for the runner
	*/
	RunnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get runner logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunnerLoggingSettingsParams) WithDefaults() *GetRunnerLoggingSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get runner logging settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunnerLoggingSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get runner logging settings params
func (o *GetRunnerLoggingSettingsParams) WithTimeout(timeout time.Duration) *GetRunnerLoggingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runner logging settings params
func (o *GetRunnerLoggingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runner logging settings params
func (o *GetRunnerLoggingSettingsParams) WithContext(ctx context.Context) *GetRunnerLoggingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runner logging settings params
func (o *GetRunnerLoggingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runner logging settings params
func (o *GetRunnerLoggingSettingsParams) WithHTTPClient(client *http.Client) *GetRunnerLoggingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runner logging settings params
func (o *GetRunnerLoggingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRunnerID adds the runnerID to the get runner logging settings params
func (o *GetRunnerLoggingSettingsParams) WithRunnerID(runnerID string) *GetRunnerLoggingSettingsParams {
	o.SetRunnerID(runnerID)
	return o
}

// SetRunnerID adds the runnerId to the get runner logging settings params
func (o *GetRunnerLoggingSettingsParams) SetRunnerID(runnerID string) {
	o.RunnerID = runnerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunnerLoggingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param runner_id
	if err := r.SetPathParam("runner_id", o.RunnerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
