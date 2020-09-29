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
)

// UserAgentTransport Deprecated wraps an http.RoundTripper and adds an User
// -Agent header to all requests  which are processed through the structure.
type UserAgentTransport struct {
	agent string
	rt    http.RoundTripper
}

// NewUserAgentTransport Deprecated initializes a new UserAgentTransport
func NewUserAgentTransport(rt http.RoundTripper, agent string) *UserAgentTransport {
	if agent == "" {
		agent = DefaultUserAgent
	}

	if rt == nil {
		rt = newDefaultTransport(0)
	}

	return &UserAgentTransport{
		agent: agent,
		rt:    rt,
	}
}

// RoundTrip Deprecated wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (ua *UserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if ua.rt == nil {
		ua.rt = newDefaultTransport(0)
	}
	if ua.agent == "" {
		ua.agent = DefaultUserAgent
	}

	req.Header.Set(userAgentHeader, ua.agent)

	return ua.rt.RoundTrip(req)
}
