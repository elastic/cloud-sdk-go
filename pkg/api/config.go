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

	multierror "github.com/hashicorp/go-multierror"
)

var (
	errEmptyAuthWriter = errors.New("api: AuthWriter must not be empty")
)

// Config contains the API config
type Config struct {
	Client     *http.Client
	AuthWriter AuthWriter
	Host       string
	Region     string

	// SkipTLSVerify will not perform any TLS/SSL verification.
	SkipTLSVerify bool

	// ErrorDevice is used to send errors to prevent cluttering the output.
	ErrorDevice io.Writer

	VerboseSettings
}

// Validate returns an error if the config is invalid
func (c Config) Validate() error {
	var err = new(multierror.Error)
	if c.Client == nil {
		err = multierror.Append(err, errors.New("api: client cannot be empty"))
	}

	if c.AuthWriter == nil {
		err = multierror.Append(err, errEmptyAuthWriter)
	}

	err = multierror.Append(err, checkHost(c.Host))
	err = multierror.Append(err, c.VerboseSettings.Validate())

	return err.ErrorOrNil()
}

// VerboseSettings define the behaviour of verbosity.
type VerboseSettings struct {
	Verbose bool
	Device  io.Writer
}

// Validate ensures the settings are usable.
func (settings VerboseSettings) Validate() error {
	var err = new(multierror.Error)
	if settings.Verbose && settings.Device == nil {
		err = multierror.Append(err, errors.New(
			"api: invalid verbose settings: output device cannot be empty when verbose is enabled",
		))
	}

	return err.ErrorOrNil()
}

func checkHost(host string) error {
	if host == "" {
		return errors.New("api: host cannot be empty")
	}

	if !strings.HasSuffix(host, "/") {
		host += "/"
	}

	_, err := url.ParseRequestURI(host)
	return err
}
