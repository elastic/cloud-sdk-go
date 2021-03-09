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

package clusters_kibana

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

// NewGetKibanaClustersParams creates a new GetKibanaClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKibanaClustersParams() *GetKibanaClustersParams {
	return &GetKibanaClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKibanaClustersParamsWithTimeout creates a new GetKibanaClustersParams object
// with the ability to set a timeout on a request.
func NewGetKibanaClustersParamsWithTimeout(timeout time.Duration) *GetKibanaClustersParams {
	return &GetKibanaClustersParams{
		timeout: timeout,
	}
}

// NewGetKibanaClustersParamsWithContext creates a new GetKibanaClustersParams object
// with the ability to set a context for a request.
func NewGetKibanaClustersParamsWithContext(ctx context.Context) *GetKibanaClustersParams {
	return &GetKibanaClustersParams{
		Context: ctx,
	}
}

// NewGetKibanaClustersParamsWithHTTPClient creates a new GetKibanaClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKibanaClustersParamsWithHTTPClient(client *http.Client) *GetKibanaClustersParams {
	return &GetKibanaClustersParams{
		HTTPClient: client,
	}
}

/* GetKibanaClustersParams contains all the parameters to send to the API endpoint
   for the get kibana clusters operation.

   Typically these are written to a http.Request.
*/
type GetKibanaClustersParams struct {

	/* From.

	   The number of clusters to skip.
	*/
	From *int64

	/* Q.

	   An optional query to filter Kibana clusters by. Maps to an Elasticsearch query_string query.
	*/
	Q *string

	/* ShowHidden.

	   Includes the hidden clusters in the response.
	*/
	ShowHidden *bool

	/* ShowMetadata.

	   Includes all of the cluster metadata in the response. NOTE: Responses can include a large amount of metadata, as well as credentials.
	*/
	ShowMetadata *bool

	/* ShowPlanDefaults.

	   When plans are shown, includes the default values in the response. NOTE: This option results in large responses.
	*/
	ShowPlanDefaults *bool

	/* ShowPlans.

	   Includes the active and pending plan information in the response. NOTE: This option can result in large responses.
	*/
	ShowPlans *bool

	/* ShowSettings.

	   Includes the cluster settings in the response.
	*/
	ShowSettings *bool

	/* Size.

	   The maximum number of clusters to include in the response. For all clusters, use -1. NOTE: This option can result in large responses.

	   Default: 100
	*/
	Size *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get kibana clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKibanaClustersParams) WithDefaults() *GetKibanaClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get kibana clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKibanaClustersParams) SetDefaults() {
	var (
		fromDefault = int64(0)

		showHiddenDefault = bool(false)

		showMetadataDefault = bool(false)

		showPlanDefaultsDefault = bool(false)

		showSettingsDefault = bool(false)

		sizeDefault = int64(100)
	)

	val := GetKibanaClustersParams{
		From:             &fromDefault,
		ShowHidden:       &showHiddenDefault,
		ShowMetadata:     &showMetadataDefault,
		ShowPlanDefaults: &showPlanDefaultsDefault,
		ShowSettings:     &showSettingsDefault,
		Size:             &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get kibana clusters params
func (o *GetKibanaClustersParams) WithTimeout(timeout time.Duration) *GetKibanaClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kibana clusters params
func (o *GetKibanaClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kibana clusters params
func (o *GetKibanaClustersParams) WithContext(ctx context.Context) *GetKibanaClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kibana clusters params
func (o *GetKibanaClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kibana clusters params
func (o *GetKibanaClustersParams) WithHTTPClient(client *http.Client) *GetKibanaClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kibana clusters params
func (o *GetKibanaClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get kibana clusters params
func (o *GetKibanaClustersParams) WithFrom(from *int64) *GetKibanaClustersParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get kibana clusters params
func (o *GetKibanaClustersParams) SetFrom(from *int64) {
	o.From = from
}

// WithQ adds the q to the get kibana clusters params
func (o *GetKibanaClustersParams) WithQ(q *string) *GetKibanaClustersParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the get kibana clusters params
func (o *GetKibanaClustersParams) SetQ(q *string) {
	o.Q = q
}

// WithShowHidden adds the showHidden to the get kibana clusters params
func (o *GetKibanaClustersParams) WithShowHidden(showHidden *bool) *GetKibanaClustersParams {
	o.SetShowHidden(showHidden)
	return o
}

// SetShowHidden adds the showHidden to the get kibana clusters params
func (o *GetKibanaClustersParams) SetShowHidden(showHidden *bool) {
	o.ShowHidden = showHidden
}

// WithShowMetadata adds the showMetadata to the get kibana clusters params
func (o *GetKibanaClustersParams) WithShowMetadata(showMetadata *bool) *GetKibanaClustersParams {
	o.SetShowMetadata(showMetadata)
	return o
}

// SetShowMetadata adds the showMetadata to the get kibana clusters params
func (o *GetKibanaClustersParams) SetShowMetadata(showMetadata *bool) {
	o.ShowMetadata = showMetadata
}

// WithShowPlanDefaults adds the showPlanDefaults to the get kibana clusters params
func (o *GetKibanaClustersParams) WithShowPlanDefaults(showPlanDefaults *bool) *GetKibanaClustersParams {
	o.SetShowPlanDefaults(showPlanDefaults)
	return o
}

// SetShowPlanDefaults adds the showPlanDefaults to the get kibana clusters params
func (o *GetKibanaClustersParams) SetShowPlanDefaults(showPlanDefaults *bool) {
	o.ShowPlanDefaults = showPlanDefaults
}

// WithShowPlans adds the showPlans to the get kibana clusters params
func (o *GetKibanaClustersParams) WithShowPlans(showPlans *bool) *GetKibanaClustersParams {
	o.SetShowPlans(showPlans)
	return o
}

// SetShowPlans adds the showPlans to the get kibana clusters params
func (o *GetKibanaClustersParams) SetShowPlans(showPlans *bool) {
	o.ShowPlans = showPlans
}

// WithShowSettings adds the showSettings to the get kibana clusters params
func (o *GetKibanaClustersParams) WithShowSettings(showSettings *bool) *GetKibanaClustersParams {
	o.SetShowSettings(showSettings)
	return o
}

// SetShowSettings adds the showSettings to the get kibana clusters params
func (o *GetKibanaClustersParams) SetShowSettings(showSettings *bool) {
	o.ShowSettings = showSettings
}

// WithSize adds the size to the get kibana clusters params
func (o *GetKibanaClustersParams) WithSize(size *int64) *GetKibanaClustersParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get kibana clusters params
func (o *GetKibanaClustersParams) SetSize(size *int64) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetKibanaClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom int64

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
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

	if o.ShowMetadata != nil {

		// query param show_metadata
		var qrShowMetadata bool

		if o.ShowMetadata != nil {
			qrShowMetadata = *o.ShowMetadata
		}
		qShowMetadata := swag.FormatBool(qrShowMetadata)
		if qShowMetadata != "" {

			if err := r.SetQueryParam("show_metadata", qShowMetadata); err != nil {
				return err
			}
		}
	}

	if o.ShowPlanDefaults != nil {

		// query param show_plan_defaults
		var qrShowPlanDefaults bool

		if o.ShowPlanDefaults != nil {
			qrShowPlanDefaults = *o.ShowPlanDefaults
		}
		qShowPlanDefaults := swag.FormatBool(qrShowPlanDefaults)
		if qShowPlanDefaults != "" {

			if err := r.SetQueryParam("show_plan_defaults", qShowPlanDefaults); err != nil {
				return err
			}
		}
	}

	if o.ShowPlans != nil {

		// query param show_plans
		var qrShowPlans bool

		if o.ShowPlans != nil {
			qrShowPlans = *o.ShowPlans
		}
		qShowPlans := swag.FormatBool(qrShowPlans)
		if qShowPlans != "" {

			if err := r.SetQueryParam("show_plans", qShowPlans); err != nil {
				return err
			}
		}
	}

	if o.ShowSettings != nil {

		// query param show_settings
		var qrShowSettings bool

		if o.ShowSettings != nil {
			qrShowSettings = *o.ShowSettings
		}
		qShowSettings := swag.FormatBool(qrShowSettings)
		if qShowSettings != "" {

			if err := r.SetQueryParam("show_settings", qShowSettings); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int64

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
