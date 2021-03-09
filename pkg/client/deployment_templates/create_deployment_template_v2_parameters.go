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

package deployment_templates

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

// NewCreateDeploymentTemplateV2Params creates a new CreateDeploymentTemplateV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDeploymentTemplateV2Params() *CreateDeploymentTemplateV2Params {
	return &CreateDeploymentTemplateV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeploymentTemplateV2ParamsWithTimeout creates a new CreateDeploymentTemplateV2Params object
// with the ability to set a timeout on a request.
func NewCreateDeploymentTemplateV2ParamsWithTimeout(timeout time.Duration) *CreateDeploymentTemplateV2Params {
	return &CreateDeploymentTemplateV2Params{
		timeout: timeout,
	}
}

// NewCreateDeploymentTemplateV2ParamsWithContext creates a new CreateDeploymentTemplateV2Params object
// with the ability to set a context for a request.
func NewCreateDeploymentTemplateV2ParamsWithContext(ctx context.Context) *CreateDeploymentTemplateV2Params {
	return &CreateDeploymentTemplateV2Params{
		Context: ctx,
	}
}

// NewCreateDeploymentTemplateV2ParamsWithHTTPClient creates a new CreateDeploymentTemplateV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDeploymentTemplateV2ParamsWithHTTPClient(client *http.Client) *CreateDeploymentTemplateV2Params {
	return &CreateDeploymentTemplateV2Params{
		HTTPClient: client,
	}
}

/* CreateDeploymentTemplateV2Params contains all the parameters to send to the API endpoint
   for the create deployment template v2 operation.

   Typically these are written to a http.Request.
*/
type CreateDeploymentTemplateV2Params struct {

	/* Body.

	   The deployment template definition.
	*/
	Body *models.DeploymentTemplateRequestBody

	/* Region.

	   Region of the deployment template
	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create deployment template v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeploymentTemplateV2Params) WithDefaults() *CreateDeploymentTemplateV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create deployment template v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDeploymentTemplateV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) WithTimeout(timeout time.Duration) *CreateDeploymentTemplateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) WithContext(ctx context.Context) *CreateDeploymentTemplateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) WithHTTPClient(client *http.Client) *CreateDeploymentTemplateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) WithBody(body *models.DeploymentTemplateRequestBody) *CreateDeploymentTemplateV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) SetBody(body *models.DeploymentTemplateRequestBody) {
	o.Body = body
}

// WithRegion adds the region to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) WithRegion(region string) *CreateDeploymentTemplateV2Params {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the create deployment template v2 params
func (o *CreateDeploymentTemplateV2Params) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeploymentTemplateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {

		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
