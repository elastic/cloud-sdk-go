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
)

// NewDeleteEsProxyRequestsParams creates a new DeleteEsProxyRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEsProxyRequestsParams() *DeleteEsProxyRequestsParams {
	return &DeleteEsProxyRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEsProxyRequestsParamsWithTimeout creates a new DeleteEsProxyRequestsParams object
// with the ability to set a timeout on a request.
func NewDeleteEsProxyRequestsParamsWithTimeout(timeout time.Duration) *DeleteEsProxyRequestsParams {
	return &DeleteEsProxyRequestsParams{
		timeout: timeout,
	}
}

// NewDeleteEsProxyRequestsParamsWithContext creates a new DeleteEsProxyRequestsParams object
// with the ability to set a context for a request.
func NewDeleteEsProxyRequestsParamsWithContext(ctx context.Context) *DeleteEsProxyRequestsParams {
	return &DeleteEsProxyRequestsParams{
		Context: ctx,
	}
}

// NewDeleteEsProxyRequestsParamsWithHTTPClient creates a new DeleteEsProxyRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEsProxyRequestsParamsWithHTTPClient(client *http.Client) *DeleteEsProxyRequestsParams {
	return &DeleteEsProxyRequestsParams{
		HTTPClient: client,
	}
}

/* DeleteEsProxyRequestsParams contains all the parameters to send to the API endpoint
   for the delete es proxy requests operation.

   Typically these are written to a http.Request.
*/
type DeleteEsProxyRequestsParams struct {

	/* XManagementRequest.

	   X-Management-Request header value. Needs to be set to true
	*/
	XManagementRequest string

	/* Body.

	   The JSON payload to proxy to the Elasticsearch cluster
	*/
	Body string

	/* ClusterID.

	   Identifier for the Elasticsearch cluster
	*/
	ClusterID string

	/* ElasticsearchPath.

	   The URL part to proxy to the Elasticsearch cluster. Example: _search or _cat/indices?v&h=i,tm&s=tm:desc
	*/
	ElasticsearchPath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete es proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEsProxyRequestsParams) WithDefaults() *DeleteEsProxyRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete es proxy requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEsProxyRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) WithTimeout(timeout time.Duration) *DeleteEsProxyRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) WithContext(ctx context.Context) *DeleteEsProxyRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) WithHTTPClient(client *http.Client) *DeleteEsProxyRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXManagementRequest adds the xManagementRequest to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) WithXManagementRequest(xManagementRequest string) *DeleteEsProxyRequestsParams {
	o.SetXManagementRequest(xManagementRequest)
	return o
}

// SetXManagementRequest adds the xManagementRequest to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) SetXManagementRequest(xManagementRequest string) {
	o.XManagementRequest = xManagementRequest
}

// WithBody adds the body to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) WithBody(body string) *DeleteEsProxyRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) SetBody(body string) {
	o.Body = body
}

// WithClusterID adds the clusterID to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) WithClusterID(clusterID string) *DeleteEsProxyRequestsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithElasticsearchPath adds the elasticsearchPath to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) WithElasticsearchPath(elasticsearchPath string) *DeleteEsProxyRequestsParams {
	o.SetElasticsearchPath(elasticsearchPath)
	return o
}

// SetElasticsearchPath adds the elasticsearchPath to the delete es proxy requests params
func (o *DeleteEsProxyRequestsParams) SetElasticsearchPath(elasticsearchPath string) {
	o.ElasticsearchPath = elasticsearchPath
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEsProxyRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Management-Request
	if err := r.SetHeaderParam("X-Management-Request", o.XManagementRequest); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param elasticsearch_path
	if err := r.SetPathParam("elasticsearch_path", o.ElasticsearchPath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
