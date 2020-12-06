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

// Unwrap Deprecated: unpacks an error message returned from a client API call.
// It checks for a few cases where the returned API error might not have been
// properly casted to its error type. It covers the following cases in order:
// * error is nil, in which case nil is returned.
// * error is a context.DeadlineExceeded error, which equals a timeout.
// * error is of type *runtime.APIError, meaning the returned API error wasn't
//   defined in the Swagger spec from which the source code has been generated
//   * HTTP code is 449, the authenticated user needs to elevate-permissions.
//   * The type wraps *http.Response, the body is read and tries json.Unmarshal
//     to *models.BasicFailedResponse and each of the BasicFailedReplyElement
//     is then added to an instance of multierror.Prefixed and returned.
//   * The error is unknown, returns "<OperationName> (status <StatusCode)".
// * error is a correctly unpacked into BasicFailedReply object which needs to
//   be unpacked from its container struct. If the error cannot be unpacked to
//   a BasicFailedReply, then a stringified json.MarshalIndent error is formed.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	return Wrap(err)
}
