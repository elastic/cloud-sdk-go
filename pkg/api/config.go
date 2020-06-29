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
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/auth"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

var (
	errEmptyAuthWriter = errors.New("auth writer must not be empty")
	errESSInvalidAuth  = errors.New(
		"apikey is the only valid authentication mechanism when targeting the Elasticsearch Service",
	)
)

// Config contains the API config
type Config struct {
	Client     *http.Client
	AuthWriter auth.Writer
	Host       string

	// SkipTLSVerify will not perform any TLS/SSL verification.
	SkipTLSVerify bool

	// SkipLogin skips validating the user / password with the instanced API
	// when AuthWriter equals *auth.UserLogin.
	SkipLogin bool

	// ErrorDevice is used to send errors to prevent cluttering the output.
	ErrorDevice io.Writer

	VerboseSettings

	// Timeout for all of the API calls performed through the API structure.
	Timeout time.Duration

	// UserAgent if specified, it sets the user agent on all outgoing requests.
	UserAgent string
}

// Validate returns an error if the config is invalid
func (c *Config) Validate() error {
	var merr = multierror.NewPrefixed("invalid api config")
	if c.Client == nil {
		merr = merr.Append(errors.New("client cannot be empty"))
	}

	if c.AuthWriter == nil {
		merr = merr.Append(errEmptyAuthWriter)
	}

	merr = merr.Append(checkHost(c.Host))
	merr = merr.Append(c.VerboseSettings.Validate())

	_, apikeyPtr := c.AuthWriter.(*auth.APIKey)
	_, apikey := c.AuthWriter.(auth.APIKey)
	if c.Host == ESSEndpoint && !(apikey || apikeyPtr) {
		merr = merr.Append(errESSInvalidAuth)
	}

	return merr.ErrorOrNil()
}

func (c *Config) fillDefaults() {
	if c.Timeout.Nanoseconds() <= 0 {
		c.Timeout = DefaultTimeout
	}

	if c.UserAgent == "" {
		c.UserAgent = DefaultUserAgent
	}

	if c.Host == "" {
		c.Host = ESSEndpoint
	}
}

// VerboseSettings define the behaviour of verbosity.
type VerboseSettings struct {
	Device  io.Writer
	Verbose bool

	// RedactAuth replaces the contents of the Authorization header with:
	// "[REDACTED]".
	RedactAuth bool
}

// Validate ensures the settings are usable.
func (settings VerboseSettings) Validate() error {
	var merr = multierror.NewPrefixed("invalid verbose settings")
	if settings.Verbose && settings.Device == nil {
		merr = merr.Append(errors.New(
			"output device cannot be empty when verbose is enabled",
		))
	}

	return merr.ErrorOrNil()
}

func checkHost(host string) error {
	if host == "" {
		return errors.New("host cannot be empty")
	}

	if !strings.HasSuffix(host, "/") {
		host += "/"
	}

	_, err := url.ParseRequestURI(host)
	return err
}
