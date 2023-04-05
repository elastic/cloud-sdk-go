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

package platform_configuration_security

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

// NewCreateActiveDirectoryConfigurationParams creates a new CreateActiveDirectoryConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateActiveDirectoryConfigurationParams() *CreateActiveDirectoryConfigurationParams {
	return &CreateActiveDirectoryConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateActiveDirectoryConfigurationParamsWithTimeout creates a new CreateActiveDirectoryConfigurationParams object
// with the ability to set a timeout on a request.
func NewCreateActiveDirectoryConfigurationParamsWithTimeout(timeout time.Duration) *CreateActiveDirectoryConfigurationParams {
	return &CreateActiveDirectoryConfigurationParams{
		timeout: timeout,
	}
}

// NewCreateActiveDirectoryConfigurationParamsWithContext creates a new CreateActiveDirectoryConfigurationParams object
// with the ability to set a context for a request.
func NewCreateActiveDirectoryConfigurationParamsWithContext(ctx context.Context) *CreateActiveDirectoryConfigurationParams {
	return &CreateActiveDirectoryConfigurationParams{
		Context: ctx,
	}
}

// NewCreateActiveDirectoryConfigurationParamsWithHTTPClient creates a new CreateActiveDirectoryConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateActiveDirectoryConfigurationParamsWithHTTPClient(client *http.Client) *CreateActiveDirectoryConfigurationParams {
	return &CreateActiveDirectoryConfigurationParams{
		HTTPClient: client,
	}
}

/*
CreateActiveDirectoryConfigurationParams contains all the parameters to send to the API endpoint

	for the create active directory configuration operation.

	Typically these are written to a http.Request.
*/
type CreateActiveDirectoryConfigurationParams struct {

	/* Body.

	   The Active Directory configuration
	*/
	Body *models.ActiveDirectorySettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create active directory configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateActiveDirectoryConfigurationParams) WithDefaults() *CreateActiveDirectoryConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create active directory configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateActiveDirectoryConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create active directory configuration params
func (o *CreateActiveDirectoryConfigurationParams) WithTimeout(timeout time.Duration) *CreateActiveDirectoryConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create active directory configuration params
func (o *CreateActiveDirectoryConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create active directory configuration params
func (o *CreateActiveDirectoryConfigurationParams) WithContext(ctx context.Context) *CreateActiveDirectoryConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create active directory configuration params
func (o *CreateActiveDirectoryConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create active directory configuration params
func (o *CreateActiveDirectoryConfigurationParams) WithHTTPClient(client *http.Client) *CreateActiveDirectoryConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create active directory configuration params
func (o *CreateActiveDirectoryConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create active directory configuration params
func (o *CreateActiveDirectoryConfigurationParams) WithBody(body *models.ActiveDirectorySettings) *CreateActiveDirectoryConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create active directory configuration params
func (o *CreateActiveDirectoryConfigurationParams) SetBody(body *models.ActiveDirectorySettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateActiveDirectoryConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
