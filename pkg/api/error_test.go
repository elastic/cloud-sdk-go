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

package api

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/go-openapi/runtime"

	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_elasticsearch"
	"github.com/elastic/cloud-sdk-go/pkg/models"
)

func newStringPointer(s string) *string { return &s }

var basicFailedReplyError = `
{
  "errors": [
    {
      "code": "clusters.cluster_plan_state_error",
      "fields": null,
      "message": "There are running instances"
    }
  ]
}`[1:]

var anotherError = `
{
  "a": "an error"
}`[1:]

type testError struct {
	Payload *testErrorPayload
}

func (e testError) Error() string {
	return e.Payload.A
}

type testErrorPayload struct {
	A string `json:"a,omitempty"`
}

type ValueError struct{}

func (e ValueError) Error() string { return "some error here" }

type PrivateResp struct {
	resp *http.Response
}

func (e PrivateResp) Error() string { return "some error here" }

func TestUnwrapError(t *testing.T) {
	var someEncapsulatedResp = &http.Response{
		StatusCode: 200,
		Body: ioutil.NopCloser(strings.NewReader(
			`{"somefield": "someerror"}`,
		)),
	}
	var privateResp = PrivateResp{resp: someEncapsulatedResp}
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"Is able to parse a standard error",
			args{
				err: errors.New("new error"),
			},
			errors.New("new error"),
		},
		{
			"Is able to parse a type that encapsulates BasicFailedReply in Payload property",
			args{
				err: &testError{
					Payload: &testErrorPayload{
						A: "an error",
					},
				},
			},
			errors.New(anotherError),
		},
		{
			"Is able to parse a type that encapsulates another unknown type",
			args{
				err: &clusters_elasticsearch.DeleteEsClusterRetryWith{
					Payload: &models.BasicFailedReply{
						Errors: []*models.BasicFailedReplyElement{
							{
								Code:    newStringPointer("clusters.cluster_plan_state_error"),
								Message: newStringPointer("There are running instances"),
							},
						},
					},
				},
			},
			errors.New(basicFailedReplyError),
		},
		{
			"Is able to parse a nil standard error",
			args{
				err: nil,
			},
			nil,
		},
		{
			name: "Can unpack a 449 error",
			args: args{
				err: &runtime.APIError{
					Code: 449,
				},
			},
			want: errors.New("the requested operation requires elevated permissions"),
		},
		{
			name: "Can unpack a nested http.Response error",
			args: args{
				err: &runtime.APIError{
					Response: privateResp,
				},
			},
			want: errors.New(`{"somefield": "someerror"}`),
		},
		{
			name: "Throws unknown error when the Response inside the APIError can't be unpacked",
			args: args{
				err: &runtime.APIError{
					Code:          400,
					OperationName: "unknown error",
					Response:      struct{}{},
				},
			},
			want: errors.New("unknown error (status 400)"),
		},
		{
			name: "Unpacks the error when it can be",
			args: args{
				err: &runtime.APIError{
					Code:          400,
					OperationName: "unknown error",
					Response:      testErrorPayload{"b"},
				},
			},
			want: errors.New(`{
  "a": "b"
}`),
		},
		{
			name: "Returns operation timed out when a context.DeadlineExceeded is received",
			args: args{err: context.DeadlineExceeded},
			want: errors.New("operation timed out"),
		},
		{
			name: "Returns the error message of a non pointer type",
			args: args{err: ValueError{}},
			want: ValueError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnwrapError(tt.args.err); !reflect.DeepEqual(err, tt.want) {
				t.Errorf("UnwrapError() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
