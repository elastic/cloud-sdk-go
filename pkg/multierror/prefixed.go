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
	"sort"
	"sync"

	"github.com/hashicorp/go-multierror"
)

type asMulti interface {
	Multierror() *Prefixed
}

// FormatFunc defines a format function which should format a slice of errors
// into a string.
type FormatFunc func(es []error) string

// Prefixed is a multierror which will prefix the error output message with the
// specified prefix. The output will be sorted so the errors can be asserted on
// test functions in case the multierror has been used in different goroutines.
// This structure is concurrently safe, avoid using value semantics.
type Prefixed struct {
	Prefix     string
	Errors     []error
	FormatFunc FormatFunc

	// SkipPrefixing when set to true subsequently appended errors needn't be
	// prefixed. This is particularly useful multierrors which contain JSON
	// marshaleable errors.
	SkipPrefixing bool

	mu sync.RWMutex
}

// NewPrefixed creates a new pointer to Prefixed.
func NewPrefixed(prefix string, errs ...error) *Prefixed {
	return &Prefixed{Prefix: prefix, Errors: unpackErrors(prefix, errs...)}
}

// NewJSONPrefixed returns a new pointer to Prefixed with the right config for
// JSON packed errors to not be prefixed.
func NewJSONPrefixed(prefix string, errs ...error) *Prefixed {
	return &Prefixed{
		Prefix:        prefix,
		Errors:        unpackErrors(prefix, errs...),
		SkipPrefixing: true,
		FormatFunc:    JSONFormatFunc,
	}
}

// Append appends a number of errors to the current instance of Prefixed. It'll
// unwrap any wrapped errors in the form of *Prefixed or *multierror.Error.
func (p *Prefixed) Append(errs ...error) *Prefixed {
	defer p.mu.Unlock()
	p.mu.Lock()

	p.Errors = append(p.Errors, unpackErrors(p.Prefix, errs...)...)
	return p
}

// ErrorOrNil either returns nil when the type is nil or when there's no Errors.
// Otherwise, the type is returned.
func (p *Prefixed) ErrorOrNil() error {
	defer p.mu.RUnlock()
	p.mu.RLock()

	if len(p.Errors) > 0 {
		return p
	}

	return nil
}

// Error returns the stored slice of error formatted using a set FormatFunc or
// multierror.ListFormatFunc when no FormatFunc is specified.
func (p *Prefixed) Error() string {
	defer p.mu.Unlock()
	p.mu.Lock()

	if len(p.Errors) == 0 {
		return ""
	}

	// Sort the errors so there's some consistency on the output.
	sort.SliceStable(p.Errors, func(i, j int) bool {
		return p.Errors[i].Error() < p.Errors[j].Error()
	})

	if p.FormatFunc == nil {
		p.FormatFunc = wrapPrefix(p.Prefix, multierror.ListFormatFunc)
	}

	return p.FormatFunc(p.Errors)
}

func wrapPrefix(prefix string, f FormatFunc) FormatFunc {
	return func(es []error) string {
		return fmt.Sprint(prefix, ": ", f(es))
	}
}

func unpackErrors(prefix string, errs ...error) []error {
	result := make([]error, 0, len(errs))
	for _, err := range errs {
		if err == nil {
			continue
		}

		result = append(result, unpack(prefix, err)...)
	}
	return result
}

func unpack(prefix string, err error) []error {
	// Handles wrapped apierror.Error, this intermediary interface is
	// needed since the apierror package imports multierror.Prefixed.
	if m, ok := err.(asMulti); ok {
		if e := m.Multierror(); e != nil {
			return handleNestedPrefixed(prefix, e)
		}
	}

	var e *Prefixed
	if errors.As(err, &e) {
		return handleNestedPrefixed(prefix, e)
	}

	var hashiErr *multierror.Error
	if errors.As(err, &hashiErr) {
		return hashiErr.Errors
	}

	return []error{err}
}

func handleNestedPrefixed(prefix string, e *Prefixed) []error {
	if prefix == e.Prefix || e.SkipPrefixing {
		return e.Errors
	}

	return prefixIndividualErrors(e)
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
