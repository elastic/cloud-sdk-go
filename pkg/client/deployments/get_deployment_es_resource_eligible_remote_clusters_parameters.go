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

// NewGetDeploymentEsResourceEligibleRemoteClustersParams creates a new GetDeploymentEsResourceEligibleRemoteClustersParams object
// with the default values initialized.
func NewGetDeploymentEsResourceEligibleRemoteClustersParams() *GetDeploymentEsResourceEligibleRemoteClustersParams {
	var ()
	return &GetDeploymentEsResourceEligibleRemoteClustersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentEsResourceEligibleRemoteClustersParamsWithTimeout creates a new GetDeploymentEsResourceEligibleRemoteClustersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentEsResourceEligibleRemoteClustersParamsWithTimeout(timeout time.Duration) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	var ()
	return &GetDeploymentEsResourceEligibleRemoteClustersParams{

		timeout: timeout,
	}
}

// NewGetDeploymentEsResourceEligibleRemoteClustersParamsWithContext creates a new GetDeploymentEsResourceEligibleRemoteClustersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentEsResourceEligibleRemoteClustersParamsWithContext(ctx context.Context) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	var ()
	return &GetDeploymentEsResourceEligibleRemoteClustersParams{

		Context: ctx,
	}
}

// NewGetDeploymentEsResourceEligibleRemoteClustersParamsWithHTTPClient creates a new GetDeploymentEsResourceEligibleRemoteClustersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentEsResourceEligibleRemoteClustersParamsWithHTTPClient(client *http.Client) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	var ()
	return &GetDeploymentEsResourceEligibleRemoteClustersParams{
		HTTPClient: client,
	}
}

/*GetDeploymentEsResourceEligibleRemoteClustersParams contains all the parameters to send to the API endpoint
for the get deployment es resource eligible remote clusters operation typically these are written to a http.Request
*/
type GetDeploymentEsResourceEligibleRemoteClustersParams struct {

	/*Body
	  (Optional) The search query to run against all deployments containing eligible remote clusters. When not specified, all the eligible deployments are matched.

	*/
	Body *models.SearchRequest
	/*DeploymentID
	  Identifier for the Deployment.

	*/
	DeploymentID string
	/*RefID
	  User-specified RefId for the Resource.

	*/
	RefID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) WithTimeout(timeout time.Duration) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) WithContext(ctx context.Context) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) WithHTTPClient(client *http.Client) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) WithBody(body *models.SearchRequest) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) SetBody(body *models.SearchRequest) {
	o.Body = body
}

// WithDeploymentID adds the deploymentID to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) WithDeploymentID(deploymentID string) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) WithRefID(refID string) *GetDeploymentEsResourceEligibleRemoteClustersParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the get deployment es resource eligible remote clusters params
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) SetRefID(refID string) {
	o.RefID = refID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentEsResourceEligibleRemoteClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
