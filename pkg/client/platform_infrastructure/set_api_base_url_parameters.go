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

// NewSetAPIBaseURLParams creates a new SetAPIBaseURLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetAPIBaseURLParams() *SetAPIBaseURLParams {
	return &SetAPIBaseURLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetAPIBaseURLParamsWithTimeout creates a new SetAPIBaseURLParams object
// with the ability to set a timeout on a request.
func NewSetAPIBaseURLParamsWithTimeout(timeout time.Duration) *SetAPIBaseURLParams {
	return &SetAPIBaseURLParams{
		timeout: timeout,
	}
}

// NewSetAPIBaseURLParamsWithContext creates a new SetAPIBaseURLParams object
// with the ability to set a context for a request.
func NewSetAPIBaseURLParamsWithContext(ctx context.Context) *SetAPIBaseURLParams {
	return &SetAPIBaseURLParams{
		Context: ctx,
	}
}

// NewSetAPIBaseURLParamsWithHTTPClient creates a new SetAPIBaseURLParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetAPIBaseURLParamsWithHTTPClient(client *http.Client) *SetAPIBaseURLParams {
	return &SetAPIBaseURLParams{
		HTTPClient: client,
	}
}

/*
SetAPIBaseURLParams contains all the parameters to send to the API endpoint

	for the set api base url operation.

	Typically these are written to a http.Request.
*/
type SetAPIBaseURLParams struct {

	/* Body.

	   Data containing the base Url to set
	*/
	Body *models.APIBaseURLData

	/* SkipCascadingOperations.

	   Whether or not to skip cascading operations, such as re-provisioning the Security Deployment.
	*/
	SkipCascadingOperations *bool

	/* Version.

	   If specified, then checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request). If not specified, will unconditionally upsert.
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set api base url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAPIBaseURLParams) WithDefaults() *SetAPIBaseURLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set api base url params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAPIBaseURLParams) SetDefaults() {
	var (
		skipCascadingOperationsDefault = bool(false)
	)

	val := SetAPIBaseURLParams{
		SkipCascadingOperations: &skipCascadingOperationsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the set api base url params
func (o *SetAPIBaseURLParams) WithTimeout(timeout time.Duration) *SetAPIBaseURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set api base url params
func (o *SetAPIBaseURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set api base url params
func (o *SetAPIBaseURLParams) WithContext(ctx context.Context) *SetAPIBaseURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set api base url params
func (o *SetAPIBaseURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set api base url params
func (o *SetAPIBaseURLParams) WithHTTPClient(client *http.Client) *SetAPIBaseURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set api base url params
func (o *SetAPIBaseURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set api base url params
func (o *SetAPIBaseURLParams) WithBody(body *models.APIBaseURLData) *SetAPIBaseURLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set api base url params
func (o *SetAPIBaseURLParams) SetBody(body *models.APIBaseURLData) {
	o.Body = body
}

// WithSkipCascadingOperations adds the skipCascadingOperations to the set api base url params
func (o *SetAPIBaseURLParams) WithSkipCascadingOperations(skipCascadingOperations *bool) *SetAPIBaseURLParams {
	o.SetSkipCascadingOperations(skipCascadingOperations)
	return o
}

// SetSkipCascadingOperations adds the skipCascadingOperations to the set api base url params
func (o *SetAPIBaseURLParams) SetSkipCascadingOperations(skipCascadingOperations *bool) {
	o.SkipCascadingOperations = skipCascadingOperations
}

// WithVersion adds the version to the set api base url params
func (o *SetAPIBaseURLParams) WithVersion(version *string) *SetAPIBaseURLParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the set api base url params
func (o *SetAPIBaseURLParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SetAPIBaseURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.SkipCascadingOperations != nil {

		// query param skip_cascading_operations
		var qrSkipCascadingOperations bool

		if o.SkipCascadingOperations != nil {
			qrSkipCascadingOperations = *o.SkipCascadingOperations
		}
		qSkipCascadingOperations := swag.FormatBool(qrSkipCascadingOperations)
		if qSkipCascadingOperations != "" {

			if err := r.SetQueryParam("skip_cascading_operations", qSkipCascadingOperations); err != nil {
				return err
			}
		}
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
