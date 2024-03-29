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

package deployments

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

// NewSearchEligibleRemoteClustersParams creates a new SearchEligibleRemoteClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchEligibleRemoteClustersParams() *SearchEligibleRemoteClustersParams {
	return &SearchEligibleRemoteClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchEligibleRemoteClustersParamsWithTimeout creates a new SearchEligibleRemoteClustersParams object
// with the ability to set a timeout on a request.
func NewSearchEligibleRemoteClustersParamsWithTimeout(timeout time.Duration) *SearchEligibleRemoteClustersParams {
	return &SearchEligibleRemoteClustersParams{
		timeout: timeout,
	}
}

// NewSearchEligibleRemoteClustersParamsWithContext creates a new SearchEligibleRemoteClustersParams object
// with the ability to set a context for a request.
func NewSearchEligibleRemoteClustersParamsWithContext(ctx context.Context) *SearchEligibleRemoteClustersParams {
	return &SearchEligibleRemoteClustersParams{
		Context: ctx,
	}
}

// NewSearchEligibleRemoteClustersParamsWithHTTPClient creates a new SearchEligibleRemoteClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchEligibleRemoteClustersParamsWithHTTPClient(client *http.Client) *SearchEligibleRemoteClustersParams {
	return &SearchEligibleRemoteClustersParams{
		HTTPClient: client,
	}
}

/*
SearchEligibleRemoteClustersParams contains all the parameters to send to the API endpoint

	for the search eligible remote clusters operation.

	Typically these are written to a http.Request.
*/
type SearchEligibleRemoteClustersParams struct {

	/* Body.

	   (Optional) The search query to run against all deployments containing eligible remote clusters. When not specified, all the eligible deployments are matched.
	*/
	Body *models.SearchRequest

	/* Version.

	   The version of the Elasticsearch cluster cluster that will potentially be configured to have remote clusters.
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search eligible remote clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchEligibleRemoteClustersParams) WithDefaults() *SearchEligibleRemoteClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search eligible remote clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchEligibleRemoteClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) WithTimeout(timeout time.Duration) *SearchEligibleRemoteClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) WithContext(ctx context.Context) *SearchEligibleRemoteClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) WithHTTPClient(client *http.Client) *SearchEligibleRemoteClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) WithBody(body *models.SearchRequest) *SearchEligibleRemoteClustersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithVersion adds the version to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) WithVersion(version string) *SearchEligibleRemoteClustersParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the search eligible remote clusters params
func (o *SearchEligibleRemoteClustersParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SearchEligibleRemoteClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {

		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
