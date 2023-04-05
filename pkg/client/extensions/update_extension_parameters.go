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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewUpdateExtensionParams creates a new UpdateExtensionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateExtensionParams() *UpdateExtensionParams {
	return &UpdateExtensionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateExtensionParamsWithTimeout creates a new UpdateExtensionParams object
// with the ability to set a timeout on a request.
func NewUpdateExtensionParamsWithTimeout(timeout time.Duration) *UpdateExtensionParams {
	return &UpdateExtensionParams{
		timeout: timeout,
	}
}

// NewUpdateExtensionParamsWithContext creates a new UpdateExtensionParams object
// with the ability to set a context for a request.
func NewUpdateExtensionParamsWithContext(ctx context.Context) *UpdateExtensionParams {
	return &UpdateExtensionParams{
		Context: ctx,
	}
}

// NewUpdateExtensionParamsWithHTTPClient creates a new UpdateExtensionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateExtensionParamsWithHTTPClient(client *http.Client) *UpdateExtensionParams {
	return &UpdateExtensionParams{
		HTTPClient: client,
	}
}

/*
UpdateExtensionParams contains all the parameters to send to the API endpoint

	for the update extension operation.

	Typically these are written to a http.Request.
*/
type UpdateExtensionParams struct {

	/* Body.

	   The extension update data.
	*/
	Body *models.UpdateExtensionRequest

	/* ExtensionID.

	   Id of an extension
	*/
	ExtensionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update extension params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateExtensionParams) WithDefaults() *UpdateExtensionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update extension params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateExtensionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update extension params
func (o *UpdateExtensionParams) WithTimeout(timeout time.Duration) *UpdateExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update extension params
func (o *UpdateExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update extension params
func (o *UpdateExtensionParams) WithContext(ctx context.Context) *UpdateExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update extension params
func (o *UpdateExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update extension params
func (o *UpdateExtensionParams) WithHTTPClient(client *http.Client) *UpdateExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update extension params
func (o *UpdateExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update extension params
func (o *UpdateExtensionParams) WithBody(body *models.UpdateExtensionRequest) *UpdateExtensionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update extension params
func (o *UpdateExtensionParams) SetBody(body *models.UpdateExtensionRequest) {
	o.Body = body
}

// WithExtensionID adds the extensionID to the update extension params
func (o *UpdateExtensionParams) WithExtensionID(extensionID string) *UpdateExtensionParams {
	o.SetExtensionID(extensionID)
	return o
}

// SetExtensionID adds the extensionId to the update extension params
func (o *UpdateExtensionParams) SetExtensionID(extensionID string) {
	o.ExtensionID = extensionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param extension_id
	if err := r.SetPathParam("extension_id", o.ExtensionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
