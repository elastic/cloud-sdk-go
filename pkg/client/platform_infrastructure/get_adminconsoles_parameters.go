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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAdminconsolesParams creates a new GetAdminconsolesParams object
// with the default values initialized.
func NewGetAdminconsolesParams() *GetAdminconsolesParams {

	return &GetAdminconsolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAdminconsolesParamsWithTimeout creates a new GetAdminconsolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAdminconsolesParamsWithTimeout(timeout time.Duration) *GetAdminconsolesParams {

	return &GetAdminconsolesParams{

		timeout: timeout,
	}
}

// NewGetAdminconsolesParamsWithContext creates a new GetAdminconsolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAdminconsolesParamsWithContext(ctx context.Context) *GetAdminconsolesParams {

	return &GetAdminconsolesParams{

		Context: ctx,
	}
}

// NewGetAdminconsolesParamsWithHTTPClient creates a new GetAdminconsolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAdminconsolesParamsWithHTTPClient(client *http.Client) *GetAdminconsolesParams {

	return &GetAdminconsolesParams{
		HTTPClient: client,
	}
}

/*GetAdminconsolesParams contains all the parameters to send to the API endpoint
for the get adminconsoles operation typically these are written to a http.Request
*/
type GetAdminconsolesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get adminconsoles params
func (o *GetAdminconsolesParams) WithTimeout(timeout time.Duration) *GetAdminconsolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get adminconsoles params
func (o *GetAdminconsolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get adminconsoles params
func (o *GetAdminconsolesParams) WithContext(ctx context.Context) *GetAdminconsolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get adminconsoles params
func (o *GetAdminconsolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get adminconsoles params
func (o *GetAdminconsolesParams) WithHTTPClient(client *http.Client) *GetAdminconsolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get adminconsoles params
func (o *GetAdminconsolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAdminconsolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
