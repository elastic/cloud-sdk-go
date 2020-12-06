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

package apierror

import (
	"errors"
)

var (
	// ErrMissingAPI is thrown when the *api.API reference is null.
	ErrMissingAPI = errors.New("api reference is required for the operation")

	// ErrDeploymentID is the message returned when a provided cluster id is not of the expected length (32 chars).
	ErrDeploymentID = errors.New("deployment id should have a length of 32 characters")

	// ErrMissingElevatedPermissions is returned when the error code is 449.
	ErrMissingElevatedPermissions = errors.New("the requested operation requires elevated permissions")
)
