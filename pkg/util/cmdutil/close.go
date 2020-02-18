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

package cmdutil

import (
	"fmt"
	"io"
)

// ErrorLoggerFunc is the callback for logging error
type ErrorLoggerFunc func(err error)

// CloseResource try to close the resource and log error if occurs.
func CloseResource(closer io.Closer, logger ErrorLoggerFunc) {
	// Check if closer is nil
	if closer == nil {
		return
	}
	// Try to close the resource
	if err := closer.Close(); err != nil {
		if logger != nil {
			// Wrap error
			logger(fmt.Errorf("unable to close resource: %w", err))
		}
	}
}
