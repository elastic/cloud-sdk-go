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

// NewCaptureDeploymentInstanceHeapDumpParams creates a new CaptureDeploymentInstanceHeapDumpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCaptureDeploymentInstanceHeapDumpParams() *CaptureDeploymentInstanceHeapDumpParams {
	return &CaptureDeploymentInstanceHeapDumpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCaptureDeploymentInstanceHeapDumpParamsWithTimeout creates a new CaptureDeploymentInstanceHeapDumpParams object
// with the ability to set a timeout on a request.
func NewCaptureDeploymentInstanceHeapDumpParamsWithTimeout(timeout time.Duration) *CaptureDeploymentInstanceHeapDumpParams {
	return &CaptureDeploymentInstanceHeapDumpParams{
		timeout: timeout,
	}
}

// NewCaptureDeploymentInstanceHeapDumpParamsWithContext creates a new CaptureDeploymentInstanceHeapDumpParams object
// with the ability to set a context for a request.
func NewCaptureDeploymentInstanceHeapDumpParamsWithContext(ctx context.Context) *CaptureDeploymentInstanceHeapDumpParams {
	return &CaptureDeploymentInstanceHeapDumpParams{
		Context: ctx,
	}
}

// NewCaptureDeploymentInstanceHeapDumpParamsWithHTTPClient creates a new CaptureDeploymentInstanceHeapDumpParams object
// with the ability to set a custom HTTPClient for a request.
func NewCaptureDeploymentInstanceHeapDumpParamsWithHTTPClient(client *http.Client) *CaptureDeploymentInstanceHeapDumpParams {
	return &CaptureDeploymentInstanceHeapDumpParams{
		HTTPClient: client,
	}
}

/* CaptureDeploymentInstanceHeapDumpParams contains all the parameters to send to the API endpoint
   for the capture deployment instance heap dump operation.

   Typically these are written to a http.Request.
*/
type CaptureDeploymentInstanceHeapDumpParams struct {

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* InstanceID.

	   The instance identifier
	*/
	InstanceID string

	/* RefID.

	   User-specified RefId for the Resource.
	*/
	RefID string

	/* ResourceKind.

	   The kind of resource (one of elasticsearch, kibana or apm).
	*/
	ResourceKind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the capture deployment instance heap dump params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CaptureDeploymentInstanceHeapDumpParams) WithDefaults() *CaptureDeploymentInstanceHeapDumpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the capture deployment instance heap dump params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CaptureDeploymentInstanceHeapDumpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) WithTimeout(timeout time.Duration) *CaptureDeploymentInstanceHeapDumpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) WithContext(ctx context.Context) *CaptureDeploymentInstanceHeapDumpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) WithHTTPClient(client *http.Client) *CaptureDeploymentInstanceHeapDumpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) WithDeploymentID(deploymentID string) *CaptureDeploymentInstanceHeapDumpParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithInstanceID adds the instanceID to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) WithInstanceID(instanceID string) *CaptureDeploymentInstanceHeapDumpParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithRefID adds the refID to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) WithRefID(refID string) *CaptureDeploymentInstanceHeapDumpParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithResourceKind adds the resourceKind to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) WithResourceKind(resourceKind string) *CaptureDeploymentInstanceHeapDumpParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the capture deployment instance heap dump params
func (o *CaptureDeploymentInstanceHeapDumpParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WriteToRequest writes these params to a swagger request
func (o *CaptureDeploymentInstanceHeapDumpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
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
