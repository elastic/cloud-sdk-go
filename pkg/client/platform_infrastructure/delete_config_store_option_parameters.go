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

// NewDeleteConfigStoreOptionParams creates a new DeleteConfigStoreOptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteConfigStoreOptionParams() *DeleteConfigStoreOptionParams {
	return &DeleteConfigStoreOptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConfigStoreOptionParamsWithTimeout creates a new DeleteConfigStoreOptionParams object
// with the ability to set a timeout on a request.
func NewDeleteConfigStoreOptionParamsWithTimeout(timeout time.Duration) *DeleteConfigStoreOptionParams {
	return &DeleteConfigStoreOptionParams{
		timeout: timeout,
	}
}

// NewDeleteConfigStoreOptionParamsWithContext creates a new DeleteConfigStoreOptionParams object
// with the ability to set a context for a request.
func NewDeleteConfigStoreOptionParamsWithContext(ctx context.Context) *DeleteConfigStoreOptionParams {
	return &DeleteConfigStoreOptionParams{
		Context: ctx,
	}
}

// NewDeleteConfigStoreOptionParamsWithHTTPClient creates a new DeleteConfigStoreOptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteConfigStoreOptionParamsWithHTTPClient(client *http.Client) *DeleteConfigStoreOptionParams {
	return &DeleteConfigStoreOptionParams{
		HTTPClient: client,
	}
}

/* DeleteConfigStoreOptionParams contains all the parameters to send to the API endpoint
   for the delete config store option operation.

   Typically these are written to a http.Request.
*/
type DeleteConfigStoreOptionParams struct {

	/* ConfigOptionID.

	   Name of the Config Store Option that you would like to delete
	*/
	ConfigOptionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete config store option params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConfigStoreOptionParams) WithDefaults() *DeleteConfigStoreOptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete config store option params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConfigStoreOptionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete config store option params
func (o *DeleteConfigStoreOptionParams) WithTimeout(timeout time.Duration) *DeleteConfigStoreOptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete config store option params
func (o *DeleteConfigStoreOptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete config store option params
func (o *DeleteConfigStoreOptionParams) WithContext(ctx context.Context) *DeleteConfigStoreOptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete config store option params
func (o *DeleteConfigStoreOptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete config store option params
func (o *DeleteConfigStoreOptionParams) WithHTTPClient(client *http.Client) *DeleteConfigStoreOptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete config store option params
func (o *DeleteConfigStoreOptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigOptionID adds the configOptionID to the delete config store option params
func (o *DeleteConfigStoreOptionParams) WithConfigOptionID(configOptionID string) *DeleteConfigStoreOptionParams {
	o.SetConfigOptionID(configOptionID)
	return o
}

// SetConfigOptionID adds the configOptionId to the delete config store option params
func (o *DeleteConfigStoreOptionParams) SetConfigOptionID(configOptionID string) {
	o.ConfigOptionID = configOptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConfigStoreOptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param config_option_id
	if err := r.SetPathParam("config_option_id", o.ConfigOptionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
