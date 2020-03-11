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
	"net/http"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
)

func TestUserAgentTransport_RoundTrip(t *testing.T) {
	var responseOK = mock.New200Response(mock.NewStringBody(`some`)).Response
	type fields struct {
		agent string
		rt    http.RoundTripper
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		want                *http.Response
		err                 error
		wantUserAgentHeader string
	}{
		{
			name: "UserAgent is set to the default value",
			fields: fields{rt: mock.NewRoundTripper(mock.New200Response(
				mock.NewStringBody(`some`),
			))},
			args:                args{req: &http.Request{Header: http.Header{}}},
			want:                &responseOK,
			wantUserAgentHeader: DefaultUserAgent,
		},
		{
			name: "UserAgent is set to a custom value",
			fields: fields{
				rt: mock.NewRoundTripper(mock.New200Response(
					mock.NewStringBody(`some`),
				)),
				agent: "someagent/1.0.0",
			},
			args:                args{req: &http.Request{Header: http.Header{}}},
			want:                &responseOK,
			wantUserAgentHeader: "someagent/1.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := &UserAgentTransport{
				agent: tt.fields.agent,
				rt:    tt.fields.rt,
			}
			got, err := ua.RoundTrip(tt.args.req)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("UserAgentTransport.RoundTrip() error = %v, wantErr %v", err, tt.err)
				return
			}
			defer got.Body.Close()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserAgentTransport.RoundTrip() = %v, want %v", got, tt.want)
			}
			actualHeader := tt.args.req.Header.Get(userAgentHeader)
			if actualHeader != tt.wantUserAgentHeader {
				t.Errorf("UserAgentTransport.RoundTrip() UserHeader = %v, want %v", actualHeader, tt.wantUserAgentHeader)
			}
		})
	}
}
