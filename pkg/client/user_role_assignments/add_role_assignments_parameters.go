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

package user_role_assignments

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

// NewAddRoleAssignmentsParams creates a new AddRoleAssignmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRoleAssignmentsParams() *AddRoleAssignmentsParams {
	return &AddRoleAssignmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRoleAssignmentsParamsWithTimeout creates a new AddRoleAssignmentsParams object
// with the ability to set a timeout on a request.
func NewAddRoleAssignmentsParamsWithTimeout(timeout time.Duration) *AddRoleAssignmentsParams {
	return &AddRoleAssignmentsParams{
		timeout: timeout,
	}
}

// NewAddRoleAssignmentsParamsWithContext creates a new AddRoleAssignmentsParams object
// with the ability to set a context for a request.
func NewAddRoleAssignmentsParamsWithContext(ctx context.Context) *AddRoleAssignmentsParams {
	return &AddRoleAssignmentsParams{
		Context: ctx,
	}
}

// NewAddRoleAssignmentsParamsWithHTTPClient creates a new AddRoleAssignmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRoleAssignmentsParamsWithHTTPClient(client *http.Client) *AddRoleAssignmentsParams {
	return &AddRoleAssignmentsParams{
		HTTPClient: client,
	}
}

/*
AddRoleAssignmentsParams contains all the parameters to send to the API endpoint

	for the add role assignments operation.

	Typically these are written to a http.Request.
*/
type AddRoleAssignmentsParams struct {

	/* Body.

	   The Role Assignments to add
	*/
	Body *models.RoleAssignments

	/* UserID.

	   Identifier for the user; include realm name and id if required
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add role assignments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRoleAssignmentsParams) WithDefaults() *AddRoleAssignmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add role assignments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRoleAssignmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add role assignments params
func (o *AddRoleAssignmentsParams) WithTimeout(timeout time.Duration) *AddRoleAssignmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add role assignments params
func (o *AddRoleAssignmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add role assignments params
func (o *AddRoleAssignmentsParams) WithContext(ctx context.Context) *AddRoleAssignmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add role assignments params
func (o *AddRoleAssignmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add role assignments params
func (o *AddRoleAssignmentsParams) WithHTTPClient(client *http.Client) *AddRoleAssignmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add role assignments params
func (o *AddRoleAssignmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add role assignments params
func (o *AddRoleAssignmentsParams) WithBody(body *models.RoleAssignments) *AddRoleAssignmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add role assignments params
func (o *AddRoleAssignmentsParams) SetBody(body *models.RoleAssignments) {
	o.Body = body
}

// WithUserID adds the userID to the add role assignments params
func (o *AddRoleAssignmentsParams) WithUserID(userID string) *AddRoleAssignmentsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the add role assignments params
func (o *AddRoleAssignmentsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AddRoleAssignmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
