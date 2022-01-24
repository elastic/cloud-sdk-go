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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// Response Wraps the response of the RoundTrip.
type Response struct {
	Response http.Response
	Error    error

	// If specified, it'll assert that the received request's fields match
	// the assertion.
	Assert *RequestAssertion
}

// NewStringBody creates an io.ReadCloser from a string.
func NewStringBody(b string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(b))
}

// NewByteBody creates an io.ReadCloser from a slice of bytes.
func NewByteBody(b []byte) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader(b))
}

// NewStructBody creates an io.ReadCloser from a structure that is attempted
// to be encoded into JSON. In case of failure, it panics.
func NewStructBody(i interface{}) io.ReadCloser {
	var b = new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(i); err != nil {
		panic(fmt.Sprintf("Failed to json.Encode structure %+v", i))
	}
	return ioutil.NopCloser(b)
}

// New200Response creates a new response with a statuscode 200
func New200Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 200,
		Body:       populateBody(body),
	}}
}

// New201Response creates a new response with a statuscode 201
func New201Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 201,
		Body:       populateBody(body),
	}}
}

// New202Response creates a new response with a statuscode 202
func New202Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 202,
		Body:       populateBody(body),
	}}
}

// New302Response creates a new response with a statuscode 302
func New302Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 302,
		Body:       populateBody(body),
	}}
}

// New400Response creates a new response with a statuscode 400
func New400Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 400,
		Body:       populateBody(body),
	}}
}

// New404Response creates a new response with a statuscode 404
func New404Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 404,
		Body:       populateBody(body),
	}}
}

// New409Response creates a new response with a statuscode 409
func New409Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 409,
		Body:       populateBody(body),
	}}
}

// New500Response creates a new response with a statuscode 500
func New500Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 500,
		Body:       populateBody(body),
	}}
}

// New501Response creates a new response with a statuscode 501
func New501Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 501,
		Body:       populateBody(body),
	}}
}

// New502Response creates a new response with a statuscode 502
func New502Response(body io.ReadCloser) Response {
	return Response{Response: http.Response{
		StatusCode: 502,
		Body:       populateBody(body),
	}}
}

// New200ResponseAssertion creates a new response with request assertion and a statuscode 200
func New200ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 200,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New201ResponseAssertion creates a new response with request assertion and a statuscode 201
func New201ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 201,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New202ResponseAssertion creates a new response with request assertion and a statuscode 202
func New202ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 202,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New302ResponseAssertion creates a new response with request assertion and a statuscode 302
func New302ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 302,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New400ResponseAssertion creates a new response with request assertion and a statuscode 400
func New400ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 400,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New404ResponseAssertion creates a new response with request assertion and a statuscode 404
func New404ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 404,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New409ResponseAssertion creates a new response with request assertion and a statuscode 409
func New409ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 409,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New500ResponseAssertion creates a new response with request assertion and a statuscode 500
func New500ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 500,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New501ResponseAssertion creates a new response with request assertion and a statuscode 501
func New501ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 501,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

// New502ResponseAssertion creates a new response with request assertion and a statuscode 502
func New502ResponseAssertion(assertion *RequestAssertion, body io.ReadCloser) Response {
	return Response{
		Response: http.Response{
			StatusCode: 502,
			Body:       populateBody(body),
		},
		Assert: assertion,
	}
}

func populateBody(body io.ReadCloser) io.ReadCloser {
	if body == nil {
		return NewStringBody("")
	}
	return body
}

// NewStructResponse takes in a  structure and creates a Response with the
// specified StatusCode.
func NewStructResponse(i interface{}, code int) Response {
	return Response{Response: http.Response{
		StatusCode: code,
		Body:       NewStructBody(i),
	}}
}

// New200StructResponse takes in a  structure and creates a 200 Response.
func New200StructResponse(i interface{}) Response {
	return NewStructResponse(i, 200)
}
