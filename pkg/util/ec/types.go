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

package ec

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
)

// Validator interface is mostly used for parameter structures that
// need to check a set of conditions and act as a gate before running
// expensive external calls
type Validator interface {
	Validate() error
}

// Bool creates a new bool pointer from a boolean
func Bool(b bool) *bool { return &b }

// String creates a new string pointer from a string
func String(s string) *string { return &s }

// Int32 creates a new int32 pointer from an int32
func Int32(i int32) *int32 { return &i }

// Int64 creates a new int64 pointer from an int64
func Int64(i int64) *int64 { return &i }

// CompareStructs two structs and return the differences
func CompareStructs(a, b interface{}) (equals bool, diff string, err error) {
	var spewConfig = spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
	}

	if reflect.DeepEqual(a, b) {
		return true, "", nil
	}

	r, l := spewConfig.Sdump(b), spewConfig.Sdump(a)
	diff, err = difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:       difflib.SplitLines(l),
		B:       difflib.SplitLines(r),
		Context: 0,
	})
	return false, diff, err
}
