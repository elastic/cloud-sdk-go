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
)

// NewSetTLSCertificateParams creates a new SetTLSCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetTLSCertificateParams() *SetTLSCertificateParams {
	return &SetTLSCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetTLSCertificateParamsWithTimeout creates a new SetTLSCertificateParams object
// with the ability to set a timeout on a request.
func NewSetTLSCertificateParamsWithTimeout(timeout time.Duration) *SetTLSCertificateParams {
	return &SetTLSCertificateParams{
		timeout: timeout,
	}
}

// NewSetTLSCertificateParamsWithContext creates a new SetTLSCertificateParams object
// with the ability to set a context for a request.
func NewSetTLSCertificateParamsWithContext(ctx context.Context) *SetTLSCertificateParams {
	return &SetTLSCertificateParams{
		Context: ctx,
	}
}

// NewSetTLSCertificateParamsWithHTTPClient creates a new SetTLSCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetTLSCertificateParamsWithHTTPClient(client *http.Client) *SetTLSCertificateParams {
	return &SetTLSCertificateParams{
		HTTPClient: client,
	}
}

/*
SetTLSCertificateParams contains all the parameters to send to the API endpoint

	for the set tls certificate operation.

	Typically these are written to a http.Request.
*/
type SetTLSCertificateParams struct {

	/* Chain.

	   New certificate chain: the PEM encoded RSA private key, followed by the server certificate, followed by the CA certificate
	*/
	Chain string

	/* ServiceName.

	   The service certificate chain to read.
	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set tls certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTLSCertificateParams) WithDefaults() *SetTLSCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set tls certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTLSCertificateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set tls certificate params
func (o *SetTLSCertificateParams) WithTimeout(timeout time.Duration) *SetTLSCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set tls certificate params
func (o *SetTLSCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set tls certificate params
func (o *SetTLSCertificateParams) WithContext(ctx context.Context) *SetTLSCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set tls certificate params
func (o *SetTLSCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set tls certificate params
func (o *SetTLSCertificateParams) WithHTTPClient(client *http.Client) *SetTLSCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set tls certificate params
func (o *SetTLSCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChain adds the chain to the set tls certificate params
func (o *SetTLSCertificateParams) WithChain(chain string) *SetTLSCertificateParams {
	o.SetChain(chain)
	return o
}

// SetChain adds the chain to the set tls certificate params
func (o *SetTLSCertificateParams) SetChain(chain string) {
	o.Chain = chain
}

// WithServiceName adds the serviceName to the set tls certificate params
func (o *SetTLSCertificateParams) WithServiceName(serviceName string) *SetTLSCertificateParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the set tls certificate params
func (o *SetTLSCertificateParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *SetTLSCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Chain); err != nil {
		return err
	}

	// path param service_name
	if err := r.SetPathParam("service_name", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
