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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShutdownDeploymentStatelessResourceParams creates a new ShutdownDeploymentStatelessResourceParams object
// with the default values initialized.
func NewShutdownDeploymentStatelessResourceParams() *ShutdownDeploymentStatelessResourceParams {
	var (
		hideDefault         = bool(false)
		skipSnapshotDefault = bool(false)
	)
	return &ShutdownDeploymentStatelessResourceParams{
		Hide:         &hideDefault,
		SkipSnapshot: &skipSnapshotDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewShutdownDeploymentStatelessResourceParamsWithTimeout creates a new ShutdownDeploymentStatelessResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShutdownDeploymentStatelessResourceParamsWithTimeout(timeout time.Duration) *ShutdownDeploymentStatelessResourceParams {
	var (
		hideDefault         = bool(false)
		skipSnapshotDefault = bool(false)
	)
	return &ShutdownDeploymentStatelessResourceParams{
		Hide:         &hideDefault,
		SkipSnapshot: &skipSnapshotDefault,

		timeout: timeout,
	}
}

// NewShutdownDeploymentStatelessResourceParamsWithContext creates a new ShutdownDeploymentStatelessResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewShutdownDeploymentStatelessResourceParamsWithContext(ctx context.Context) *ShutdownDeploymentStatelessResourceParams {
	var (
		hideDefault         = bool(false)
		skipSnapshotDefault = bool(false)
	)
	return &ShutdownDeploymentStatelessResourceParams{
		Hide:         &hideDefault,
		SkipSnapshot: &skipSnapshotDefault,

		Context: ctx,
	}
}

// NewShutdownDeploymentStatelessResourceParamsWithHTTPClient creates a new ShutdownDeploymentStatelessResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShutdownDeploymentStatelessResourceParamsWithHTTPClient(client *http.Client) *ShutdownDeploymentStatelessResourceParams {
	var (
		hideDefault         = bool(false)
		skipSnapshotDefault = bool(false)
	)
	return &ShutdownDeploymentStatelessResourceParams{
		Hide:         &hideDefault,
		SkipSnapshot: &skipSnapshotDefault,
		HTTPClient:   client,
	}
}

/*ShutdownDeploymentStatelessResourceParams contains all the parameters to send to the API endpoint
for the shutdown deployment stateless resource operation typically these are written to a http.Request
*/
type ShutdownDeploymentStatelessResourceParams struct {

	/*DeploymentID
	  Identifier for the Deployment

	*/
	DeploymentID string
	/*Hide
	  Hide cluster on shutdown. Hidden clusters are not listed by default

	*/
	Hide *bool
	/*RefID
	  User-specified RefId for the Resource

	*/
	RefID string
	/*SkipSnapshot
	  If true, will skip taking a snapshot of the cluster before shutting the cluster down (if even possible)

	*/
	SkipSnapshot *bool
	/*StatelessResourceKind
	  The kind of stateless resource

	*/
	StatelessResourceKind string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) WithTimeout(timeout time.Duration) *ShutdownDeploymentStatelessResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) WithContext(ctx context.Context) *ShutdownDeploymentStatelessResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) WithHTTPClient(client *http.Client) *ShutdownDeploymentStatelessResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) WithDeploymentID(deploymentID string) *ShutdownDeploymentStatelessResourceParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithHide adds the hide to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) WithHide(hide *bool) *ShutdownDeploymentStatelessResourceParams {
	o.SetHide(hide)
	return o
}

// SetHide adds the hide to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) SetHide(hide *bool) {
	o.Hide = hide
}

// WithRefID adds the refID to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) WithRefID(refID string) *ShutdownDeploymentStatelessResourceParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithSkipSnapshot adds the skipSnapshot to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) WithSkipSnapshot(skipSnapshot *bool) *ShutdownDeploymentStatelessResourceParams {
	o.SetSkipSnapshot(skipSnapshot)
	return o
}

// SetSkipSnapshot adds the skipSnapshot to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) SetSkipSnapshot(skipSnapshot *bool) {
	o.SkipSnapshot = skipSnapshot
}

// WithStatelessResourceKind adds the statelessResourceKind to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) WithStatelessResourceKind(statelessResourceKind string) *ShutdownDeploymentStatelessResourceParams {
	o.SetStatelessResourceKind(statelessResourceKind)
	return o
}

// SetStatelessResourceKind adds the statelessResourceKind to the shutdown deployment stateless resource params
func (o *ShutdownDeploymentStatelessResourceParams) SetStatelessResourceKind(statelessResourceKind string) {
	o.StatelessResourceKind = statelessResourceKind
}

// WriteToRequest writes these params to a swagger request
func (o *ShutdownDeploymentStatelessResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	if o.Hide != nil {

		// query param hide
		var qrHide bool
		if o.Hide != nil {
			qrHide = *o.Hide
		}
		qHide := swag.FormatBool(qrHide)
		if qHide != "" {
			if err := r.SetQueryParam("hide", qHide); err != nil {
				return err
			}
		}

	}

	// path param ref_id
	if err := r.SetPathParam("ref_id", o.RefID); err != nil {
		return err
	}

	if o.SkipSnapshot != nil {

		// query param skip_snapshot
		var qrSkipSnapshot bool
		if o.SkipSnapshot != nil {
			qrSkipSnapshot = *o.SkipSnapshot
		}
		qSkipSnapshot := swag.FormatBool(qrSkipSnapshot)
		if qSkipSnapshot != "" {
			if err := r.SetQueryParam("skip_snapshot", qSkipSnapshot); err != nil {
				return err
			}
		}

	}

	// path param stateless_resource_kind
	if err := r.SetPathParam("stateless_resource_kind", o.StatelessResourceKind); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
