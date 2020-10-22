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
	"time"

	runtimeclient "github.com/go-openapi/runtime/client"

	"github.com/elastic/cloud-sdk-go/pkg/auth"
	"github.com/elastic/cloud-sdk-go/pkg/client"
)

// API contains all of the API clients and authentication objects necessary
// for the EC API
type API struct {
	V1API      *client.Rest
	AuthWriter auth.Writer
}

// AuthWriter wraps the runtime.ClientAuthInfoWriter interface adding a method
// to Auth generic http.Request.
// This type alias is used to maintain API compatibility.
type AuthWriter auth.Writer

// NewAPI initializes the API clients from an API config that it receives
func NewAPI(c Config) (*API, error) {
	c.fillDefaults()
	if err := c.Validate(); err != nil {
		return nil, err
	}

	t, err := NewTransport(c.Client.Transport, TransportConfig{
		SkipTLSVerify:   c.SkipTLSVerify,
		ErrorDevice:     c.ErrorDevice,
		VerboseSettings: c.VerboseSettings,
		Timeout:         c.Timeout,
		UserAgent:       c.UserAgent,
		Retries:         c.Retries,
		RetryBackoff:    c.RetryBackoff,
	})
	if err != nil {
		return nil, err
	}
	c.Client.Transport = t

	// Sadly, all the client parameters take the DefaultTimeout from the runtime
	// client if not specified in the call as a query parameter, modifying this
	// value effectively affects all of the related clients.
	runtimeclient.DefaultTimeout = c.Timeout

	// If retries are set, then we don't want the framework cancellation to kick in
	// while we're still retrying requests, to enable that, we're multiplying the
	// specified timeout by the # of retries, accounting for the backoffs as well.
	if c.Retries > 0 {
		if c.RetryBackoff.Microseconds() <= 0 {
			c.RetryBackoff = defaultBackoff
		}
		runtimeclient.DefaultTimeout = c.Timeout*time.Duration(c.Retries) + (c.RetryBackoff * time.Duration(c.Retries))
		// We leave the http client with no timeout, since we don't want yet another
		// layer of context cancellatins to step in.
		c.Client.Timeout = 0
	}

	transport, err := NewCloudClientRuntime(c)
	if err != nil {
		return nil, err
	}

	var api = API{AuthWriter: c.AuthWriter, V1API: client.New(transport, nil)}
	if !c.SkipLogin {
		if err := LoginUser(&api, c.ErrorDevice); err != nil {
			return nil, err
		}
	}

	return &api, nil
}
