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

package deployments

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

// NewSetDeploymentEsResourceRemoteClustersParams creates a new SetDeploymentEsResourceRemoteClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDeploymentEsResourceRemoteClustersParams() *SetDeploymentEsResourceRemoteClustersParams {
	return &SetDeploymentEsResourceRemoteClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDeploymentEsResourceRemoteClustersParamsWithTimeout creates a new SetDeploymentEsResourceRemoteClustersParams object
// with the ability to set a timeout on a request.
func NewSetDeploymentEsResourceRemoteClustersParamsWithTimeout(timeout time.Duration) *SetDeploymentEsResourceRemoteClustersParams {
	return &SetDeploymentEsResourceRemoteClustersParams{
		timeout: timeout,
	}
}

// NewSetDeploymentEsResourceRemoteClustersParamsWithContext creates a new SetDeploymentEsResourceRemoteClustersParams object
// with the ability to set a context for a request.
func NewSetDeploymentEsResourceRemoteClustersParamsWithContext(ctx context.Context) *SetDeploymentEsResourceRemoteClustersParams {
	return &SetDeploymentEsResourceRemoteClustersParams{
		Context: ctx,
	}
}

// NewSetDeploymentEsResourceRemoteClustersParamsWithHTTPClient creates a new SetDeploymentEsResourceRemoteClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDeploymentEsResourceRemoteClustersParamsWithHTTPClient(client *http.Client) *SetDeploymentEsResourceRemoteClustersParams {
	return &SetDeploymentEsResourceRemoteClustersParams{
		HTTPClient: client,
	}
}

/*
SetDeploymentEsResourceRemoteClustersParams contains all the parameters to send to the API endpoint

	for the set deployment es resource remote clusters operation.

	Typically these are written to a http.Request.
*/
type SetDeploymentEsResourceRemoteClustersParams struct {

	/* Body.

	   List of certificate based remote clusters for the resource
	*/
	Body *models.RemoteResources

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set deployment es resource remote clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDeploymentEsResourceRemoteClustersParams) WithDefaults() *SetDeploymentEsResourceRemoteClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set deployment es resource remote clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDeploymentEsResourceRemoteClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) WithTimeout(timeout time.Duration) *SetDeploymentEsResourceRemoteClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) WithContext(ctx context.Context) *SetDeploymentEsResourceRemoteClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) WithHTTPClient(client *http.Client) *SetDeploymentEsResourceRemoteClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) WithBody(body *models.RemoteResources) *SetDeploymentEsResourceRemoteClustersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) SetBody(body *models.RemoteResources) {
	o.Body = body
}

// WithDeploymentID adds the deploymentID to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) WithDeploymentID(deploymentID string) *SetDeploymentEsResourceRemoteClustersParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) WithRefID(refID string) *SetDeploymentEsResourceRemoteClustersParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the set deployment es resource remote clusters params
func (o *SetDeploymentEsResourceRemoteClustersParams) SetRefID(refID string) {
	o.RefID = refID
}

// WriteToRequest writes these params to a swagger request
func (o *SetDeploymentEsResourceRemoteClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	// path param ref_id
	if err := r.SetPathParam("ref_id", o.RefID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
