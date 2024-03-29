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

// NewGetProxiesFilteredGroupHealthParams creates a new GetProxiesFilteredGroupHealthParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProxiesFilteredGroupHealthParams() *GetProxiesFilteredGroupHealthParams {
	return &GetProxiesFilteredGroupHealthParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxiesFilteredGroupHealthParamsWithTimeout creates a new GetProxiesFilteredGroupHealthParams object
// with the ability to set a timeout on a request.
func NewGetProxiesFilteredGroupHealthParamsWithTimeout(timeout time.Duration) *GetProxiesFilteredGroupHealthParams {
	return &GetProxiesFilteredGroupHealthParams{
		timeout: timeout,
	}
}

// NewGetProxiesFilteredGroupHealthParamsWithContext creates a new GetProxiesFilteredGroupHealthParams object
// with the ability to set a context for a request.
func NewGetProxiesFilteredGroupHealthParamsWithContext(ctx context.Context) *GetProxiesFilteredGroupHealthParams {
	return &GetProxiesFilteredGroupHealthParams{
		Context: ctx,
	}
}

// NewGetProxiesFilteredGroupHealthParamsWithHTTPClient creates a new GetProxiesFilteredGroupHealthParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProxiesFilteredGroupHealthParamsWithHTTPClient(client *http.Client) *GetProxiesFilteredGroupHealthParams {
	return &GetProxiesFilteredGroupHealthParams{
		HTTPClient: client,
	}
}

/*
GetProxiesFilteredGroupHealthParams contains all the parameters to send to the API endpoint

	for the get proxies filtered group health operation.

	Typically these are written to a http.Request.
*/
type GetProxiesFilteredGroupHealthParams struct {

	/* ExpectStatus.

	   The expected status
	*/
	ExpectStatus *string

	/* ProxiesFilteredGroupID.

	   "The identifier for the filtered group of proxies
	*/
	ProxiesFilteredGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get proxies filtered group health params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProxiesFilteredGroupHealthParams) WithDefaults() *GetProxiesFilteredGroupHealthParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get proxies filtered group health params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProxiesFilteredGroupHealthParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) WithTimeout(timeout time.Duration) *GetProxiesFilteredGroupHealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) WithContext(ctx context.Context) *GetProxiesFilteredGroupHealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) WithHTTPClient(client *http.Client) *GetProxiesFilteredGroupHealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpectStatus adds the expectStatus to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) WithExpectStatus(expectStatus *string) *GetProxiesFilteredGroupHealthParams {
	o.SetExpectStatus(expectStatus)
	return o
}

// SetExpectStatus adds the expectStatus to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) SetExpectStatus(expectStatus *string) {
	o.ExpectStatus = expectStatus
}

// WithProxiesFilteredGroupID adds the proxiesFilteredGroupID to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) WithProxiesFilteredGroupID(proxiesFilteredGroupID string) *GetProxiesFilteredGroupHealthParams {
	o.SetProxiesFilteredGroupID(proxiesFilteredGroupID)
	return o
}

// SetProxiesFilteredGroupID adds the proxiesFilteredGroupId to the get proxies filtered group health params
func (o *GetProxiesFilteredGroupHealthParams) SetProxiesFilteredGroupID(proxiesFilteredGroupID string) {
	o.ProxiesFilteredGroupID = proxiesFilteredGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxiesFilteredGroupHealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExpectStatus != nil {

		// query param expect_status
		var qrExpectStatus string

		if o.ExpectStatus != nil {
			qrExpectStatus = *o.ExpectStatus
		}
		qExpectStatus := qrExpectStatus
		if qExpectStatus != "" {

			if err := r.SetQueryParam("expect_status", qExpectStatus); err != nil {
				return err
			}
		}
	}

	// path param proxies_filtered_group_id
	if err := r.SetPathParam("proxies_filtered_group_id", o.ProxiesFilteredGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
