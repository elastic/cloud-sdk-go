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

package auth

import (
	"errors"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// Writer wraps the runtime.ClientAuthInfoWriter interface adding a method
// to Auth generic http.Request.
type Writer interface {
	runtime.ClientAuthInfoWriter
	AuthRequest(req *http.Request) *http.Request
}

// Config to create e new AuthWriters
type Config struct {
	APIKey   string
	Password string
	Username string
}

// Validate ensures that the config is usable.
func (c Config) Validate() error {
	var merr = multierror.NewPrefixed("authwriter")
	var emptyAPIKey = c.APIKey == ""
	var emptyUser = c.Username == ""
	var emptyPass = c.Password == ""

	var emptyCreds = emptyAPIKey && emptyUser && emptyPass
	if emptyCreds {
		merr = merr.Append(
			errors.New("one of apikey or username and password must be specified"),
		)
	}

	var allCreds = !emptyAPIKey && (!emptyUser || !emptyPass)
	if allCreds {
		merr = merr.Append(
			errors.New("only one of of apikey or username and password can be specified"),
		)
	}

	return merr.ErrorOrNil()
}

// NewAuthWriter creates a new instance of one of the implementations of Writer
// *APIKey or *UserLogin.
func NewAuthWriter(c Config) (Writer, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	if c.APIKey != "" {
		return NewAPIKey(c.APIKey)
	}

	return NewUserLogin(c.Username, c.Password)
}
