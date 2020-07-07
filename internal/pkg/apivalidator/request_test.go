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

package apivalidator

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/go-openapi/spec"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

var testSpec = spec.Swagger{
	SwaggerProps: spec.SwaggerProps{
		Paths: &spec.Paths{
			Paths: map[string]spec.PathItem{
				"/deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}": {
					PathItemProps: spec.PathItemProps{
						Put: &spec.Operation{
							OperationProps: spec.OperationProps{
								Description: "Puts a thing",
							},
						},
						Post: &spec.Operation{
							OperationProps: spec.OperationProps{
								Description: "Posts a thing",
							},
						},
					},
				},
				"/deployments/{deployment_id}/{resource_kind}/{ref_id}": {
					PathItemProps: spec.PathItemProps{
						Get: &spec.Operation{
							OperationProps: spec.OperationProps{
								Description: "Gets a thing",
							},
						},
						Delete: &spec.Operation{
							OperationProps: spec.OperationProps{
								Description: "Deletes a thing",
							},
						},
					},
				},
				"/users/auth/keys": {
					PathItemProps: spec.PathItemProps{
						Head: &spec.Operation{
							OperationProps: spec.OperationProps{
								Description: "Heads a thing",
							},
						},
						Patch: &spec.Operation{
							OperationProps: spec.OperationProps{
								Description: "Patches a thing",
							},
						},
					},
				},
			},
		},
	},
}

type invalidResponse struct{}

func TestNewHTTPRequests(t *testing.T) {
	type args struct {
		host      string
		client    *http.Client
		cloudSpec *spec.Swagger
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Fails due to empty client",
			args: args{
				host:      "http://0.0.0.0",
				cloudSpec: &testSpec,
			},
			wantErr: errors.New("requests to api: an http client must be specified"),
		},
		{
			// This test is expected to fail as we are not connecting to an api validation proxy.
			// We are checking to see if all requests have been built correctly.
			name: "Succeeds in creating all http requests, but fails due to connection error",
			args: args{
				host:      "http://0.0.0.0",
				client:    new(http.Client),
				cloudSpec: &testSpec,
			},
			wantErr: multierror.NewPrefixed("api spec validation",
				&url.Error{
					Op:  "Delete",
					URL: "http://0.0.0.0/deployments/%7Bdeployment_id%7D/kibana/%7Bref_id%7D",
					Err: errors.New("dial tcp 0.0.0.0:80: connect: connection refused"),
				},
				&url.Error{
					Op:  "Get",
					URL: "http://0.0.0.0/deployments/%7Bdeployment_id%7D/kibana/%7Bref_id%7D",
					Err: errors.New("dial tcp 0.0.0.0:80: connect: connection refused"),
				},
				&url.Error{
					Op:  "Head",
					URL: "http://0.0.0.0/users/auth/keys",
					Err: errors.New("dial tcp 0.0.0.0:80: connect: connection refused"),
				},
				&url.Error{
					Op:  "Patch",
					URL: "http://0.0.0.0/users/auth/keys",
					Err: errors.New("dial tcp 0.0.0.0:80: connect: connection refused"),
				},
				&url.Error{
					Op:  "Post",
					URL: "http://0.0.0.0/deployments/%7Bdeployment_id%7D/apm/%7Bref_id%7D",
					Err: errors.New("dial tcp 0.0.0.0:80: connect: connection refused"),
				},
				&url.Error{
					Op:  "Put",
					URL: "http://0.0.0.0/deployments/%7Bdeployment_id%7D/apm/%7Bref_id%7D",
					Err: errors.New("dial tcp 0.0.0.0:80: connect: connection refused"),
				},
			),
		},
		{
			name: "Fails due to json decoding errors",
			args: args{
				host: "http://0.0.0.0",
				client: mock.NewClient(
					mock.New200Response(mock.NewStructBody(invalidResponse{})),
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
				),
				cloudSpec: &testSpec,
			},
			wantErr: multierror.NewPrefixed("api spec validation",
				errors.New("json: cannot unmarshal string into Go value of type apivalidator.prismResponseBody"),
			),
		},
		{
			name: "Succeeds in creating all http requests and returns no error as no prism discrepancies were found",
			args: args{
				host: "http://0.0.0.0",
				client: mock.NewClient(
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
					mock.New200Response(mock.NewStringBody(`{}`)),
				),
				cloudSpec: &testSpec,
			},
		},
		{
			name: "Succeeds in returning an error when prism found an un-existing endpoint",
			args: args{
				host: "http://0.0.0.0",
				client: mock.NewClient(
					mock.New404Response(mock.NewStringBody(`{}`)),
					mock.New404Response(mock.NewStringBody(`{}`)),
					mock.New404Response(mock.NewStringBody(`{}`)),
					mock.New404Response(mock.NewStringBody(`{}`)),
					mock.New404Response(mock.NewStringBody(`{}`)),
					mock.New404Response(mock.NewStringBody(`{"type":"https://stoplight.io/prism/errors#NO_PATH_MATCHED_ERROR","title":"Route not resolved, no path matched","status":404,"detail":"The route /deploymentsdfgdsf hasn't been found in the specification file"}`)),
				),
				cloudSpec: &testSpec,
			},
			wantErr: multierror.NewPrefixed("api spec validation",
				errors.New("prism error: Type: https://stoplight.io/prism/errors#NO_PATH_MATCHED_ERROR, Title: Route not resolved, no path matched, Status: 404, Detail: The route /deploymentsdfgdsf hasn't been found in the specification file"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewHTTPRequests(tt.args.host, tt.args.client, tt.args.cloudSpec); err != nil && (err.Error() != tt.wantErr.Error()) {
				t.Errorf("NewHTTPRequests() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateRequest(t *testing.T) {
	type args struct {
		request *http.Request
		client  *http.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "fails due to invalid request",
			args: args{
				request: &http.Request{
					Method: "el chavo",
					URL:    &url.URL{},
				},
				client: new(http.Client),
			},
			wantErr: errors.New(`net/http: invalid method "el chavo"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateRequest(tt.args.request, tt.args.client); err.Error() != tt.wantErr.Error() {
				t.Errorf("validateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
