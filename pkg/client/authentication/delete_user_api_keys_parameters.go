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

package authentication

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

// NewDeleteUserAPIKeysParams creates a new DeleteUserAPIKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserAPIKeysParams() *DeleteUserAPIKeysParams {
	return &DeleteUserAPIKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserAPIKeysParamsWithTimeout creates a new DeleteUserAPIKeysParams object
// with the ability to set a timeout on a request.
func NewDeleteUserAPIKeysParamsWithTimeout(timeout time.Duration) *DeleteUserAPIKeysParams {
	return &DeleteUserAPIKeysParams{
		timeout: timeout,
	}
}

// NewDeleteUserAPIKeysParamsWithContext creates a new DeleteUserAPIKeysParams object
// with the ability to set a context for a request.
func NewDeleteUserAPIKeysParamsWithContext(ctx context.Context) *DeleteUserAPIKeysParams {
	return &DeleteUserAPIKeysParams{
		Context: ctx,
	}
}

// NewDeleteUserAPIKeysParamsWithHTTPClient creates a new DeleteUserAPIKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserAPIKeysParamsWithHTTPClient(client *http.Client) *DeleteUserAPIKeysParams {
	return &DeleteUserAPIKeysParams{
		HTTPClient: client,
	}
}

/*
DeleteUserAPIKeysParams contains all the parameters to send to the API endpoint

	for the delete user api keys operation.

	Typically these are written to a http.Request.
*/
type DeleteUserAPIKeysParams struct {

	/* UserID.

	   The user ID.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserAPIKeysParams) WithDefaults() *DeleteUserAPIKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user api keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserAPIKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user api keys params
func (o *DeleteUserAPIKeysParams) WithTimeout(timeout time.Duration) *DeleteUserAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user api keys params
func (o *DeleteUserAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user api keys params
func (o *DeleteUserAPIKeysParams) WithContext(ctx context.Context) *DeleteUserAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user api keys params
func (o *DeleteUserAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user api keys params
func (o *DeleteUserAPIKeysParams) WithHTTPClient(client *http.Client) *DeleteUserAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user api keys params
func (o *DeleteUserAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user api keys params
func (o *DeleteUserAPIKeysParams) WithUserID(userID string) *DeleteUserAPIKeysParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user api keys params
func (o *DeleteUserAPIKeysParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
