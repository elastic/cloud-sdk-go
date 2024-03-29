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

package comments

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

// NewListCommentParams creates a new ListCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListCommentParams() *ListCommentParams {
	return &ListCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListCommentParamsWithTimeout creates a new ListCommentParams object
// with the ability to set a timeout on a request.
func NewListCommentParamsWithTimeout(timeout time.Duration) *ListCommentParams {
	return &ListCommentParams{
		timeout: timeout,
	}
}

// NewListCommentParamsWithContext creates a new ListCommentParams object
// with the ability to set a context for a request.
func NewListCommentParamsWithContext(ctx context.Context) *ListCommentParams {
	return &ListCommentParams{
		Context: ctx,
	}
}

// NewListCommentParamsWithHTTPClient creates a new ListCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewListCommentParamsWithHTTPClient(client *http.Client) *ListCommentParams {
	return &ListCommentParams{
		HTTPClient: client,
	}
}

/*
ListCommentParams contains all the parameters to send to the API endpoint

	for the list comment operation.

	Typically these are written to a http.Request.
*/
type ListCommentParams struct {

	/* ResourceID.

	   Id of the Resource that a Comment belongs to.
	*/
	ResourceID string

	/* ResourceType.

	   The kind of Resource that a Comment belongs to. Should be one of [elasticsearch, kibana, apm, appsearch, enterprise_search, integrations_server, allocator, constructor, runner, proxy].
	*/
	ResourceType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCommentParams) WithDefaults() *ListCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list comment params
func (o *ListCommentParams) WithTimeout(timeout time.Duration) *ListCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list comment params
func (o *ListCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list comment params
func (o *ListCommentParams) WithContext(ctx context.Context) *ListCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list comment params
func (o *ListCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list comment params
func (o *ListCommentParams) WithHTTPClient(client *http.Client) *ListCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list comment params
func (o *ListCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the list comment params
func (o *ListCommentParams) WithResourceID(resourceID string) *ListCommentParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the list comment params
func (o *ListCommentParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the list comment params
func (o *ListCommentParams) WithResourceType(resourceType string) *ListCommentParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the list comment params
func (o *ListCommentParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *ListCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_id
	if err := r.SetPathParam("resource_id", o.ResourceID); err != nil {
		return err
	}

	// path param resource_type
	if err := r.SetPathParam("resource_type", o.ResourceType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
