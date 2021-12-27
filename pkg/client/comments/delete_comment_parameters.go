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

// NewDeleteCommentParams creates a new DeleteCommentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteCommentParams() *DeleteCommentParams {
	return &DeleteCommentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCommentParamsWithTimeout creates a new DeleteCommentParams object
// with the ability to set a timeout on a request.
func NewDeleteCommentParamsWithTimeout(timeout time.Duration) *DeleteCommentParams {
	return &DeleteCommentParams{
		timeout: timeout,
	}
}

// NewDeleteCommentParamsWithContext creates a new DeleteCommentParams object
// with the ability to set a context for a request.
func NewDeleteCommentParamsWithContext(ctx context.Context) *DeleteCommentParams {
	return &DeleteCommentParams{
		Context: ctx,
	}
}

// NewDeleteCommentParamsWithHTTPClient creates a new DeleteCommentParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteCommentParamsWithHTTPClient(client *http.Client) *DeleteCommentParams {
	return &DeleteCommentParams{
		HTTPClient: client,
	}
}

/* DeleteCommentParams contains all the parameters to send to the API endpoint
   for the delete comment operation.

   Typically these are written to a http.Request.
*/
type DeleteCommentParams struct {

	/* CommentID.

	   Id of a Comment
	*/
	CommentID string

	/* ResourceID.

	   Id of the Resource that a Comment belongs to.
	*/
	ResourceID string

	/* ResourceType.

	   The kind of Resource that a Comment belongs to. Should be one of [elasticsearch, kibana, apm, appsearch, enterprise_search, integrations_server, allocator, constructor, runner, proxy].
	*/
	ResourceType string

	/* Version.

	   If specified then checks for conflicts against the version stored in the persistent store (returned in 'x-cloud-resource-version' of the GET request)
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCommentParams) WithDefaults() *DeleteCommentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete comment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteCommentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete comment params
func (o *DeleteCommentParams) WithTimeout(timeout time.Duration) *DeleteCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete comment params
func (o *DeleteCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete comment params
func (o *DeleteCommentParams) WithContext(ctx context.Context) *DeleteCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete comment params
func (o *DeleteCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete comment params
func (o *DeleteCommentParams) WithHTTPClient(client *http.Client) *DeleteCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete comment params
func (o *DeleteCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentID adds the commentID to the delete comment params
func (o *DeleteCommentParams) WithCommentID(commentID string) *DeleteCommentParams {
	o.SetCommentID(commentID)
	return o
}

// SetCommentID adds the commentId to the delete comment params
func (o *DeleteCommentParams) SetCommentID(commentID string) {
	o.CommentID = commentID
}

// WithResourceID adds the resourceID to the delete comment params
func (o *DeleteCommentParams) WithResourceID(resourceID string) *DeleteCommentParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the delete comment params
func (o *DeleteCommentParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the delete comment params
func (o *DeleteCommentParams) WithResourceType(resourceType string) *DeleteCommentParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the delete comment params
func (o *DeleteCommentParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithVersion adds the version to the delete comment params
func (o *DeleteCommentParams) WithVersion(version *string) *DeleteCommentParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete comment params
func (o *DeleteCommentParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param comment_id
	if err := r.SetPathParam("comment_id", o.CommentID); err != nil {
		return err
	}

	// path param resource_id
	if err := r.SetPathParam("resource_id", o.ResourceID); err != nil {
		return err
	}

	// path param resource_type
	if err := r.SetPathParam("resource_type", o.ResourceType); err != nil {
		return err
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
