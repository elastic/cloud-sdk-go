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
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// APIKey represents an APIKey used in the Authorization header as means of
// authentication. It is the preferred method of authentication.
type APIKey string

// NewAPIKey constructs a new APIKey, returns an error if the key is invalid.
func NewAPIKey(key string) (*APIKey, error) {
	var k = APIKey(key)

	if err := k.Validate(); err != nil {
		return nil, err
	}

	return &k, nil
}

// Validate ensures the validity of the data container.
func (k APIKey) Validate() error {
	if k == "" {
		return errors.New("auth: APIKey must not be empty")
	}
	return nil
}

// AuthenticateRequest authenticates a runtime.ClientRequest. Implements the
// runtime.ClientAuthInfoWriter interface.
func (k APIKey) AuthenticateRequest(c runtime.ClientRequest, r strfmt.Registry) error {
	return httptransport.APIKeyAuth("Authorization", "header", "ApiKey "+k.String()).
		AuthenticateRequest(c, r)
}

// AuthRequest adds the Authorization header to an http.Request
func (k APIKey) AuthRequest(req *http.Request) *http.Request {
	req.Header.Add("Authorization", "ApiKey "+k.String())
	return req
}

func (k APIKey) String() string { return string(k) }
