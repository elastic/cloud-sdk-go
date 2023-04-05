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

// NewStartDeploymentResourceInstancesAllMaintenanceModeParams creates a new StartDeploymentResourceInstancesAllMaintenanceModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartDeploymentResourceInstancesAllMaintenanceModeParams() *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	return &StartDeploymentResourceInstancesAllMaintenanceModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartDeploymentResourceInstancesAllMaintenanceModeParamsWithTimeout creates a new StartDeploymentResourceInstancesAllMaintenanceModeParams object
// with the ability to set a timeout on a request.
func NewStartDeploymentResourceInstancesAllMaintenanceModeParamsWithTimeout(timeout time.Duration) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	return &StartDeploymentResourceInstancesAllMaintenanceModeParams{
		timeout: timeout,
	}
}

// NewStartDeploymentResourceInstancesAllMaintenanceModeParamsWithContext creates a new StartDeploymentResourceInstancesAllMaintenanceModeParams object
// with the ability to set a context for a request.
func NewStartDeploymentResourceInstancesAllMaintenanceModeParamsWithContext(ctx context.Context) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	return &StartDeploymentResourceInstancesAllMaintenanceModeParams{
		Context: ctx,
	}
}

// NewStartDeploymentResourceInstancesAllMaintenanceModeParamsWithHTTPClient creates a new StartDeploymentResourceInstancesAllMaintenanceModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartDeploymentResourceInstancesAllMaintenanceModeParamsWithHTTPClient(client *http.Client) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	return &StartDeploymentResourceInstancesAllMaintenanceModeParams{
		HTTPClient: client,
	}
}

/*
StartDeploymentResourceInstancesAllMaintenanceModeParams contains all the parameters to send to the API endpoint

	for the start deployment resource instances all maintenance mode operation.

	Typically these are written to a http.Request.
*/
type StartDeploymentResourceInstancesAllMaintenanceModeParams struct {

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	/* ResourceKind.

	   The kind of resource (one of elasticsearch, kibana, apm, or integrations_server).
	*/
	ResourceKind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start deployment resource instances all maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) WithDefaults() *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start deployment resource instances all maintenance mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) WithTimeout(timeout time.Duration) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) WithContext(ctx context.Context) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) WithHTTPClient(client *http.Client) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) WithDeploymentID(deploymentID string) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) WithRefID(refID string) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithResourceKind adds the resourceKind to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) WithResourceKind(resourceKind string) *StartDeploymentResourceInstancesAllMaintenanceModeParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the start deployment resource instances all maintenance mode params
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WriteToRequest writes these params to a swagger request
func (o *StartDeploymentResourceInstancesAllMaintenanceModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resource_kind
	if err := r.SetPathParam("resource_kind", o.ResourceKind); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
