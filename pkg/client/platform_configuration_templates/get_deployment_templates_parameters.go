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

package platform_configuration_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeploymentTemplatesParams creates a new GetDeploymentTemplatesParams object
// with the default values initialized.
func NewGetDeploymentTemplatesParams() *GetDeploymentTemplatesParams {
	var (
		formatDefault                     = string("cluster")
		showHiddenDefault                 = bool(false)
		showInstanceConfigurationsDefault = bool(false)
	)
	return &GetDeploymentTemplatesParams{
		Format:                     &formatDefault,
		ShowHidden:                 &showHiddenDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentTemplatesParamsWithTimeout creates a new GetDeploymentTemplatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentTemplatesParamsWithTimeout(timeout time.Duration) *GetDeploymentTemplatesParams {
	var (
		formatDefault                     = string("cluster")
		showHiddenDefault                 = bool(false)
		showInstanceConfigurationsDefault = bool(false)
	)
	return &GetDeploymentTemplatesParams{
		Format:                     &formatDefault,
		ShowHidden:                 &showHiddenDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,

		timeout: timeout,
	}
}

// NewGetDeploymentTemplatesParamsWithContext creates a new GetDeploymentTemplatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentTemplatesParamsWithContext(ctx context.Context) *GetDeploymentTemplatesParams {
	var (
		formatDefault                     = string("cluster")
		showHiddenDefault                 = bool(false)
		showInstanceConfigurationsDefault = bool(false)
	)
	return &GetDeploymentTemplatesParams{
		Format:                     &formatDefault,
		ShowHidden:                 &showHiddenDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,

		Context: ctx,
	}
}

// NewGetDeploymentTemplatesParamsWithHTTPClient creates a new GetDeploymentTemplatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentTemplatesParamsWithHTTPClient(client *http.Client) *GetDeploymentTemplatesParams {
	var (
		formatDefault                     = string("cluster")
		showHiddenDefault                 = bool(false)
		showInstanceConfigurationsDefault = bool(false)
	)
	return &GetDeploymentTemplatesParams{
		Format:                     &formatDefault,
		ShowHidden:                 &showHiddenDefault,
		ShowInstanceConfigurations: &showInstanceConfigurationsDefault,
		HTTPClient:                 client,
	}
}

/*GetDeploymentTemplatesParams contains all the parameters to send to the API endpoint
for the get deployment templates operation typically these are written to a http.Request
*/
type GetDeploymentTemplatesParams struct {

	/*Format
	  If cluster is specified populates cluster_template in the response, if deployment is specified populates deployment_template in the response

	*/
	Format *string
	/*Metadata
	  An optional key/value pair in the form of (key:value) that will act as a filter and exclude any templates that do not have a matching metadata item associated.

	*/
	Metadata *string
	/*ShowHidden
	  If true, templates flagged as hidden will be returned.

	*/
	ShowHidden *bool
	/*ShowInstanceConfigurations
	  If true, will return details for each instance configuration referenced by the template.

	*/
	ShowInstanceConfigurations *bool
	/*StackVersion
	  If present, it will cause the returned deployment templates to be adapted to return only the elements allowed in that version.

	*/
	StackVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment templates params
func (o *GetDeploymentTemplatesParams) WithTimeout(timeout time.Duration) *GetDeploymentTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment templates params
func (o *GetDeploymentTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment templates params
func (o *GetDeploymentTemplatesParams) WithContext(ctx context.Context) *GetDeploymentTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment templates params
func (o *GetDeploymentTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment templates params
func (o *GetDeploymentTemplatesParams) WithHTTPClient(client *http.Client) *GetDeploymentTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment templates params
func (o *GetDeploymentTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFormat adds the format to the get deployment templates params
func (o *GetDeploymentTemplatesParams) WithFormat(format *string) *GetDeploymentTemplatesParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get deployment templates params
func (o *GetDeploymentTemplatesParams) SetFormat(format *string) {
	o.Format = format
}

// WithMetadata adds the metadata to the get deployment templates params
func (o *GetDeploymentTemplatesParams) WithMetadata(metadata *string) *GetDeploymentTemplatesParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the get deployment templates params
func (o *GetDeploymentTemplatesParams) SetMetadata(metadata *string) {
	o.Metadata = metadata
}

// WithShowHidden adds the showHidden to the get deployment templates params
func (o *GetDeploymentTemplatesParams) WithShowHidden(showHidden *bool) *GetDeploymentTemplatesParams {
	o.SetShowHidden(showHidden)
	return o
}

// SetShowHidden adds the showHidden to the get deployment templates params
func (o *GetDeploymentTemplatesParams) SetShowHidden(showHidden *bool) {
	o.ShowHidden = showHidden
}

// WithShowInstanceConfigurations adds the showInstanceConfigurations to the get deployment templates params
func (o *GetDeploymentTemplatesParams) WithShowInstanceConfigurations(showInstanceConfigurations *bool) *GetDeploymentTemplatesParams {
	o.SetShowInstanceConfigurations(showInstanceConfigurations)
	return o
}

// SetShowInstanceConfigurations adds the showInstanceConfigurations to the get deployment templates params
func (o *GetDeploymentTemplatesParams) SetShowInstanceConfigurations(showInstanceConfigurations *bool) {
	o.ShowInstanceConfigurations = showInstanceConfigurations
}

// WithStackVersion adds the stackVersion to the get deployment templates params
func (o *GetDeploymentTemplatesParams) WithStackVersion(stackVersion *string) *GetDeploymentTemplatesParams {
	o.SetStackVersion(stackVersion)
	return o
}

// SetStackVersion adds the stackVersion to the get deployment templates params
func (o *GetDeploymentTemplatesParams) SetStackVersion(stackVersion *string) {
	o.StackVersion = stackVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	if o.Metadata != nil {

		// query param metadata
		var qrMetadata string
		if o.Metadata != nil {
			qrMetadata = *o.Metadata
		}
		qMetadata := qrMetadata
		if qMetadata != "" {
			if err := r.SetQueryParam("metadata", qMetadata); err != nil {
				return err
			}
		}

	}

	if o.ShowHidden != nil {

		// query param show_hidden
		var qrShowHidden bool
		if o.ShowHidden != nil {
			qrShowHidden = *o.ShowHidden
		}
		qShowHidden := swag.FormatBool(qrShowHidden)
		if qShowHidden != "" {
			if err := r.SetQueryParam("show_hidden", qShowHidden); err != nil {
				return err
			}
		}

	}

	if o.ShowInstanceConfigurations != nil {

		// query param show_instance_configurations
		var qrShowInstanceConfigurations bool
		if o.ShowInstanceConfigurations != nil {
			qrShowInstanceConfigurations = *o.ShowInstanceConfigurations
		}
		qShowInstanceConfigurations := swag.FormatBool(qrShowInstanceConfigurations)
		if qShowInstanceConfigurations != "" {
			if err := r.SetQueryParam("show_instance_configurations", qShowInstanceConfigurations); err != nil {
				return err
			}
		}

	}

	if o.StackVersion != nil {

		// query param stack_version
		var qrStackVersion string
		if o.StackVersion != nil {
			qrStackVersion = *o.StackVersion
		}
		qStackVersion := qrStackVersion
		if qStackVersion != "" {
			if err := r.SetQueryParam("stack_version", qStackVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
