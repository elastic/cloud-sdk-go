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
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// RequestAssertion is used to assert the contents of the request passed to an
// http.RoundTripper.
type RequestAssertion struct {
	Body   io.ReadCloser
	Header http.Header
	Method string
	Path   string
	Host   string
	Query  url.Values
}

// AssertRequest ensures that a RequestAssertion matches certain *http.Request
// fields. If they do not match, an error is return.
func AssertRequest(want *RequestAssertion, req *http.Request) error {
	var merr = multierror.NewPrefixed("request assertion")
	if req.Body != nil || want.Body != nil {
		if !reflect.DeepEqual(want.Body, req.Body) {
			var wantB []byte
			if want.Body != nil {
				wantB, _ = io.ReadAll(
					io.TeeReader(want.Body, new(bytes.Buffer)),
				)
			}
			var gotB []byte
			if req.Body != nil {
				gotB, _ = io.ReadAll(
					io.TeeReader(req.Body, new(bytes.Buffer)),
				)
			}
			if !reflect.DeepEqual(wantB, gotB) {
				merr = merr.Append(
					fmt.Errorf("actual body %s, expected %s", gotB, wantB),
				)
			}
		}
	}

	if !reflect.DeepEqual(want.Header, req.Header) {
		merr = merr.Append(fmt.Errorf(
			"headers do not match (expected != actual): %v != %v", want.Header, req.Header),
		)
	}

	if !reflect.DeepEqual(want.Method, req.Method) {
		merr = merr.Append(fmt.Errorf(
			"methods do not match (expected != actual): %s != %s", want.Method, req.Method),
		)
	}

	if req.URL != nil {
		if !reflect.DeepEqual(want.Path, req.URL.Path) {
			merr = merr.Append(fmt.Errorf(
				"paths do not match (expected != actual): %s != %s", want.Path, req.URL.Path),
			)
		}
		if (len(req.URL.Query()) > 0 || len(want.Query) > 0) && !reflect.DeepEqual(want.Query, req.URL.Query()) {
			merr = merr.Append(fmt.Errorf(
				"query does not match (expected != actual): %s != %s", want.Query, req.URL.Query()),
			)
		}
	}

	if !reflect.DeepEqual(want.Host, req.Host) {
		merr = merr.Append(fmt.Errorf(
			"hosts do not match (expected != actual): %s != %s", want.Host, req.Host),
		)
	}

	return merr.ErrorOrNil()
}
