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
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	// DefaultTimeout is used when TransportConfig.Transport is not specified.
	DefaultTimeout = 30 * time.Second
)

// TransportConfig is meant to be used so an http.RoundTripper is constructed
// with the appropriate settings.
type TransportConfig struct {
	// When SkipTLSVerify the TLS verification is completely skipped.
	SkipTLSVerify bool

	// ErrorDevice where any error or notices will be sent.
	ErrorDevice io.Writer

	// Can enable a debug RoundTripper which dumps the request and responses to
	// the configured device.
	VerboseSettings

	// Timeout for the Transport net.Dialer.
	Timeout time.Duration

	// UserAgent if specified, it sets the user agent on all outgoing requests.
	UserAgent string

	// Number of retries to perform on request timeout.
	Retries int

	// Cooldown time between retries.
	RetryBackoff time.Duration
}

func newDefaultTransport(timeout time.Duration) *http.Transport {
	if timeout.Seconds() <= 0 {
		timeout = DefaultTimeout
	}

	var transport = http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = (&net.Dialer{
		Timeout:   timeout,
		KeepAlive: 30 * time.Second,
	}).DialContext

	return transport
}

// NewTransport constructs a new http.RoundTripper from its config.
func NewTransport(rt http.RoundTripper, cfg TransportConfig) (http.RoundTripper, error) {
	if rt == nil {
		rt = newDefaultTransport(cfg.Timeout)
	}

	if t, ok := rt.(*http.Transport); ok {
		if t.TLSClientConfig == nil {
			t.TLSClientConfig = new(tls.Config)
		}
		t.TLSClientConfig.InsecureSkipVerify = cfg.SkipTLSVerify
	}

	if t, ok := rt.(*CustomTransport); ok {
		return t, nil
	}

	return NewCustomTransport(CustomTransportCfg{
		RoundTripper: rt,
		UserAgent:    cfg.UserAgent,
		Retries:      cfg.Retries,
		RedactAuth:   cfg.RedactAuth,
		Verbose:      cfg.Verbose,
		Writer:       cfg.Device,
		Backoff:      cfg.RetryBackoff,
	})
}
