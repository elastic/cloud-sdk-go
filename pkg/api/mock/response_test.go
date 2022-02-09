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
	"net/url"
	"reflect"
	"testing"
)

var mockRequestAssertion = &RequestAssertion{
	Header: map[string][]string{"Accept": {"application/json"}},
	Method: "DELETE",
	Host:   "mock.elastic.co",
	Path:   "/api/v1",
	Query: url.Values{
		"some_value": []string{"false"},
	},
	Body: NewStringBody(`{}` + "\n"),
}

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

func TestNew302Response(t *testing.T) {
	bodyBuffer := NewStringBody("302")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 302",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 302,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 302 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 302,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New302Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New302Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew400Response(t *testing.T) {
	bodyBuffer := NewStringBody("400")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 400",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 400,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 400 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 400,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New400Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New400Response() = %v, want %v", got, tt.want)
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

func TestNew409Response(t *testing.T) {
	bodyBuffer := NewStringBody("409")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 409",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 409,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 409 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 409,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New409Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New409Response() = %v, want %v", got, tt.want)
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

func TestNew501Response(t *testing.T) {
	bodyBuffer := NewStringBody("501")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 501",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 501,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 501 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 501,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New501Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New501Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew502Response(t *testing.T) {
	bodyBuffer := NewStringBody("502")
	type args struct {
		body io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 502",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 502,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 502 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 502,
				Body:       bodyBuffer,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New502Response(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New502Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew200ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("200")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
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
		{
			name: "Returns statuscode 200 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 200,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New200ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New200Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew201ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("201")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
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
		{
			name: "Returns statuscode 201 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 201,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New201ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New201Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew202ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("202")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
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
		{
			name: "Returns statuscode 202 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 202,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New202ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New202Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew302ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("302")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 302",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 302,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 302 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 302,
				Body:       bodyBuffer,
			}},
		},
		{
			name: "Returns statuscode 302 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 302,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New302ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New302Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew400ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("400")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 400",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 400,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 400 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 400,
				Body:       bodyBuffer,
			}},
		},
		{
			name: "Returns statuscode 400 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 400,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New400ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New400Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew404ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("404")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
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
		{
			name: "Returns statuscode 404 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 404,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New404ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New404Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew409ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("409")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 409",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 409,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 409 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 409,
				Body:       bodyBuffer,
			}},
		},
		{
			name: "Returns statuscode 409 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 409,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New409ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New409Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew500ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("500")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
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
		{
			name: "Returns statuscode 500 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 500,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New500ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New500Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew501ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("501")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 501",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 501,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 501 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 501,
				Body:       bodyBuffer,
			}},
		},
		{
			name: "Returns statuscode 501 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 501,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New501ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New501Response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew502ResponseAssertion(t *testing.T) {
	bodyBuffer := NewStringBody("502")
	type args struct {
		assertion *RequestAssertion
		body      io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Returns statuscode 502",
			args: args{},
			want: Response{Response: http.Response{
				StatusCode: 502,
				Body:       NewStringBody(""),
			}},
		},
		{
			name: "Returns statuscode 502 with body",
			args: args{
				body: bodyBuffer,
			},
			want: Response{Response: http.Response{
				StatusCode: 502,
				Body:       bodyBuffer,
			}},
		},
		{
			name: "Returns statuscode 502 with body and request assertion",
			args: args{
				assertion: mockRequestAssertion,
				body:      bodyBuffer,
			},
			want: Response{
				Response: http.Response{
					StatusCode: 502,
					Body:       bodyBuffer,
				},
				Assert: mockRequestAssertion,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New502ResponseAssertion(tt.args.assertion, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New502Response() = %v, want %v", got, tt.want)
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
			name: "get a 409 response",
			args: args{i: S{}, code: 409},
			want: Response{Response: http.Response{
				StatusCode: 409,
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
