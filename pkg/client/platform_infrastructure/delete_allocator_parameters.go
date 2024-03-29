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
)

// NewDeleteAllocatorParams creates a new DeleteAllocatorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAllocatorParams() *DeleteAllocatorParams {
	return &DeleteAllocatorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllocatorParamsWithTimeout creates a new DeleteAllocatorParams object
// with the ability to set a timeout on a request.
func NewDeleteAllocatorParamsWithTimeout(timeout time.Duration) *DeleteAllocatorParams {
	return &DeleteAllocatorParams{
		timeout: timeout,
	}
}

// NewDeleteAllocatorParamsWithContext creates a new DeleteAllocatorParams object
// with the ability to set a context for a request.
func NewDeleteAllocatorParamsWithContext(ctx context.Context) *DeleteAllocatorParams {
	return &DeleteAllocatorParams{
		Context: ctx,
	}
}

// NewDeleteAllocatorParamsWithHTTPClient creates a new DeleteAllocatorParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAllocatorParamsWithHTTPClient(client *http.Client) *DeleteAllocatorParams {
	return &DeleteAllocatorParams{
		HTTPClient: client,
	}
}

/*
DeleteAllocatorParams contains all the parameters to send to the API endpoint

	for the delete allocator operation.

	Typically these are written to a http.Request.
*/
type DeleteAllocatorParams struct {

	/* AllocatorID.

	   The allocator identifier.
	*/
	AllocatorID string

	/* RemoveInstances.

	   Removes the instances from the deleted allocator.
	*/
	RemoveInstances *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete allocator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllocatorParams) WithDefaults() *DeleteAllocatorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete allocator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllocatorParams) SetDefaults() {
	var (
		removeInstancesDefault = bool(false)
	)

	val := DeleteAllocatorParams{
		RemoveInstances: &removeInstancesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete allocator params
func (o *DeleteAllocatorParams) WithTimeout(timeout time.Duration) *DeleteAllocatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete allocator params
func (o *DeleteAllocatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete allocator params
func (o *DeleteAllocatorParams) WithContext(ctx context.Context) *DeleteAllocatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete allocator params
func (o *DeleteAllocatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete allocator params
func (o *DeleteAllocatorParams) WithHTTPClient(client *http.Client) *DeleteAllocatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete allocator params
func (o *DeleteAllocatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorID adds the allocatorID to the delete allocator params
func (o *DeleteAllocatorParams) WithAllocatorID(allocatorID string) *DeleteAllocatorParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the delete allocator params
func (o *DeleteAllocatorParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WithRemoveInstances adds the removeInstances to the delete allocator params
func (o *DeleteAllocatorParams) WithRemoveInstances(removeInstances *bool) *DeleteAllocatorParams {
	o.SetRemoveInstances(removeInstances)
	return o
}

// SetRemoveInstances adds the removeInstances to the delete allocator params
func (o *DeleteAllocatorParams) SetRemoveInstances(removeInstances *bool) {
	o.RemoveInstances = removeInstances
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllocatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param allocator_id
	if err := r.SetPathParam("allocator_id", o.AllocatorID); err != nil {
		return err
	}

	if o.RemoveInstances != nil {

		// query param remove_instances
		var qrRemoveInstances bool

		if o.RemoveInstances != nil {
			qrRemoveInstances = *o.RemoveInstances
		}
		qRemoveInstances := swag.FormatBool(qrRemoveInstances)
		if qRemoveInstances != "" {

			if err := r.SetQueryParam("remove_instances", qRemoveInstances); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
