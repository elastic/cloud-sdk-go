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

// NewSearchRunnersParams creates a new SearchRunnersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchRunnersParams() *SearchRunnersParams {
	return &SearchRunnersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchRunnersParamsWithTimeout creates a new SearchRunnersParams object
// with the ability to set a timeout on a request.
func NewSearchRunnersParamsWithTimeout(timeout time.Duration) *SearchRunnersParams {
	return &SearchRunnersParams{
		timeout: timeout,
	}
}

// NewSearchRunnersParamsWithContext creates a new SearchRunnersParams object
// with the ability to set a context for a request.
func NewSearchRunnersParamsWithContext(ctx context.Context) *SearchRunnersParams {
	return &SearchRunnersParams{
		Context: ctx,
	}
}

// NewSearchRunnersParamsWithHTTPClient creates a new SearchRunnersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchRunnersParamsWithHTTPClient(client *http.Client) *SearchRunnersParams {
	return &SearchRunnersParams{
		HTTPClient: client,
	}
}

/*
SearchRunnersParams contains all the parameters to send to the API endpoint

	for the search runners operation.

	Typically these are written to a http.Request.
*/
type SearchRunnersParams struct {

	/* Body.

	   The optional search request to execute. If not supplied then all runners are matched.
	*/
	Body *models.SearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search runners params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRunnersParams) WithDefaults() *SearchRunnersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search runners params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchRunnersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search runners params
func (o *SearchRunnersParams) WithTimeout(timeout time.Duration) *SearchRunnersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search runners params
func (o *SearchRunnersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search runners params
func (o *SearchRunnersParams) WithContext(ctx context.Context) *SearchRunnersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search runners params
func (o *SearchRunnersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search runners params
func (o *SearchRunnersParams) WithHTTPClient(client *http.Client) *SearchRunnersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search runners params
func (o *SearchRunnersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search runners params
func (o *SearchRunnersParams) WithBody(body *models.SearchRequest) *SearchRunnersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search runners params
func (o *SearchRunnersParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchRunnersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
