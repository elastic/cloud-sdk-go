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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewSetAllInstancesSettingsOverridesParams creates a new SetAllInstancesSettingsOverridesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetAllInstancesSettingsOverridesParams() *SetAllInstancesSettingsOverridesParams {
	return &SetAllInstancesSettingsOverridesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetAllInstancesSettingsOverridesParamsWithTimeout creates a new SetAllInstancesSettingsOverridesParams object
// with the ability to set a timeout on a request.
func NewSetAllInstancesSettingsOverridesParamsWithTimeout(timeout time.Duration) *SetAllInstancesSettingsOverridesParams {
	return &SetAllInstancesSettingsOverridesParams{
		timeout: timeout,
	}
}

// NewSetAllInstancesSettingsOverridesParamsWithContext creates a new SetAllInstancesSettingsOverridesParams object
// with the ability to set a context for a request.
func NewSetAllInstancesSettingsOverridesParamsWithContext(ctx context.Context) *SetAllInstancesSettingsOverridesParams {
	return &SetAllInstancesSettingsOverridesParams{
		Context: ctx,
	}
}

// NewSetAllInstancesSettingsOverridesParamsWithHTTPClient creates a new SetAllInstancesSettingsOverridesParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetAllInstancesSettingsOverridesParamsWithHTTPClient(client *http.Client) *SetAllInstancesSettingsOverridesParams {
	return &SetAllInstancesSettingsOverridesParams{
		HTTPClient: client,
	}
}

/* SetAllInstancesSettingsOverridesParams contains all the parameters to send to the API endpoint
   for the set all instances settings overrides operation.

   Typically these are written to a http.Request.
*/
type SetAllInstancesSettingsOverridesParams struct {

	/* Body.

	   The overrides to apply to all instances. Capacity overrides the RAM size in MB of the instance, and storage multipler overrides the multiplier of the instance RAM size that determines the storage quota of the instance. Capacity must be in the range [1024, 65536] and storage multiplier must be in the range [1.0, 1000.0].
	*/
	Body *models.InstanceOverrides

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	/* ResourceKind.

	   The kind of resource. Only Elasticsearch resources are currently supported.
	*/
	ResourceKind string

	/* RestartAfterUpdate.

	   After overrides are set, restarts the instance to apply the changes.
	*/
	RestartAfterUpdate *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set all instances settings overrides params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAllInstancesSettingsOverridesParams) WithDefaults() *SetAllInstancesSettingsOverridesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set all instances settings overrides params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAllInstancesSettingsOverridesParams) SetDefaults() {
	var (
		restartAfterUpdateDefault = bool(false)
	)

	val := SetAllInstancesSettingsOverridesParams{
		RestartAfterUpdate: &restartAfterUpdateDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) WithTimeout(timeout time.Duration) *SetAllInstancesSettingsOverridesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) WithContext(ctx context.Context) *SetAllInstancesSettingsOverridesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) WithHTTPClient(client *http.Client) *SetAllInstancesSettingsOverridesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) WithBody(body *models.InstanceOverrides) *SetAllInstancesSettingsOverridesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) SetBody(body *models.InstanceOverrides) {
	o.Body = body
}

// WithDeploymentID adds the deploymentID to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) WithDeploymentID(deploymentID string) *SetAllInstancesSettingsOverridesParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithRefID adds the refID to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) WithRefID(refID string) *SetAllInstancesSettingsOverridesParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithResourceKind adds the resourceKind to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) WithResourceKind(resourceKind string) *SetAllInstancesSettingsOverridesParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WithRestartAfterUpdate adds the restartAfterUpdate to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) WithRestartAfterUpdate(restartAfterUpdate *bool) *SetAllInstancesSettingsOverridesParams {
	o.SetRestartAfterUpdate(restartAfterUpdate)
	return o
}

// SetRestartAfterUpdate adds the restartAfterUpdate to the set all instances settings overrides params
func (o *SetAllInstancesSettingsOverridesParams) SetRestartAfterUpdate(restartAfterUpdate *bool) {
	o.RestartAfterUpdate = restartAfterUpdate
}

// WriteToRequest writes these params to a swagger request
func (o *SetAllInstancesSettingsOverridesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param resource_kind
	if err := r.SetPathParam("resource_kind", o.ResourceKind); err != nil {
		return err
	}

	if o.RestartAfterUpdate != nil {

		// query param restart_after_update
		var qrRestartAfterUpdate bool

		if o.RestartAfterUpdate != nil {
			qrRestartAfterUpdate = *o.RestartAfterUpdate
		}
		qRestartAfterUpdate := swag.FormatBool(qrRestartAfterUpdate)
		if qRestartAfterUpdate != "" {

			if err := r.SetQueryParam("restart_after_update", qRestartAfterUpdate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
