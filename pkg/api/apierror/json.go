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

// JSONError wraps any incoming error inside this struct so that it can be
// correctly  marshaled to JSON. If Error() is called, the error is still
// returned in a string format.
type JSONError struct {
	Message string `json:"message,omitempty"`
}

// Error complies with the error interface
func (me JSONError) Error() string { return me.Message }

// NewJSONError creates a marshaleable error from an error.
func NewJSONError(err error) error {
	return JSONError{Message: err.Error()}
}
