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

package platform_configuration_trust_relationships

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
)

// NewGetTrustRelationshipParams creates a new GetTrustRelationshipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTrustRelationshipParams() *GetTrustRelationshipParams {
	return &GetTrustRelationshipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTrustRelationshipParamsWithTimeout creates a new GetTrustRelationshipParams object
// with the ability to set a timeout on a request.
func NewGetTrustRelationshipParamsWithTimeout(timeout time.Duration) *GetTrustRelationshipParams {
	return &GetTrustRelationshipParams{
		timeout: timeout,
	}
}

// NewGetTrustRelationshipParamsWithContext creates a new GetTrustRelationshipParams object
// with the ability to set a context for a request.
func NewGetTrustRelationshipParamsWithContext(ctx context.Context) *GetTrustRelationshipParams {
	return &GetTrustRelationshipParams{
		Context: ctx,
	}
}

// NewGetTrustRelationshipParamsWithHTTPClient creates a new GetTrustRelationshipParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTrustRelationshipParamsWithHTTPClient(client *http.Client) *GetTrustRelationshipParams {
	return &GetTrustRelationshipParams{
		HTTPClient: client,
	}
}

/*
GetTrustRelationshipParams contains all the parameters to send to the API endpoint

	for the get trust relationship operation.

	Typically these are written to a http.Request.
*/
type GetTrustRelationshipParams struct {

	/* IncludeCertificate.

	   Whether to include the public CA certificates in the response.
	*/
	IncludeCertificate *bool

	/* TrustRelationshipID.

	   Identifier for the trust relationship
	*/
	TrustRelationshipID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get trust relationship params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrustRelationshipParams) WithDefaults() *GetTrustRelationshipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get trust relationship params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTrustRelationshipParams) SetDefaults() {
	var (
		includeCertificateDefault = bool(false)
	)

	val := GetTrustRelationshipParams{
		IncludeCertificate: &includeCertificateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get trust relationship params
func (o *GetTrustRelationshipParams) WithTimeout(timeout time.Duration) *GetTrustRelationshipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trust relationship params
func (o *GetTrustRelationshipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trust relationship params
func (o *GetTrustRelationshipParams) WithContext(ctx context.Context) *GetTrustRelationshipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trust relationship params
func (o *GetTrustRelationshipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trust relationship params
func (o *GetTrustRelationshipParams) WithHTTPClient(client *http.Client) *GetTrustRelationshipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trust relationship params
func (o *GetTrustRelationshipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncludeCertificate adds the includeCertificate to the get trust relationship params
func (o *GetTrustRelationshipParams) WithIncludeCertificate(includeCertificate *bool) *GetTrustRelationshipParams {
	o.SetIncludeCertificate(includeCertificate)
	return o
}

// SetIncludeCertificate adds the includeCertificate to the get trust relationship params
func (o *GetTrustRelationshipParams) SetIncludeCertificate(includeCertificate *bool) {
	o.IncludeCertificate = includeCertificate
}

// WithTrustRelationshipID adds the trustRelationshipID to the get trust relationship params
func (o *GetTrustRelationshipParams) WithTrustRelationshipID(trustRelationshipID string) *GetTrustRelationshipParams {
	o.SetTrustRelationshipID(trustRelationshipID)
	return o
}

// SetTrustRelationshipID adds the trustRelationshipId to the get trust relationship params
func (o *GetTrustRelationshipParams) SetTrustRelationshipID(trustRelationshipID string) {
	o.TrustRelationshipID = trustRelationshipID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTrustRelationshipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IncludeCertificate != nil {

		// query param include_certificate
		var qrIncludeCertificate bool

		if o.IncludeCertificate != nil {
			qrIncludeCertificate = *o.IncludeCertificate
		}
		qIncludeCertificate := swag.FormatBool(qrIncludeCertificate)
		if qIncludeCertificate != "" {

			if err := r.SetQueryParam("include_certificate", qIncludeCertificate); err != nil {
				return err
			}
		}
	}

	// path param trust_relationship_id
	if err := r.SetPathParam("trust_relationship_id", o.TrustRelationshipID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
