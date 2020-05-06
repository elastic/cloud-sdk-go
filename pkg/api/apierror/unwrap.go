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

// Unwrap unpacks an error message returned from a client API call.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if reflect.Ptr != reflect.ValueOf(err).Kind() {
		if err == context.DeadlineExceeded {
			return errors.New("operation timed out")
		}
		return err
	}
	if e, ok := err.(*runtime.APIError); ok {
		if e.Code == 449 {
			return errors.New("the requested operation requires elevated permissions")
		}

		var defaultError = fmt.Errorf("%s (status %d)", e.OperationName, e.Code)
		if e.Response != nil {
			if v := reflect.ValueOf(e.Response); v.IsValid() {
				if resp := v.FieldByName("resp"); resp.IsValid() {
					ptr := unsafe.Pointer(resp.Pointer())
					res := (*http.Response)(ptr)
					b, err := ioutil.ReadAll(res.Body)
					if err != nil {
						return fmt.Errorf(
							"failed reading error body: %s", defaultError,
						)
					}
					defer res.Body.Close()
					return errors.New(string(b))
				}
			}
		}

		if res, _ := json.MarshalIndent(e.Response, "", "  "); !bytes.Equal(res, []byte("{}")) {
			return errors.New(string(res))
		}
		return defaultError
	}

	payload := reflect.ValueOf(err).Elem().FieldByName("Payload")
	if payload.IsValid() {
		if r, ok := payload.Interface().(*models.BasicFailedReply); ok {
			merr := multierror.NewPrefixed("api error")
			for _, e := range r.Errors {
				merr = merr.Append(fmt.Errorf("%s: %s", *e.Code, *e.Message))
			}
			return merr.ErrorOrNil()
		}
		res, _ := json.MarshalIndent(payload.Interface(), "", "  ")
		return errors.New(string(res))
	}

	return err
}
