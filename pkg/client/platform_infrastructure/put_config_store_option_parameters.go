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

// NewPutConfigStoreOptionParams creates a new PutConfigStoreOptionParams object
// with the default values initialized.
func NewPutConfigStoreOptionParams() *PutConfigStoreOptionParams {
	var ()
	return &PutConfigStoreOptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutConfigStoreOptionParamsWithTimeout creates a new PutConfigStoreOptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutConfigStoreOptionParamsWithTimeout(timeout time.Duration) *PutConfigStoreOptionParams {
	var ()
	return &PutConfigStoreOptionParams{

		timeout: timeout,
	}
}

// NewPutConfigStoreOptionParamsWithContext creates a new PutConfigStoreOptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutConfigStoreOptionParamsWithContext(ctx context.Context) *PutConfigStoreOptionParams {
	var ()
	return &PutConfigStoreOptionParams{

		Context: ctx,
	}
}

// NewPutConfigStoreOptionParamsWithHTTPClient creates a new PutConfigStoreOptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutConfigStoreOptionParamsWithHTTPClient(client *http.Client) *PutConfigStoreOptionParams {
	var ()
	return &PutConfigStoreOptionParams{
		HTTPClient: client,
	}
}

/*PutConfigStoreOptionParams contains all the parameters to send to the API endpoint
for the put config store option operation typically these are written to a http.Request
*/
type PutConfigStoreOptionParams struct {

	/*Body
	  The Config Store Option definition

	*/
	Body *models.ConfigStoreOptionData
	/*ConfigOptionID
	  Name of the Config Store Option that you would like to modify

	*/
	ConfigOptionID string
	/*Version
	  If specified then checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put config store option params
func (o *PutConfigStoreOptionParams) WithTimeout(timeout time.Duration) *PutConfigStoreOptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put config store option params
func (o *PutConfigStoreOptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put config store option params
func (o *PutConfigStoreOptionParams) WithContext(ctx context.Context) *PutConfigStoreOptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put config store option params
func (o *PutConfigStoreOptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put config store option params
func (o *PutConfigStoreOptionParams) WithHTTPClient(client *http.Client) *PutConfigStoreOptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put config store option params
func (o *PutConfigStoreOptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put config store option params
func (o *PutConfigStoreOptionParams) WithBody(body *models.ConfigStoreOptionData) *PutConfigStoreOptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put config store option params
func (o *PutConfigStoreOptionParams) SetBody(body *models.ConfigStoreOptionData) {
	o.Body = body
}

// WithConfigOptionID adds the configOptionID to the put config store option params
func (o *PutConfigStoreOptionParams) WithConfigOptionID(configOptionID string) *PutConfigStoreOptionParams {
	o.SetConfigOptionID(configOptionID)
	return o
}

// SetConfigOptionID adds the configOptionId to the put config store option params
func (o *PutConfigStoreOptionParams) SetConfigOptionID(configOptionID string) {
	o.ConfigOptionID = configOptionID
}

// WithVersion adds the version to the put config store option params
func (o *PutConfigStoreOptionParams) WithVersion(version *string) *PutConfigStoreOptionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the put config store option params
func (o *PutConfigStoreOptionParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PutConfigStoreOptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param config_option_id
	if err := r.SetPathParam("config_option_id", o.ConfigOptionID); err != nil {
		return err
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
