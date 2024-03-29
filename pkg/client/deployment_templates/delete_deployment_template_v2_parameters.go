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
)

// NewDeleteDeploymentTemplateV2Params creates a new DeleteDeploymentTemplateV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeploymentTemplateV2Params() *DeleteDeploymentTemplateV2Params {
	return &DeleteDeploymentTemplateV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentTemplateV2ParamsWithTimeout creates a new DeleteDeploymentTemplateV2Params object
// with the ability to set a timeout on a request.
func NewDeleteDeploymentTemplateV2ParamsWithTimeout(timeout time.Duration) *DeleteDeploymentTemplateV2Params {
	return &DeleteDeploymentTemplateV2Params{
		timeout: timeout,
	}
}

// NewDeleteDeploymentTemplateV2ParamsWithContext creates a new DeleteDeploymentTemplateV2Params object
// with the ability to set a context for a request.
func NewDeleteDeploymentTemplateV2ParamsWithContext(ctx context.Context) *DeleteDeploymentTemplateV2Params {
	return &DeleteDeploymentTemplateV2Params{
		Context: ctx,
	}
}

// NewDeleteDeploymentTemplateV2ParamsWithHTTPClient creates a new DeleteDeploymentTemplateV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeploymentTemplateV2ParamsWithHTTPClient(client *http.Client) *DeleteDeploymentTemplateV2Params {
	return &DeleteDeploymentTemplateV2Params{
		HTTPClient: client,
	}
}

/*
DeleteDeploymentTemplateV2Params contains all the parameters to send to the API endpoint

	for the delete deployment template v2 operation.

	Typically these are written to a http.Request.
*/
type DeleteDeploymentTemplateV2Params struct {

	/* Region.

	   Region of the deployment template
	*/
	Region string

	/* TemplateID.

	   The identifier for the deployment template.
	*/
	TemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete deployment template v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeploymentTemplateV2Params) WithDefaults() *DeleteDeploymentTemplateV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete deployment template v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeploymentTemplateV2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) WithTimeout(timeout time.Duration) *DeleteDeploymentTemplateV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) WithContext(ctx context.Context) *DeleteDeploymentTemplateV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) WithHTTPClient(client *http.Client) *DeleteDeploymentTemplateV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegion adds the region to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) WithRegion(region string) *DeleteDeploymentTemplateV2Params {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) SetRegion(region string) {
	o.Region = region
}

// WithTemplateID adds the templateID to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) WithTemplateID(templateID string) *DeleteDeploymentTemplateV2Params {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the delete deployment template v2 params
func (o *DeleteDeploymentTemplateV2Params) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentTemplateV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param region
	qrRegion := o.Region
	qRegion := qrRegion
	if qRegion != "" {

		if err := r.SetQueryParam("region", qRegion); err != nil {
			return err
		}
	}

	// path param template_id
	if err := r.SetPathParam("template_id", o.TemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
