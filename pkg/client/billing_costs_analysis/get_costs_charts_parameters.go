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

package billing_costs_analysis

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

// NewGetCostsChartsParams creates a new GetCostsChartsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCostsChartsParams() *GetCostsChartsParams {
	return &GetCostsChartsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCostsChartsParamsWithTimeout creates a new GetCostsChartsParams object
// with the ability to set a timeout on a request.
func NewGetCostsChartsParamsWithTimeout(timeout time.Duration) *GetCostsChartsParams {
	return &GetCostsChartsParams{
		timeout: timeout,
	}
}

// NewGetCostsChartsParamsWithContext creates a new GetCostsChartsParams object
// with the ability to set a context for a request.
func NewGetCostsChartsParamsWithContext(ctx context.Context) *GetCostsChartsParams {
	return &GetCostsChartsParams{
		Context: ctx,
	}
}

// NewGetCostsChartsParamsWithHTTPClient creates a new GetCostsChartsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCostsChartsParamsWithHTTPClient(client *http.Client) *GetCostsChartsParams {
	return &GetCostsChartsParams{
		HTTPClient: client,
	}
}

/* GetCostsChartsParams contains all the parameters to send to the API endpoint
   for the get costs charts operation.

   Typically these are written to a http.Request.
*/
type GetCostsChartsParams struct {

	/* BucketingStrategy.

	   The desired bucketing strategy for the charts. Defaults to `daily`.

	   Default: "Daily"
	*/
	BucketingStrategy *string

	/* From.

	   A datetime for the beginning of the desired range for which to fetch costs. Defaults to start of current month.
	*/
	From *string

	/* OrganizationID.

	   Identifier for the organization
	*/
	OrganizationID string

	/* To.

	   A datetime for the end of the desired range for which to fetch costs. Defaults to the current date.
	*/
	To *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get costs charts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostsChartsParams) WithDefaults() *GetCostsChartsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get costs charts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostsChartsParams) SetDefaults() {
	var (
		bucketingStrategyDefault = string("Daily")
	)

	val := GetCostsChartsParams{
		BucketingStrategy: &bucketingStrategyDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get costs charts params
func (o *GetCostsChartsParams) WithTimeout(timeout time.Duration) *GetCostsChartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get costs charts params
func (o *GetCostsChartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get costs charts params
func (o *GetCostsChartsParams) WithContext(ctx context.Context) *GetCostsChartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get costs charts params
func (o *GetCostsChartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get costs charts params
func (o *GetCostsChartsParams) WithHTTPClient(client *http.Client) *GetCostsChartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get costs charts params
func (o *GetCostsChartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketingStrategy adds the bucketingStrategy to the get costs charts params
func (o *GetCostsChartsParams) WithBucketingStrategy(bucketingStrategy *string) *GetCostsChartsParams {
	o.SetBucketingStrategy(bucketingStrategy)
	return o
}

// SetBucketingStrategy adds the bucketingStrategy to the get costs charts params
func (o *GetCostsChartsParams) SetBucketingStrategy(bucketingStrategy *string) {
	o.BucketingStrategy = bucketingStrategy
}

// WithFrom adds the from to the get costs charts params
func (o *GetCostsChartsParams) WithFrom(from *string) *GetCostsChartsParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get costs charts params
func (o *GetCostsChartsParams) SetFrom(from *string) {
	o.From = from
}

// WithOrganizationID adds the organizationID to the get costs charts params
func (o *GetCostsChartsParams) WithOrganizationID(organizationID string) *GetCostsChartsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get costs charts params
func (o *GetCostsChartsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithTo adds the to to the get costs charts params
func (o *GetCostsChartsParams) WithTo(to *string) *GetCostsChartsParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get costs charts params
func (o *GetCostsChartsParams) SetTo(to *string) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetCostsChartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BucketingStrategy != nil {

		// query param bucketing_strategy
		var qrBucketingStrategy string

		if o.BucketingStrategy != nil {
			qrBucketingStrategy = *o.BucketingStrategy
		}
		qBucketingStrategy := qrBucketingStrategy
		if qBucketingStrategy != "" {

			if err := r.SetQueryParam("bucketing_strategy", qBucketingStrategy); err != nil {
				return err
			}
		}
	}

	if o.From != nil {

		// query param from
		var qrFrom string

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if o.To != nil {

		// query param to
		var qrTo string

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
