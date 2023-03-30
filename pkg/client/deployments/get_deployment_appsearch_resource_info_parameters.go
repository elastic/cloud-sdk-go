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
	"github.com/go-openapi/swag"
)

// NewGetDeploymentAppsearchResourceInfoParams creates a new GetDeploymentAppsearchResourceInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentAppsearchResourceInfoParams() *GetDeploymentAppsearchResourceInfoParams {
	return &GetDeploymentAppsearchResourceInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentAppsearchResourceInfoParamsWithTimeout creates a new GetDeploymentAppsearchResourceInfoParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentAppsearchResourceInfoParamsWithTimeout(timeout time.Duration) *GetDeploymentAppsearchResourceInfoParams {
	return &GetDeploymentAppsearchResourceInfoParams{
		timeout: timeout,
	}
}

// NewGetDeploymentAppsearchResourceInfoParamsWithContext creates a new GetDeploymentAppsearchResourceInfoParams object
// with the ability to set a context for a request.
func NewGetDeploymentAppsearchResourceInfoParamsWithContext(ctx context.Context) *GetDeploymentAppsearchResourceInfoParams {
	return &GetDeploymentAppsearchResourceInfoParams{
		Context: ctx,
	}
}

// NewGetDeploymentAppsearchResourceInfoParamsWithHTTPClient creates a new GetDeploymentAppsearchResourceInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentAppsearchResourceInfoParamsWithHTTPClient(client *http.Client) *GetDeploymentAppsearchResourceInfoParams {
	return &GetDeploymentAppsearchResourceInfoParams{
		HTTPClient: client,
	}
}

/* GetDeploymentAppsearchResourceInfoParams contains all the parameters to send to the API endpoint
   for the get deployment appsearch resource info operation.

   Typically these are written to a http.Request.
*/
type GetDeploymentAppsearchResourceInfoParams struct {

	/* DeploymentID.

	   Identifier for the Deployment
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	/* ShowMetadata.

	   Whether to include the full cluster metadata in the response - can be large per cluster and also include credentials.
	*/
	ShowMetadata *bool

	/* ShowPlanDefaults.

	   If showing plans, whether to show values that are left at their default value (less readable but more informative).
	*/
	ShowPlanDefaults *bool

	/* ShowPlanHistory.

	   Whether to include with the current and pending plan information the plan history- can be very large per cluster.
	*/
	ShowPlanHistory *bool

	/* ShowPlanLogs.

	   Whether to include with the current and pending plan information the attempt log - can be very large per cluster.
	*/
	ShowPlanLogs *bool

	/* ShowPlans.

	   Whether to include the full current and pending plan information in the response - can be large per cluster.

	   Default: true
	*/
	ShowPlans *bool

	/* ShowSettings.

	   Whether to show cluster settings in the response.
	*/
	ShowSettings *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployment appsearch resource info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentAppsearchResourceInfoParams) WithDefaults() *GetDeploymentAppsearchResourceInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment appsearch resource info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentAppsearchResourceInfoParams) SetDefaults() {
	var (
		showMetadataDefault = bool(false)

		showPlanDefaultsDefault = bool(false)

		showPlanHistoryDefault = bool(false)

		showPlanLogsDefault = bool(false)

		showPlansDefault = bool(true)

		showSettingsDefault = bool(false)
	)

	val := GetDeploymentAppsearchResourceInfoParams{
		ShowMetadata:     &showMetadataDefault,
		ShowPlanDefaults: &showPlanDefaultsDefault,
		ShowPlanHistory:  &showPlanHistoryDefault,
		ShowPlanLogs:     &showPlanLogsDefault,
		ShowPlans:        &showPlansDefault,
		ShowSettings:     &showSettingsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithTimeout(timeout time.Duration) *GetDeploymentAppsearchResourceInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithContext(ctx context.Context) *GetDeploymentAppsearchResourceInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithHTTPClient(client *http.Client) *GetDeploymentAppsearchResourceInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithDeploymentID(deploymentID string) *GetDeploymentAppsearchResourceInfoParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithRefID(refID string) *GetDeploymentAppsearchResourceInfoParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithShowMetadata adds the showMetadata to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithShowMetadata(showMetadata *bool) *GetDeploymentAppsearchResourceInfoParams {
	o.SetShowMetadata(showMetadata)
	return o
}

// SetShowMetadata adds the showMetadata to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetShowMetadata(showMetadata *bool) {
	o.ShowMetadata = showMetadata
}

// WithShowPlanDefaults adds the showPlanDefaults to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithShowPlanDefaults(showPlanDefaults *bool) *GetDeploymentAppsearchResourceInfoParams {
	o.SetShowPlanDefaults(showPlanDefaults)
	return o
}

// SetShowPlanDefaults adds the showPlanDefaults to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetShowPlanDefaults(showPlanDefaults *bool) {
	o.ShowPlanDefaults = showPlanDefaults
}

// WithShowPlanHistory adds the showPlanHistory to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithShowPlanHistory(showPlanHistory *bool) *GetDeploymentAppsearchResourceInfoParams {
	o.SetShowPlanHistory(showPlanHistory)
	return o
}

// SetShowPlanHistory adds the showPlanHistory to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetShowPlanHistory(showPlanHistory *bool) {
	o.ShowPlanHistory = showPlanHistory
}

// WithShowPlanLogs adds the showPlanLogs to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithShowPlanLogs(showPlanLogs *bool) *GetDeploymentAppsearchResourceInfoParams {
	o.SetShowPlanLogs(showPlanLogs)
	return o
}

// SetShowPlanLogs adds the showPlanLogs to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetShowPlanLogs(showPlanLogs *bool) {
	o.ShowPlanLogs = showPlanLogs
}

// WithShowPlans adds the showPlans to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithShowPlans(showPlans *bool) *GetDeploymentAppsearchResourceInfoParams {
	o.SetShowPlans(showPlans)
	return o
}

// SetShowPlans adds the showPlans to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetShowPlans(showPlans *bool) {
	o.ShowPlans = showPlans
}

// WithShowSettings adds the showSettings to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) WithShowSettings(showSettings *bool) *GetDeploymentAppsearchResourceInfoParams {
	o.SetShowSettings(showSettings)
	return o
}

// SetShowSettings adds the showSettings to the get deployment appsearch resource info params
func (o *GetDeploymentAppsearchResourceInfoParams) SetShowSettings(showSettings *bool) {
	o.ShowSettings = showSettings
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentAppsearchResourceInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ShowPlanHistory != nil {

		// query param show_plan_history
		var qrShowPlanHistory bool

		if o.ShowPlanHistory != nil {
			qrShowPlanHistory = *o.ShowPlanHistory
		}
		qShowPlanHistory := swag.FormatBool(qrShowPlanHistory)
		if qShowPlanHistory != "" {

			if err := r.SetQueryParam("show_plan_history", qShowPlanHistory); err != nil {
				return err
			}
		}
	}

	if o.ShowPlanLogs != nil {

		// query param show_plan_logs
		var qrShowPlanLogs bool

		if o.ShowPlanLogs != nil {
			qrShowPlanLogs = *o.ShowPlanLogs
		}
		qShowPlanLogs := swag.FormatBool(qrShowPlanLogs)
		if qShowPlanLogs != "" {

			if err := r.SetQueryParam("show_plan_logs", qShowPlanLogs); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
