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
)

// NewGetDeploymentEsResourceRemoteClustersParams creates a new GetDeploymentEsResourceRemoteClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentEsResourceRemoteClustersParams() *GetDeploymentEsResourceRemoteClustersParams {
	return &GetDeploymentEsResourceRemoteClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentEsResourceRemoteClustersParamsWithTimeout creates a new GetDeploymentEsResourceRemoteClustersParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentEsResourceRemoteClustersParamsWithTimeout(timeout time.Duration) *GetDeploymentEsResourceRemoteClustersParams {
	return &GetDeploymentEsResourceRemoteClustersParams{
		timeout: timeout,
	}
}

// NewGetDeploymentEsResourceRemoteClustersParamsWithContext creates a new GetDeploymentEsResourceRemoteClustersParams object
// with the ability to set a context for a request.
func NewGetDeploymentEsResourceRemoteClustersParamsWithContext(ctx context.Context) *GetDeploymentEsResourceRemoteClustersParams {
	return &GetDeploymentEsResourceRemoteClustersParams{
		Context: ctx,
	}
}

// NewGetDeploymentEsResourceRemoteClustersParamsWithHTTPClient creates a new GetDeploymentEsResourceRemoteClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentEsResourceRemoteClustersParamsWithHTTPClient(client *http.Client) *GetDeploymentEsResourceRemoteClustersParams {
	return &GetDeploymentEsResourceRemoteClustersParams{
		HTTPClient: client,
	}
}

/* GetDeploymentEsResourceRemoteClustersParams contains all the parameters to send to the API endpoint
   for the get deployment es resource remote clusters operation.

   Typically these are written to a http.Request.
*/
type GetDeploymentEsResourceRemoteClustersParams struct {

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource.
	*/
	RefID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployment es resource remote clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentEsResourceRemoteClustersParams) WithDefaults() *GetDeploymentEsResourceRemoteClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment es resource remote clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentEsResourceRemoteClustersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) WithTimeout(timeout time.Duration) *GetDeploymentEsResourceRemoteClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) WithContext(ctx context.Context) *GetDeploymentEsResourceRemoteClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) WithHTTPClient(client *http.Client) *GetDeploymentEsResourceRemoteClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) WithDeploymentID(deploymentID string) *GetDeploymentEsResourceRemoteClustersParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) WithRefID(refID string) *GetDeploymentEsResourceRemoteClustersParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the get deployment es resource remote clusters params
func (o *GetDeploymentEsResourceRemoteClustersParams) SetRefID(refID string) {
	o.RefID = refID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentEsResourceRemoteClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
