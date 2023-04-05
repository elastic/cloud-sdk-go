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

// NewRestartDeploymentEsResourceParams creates a new RestartDeploymentEsResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestartDeploymentEsResourceParams() *RestartDeploymentEsResourceParams {
	return &RestartDeploymentEsResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestartDeploymentEsResourceParamsWithTimeout creates a new RestartDeploymentEsResourceParams object
// with the ability to set a timeout on a request.
func NewRestartDeploymentEsResourceParamsWithTimeout(timeout time.Duration) *RestartDeploymentEsResourceParams {
	return &RestartDeploymentEsResourceParams{
		timeout: timeout,
	}
}

// NewRestartDeploymentEsResourceParamsWithContext creates a new RestartDeploymentEsResourceParams object
// with the ability to set a context for a request.
func NewRestartDeploymentEsResourceParamsWithContext(ctx context.Context) *RestartDeploymentEsResourceParams {
	return &RestartDeploymentEsResourceParams{
		Context: ctx,
	}
}

// NewRestartDeploymentEsResourceParamsWithHTTPClient creates a new RestartDeploymentEsResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestartDeploymentEsResourceParamsWithHTTPClient(client *http.Client) *RestartDeploymentEsResourceParams {
	return &RestartDeploymentEsResourceParams{
		HTTPClient: client,
	}
}

/*
RestartDeploymentEsResourceParams contains all the parameters to send to the API endpoint

	for the restart deployment es resource operation.

	Typically these are written to a http.Request.
*/
type RestartDeploymentEsResourceParams struct {

	/* CancelPending.

	   If true, cancels any pending plans before restarting. If false and there are pending plans, returns an error.
	*/
	CancelPending *bool

	/* DeploymentID.

	   Identifier for the Deployment.
	*/
	DeploymentID string

	/* GroupAttribute.

	   Indicates the property or properties used to divide the list of instances to restart in groups. Valid options are: '\_\_all\_\_' (restart all at once), '\_\_zone\_\_' by logical zone, '\_\_name\_\_' one instance at a time, or a comma-separated list of attributes of the instances

	   Default: "__zone__"
	*/
	GroupAttribute *string

	/* RefID.

	   User-specified RefId for the Resource (or '_main' if there is only one).
	*/
	RefID string

	/* RestoreSnapshot.

	   When set to true and restoring from shutdown, then will restore the cluster from the last snapshot (if available).

	   Default: true
	*/
	RestoreSnapshot *bool

	/* ShardInitWaitTime.

	   The time, in seconds, to wait for shards that show no progress of initializing, before rolling the next group (default: 10 minutes)

	   Default: 600
	*/
	ShardInitWaitTime *int64

	/* SkipSnapshot.

	   If true, will not take a snapshot of the cluster before restarting.

	   Default: true
	*/
	SkipSnapshot *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restart deployment es resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartDeploymentEsResourceParams) WithDefaults() *RestartDeploymentEsResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restart deployment es resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartDeploymentEsResourceParams) SetDefaults() {
	var (
		cancelPendingDefault = bool(false)

		groupAttributeDefault = string("__zone__")

		restoreSnapshotDefault = bool(true)

		shardInitWaitTimeDefault = int64(600)

		skipSnapshotDefault = bool(true)
	)

	val := RestartDeploymentEsResourceParams{
		CancelPending:     &cancelPendingDefault,
		GroupAttribute:    &groupAttributeDefault,
		RestoreSnapshot:   &restoreSnapshotDefault,
		ShardInitWaitTime: &shardInitWaitTimeDefault,
		SkipSnapshot:      &skipSnapshotDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithTimeout(timeout time.Duration) *RestartDeploymentEsResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithContext(ctx context.Context) *RestartDeploymentEsResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithHTTPClient(client *http.Client) *RestartDeploymentEsResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCancelPending adds the cancelPending to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithCancelPending(cancelPending *bool) *RestartDeploymentEsResourceParams {
	o.SetCancelPending(cancelPending)
	return o
}

// SetCancelPending adds the cancelPending to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetCancelPending(cancelPending *bool) {
	o.CancelPending = cancelPending
}

// WithDeploymentID adds the deploymentID to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithDeploymentID(deploymentID string) *RestartDeploymentEsResourceParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetDeploymentID(deploymentID string) {
	o.DeploymentID = deploymentID
}

// WithGroupAttribute adds the groupAttribute to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithGroupAttribute(groupAttribute *string) *RestartDeploymentEsResourceParams {
	o.SetGroupAttribute(groupAttribute)
	return o
}

// SetGroupAttribute adds the groupAttribute to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetGroupAttribute(groupAttribute *string) {
	o.GroupAttribute = groupAttribute
}

// WithRefID adds the refID to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithRefID(refID string) *RestartDeploymentEsResourceParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithRestoreSnapshot adds the restoreSnapshot to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithRestoreSnapshot(restoreSnapshot *bool) *RestartDeploymentEsResourceParams {
	o.SetRestoreSnapshot(restoreSnapshot)
	return o
}

// SetRestoreSnapshot adds the restoreSnapshot to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetRestoreSnapshot(restoreSnapshot *bool) {
	o.RestoreSnapshot = restoreSnapshot
}

// WithShardInitWaitTime adds the shardInitWaitTime to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithShardInitWaitTime(shardInitWaitTime *int64) *RestartDeploymentEsResourceParams {
	o.SetShardInitWaitTime(shardInitWaitTime)
	return o
}

// SetShardInitWaitTime adds the shardInitWaitTime to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetShardInitWaitTime(shardInitWaitTime *int64) {
	o.ShardInitWaitTime = shardInitWaitTime
}

// WithSkipSnapshot adds the skipSnapshot to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) WithSkipSnapshot(skipSnapshot *bool) *RestartDeploymentEsResourceParams {
	o.SetSkipSnapshot(skipSnapshot)
	return o
}

// SetSkipSnapshot adds the skipSnapshot to the restart deployment es resource params
func (o *RestartDeploymentEsResourceParams) SetSkipSnapshot(skipSnapshot *bool) {
	o.SkipSnapshot = skipSnapshot
}

// WriteToRequest writes these params to a swagger request
func (o *RestartDeploymentEsResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CancelPending != nil {

		// query param cancel_pending
		var qrCancelPending bool

		if o.CancelPending != nil {
			qrCancelPending = *o.CancelPending
		}
		qCancelPending := swag.FormatBool(qrCancelPending)
		if qCancelPending != "" {

			if err := r.SetQueryParam("cancel_pending", qCancelPending); err != nil {
				return err
			}
		}
	}

	// path param deployment_id
	if err := r.SetPathParam("deployment_id", o.DeploymentID); err != nil {
		return err
	}

	if o.GroupAttribute != nil {

		// query param group_attribute
		var qrGroupAttribute string

		if o.GroupAttribute != nil {
			qrGroupAttribute = *o.GroupAttribute
		}
		qGroupAttribute := qrGroupAttribute
		if qGroupAttribute != "" {

			if err := r.SetQueryParam("group_attribute", qGroupAttribute); err != nil {
				return err
			}
		}
	}

	// path param ref_id
	if err := r.SetPathParam("ref_id", o.RefID); err != nil {
		return err
	}

	if o.RestoreSnapshot != nil {

		// query param restore_snapshot
		var qrRestoreSnapshot bool

		if o.RestoreSnapshot != nil {
			qrRestoreSnapshot = *o.RestoreSnapshot
		}
		qRestoreSnapshot := swag.FormatBool(qrRestoreSnapshot)
		if qRestoreSnapshot != "" {

			if err := r.SetQueryParam("restore_snapshot", qRestoreSnapshot); err != nil {
				return err
			}
		}
	}

	if o.ShardInitWaitTime != nil {

		// query param shard_init_wait_time
		var qrShardInitWaitTime int64

		if o.ShardInitWaitTime != nil {
			qrShardInitWaitTime = *o.ShardInitWaitTime
		}
		qShardInitWaitTime := swag.FormatInt64(qrShardInitWaitTime)
		if qShardInitWaitTime != "" {

			if err := r.SetQueryParam("shard_init_wait_time", qShardInitWaitTime); err != nil {
				return err
			}
		}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
