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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"unsafe"

	"github.com/go-openapi/runtime"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

var (
	// ErrTimedOut is returned when the error is context.DeadlineExceeded.
	ErrTimedOut = errors.New("operation timed out")

	// ErrMissingElevatedPermissions is returned when the error code is 449.
	ErrMissingElevatedPermissions = errors.New("the requested operation requires elevated permissions")
)

// Unwrap unpacks an error message returned from a client API call.
// It checks for a few cases where the returned API error might not have been
// properly casted to its error type. It covers the following cases in order:
// * error is nil, in which case nil is returned.
// * error is a context.DeadlineExceeded error, which equals a timeout.
// * error is of type *runtime.APIError, meaning the returned API error wasn't
//   defined in the Swagger spec from which the source code has been generated
//   * HTTP code is 449, the authenticated user needs to elevate-permissions.
//   * The type wraps *http.Response, the body is read and tries json.Unmarshal
//     to *models.BasicFailedResponse and each of the BasicFailedReplyElement
//     is then added to an instance ofmultierror.Prefixed and returned.
//   * The error is unknown, returns "<OperationName> (status <StatusCode)".
// * error is a correctly unpacked into BasicFailedReply object which needs to
//   be unpacked from its container struct. If the error cannot be unpacked to
//   a BasicFailedReply, then a stringified json.MarshalIndent error is formed.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if reflect.Ptr != reflect.ValueOf(err).Kind() {
		if err == context.DeadlineExceeded {
			return ErrTimedOut
		}
		return err
	}

	if err := unwrapRuntimeAPIError(err); err != nil {
		return err
	}

	if err := unwrapBasicFailedReply(err); err != nil {
		return err
	}

	return err
}

func unwrapBasicFailedReply(err error) error {
	payload := reflect.ValueOf(err).Elem().FieldByName("Payload")
	if !payload.IsValid() {
		return nil
	}

	if r, ok := payload.Interface().(*models.BasicFailedReply); ok {
		return newMultierror(r)
	}

	res, _ := json.MarshalIndent(payload.Interface(), "", "  ")
	if err := unmarshalBasicFailedReply(res); err != nil {
		return err
	}
	return errors.New(string(res))
}

func unwrapRuntimeAPIError(err error) error {
	apiErr, ok := err.(*runtime.APIError)
	if !ok {
		return nil
	}

	if apiErr.Code == 449 {
		return ErrMissingElevatedPermissions
	}

	if err := tryWrappedResponse(apiErr.Response); err != nil {
		return err
	}

	if res, _ := json.MarshalIndent(apiErr.Response, "", "  "); !bytes.Equal(res, []byte("{}")) {
		if err := unmarshalBasicFailedReply(res); err != nil {
			return err
		}
		return errors.New(string(res))
	}

	return fmt.Errorf("%s (status %d)", apiErr.OperationName, apiErr.Code)
}

func tryWrappedResponse(apiError interface{}) error {
	v := reflect.ValueOf(apiError)
	if !v.IsValid() {
		return nil
	}

	resp := v.FieldByName("resp")
	if !resp.IsValid() {
		return nil
	}

	ptr := unsafe.Pointer(resp.Pointer())
	res := (*http.Response)(ptr)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed reading error body")
	}
	defer res.Body.Close()

	if err := unmarshalBasicFailedReply(b); err != nil {
		return err
	}

	return errors.New(string(b))
}

func unmarshalBasicFailedReply(b []byte) error {
	var basicFailedReply models.BasicFailedReply
	if err := json.Unmarshal(b, &basicFailedReply); err == nil {
		if err := newMultierror(&basicFailedReply); err != nil {
			return err
		}
	}
	return nil
}
