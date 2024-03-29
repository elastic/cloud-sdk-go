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

// NewCreateBlueprinterRoleParams creates a new CreateBlueprinterRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBlueprinterRoleParams() *CreateBlueprinterRoleParams {
	return &CreateBlueprinterRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprinterRoleParamsWithTimeout creates a new CreateBlueprinterRoleParams object
// with the ability to set a timeout on a request.
func NewCreateBlueprinterRoleParamsWithTimeout(timeout time.Duration) *CreateBlueprinterRoleParams {
	return &CreateBlueprinterRoleParams{
		timeout: timeout,
	}
}

// NewCreateBlueprinterRoleParamsWithContext creates a new CreateBlueprinterRoleParams object
// with the ability to set a context for a request.
func NewCreateBlueprinterRoleParamsWithContext(ctx context.Context) *CreateBlueprinterRoleParams {
	return &CreateBlueprinterRoleParams{
		Context: ctx,
	}
}

// NewCreateBlueprinterRoleParamsWithHTTPClient creates a new CreateBlueprinterRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBlueprinterRoleParamsWithHTTPClient(client *http.Client) *CreateBlueprinterRoleParams {
	return &CreateBlueprinterRoleParams{
		HTTPClient: client,
	}
}

/*
CreateBlueprinterRoleParams contains all the parameters to send to the API endpoint

	for the create blueprinter role operation.

	Typically these are written to a http.Request.
*/
type CreateBlueprinterRoleParams struct {

	/* Body.

	   The data you want to use for creating a role.
	*/
	Body *models.RoleAggregateCreateData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create blueprinter role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprinterRoleParams) WithDefaults() *CreateBlueprinterRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create blueprinter role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBlueprinterRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create blueprinter role params
func (o *CreateBlueprinterRoleParams) WithTimeout(timeout time.Duration) *CreateBlueprinterRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprinter role params
func (o *CreateBlueprinterRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprinter role params
func (o *CreateBlueprinterRoleParams) WithContext(ctx context.Context) *CreateBlueprinterRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprinter role params
func (o *CreateBlueprinterRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprinter role params
func (o *CreateBlueprinterRoleParams) WithHTTPClient(client *http.Client) *CreateBlueprinterRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprinter role params
func (o *CreateBlueprinterRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create blueprinter role params
func (o *CreateBlueprinterRoleParams) WithBody(body *models.RoleAggregateCreateData) *CreateBlueprinterRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create blueprinter role params
func (o *CreateBlueprinterRoleParams) SetBody(body *models.RoleAggregateCreateData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprinterRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
