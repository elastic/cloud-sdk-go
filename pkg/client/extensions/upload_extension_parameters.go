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

package extensions

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

// NewUploadExtensionParams creates a new UploadExtensionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadExtensionParams() *UploadExtensionParams {
	return &UploadExtensionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadExtensionParamsWithTimeout creates a new UploadExtensionParams object
// with the ability to set a timeout on a request.
func NewUploadExtensionParamsWithTimeout(timeout time.Duration) *UploadExtensionParams {
	return &UploadExtensionParams{
		timeout: timeout,
	}
}

// NewUploadExtensionParamsWithContext creates a new UploadExtensionParams object
// with the ability to set a context for a request.
func NewUploadExtensionParamsWithContext(ctx context.Context) *UploadExtensionParams {
	return &UploadExtensionParams{
		Context: ctx,
	}
}

// NewUploadExtensionParamsWithHTTPClient creates a new UploadExtensionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadExtensionParamsWithHTTPClient(client *http.Client) *UploadExtensionParams {
	return &UploadExtensionParams{
		HTTPClient: client,
	}
}

/* UploadExtensionParams contains all the parameters to send to the API endpoint
   for the upload extension operation.

   Typically these are written to a http.Request.
*/
type UploadExtensionParams struct {

	/* ExtensionID.

	   Id of an extension
	*/
	ExtensionID string

	/* File.

	   Zip file that contains the extension
	*/
	File runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload extension params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadExtensionParams) WithDefaults() *UploadExtensionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload extension params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadExtensionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload extension params
func (o *UploadExtensionParams) WithTimeout(timeout time.Duration) *UploadExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload extension params
func (o *UploadExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload extension params
func (o *UploadExtensionParams) WithContext(ctx context.Context) *UploadExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload extension params
func (o *UploadExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload extension params
func (o *UploadExtensionParams) WithHTTPClient(client *http.Client) *UploadExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload extension params
func (o *UploadExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtensionID adds the extensionID to the upload extension params
func (o *UploadExtensionParams) WithExtensionID(extensionID string) *UploadExtensionParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the upload extension params
func (o *UploadExtensionParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WithFile adds the file to the upload extension params
func (o *UploadExtensionParams) WithFile(file runtime.NamedReadCloser) *UploadExtensionParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload extension params
func (o *UploadExtensionParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *UploadExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extension_id
	if err := r.SetPathParam("extension_id", o.ExtensionID); err != nil {
		return err
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
