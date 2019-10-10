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
	"reflect"
	"testing"
)

func TestRoundTripper_RoundTrip(t *testing.T) {
	type fields struct {
		Responses []Response
		iteration int32
	}
	type args struct {
		req []*http.Request
	}
	type want struct {
		want *http.Response
		err  error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []want
	}{
		{
			name: "Single request mock with no body",
			fields: fields{Responses: []Response{
				{
					Response: http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
			}},
			args: args{req: []*http.Request{
				{},
			}},
			want: []want{
				{
					want: &http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
			},
		},
		{
			name: "Single request mock with body",
			fields: fields{Responses: []Response{
				{
					Response: http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
			}},
			args: args{req: []*http.Request{
				{Body: NewStringBody(`{"some":"body"}`)},
			}},
			want: []want{
				{
					want: &http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
			},
		},
		{
			name: "Single request mock with body returns an error",
			fields: fields{Responses: []Response{
				{
					Error: errors.New("some error"),
				},
			}},
			args: args{req: []*http.Request{
				{Body: NewStringBody(`{"some":"body"}`)},
			}},
			want: []want{
				{err: errors.New("some error")},
			},
		},
		{
			name: "Multiple request mock with no body",
			fields: fields{Responses: []Response{
				{
					Response: http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
				{
					Response: http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something something"),
					},
				},
			}},
			args: args{req: []*http.Request{
				{},
				{},
			}},
			want: []want{
				{
					want: &http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
				{
					want: &http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something something"),
					},
				},
			},
		},
		{
			name: "Multiple request mock with body",
			fields: fields{Responses: []Response{
				{
					Response: http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
				{
					Response: http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something something something"),
					},
				},
			}},
			args: args{req: []*http.Request{
				{Body: NewStringBody(`{"some":"body"}`)},
				{Body: NewStringBody(`{"some":"other body"}`)},
			}},
			want: []want{
				{
					want: &http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
				{
					want: &http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something something something"),
					},
				},
			},
		},
		{
			name: "Multiple request mock with body returns an error",
			fields: fields{Responses: []Response{
				{
					Response: http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
				{
					Error: errors.New("some error"),
				},
			}},
			args: args{req: []*http.Request{
				{Body: NewStringBody(`{"some":"body"}`)},
				{Body: NewStringBody(`{"some":"some body"}`)},
			}},
			want: []want{
				{
					want: &http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
				{err: errors.New("some error")},
			},
		},
		{
			name: "Multiple request mock when there's only one response return an error",
			fields: fields{Responses: []Response{
				{
					Response: http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
			}},
			args: args{req: []*http.Request{
				{Body: NewStringBody(`{"some":"body"}`)},
				{
					Body:   NewStringBody(`{"some":"some body"}`),
					Method: "POST",
					URL: &url.URL{
						Scheme: "https",
						Host:   "localhost",
						Path:   "some/path",
					},
				},
			}},
			want: []want{
				{
					want: &http.Response{
						Status:     http.StatusText(http.StatusOK),
						StatusCode: http.StatusOK,
						Body:       NewStringBody("something"),
					},
				},
				{err: errors.New("failed to obtain response in iteration 2: POST https://localhost/some/path")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := &RoundTripper{
				Responses: tt.fields.Responses,
				iteration: tt.fields.iteration,
			}
			var gotRes []want
			for _, req := range tt.args.req {
				got, err := rt.RoundTrip(req)
				if got != nil {
					defer got.Body.Close()
				}
				gotRes = append(gotRes, want{
					want: got,
					err:  err,
				})
			}
			if !reflect.DeepEqual(gotRes, tt.want) {
				t.Errorf("RoundTripper.RoundTrip() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}

func TestRoundTripper_RoundTrip_Concurrent(t *testing.T) {
	// This test ensures that no contention or data races happen on
	// different goroutines (As long as the tests are run with -race).
	rt := new(RoundTripper)
	var defaultResponse = http.Response{
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       NewStringBody("something"),
	}

	go func() {
		rt.Add(Response{Response: defaultResponse})
		rt.Add(Response{Response: defaultResponse})
		resp, _ := rt.RoundTrip(&http.Request{})
		defer resp.Body.Close()
		resp2, _ := rt.RoundTrip(&http.Request{})
		defer resp2.Body.Close()
	}()

	rt.Add(Response{Response: defaultResponse})
	rt.Add(Response{Response: defaultResponse})
	resp, _ := rt.RoundTrip(&http.Request{})
	defer resp.Body.Close()
	resp2, _ := rt.RoundTrip(&http.Request{})
	defer resp2.Body.Close()
}
