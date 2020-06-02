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

package mock

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestAssertRequest(t *testing.T) {
	type args struct {
		want *RequestAssertion
		req  *http.Request
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "expects a body but finds none",
			args: args{
				want: &RequestAssertion{
					Body: NewStringBody(`{"some field":1}`),
				},
				req: &http.Request{},
			},
			err: multierror.NewPrefixed("request assertion",
				errors.New(`got body , want {"some field":1}`),
			),
		},
		{
			name: "matches all fields",
			args: args{
				want: &RequestAssertion{
					Body: NewStringBody(`{"some field":1}`),
					Header: map[string][]string{
						"Authorization": {"Apikey Someapikey"},
						"Content-Type":  {"application/json"},
					},
					Method: "POST",
					Host:   "somehost",
					Path:   "/somepath/somesubpath",
				},
				req: &http.Request{
					Header: map[string][]string{
						"Authorization": {"Apikey Someapikey"},
						"Content-Type":  {"application/json"},
					},
					Body: NewStringBody(`{"some field":1}`),
					URL: &url.URL{
						Path: "/somepath/somesubpath",
						Host: "somehost",
					},
					Host:   "somehost",
					Method: "POST",
				},
			},
		},
		{
			name: "matches all fields and query",
			args: args{
				want: &RequestAssertion{
					Body: NewStringBody(`{"some field":1}`),
					Header: map[string][]string{
						"Authorization": {"Apikey Someapikey"},
						"Content-Type":  {"application/json"},
					},
					Method: "POST",
					Host:   "somehost",
					Path:   "/somepath/somesubpath",
					Query: url.Values{
						"cutoff":  []string{"14d"},
						"timeout": []string{"120s"},
					},
				},
				req: &http.Request{
					Header: map[string][]string{
						"Authorization": {"Apikey Someapikey"},
						"Content-Type":  {"application/json"},
					},
					Body: NewStringBody(`{"some field":1}`),
					URL: &url.URL{
						Path:     "/somepath/somesubpath",
						Host:     "somehost",
						RawQuery: "cutoff=14d&timeout=120s",
					},
					Host:   "somehost",
					Method: "POST",
				},
			},
		},
		{
			name: "matches no fields, returns an error",
			args: args{
				want: &RequestAssertion{
					Body: NewStringBody(`{"some field":2}`),
					Header: map[string][]string{
						"Content-Type": {"application/json"},
					},
					Method: "GET",
					Host:   "someotherhost",
					Path:   "/someotherpath/somesubpath",
				},
				req: &http.Request{
					Header: map[string][]string{
						"Authorization": {"Apikey Someapikey"},
						"Content-Type":  {"application/json"},
					},
					Body: NewStringBody(`{"some field":1}`),
					URL: &url.URL{
						Path:     "/somepath/somesubpath",
						Host:     "somehost",
						RawQuery: "cutoff=14d&timeout=120s",
					},
					Host:   "somehost",
					Method: "POST",
				},
			},
			err: multierror.NewPrefixed("request assertion",
				errors.New(`got body {"some field":1}, want {"some field":2}`),
				errors.New(`headers do not match: map[Content-Type:[application/json]] != map[Authorization:[Apikey Someapikey] Content-Type:[application/json]]`),
				errors.New(`methods do not match: GET != POST`),
				errors.New(`paths do not match: /someotherpath/somesubpath != /somepath/somesubpath`),
				errors.New("query does not match: map[] != map[cutoff:[14d] timeout:[120s]]"),
				errors.New(`hosts do not match: someotherhost != somehost`),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AssertRequest(tt.args.want, tt.args.req)
			var errString string
			if tt.err != nil {
				errString = tt.err.Error()
			}
			if err != nil || errString != "" {
				assert.EqualError(t, err, errString)
				return
			}

			assert.NoError(t, err)
		})
	}
}
