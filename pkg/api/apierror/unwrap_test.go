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

package apierror

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments_traffic_filter"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

type testErrorPayload struct {
	A string `json:"a,omitempty"`
}

type ValueError struct{}

func (e ValueError) Error() string { return "some error here" }

type PrivateResp struct {
	resp *http.Response
}

func (e PrivateResp) Error() string { return "some error here" }

func TestUnwrap(t *testing.T) {
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
			name: "Is able to parse a standard error",
			args: args{err: errors.New("new error")},
			want: errors.New("new error"),
		},
		{
			name: "Is able to parse a type that encapsulates another unknown type",
			args: args{err: &deployments.DeleteDeploymentStatelessResourceRetryWith{
				Payload: &models.BasicFailedReply{
					Errors: []*models.BasicFailedReplyElement{
						{
							Code:    ec.String("clusters.cluster_plan_state_error"),
							Message: ec.String("There are running instances"),
						},
					},
				},
			}},
			want: multierror.NewPrefixed(
				"api error", errors.New("clusters.cluster_plan_state_error: There are running instances"),
			),
		},
		{
			name: "Is able to parse a nil standard error",
			args: args{err: nil},
		},
		{
			name: "Can unpack a 449 error",
			args: args{err: &runtime.APIError{Code: 449}},
			want: ErrMissingElevatedPermissions,
		},
		{
			name: "Can unpack a nested http.Response error",
			args: args{
				err: &runtime.APIError{Response: privateResp},
			},
			want: errors.New(`{"somefield": "someerror"}`),
		},
		{
			name: "Throws unknown error when the Response inside the APIError can't be unpacked",
			args: args{err: &runtime.APIError{
				Code:          400,
				OperationName: "unknown error",
				Response:      struct{}{},
			}},
			want: errors.New("unknown error (status 400)"),
		},
		{
			name: "Unpacks the error when it can be",
			args: args{err: &runtime.APIError{
				Code:          400,
				OperationName: "unknown error",
				Response:      testErrorPayload{"b"},
			}},
			want: errors.New("{\n  \"a\": \"b\"\n}"),
		},
		{
			name: "Returns operation timed out when a context.DeadlineExceeded is received",
			args: args{err: context.DeadlineExceeded},
			want: errors.New(ErrTimedOutMsg),
		},
		{
			name: "Returns the error message of a non pointer type",
			args: args{err: ValueError{}},
			want: ValueError{},
		},
		{
			name: "Is able to parse a type that encapsulates a BasicFailedReply with fields",
			args: args{err: &deployments.DeleteDeploymentStatelessResourceRetryWith{
				Payload: &models.BasicFailedReply{
					Errors: []*models.BasicFailedReplyElement{
						{
							Code:    ec.String("clusters.cluster_plan_state_error"),
							Message: ec.String("There are running instances"),
						},
						{
							Code:    ec.String("auth.invalid_password"),
							Fields:  []string{"body.password"},
							Message: ec.String("request password doesn't match the user's password"),
						},
					},
				},
			}},
			want: multierror.NewPrefixed("api error",
				errors.New("clusters.cluster_plan_state_error: There are running instances"),
				errors.New("auth.invalid_password: request password doesn't match the user's password (body.password)"),
			),
		},
		{
			name: "Is able to parse a type that encapsulates a BasicFailedReply with fields",
			args: args{err: &deployments_traffic_filter.GetTrafficFilterRulesetNotFound{
				Payload: &models.BasicFailedReply{Errors: []*models.BasicFailedReplyElement{
					{
						Code:    ec.String("root.not_found"),
						Message: ec.String("traffic filter rule not found"),
					},
				}},
			}},
			want: multierror.NewPrefixed("api error",
				errors.New("root.not_found: traffic filter rule not found"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Unwrap(tt.args.err)

			var wantMsg string
			if tt.want != nil {
				wantMsg = tt.want.Error()
			}
			if err != nil && !assert.EqualError(t, err, wantMsg) {
				t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
