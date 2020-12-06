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
	"bytes"
	"encoding/json"
	"errors"

	"github.com/go-openapi/runtime"
)

// IsRuntimeStatusCode returns true when the error is of *runtime.APIError type
// and the error status code matches the one passed as the argument.
func IsRuntimeStatusCode(err error, code int) bool {
	var rtError *runtime.APIError
	rtCode := -1

	if errors.As(err, &rtError) {
		rtCode = rtError.Code
	}

	return rtCode == code
}

// IsRuntimeStatusCodeHasString returns true when the error is *runtime.APIError
// and the error status code matches the one passed as the argument.
func IsRuntimeStatusCodeHasString(err error, code int, msg string) bool {
	var rtError *runtime.APIError
	var strMatch bool
	rtCode := -1

	if errors.As(err, &rtError) {
		rtCode = rtError.Code
		b, _ := json.Marshal(rtError.Response)
		if bytes.Contains(b, []byte(msg)) {
			strMatch = true
		}
	}

	return rtCode == code && strMatch
}
