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
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestNew200Response(t *testing.T) {
	bodyBuffer := NewStringBody("200")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 200",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 200,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 200 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 200,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New200Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New200Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew201Response(t *testing.T) {
	bodyBuffer := NewStringBody("201")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 201",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 201,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 201 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 201,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New201Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New201Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew202Response(t *testing.T) {
	bodyBuffer := NewStringBody("202")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 202",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 202,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 202 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 202,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New202Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New202Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew404Response(t *testing.T) {
	bodyBuffer := NewStringBody("404")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 404",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 404,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 404 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 404,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New404Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New404Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew500Response(t *testing.T) {
	bodyBuffer := NewStringBody("500")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 500",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 500,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 500 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 500,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New500Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New500Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew200StructResponse(t *testing.T) {
	type S struct {
		Something string
	}
	structBody := NewStructBody(S{})
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "get a 200 response",
			args: args{i: S{}},
			want: Response{Response: http.Response{
				StatusCode: 200,
				Body:       structBody,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New200StructResponse(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New200StructResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStructResponse(t *testing.T) {
	type S struct {
		Something string
	}
	structBody := NewStructBody(S{})
	type args struct {
		i    interface{}
		code int
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "get a 200 response",
			args: args{i: S{}, code: 200},
			want: Response{Response: http.Response{
				StatusCode: 200,
				Body:       structBody,
			}},
		},
		{
			name: "get a 201 response",
			args: args{i: S{}, code: 201},
			want: Response{Response: http.Response{
				StatusCode: 201,
				Body:       structBody,
			}},
		},
		{
			name: "get a 202 response",
			args: args{i: S{}, code: 202},
			want: Response{Response: http.Response{
				StatusCode: 202,
				Body:       structBody,
			}},
		},
		{
			name: "get a 404 response",
			args: args{i: S{}, code: 404},
			want: Response{Response: http.Response{
				StatusCode: 404,
				Body:       structBody,
			}},
		},
		{
			name: "get a 500 response",
			args: args{i: S{}, code: 500},
			want: Response{Response: http.Response{
				StatusCode: 500,
				Body:       structBody,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStructResponse(tt.args.i, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStructResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
