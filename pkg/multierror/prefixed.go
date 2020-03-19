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
	"errors"
	"fmt"
	"runtime"

	"github.com/hashicorp/go-multierror"
)

// FormatFunc defines a format function which should format a slice of errors
// into a string.
type FormatFunc func(es []error) string

// Prefixed is a multierror which will prefix the error output message with the
// specified prefix.
type Prefixed struct {
	Prefix     string
	Errors     []error
	FormatFunc FormatFunc
}

// NewPrefixed creates a new pointer to Prefixed w
func NewPrefixed(prefix string, errs ...error) *Prefixed {
	return &Prefixed{Prefix: prefix, Errors: unpackErrors(errs...)}
}

// Append appends a number of errors to the current instance of Prefixed. It'll
// unwrap any wrapped errors in the form of *Prefixed or *multierror.Error.
func (p *Prefixed) Append(errs ...error) *Prefixed {
	p.Errors = append(p.Errors, unpackErrors(errs...)...)
	return p
}

// ErrorOrNil either returns nil when the type is nil or when there's no Errors.
// Otherwise, the type is returned.
func (p *Prefixed) ErrorOrNil() error {
	if len(p.Errors) > 0 {
		return p
	}

	return nil
}

func (p *Prefixed) Error() string {
	if len(p.Errors) == 0 {
		return ""
	}

	if p.FormatFunc == nil {
		p.FormatFunc = multierror.ListFormatFunc
	}

	return fmt.Sprint(p.Prefix, ": ", p.FormatFunc(p.Errors))
}

func unpackErrors(errs ...error) []error {
	var result = make([]error, 0, len(errs))
	for _, err := range errs {
		if err == nil {
			continue
		}

		if e, ok := err.(*Prefixed); ok {
			result = append(result, prefixIndividualErrors(e)...)
			continue
		}
		if e, ok := err.(*multierror.Error); ok {
			result = append(result, e.Errors...)
			continue
		}
		result = append(result, err)
	}
	return result
}

func prefixIndividualErrors(prefixed *Prefixed) []error {
	var result = make([]error, 0, len(prefixed.Errors))
	for _, errElement := range prefixed.Errors {
		if prefixed.Prefix == "" {
			// Calling it with skip 3 since there's 3 internal calls which must
			// be skipped until the external call can surface.
			if pc, _, _, ok := runtime.Caller(3); ok {
				if details := runtime.FuncForPC(pc); details != nil {
					prefixed.Prefix = details.Name()
				}
			}
		}
		result = append(result, errors.New(
			fmt.Sprint(prefixed.Prefix, ": ", errElement.Error()),
		))
	}
	return result
}
