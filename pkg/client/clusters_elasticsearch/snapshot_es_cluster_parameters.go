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

package clusters_elasticsearch

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

// NewSnapshotEsClusterParams creates a new SnapshotEsClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotEsClusterParams() *SnapshotEsClusterParams {
	return &SnapshotEsClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotEsClusterParamsWithTimeout creates a new SnapshotEsClusterParams object
// with the ability to set a timeout on a request.
func NewSnapshotEsClusterParamsWithTimeout(timeout time.Duration) *SnapshotEsClusterParams {
	return &SnapshotEsClusterParams{
		timeout: timeout,
	}
}

// NewSnapshotEsClusterParamsWithContext creates a new SnapshotEsClusterParams object
// with the ability to set a context for a request.
func NewSnapshotEsClusterParamsWithContext(ctx context.Context) *SnapshotEsClusterParams {
	return &SnapshotEsClusterParams{
		Context: ctx,
	}
}

// NewSnapshotEsClusterParamsWithHTTPClient creates a new SnapshotEsClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotEsClusterParamsWithHTTPClient(client *http.Client) *SnapshotEsClusterParams {
	return &SnapshotEsClusterParams{
		HTTPClient: client,
	}
}

/* SnapshotEsClusterParams contains all the parameters to send to the API endpoint
   for the snapshot es cluster operation.

   Typically these are written to a http.Request.
*/
type SnapshotEsClusterParams struct {

	/* Body.

	   Overrides default settings for the snapshot
	*/
	Body *models.ClusterSnapshotRequest

	/* ClusterID.

	   The Elasticsearch cluster identifier.
	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot es cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotEsClusterParams) WithDefaults() *SnapshotEsClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot es cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotEsClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapshot es cluster params
func (o *SnapshotEsClusterParams) WithTimeout(timeout time.Duration) *SnapshotEsClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot es cluster params
func (o *SnapshotEsClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot es cluster params
func (o *SnapshotEsClusterParams) WithContext(ctx context.Context) *SnapshotEsClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot es cluster params
func (o *SnapshotEsClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot es cluster params
func (o *SnapshotEsClusterParams) WithHTTPClient(client *http.Client) *SnapshotEsClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot es cluster params
func (o *SnapshotEsClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the snapshot es cluster params
func (o *SnapshotEsClusterParams) WithBody(body *models.ClusterSnapshotRequest) *SnapshotEsClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the snapshot es cluster params
func (o *SnapshotEsClusterParams) SetBody(body *models.ClusterSnapshotRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the snapshot es cluster params
func (o *SnapshotEsClusterParams) WithClusterID(clusterID string) *SnapshotEsClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the snapshot es cluster params
func (o *SnapshotEsClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotEsClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
