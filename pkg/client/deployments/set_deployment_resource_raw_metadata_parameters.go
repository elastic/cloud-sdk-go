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

// NewSetDeploymentResourceRawMetadataParams creates a new SetDeploymentResourceRawMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDeploymentResourceRawMetadataParams() *SetDeploymentResourceRawMetadataParams {
	return &SetDeploymentResourceRawMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDeploymentResourceRawMetadataParamsWithTimeout creates a new SetDeploymentResourceRawMetadataParams object
// with the ability to set a timeout on a request.
func NewSetDeploymentResourceRawMetadataParamsWithTimeout(timeout time.Duration) *SetDeploymentResourceRawMetadataParams {
	return &SetDeploymentResourceRawMetadataParams{
		timeout: timeout,
	}
}

// NewSetDeploymentResourceRawMetadataParamsWithContext creates a new SetDeploymentResourceRawMetadataParams object
// with the ability to set a context for a request.
func NewSetDeploymentResourceRawMetadataParamsWithContext(ctx context.Context) *SetDeploymentResourceRawMetadataParams {
	return &SetDeploymentResourceRawMetadataParams{
		Context: ctx,
	}
}

// NewSetDeploymentResourceRawMetadataParamsWithHTTPClient creates a new SetDeploymentResourceRawMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDeploymentResourceRawMetadataParamsWithHTTPClient(client *http.Client) *SetDeploymentResourceRawMetadataParams {
	return &SetDeploymentResourceRawMetadataParams{
		HTTPClient: client,
	}
}

/* SetDeploymentResourceRawMetadataParams contains all the parameters to send to the API endpoint
   for the set deployment resource raw metadata operation.

   Typically these are written to a http.Request.
*/
type SetDeploymentResourceRawMetadataParams struct {

	/* Body.

	   The freeform JSON for the cluster (should always be based on the current version retrieved from the GET)
	*/
	Body string

	/* DeploymentID.

	   Identifier for the Deployment
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one)
	*/
	RefID string

	/* ResourceKind.

	   The kind of resource
	*/
	ResourceKind string

	/* Version.

	   If specified, checks for conflicts against the metadata version (returned in 'x-cloud-resource-version' of the GET request)
	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set deployment resource raw metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDeploymentResourceRawMetadataParams) WithDefaults() *SetDeploymentResourceRawMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set deployment resource raw metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDeploymentResourceRawMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) WithTimeout(timeout time.Duration) *SetDeploymentResourceRawMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) WithContext(ctx context.Context) *SetDeploymentResourceRawMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) WithHTTPClient(client *http.Client) *SetDeploymentResourceRawMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) WithBody(body string) *SetDeploymentResourceRawMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) SetBody(body string) {
	o.Body = body
}

// WithDeploymentID adds the deploymentID to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) WithDeploymentID(deploymentID string) *SetDeploymentResourceRawMetadataParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) WithRefID(refID string) *SetDeploymentResourceRawMetadataParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithResourceKind adds the resourceKind to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) WithResourceKind(resourceKind string) *SetDeploymentResourceRawMetadataParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WithVersion adds the version to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) WithVersion(version *string) *SetDeploymentResourceRawMetadataParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the set deployment resource raw metadata params
func (o *SetDeploymentResourceRawMetadataParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SetDeploymentResourceRawMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

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
