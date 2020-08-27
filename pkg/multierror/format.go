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

package multierror

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/go-multierror"
)

// JSONError wraps a list of errors to be encoded in JSON format.
type JSONError struct {
	Errors []interface{} `json:"errors,omitempty"`
}

// WithFormatFunc takes in an error and tries to set the ErrorFormatFunc to the
// passed function if the error is of type *Prefixed or *multierror.Error,
// otherwise it returns the error as is.
func WithFormatFunc(err error, ff FormatFunc) error {
	var merr *Prefixed
	if errors.As(err, &merr) {
		merr.FormatFunc = ff
	}

	var hashiErr *multierror.Error
	if errors.As(err, &hashiErr) {
		hashiErr.ErrorFormat = multierror.ErrorFormatFunc(ff)
	}

	return err
}

// WithFormat is convenient a helper to modify the multierror format function
// to the JSON format, if the format isn't "json", then the unmodified error is
// returned.
func WithFormat(err error, format string) error {
	if !strings.EqualFold(format, "json") {
		return err
	}

	return WithFormatFunc(err, JSONFormatFunc)
}

// JSONFormatFunc takes in a list of errors and encodes each error to JSON.
// If the error is not encodable to JSON, then the error is transformed to the
// json format: `{"message": "err.Error()"}`.
func JSONFormatFunc(es []error) string {
	var buf = new(bytes.Buffer)
	var enc = json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	var errs = make([]interface{}, 0, len(es))
	for _, e := range es {
		var t = make(map[string]interface{})
		if err := swag.FromDynamicJSON(e, &t); len(t) == 0 || err != nil {
			t["message"] = e.Error()
		}
		errs = append(errs, t)
	}

	_ = enc.Encode(JSONError{Errors: errs})

	return buf.String()
}
