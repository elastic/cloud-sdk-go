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

// NewStartDeploymentResourceInstancesParams creates a new StartDeploymentResourceInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartDeploymentResourceInstancesParams() *StartDeploymentResourceInstancesParams {
	return &StartDeploymentResourceInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartDeploymentResourceInstancesParamsWithTimeout creates a new StartDeploymentResourceInstancesParams object
// with the ability to set a timeout on a request.
func NewStartDeploymentResourceInstancesParamsWithTimeout(timeout time.Duration) *StartDeploymentResourceInstancesParams {
	return &StartDeploymentResourceInstancesParams{
		timeout: timeout,
	}
}

// NewStartDeploymentResourceInstancesParamsWithContext creates a new StartDeploymentResourceInstancesParams object
// with the ability to set a context for a request.
func NewStartDeploymentResourceInstancesParamsWithContext(ctx context.Context) *StartDeploymentResourceInstancesParams {
	return &StartDeploymentResourceInstancesParams{
		Context: ctx,
	}
}

// NewStartDeploymentResourceInstancesParamsWithHTTPClient creates a new StartDeploymentResourceInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartDeploymentResourceInstancesParamsWithHTTPClient(client *http.Client) *StartDeploymentResourceInstancesParams {
	return &StartDeploymentResourceInstancesParams{
		HTTPClient: client,
	}
}

/*
StartDeploymentResourceInstancesParams contains all the parameters to send to the API endpoint

	for the start deployment resource instances operation.

	Typically these are written to a http.Request.
*/
type StartDeploymentResourceInstancesParams struct {

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* IgnoreMissing.

	   If true and the instance does not exist then quietly proceed to the next instance, otherwise treated as an error
	*/
	IgnoreMissing *bool

	/* InstanceIds.

	   A comma-separated list of instance identifiers.
	*/
	InstanceIds []string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	/* ResourceKind.

	   The kind of resource (one of elasticsearch, kibana, apm, or integrations_server).
	*/
	ResourceKind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start deployment resource instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartDeploymentResourceInstancesParams) WithDefaults() *StartDeploymentResourceInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start deployment resource instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartDeploymentResourceInstancesParams) SetDefaults() {
	var (
		ignoreMissingDefault = bool(false)
	)

	val := StartDeploymentResourceInstancesParams{
		IgnoreMissing: &ignoreMissingDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) WithTimeout(timeout time.Duration) *StartDeploymentResourceInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) WithContext(ctx context.Context) *StartDeploymentResourceInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) WithHTTPClient(client *http.Client) *StartDeploymentResourceInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) WithDeploymentID(deploymentID string) *StartDeploymentResourceInstancesParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithIgnoreMissing adds the ignoreMissing to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) WithIgnoreMissing(ignoreMissing *bool) *StartDeploymentResourceInstancesParams {
	o.SetIgnoreMissing(ignoreMissing)
	return o
}

// SetIgnoreMissing adds the ignoreMissing to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) SetIgnoreMissing(ignoreMissing *bool) {
	o.IgnoreMissing = ignoreMissing
}

// WithInstanceIds adds the instanceIds to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) WithInstanceIds(instanceIds []string) *StartDeploymentResourceInstancesParams {
	o.SetInstanceIds(instanceIds)
	return o
}

// SetInstanceIds adds the instanceIds to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) SetInstanceIds(instanceIds []string) {
	o.InstanceIds = instanceIds
}

// WithRefID adds the refID to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) WithRefID(refID string) *StartDeploymentResourceInstancesParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithResourceKind adds the resourceKind to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) WithResourceKind(resourceKind string) *StartDeploymentResourceInstancesParams {
	o.SetResourceKind(resourceKind)
	return o
}

// SetResourceKind adds the resourceKind to the start deployment resource instances params
func (o *StartDeploymentResourceInstancesParams) SetResourceKind(resourceKind string) {
	o.ResourceKind = resourceKind
}

// WriteToRequest writes these params to a swagger request
func (o *StartDeploymentResourceInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	if o.IgnoreMissing != nil {

		// query param ignore_missing
		var qrIgnoreMissing bool

		if o.IgnoreMissing != nil {
			qrIgnoreMissing = *o.IgnoreMissing
		}
		qIgnoreMissing := swag.FormatBool(qrIgnoreMissing)
		if qIgnoreMissing != "" {

			if err := r.SetQueryParam("ignore_missing", qIgnoreMissing); err != nil {
				return err
			}
		}
	}

	if o.InstanceIds != nil {

		// binding items for instance_ids
		joinedInstanceIds := o.bindParamInstanceIds(reg)

		// path array param instance_ids
		// SetPathParam does not support variadic arguments, since we used JoinByFormat
		// we can send the first item in the array as it's all the items of the previous
		// array joined together
		if len(joinedInstanceIds) > 0 {
			if err := r.SetPathParam("instance_ids", joinedInstanceIds[0]); err != nil {
				return err
			}
		}
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

// bindParamStartDeploymentResourceInstances binds the parameter instance_ids
func (o *StartDeploymentResourceInstancesParams) bindParamInstanceIds(formats strfmt.Registry) []string {
	instanceIdsIR := o.InstanceIds

	var instanceIdsIC []string
	for _, instanceIdsIIR := range instanceIdsIR { // explode []string

		instanceIdsIIV := instanceIdsIIR // string as string
		instanceIdsIC = append(instanceIdsIC, instanceIdsIIV)
	}

	// items.CollectionFormat: "csv"
	instanceIdsIS := swag.JoinByFormat(instanceIdsIC, "csv")

	return instanceIdsIS
}
