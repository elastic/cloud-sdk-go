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

import "errors"

var (
	// ErrNoTokenAvailable should be returned by any implementations which
	// fail to find any persisted token for the user.
	ErrNoTokenAvailable = errors.New("auth: no persisted token available for the user")
)

// TokenHandler is an interface which abstracts the application management of
// JWT Bearer tokens. Lightweight on purpose to allow loose implementations.
type TokenHandler interface {
	// Load returns a persisted token scoped to the current authenticated user.
	Load() (string, error)

	// Update replaces the token with a new one.
	Update(string) error

	// Token returns current token.
	Token() string
}
