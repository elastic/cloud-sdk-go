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

package platform_configuration_instances

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
	"github.com/go-openapi/swag"
)

// NewGetInstanceConfigurationParams creates a new GetInstanceConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstanceConfigurationParams() *GetInstanceConfigurationParams {
	return &GetInstanceConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstanceConfigurationParamsWithTimeout creates a new GetInstanceConfigurationParams object
// with the ability to set a timeout on a request.
func NewGetInstanceConfigurationParamsWithTimeout(timeout time.Duration) *GetInstanceConfigurationParams {
	return &GetInstanceConfigurationParams{
		timeout: timeout,
	}
}

// NewGetInstanceConfigurationParamsWithContext creates a new GetInstanceConfigurationParams object
// with the ability to set a context for a request.
func NewGetInstanceConfigurationParamsWithContext(ctx context.Context) *GetInstanceConfigurationParams {
	return &GetInstanceConfigurationParams{
		Context: ctx,
	}
}

// NewGetInstanceConfigurationParamsWithHTTPClient creates a new GetInstanceConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstanceConfigurationParamsWithHTTPClient(client *http.Client) *GetInstanceConfigurationParams {
	return &GetInstanceConfigurationParams{
		HTTPClient: client,
	}
}

/*
GetInstanceConfigurationParams contains all the parameters to send to the API endpoint

	for the get instance configuration operation.

	Typically these are written to a http.Request.
*/
type GetInstanceConfigurationParams struct {

	/* ConfigVersion.

	   Optionally retrieve the specified config version of the IC (otherwise retrieves the latest/only version)
	*/
	ConfigVersion *int64

	/* ID.

	   ID of the instance configuration
	*/
	ID string

	/* ShowDeleted.

	   If true, if the instance configuration has been marked for deletion it is still returned. Otherwise, instance configurations marked for deletion generate a 404
	*/
	ShowDeleted *bool

	/* ShowMaxZones.

	   If true, will populate the max_zones field of the instance configuration.
	*/
	ShowMaxZones *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get instance configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstanceConfigurationParams) WithDefaults() *GetInstanceConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get instance configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstanceConfigurationParams) SetDefaults() {
	var (
		showMaxZonesDefault = bool(false)
	)

	val := GetInstanceConfigurationParams{
		ShowMaxZones: &showMaxZonesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithTimeout(timeout time.Duration) *GetInstanceConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithContext(ctx context.Context) *GetInstanceConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithHTTPClient(client *http.Client) *GetInstanceConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigVersion adds the configVersion to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithConfigVersion(configVersion *int64) *GetInstanceConfigurationParams {
	o.SetConfigVersion(configVersion)
	return o
}

// SetConfigVersion adds the configVersion to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetConfigVersion(configVersion *int64) {
	o.ConfigVersion = configVersion
}

// WithID adds the id to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithID(id string) *GetInstanceConfigurationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetID(id string) {
	o.ID = id
}

// WithShowDeleted adds the showDeleted to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithShowDeleted(showDeleted *bool) *GetInstanceConfigurationParams {
	o.SetShowDeleted(showDeleted)
	return o
}

// SetShowDeleted adds the showDeleted to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetShowDeleted(showDeleted *bool) {
	o.ShowDeleted = showDeleted
}

// WithShowMaxZones adds the showMaxZones to the get instance configuration params
func (o *GetInstanceConfigurationParams) WithShowMaxZones(showMaxZones *bool) *GetInstanceConfigurationParams {
	o.SetShowMaxZones(showMaxZones)
	return o
}

// SetShowMaxZones adds the showMaxZones to the get instance configuration params
func (o *GetInstanceConfigurationParams) SetShowMaxZones(showMaxZones *bool) {
	o.ShowMaxZones = showMaxZones
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstanceConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConfigVersion != nil {

		// query param config_version
		var qrConfigVersion int64

		if o.ConfigVersion != nil {
			qrConfigVersion = *o.ConfigVersion
		}
		qConfigVersion := swag.FormatInt64(qrConfigVersion)
		if qConfigVersion != "" {

			if err := r.SetQueryParam("config_version", qConfigVersion); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.ShowDeleted != nil {

		// query param show_deleted
		var qrShowDeleted bool

		if o.ShowDeleted != nil {
			qrShowDeleted = *o.ShowDeleted
		}
		qShowDeleted := swag.FormatBool(qrShowDeleted)
		if qShowDeleted != "" {

			if err := r.SetQueryParam("show_deleted", qShowDeleted); err != nil {
				return err
			}
		}
	}

	if o.ShowMaxZones != nil {

		// query param show_max_zones
		var qrShowMaxZones bool

		if o.ShowMaxZones != nil {
			qrShowMaxZones = *o.ShowMaxZones
		}
		qShowMaxZones := swag.FormatBool(qrShowMaxZones)
		if qShowMaxZones != "" {

			if err := r.SetQueryParam("show_max_zones", qShowMaxZones); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
