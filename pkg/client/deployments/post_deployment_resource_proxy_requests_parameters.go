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

// NewPostDeploymentResourceProxyRequestsParams creates a new PostDeploymentResourceProxyRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDeploymentResourceProxyRequestsParams() *PostDeploymentResourceProxyRequestsParams {
	return &PostDeploymentResourceProxyRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDeploymentResourceProxyRequestsParamsWithTimeout creates a new PostDeploymentResourceProxyRequestsParams object
// with the ability to set a timeout on a request.
func NewPostDeploymentResourceProxyRequestsParamsWithTimeout(timeout time.Duration) *PostDeploymentResourceProxyRequestsParams {
	return &PostDeploymentResourceProxyRequestsParams{
		timeout: timeout,
	}
}

// NewPostDeploymentResourceProxyRequestsParamsWithContext creates a new PostDeploymentResourceProxyRequestsParams object
// with the ability to set a context for a request.
func NewPostDeploymentResourceProxyRequestsParamsWithContext(ctx context.Context) *PostDeploymentResourceProxyRequestsParams {
	return &PostDeploymentResourceProxyRequestsParams{
		Context: ctx,
	}
}

// NewPostDeploymentResourceProxyRequestsParamsWithHTTPClient creates a new PostDeploymentResourceProxyRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDeploymentResourceProxyRequestsParamsWithHTTPClient(client *http.Client) *PostDeploymentResourceProxyRequestsParams {
	return &PostDeploymentResourceProxyRequestsParams{
		HTTPClient: client,
	}
}

/*
PostDeploymentResourceProxyRequestsParams contains all the parameters to send to the API endpoint

	for the post deployment resource proxy requests operation.

	Typically these are written to a http.Request.
*/
type PostDeploymentResourceProxyRequestsParams struct {

	/* XManagementRequest.

	   You must specify the `X-Management-Request` HTTP header with value `true`. NOTE: Use this endpoint for management purposes. It does not provide high performance.
	*/
	XManagementRequest string

	/* Body.

	   The JSON payload to proxy to the deployment resource.
	*/
	Body string

	/* DeploymentID.

	   Identifier for the Deployment
	*/
	DeploymentID string

	/* ProxyPath.

	   The URL part to proxy to the deployment resource. Example: _cat/indices, /api/spaces/space or /api/ent/v1/internal/health
	*/
	ProxyPath string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one)
	*/
	RefID string

	/* ResourceKind.

	   The kind of resource
	*/
	ResourceKind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post deployment resource proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDeploymentResourceProxyRequestsParams) WithDefaults() *PostDeploymentResourceProxyRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post deployment resource proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDeploymentResourceProxyRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithTimeout(timeout time.Duration) *PostDeploymentResourceProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithContext(ctx context.Context) *PostDeploymentResourceProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithHTTPClient(client *http.Client) *PostDeploymentResourceProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *PostDeploymentResourceProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithBody adds the body to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithBody(body string) *PostDeploymentResourceProxyRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetBody(body string) {
	o.Body = body
}

// WithDeploymentID adds the deploymentID to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithDeploymentID(deploymentID string) *PostDeploymentResourceProxyRequestsParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithProxyPath adds the proxyPath to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithProxyPath(proxyPath string) *PostDeploymentResourceProxyRequestsParams {
	o.SetProxyPath(proxyPath)
	return o
}

// SetProxyPath adds the proxyPath to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetProxyPath(proxyPath string) {
	o.ProxyPath = proxyPath
}

// WithRefID adds the refID to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithRefID(refID string) *PostDeploymentResourceProxyRequestsParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithResourceKind adds the resourceKind to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) WithResourceKind(resourceKind string) *PostDeploymentResourceProxyRequestsParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the post deployment resource proxy requests params
func (o *PostDeploymentResourceProxyRequestsParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WriteToRequest writes these params to a swagger request
func (o *PostDeploymentResourceProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	// path param proxy_path
	if err := r.SetPathParam("proxy_path", o.ProxyPath); err != nil {
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
