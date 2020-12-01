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
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

const (
	// ErrTimedOutMsg is returned when the error is context.DeadlineExceeded.
	ErrTimedOutMsg = "operation timed out"

	errPrefix = "api error"
)

type knownError interface {
	GetPayload() *models.BasicFailedReply
}

// Error wraps a API error and implements both the Error and Unwrap interfaces.
type Error struct {
	// Err is the unmodified wrapped error.
	Err error

	// merr is returned in the case that multiple errors are found as a
	// *models.BasicFailedReply.
	merr *multierror.Prefixed

	// err will only be populated when merr is not.
	err error
}

// Wrap creates a new Error from the passed error, it doesn't modify the
// original error, but instead, tries to unwrap and extract the wrapped
// errors form the go-openapi/runtime operations. The original error type
// can be accessed by calling Unwrap() or errors.Unwrap/Is/As.
//
// If the error is of *models.BasicFailedReply type, then the errors are
// unpacked into a *multierror.Prefixed which can be obtained by calling
// Multierror().
// The other possible case is when the API returns an error with an unexpected
// status code which is not present in the swagger spec, in which case,
// the same operation will be attempted (unwrapping a BasicFailedReply)
// with a fallback to unmarshal any JSON which can be read:
// * error is nil, in which case nil is returned.
// * error is a context.DeadlineExceeded error, which equals a timeout.
// * error is of type *runtime.APIError, meaning the returned API error wasn't
//   defined in the Swagger spec from which the source code has been generated
//   * HTTP code is 449, the authenticated user needs to elevate-permissions.
//   * The type wraps *http.Response, the body is read and tries json.Unmarshal
//     to *models.BasicFailedResponse and each of the BasicFailedReplyElement
//     is then added to an instance of multierror.Prefixed and returned.
//   * The error is unknown, returns "<OperationName> (status <StatusCode)".
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	result := Error{Err: err}

	if bfr, ok := err.(knownError); ok {
		m := newBasicFailedReplyMultierror(errPrefix, bfr.GetPayload())
		if m.ErrorOrNil() != nil {
			result.merr = m
			return &result
		}
	}

	if rtimeErr := unwrapRuntimeAPIError(errPrefix, err); rtimeErr != nil {
		var p *multierror.Prefixed
		if errors.As(rtimeErr, &p) {
			result.merr = p
			return &result
		}
		result.err = rtimeErr
	}

	return &result
}

// Unwrap returns the wrapped error.
func (e *Error) Unwrap() error {
	if e == nil || e.Err == nil {
		return nil
	}

	return e.Err
}

// Is implements errors.Is by comparing the current value directly.
func (e *Error) Is(target error) bool {
	return errors.Is(e.Unwrap(), target)
}

func (e *Error) Error() string {
	if e == nil || e.Err == nil {
		return ""
	}

	if reflect.Ptr != reflect.ValueOf(e.Err).Kind() {
		if e.Err == context.DeadlineExceeded {
			return ErrTimedOutMsg
		}
		return e.Err.Error()
	}

	if e.merr != nil && e.merr.ErrorOrNil() != nil {
		return e.merr.Error()
	}

	if e.err != nil {
		return e.err.Error()
	}

	return e.Err.Error()
}

// Multierror returns a multierror if there's one.
func (e *Error) Multierror() *multierror.Prefixed {
	return e.merr
}

func unwrapRuntimeAPIError(prefix string, err error) error {
	apiErr, ok := err.(*runtime.APIError)
	if !ok {
		return nil
	}

	if apiErr.Code == 449 {
		return ErrMissingElevatedPermissions
	}

	if err := tryWrappedResponse(prefix, apiErr.Response); err != nil {
		return err
	}

	if res, _ := json.MarshalIndent(apiErr.Response, "", "  "); !bytes.Equal(res, []byte("{}")) {
		if err := unmarshalBasicFailedReply(prefix, res); err != nil {
			return err
		}
		return errors.New(string(res))
	}

	return fmt.Errorf("%s (status %d)", apiErr.OperationName, apiErr.Code)
}

func tryWrappedResponse(prefix string, apiError interface{}) error {
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
		return fmt.Errorf("failed reading error body: %w", err)
	}
	defer res.Body.Close()

	if err := unmarshalBasicFailedReply(prefix, b); err != nil {
		return err
	}

	return errors.New(string(b))
}

func unmarshalBasicFailedReply(prefix string, b []byte) error {
	var basicFailedReply models.BasicFailedReply
	if err := json.Unmarshal(b, &basicFailedReply); err == nil {
		if err := newBasicFailedReplyMultierror(prefix, &basicFailedReply); err != nil {
			return err.ErrorOrNil()
		}
	}
	return nil
}
