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
	"errors"
	"net/http"
	"net/http/httputil"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
)

// nolint
func TestErrCatchTransport_RoundTrip(t *testing.T) {
	var responseOK = mock.New200Response(mock.NewStringBody(`some`)).Response
	httputil.DumpResponse(&responseOK, responseOK.Body != nil)
	var responseNotFound = mock.New404Response(mock.NewStringBody(`notfound`)).Response
	httputil.DumpResponse(&responseNotFound, responseNotFound.Body != nil)
	type fields struct {
		rt http.RoundTripper
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *http.Response
		err    error
	}{
		{
			name: "returns a response",
			fields: fields{rt: mock.NewRoundTripper(mock.New200Response(
				mock.NewStringBody(`some`),
			))},
			args: args{req: &http.Request{}},
			want: &responseOK,
		},
		{
			name: "returns another response",
			fields: fields{rt: mock.NewRoundTripper(mock.New404Response(
				mock.NewStringBody(`notfound`),
			))},
			args: args{req: &http.Request{}},
			want: &responseNotFound,
		},
		{
			name: "returns an error and doesn't panic",
			fields: fields{rt: mock.NewRoundTripper(mock.Response{
				Error: errors.New("errored out"),
			})},
			args: args{req: &http.Request{}},
			err:  errors.New("errored out"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrCatchTransport{
				rt: tt.fields.rt,
			}
			got, err := e.RoundTrip(tt.args.req)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("ErrCatchTransport.RoundTrip() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrCatchTransport.RoundTrip() = %v, want %v", got, tt.want)
			}
		})
	}
}
