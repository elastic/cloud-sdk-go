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
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
)

// NewClient returns a pointer to http.Client with the mocked Transport.
func NewClient(r ...Response) *http.Client {
	return &http.Client{
		Transport: NewRoundTripper(r...),
	}
}

// NewRoundTripper initializes a new roundtripper and accepts multiple Response
// structures as variadric arguments.
func NewRoundTripper(r ...Response) *RoundTripper {
	return &RoundTripper{
		Responses: r,
	}
}

// RoundTripper is aimed to be used as the Transport property in an http.Client
// in order to mock the responses that it would return in the normal execution.
// If the number of responses that are mocked are not enough, an error with the
// request iteration ID, method and full URL is returned.
type RoundTripper struct {
	Responses []Response

	iteration int32
	mu        sync.RWMutex
}

// Add accepts multiple Response structures as variadric arguments and appends
// those to the current list of Responses.
func (rt *RoundTripper) Add(res ...Response) *RoundTripper {
	rt.mu.Lock()
	defer rt.mu.Unlock()

	rt.Responses = append(rt.Responses, res...)
	return rt
}

// RoundTrip executes a single HTTP transaction, returning
// a Response for the provided Request.
//
// RoundTrip should not attempt to interpret the response. In
// particular, RoundTrip must return err == nil if it obtained
// a response, regardless of the response's HTTP status code.
// A non-nil err should be reserved for failure to obtain a
// response. Similarly, RoundTrip should not attempt to
// handle higher-level protocol details such as redirects,
// authentication, or cookies.
//
// RoundTrip should not modify the request, except for
// consuming and closing the Request's Body. RoundTrip may
// read fields of the request in a separate goroutine. Callers
// should not mutate or reuse the request until the Response's
// Body has been closed.
//
// RoundTrip must always close the body, including on errors,
// but depending on the implementation may do so in a separate
// goroutine even after RoundTrip returns. This means that
// callers wanting to reuse the body for subsequent requests
// must arrange to wait for the Close call before doing so.
//
// The Request's URL and Header fields must be initialized.
func (rt *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.mu.RLock()
	defer func() {
		atomic.AddInt32(&rt.iteration, 1)
		rt.mu.RUnlock()
	}()

	var iteration = atomic.LoadInt32(&rt.iteration)
	if int(iteration) > len(rt.Responses)-1 {
		return nil, fmt.Errorf(
			"failed to obtain response in iteration %d: %s %s",
			iteration+1, req.Method, req.URL,
		)
	}

	// Consume and close the body.
	if req.Body != nil {
		//nolint
		ioutil.ReadAll(req.Body)
		defer req.Body.Close()
	}

	r := rt.Responses[iteration]
	if r.Error != nil {
		return nil, r.Error
	}

	return &r.Response, nil
}
