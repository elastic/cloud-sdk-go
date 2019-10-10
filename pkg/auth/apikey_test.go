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

package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestAPIKey_AuthRequest(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		k    APIKey
		args args
		want *http.Request
	}{
		{
			name: "auths the request",
			k:    APIKey("some"),
			args: args{req: &http.Request{
				Header: make(http.Header),
			}},
			want: &http.Request{
				Header: http.Header{
					"Authorization": []string{"ApiKey some"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.AuthRequest(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("APIKey.AuthRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPIKey(t *testing.T) {
	var dummyAPIKey = APIKey("dummy")
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want *APIKey
		err  error
	}{
		{
			name: "succeeds",
			args: args{key: "dummy"},
			want: &dummyAPIKey,
		},
		{
			name: "fails on empty",
			args: args{key: ""},
			err:  errors.New("auth: APIKey must not be empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAPIKey(tt.args.key)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewAPIKey() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
