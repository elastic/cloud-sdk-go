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
	"testing"
)

// -----------------------------------------------------------------------------
type mockCloser struct {
	raiseError bool
}

func (m *mockCloser) Close() error {
	if m.raiseError {
		return fmt.Errorf("foo")
	}

	// No error
	return nil
}

// -----------------------------------------------------------------------------

var closerLoggerTestCases = []struct {
	desc    string
	closer  io.Closer
	wantErr bool
}{
	{
		desc:    "closer nil",
		closer:  nil,
		wantErr: false,
	},
	{
		desc:    "closer error",
		closer:  &mockCloser{raiseError: true},
		wantErr: true,
	},
	{
		desc:    "success",
		closer:  &mockCloser{raiseError: false},
		wantErr: false,
	},
}

func TestCloser_CloseResource_NilLogger(t *testing.T) {
	for _, tC := range closerLoggerTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			CloseResource(tC.closer, nil)
			// Nothing to expect
		})
	}
}

func TestCloser_CloseResource(t *testing.T) {
	for _, tC := range closerLoggerTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			CloseResource(tC.closer, func(err error) {
				if tC.wantErr != (err != nil) {
					t.Errorf("unexpected error occurs, got %v", err)
				}
			})
		})
	}
}
