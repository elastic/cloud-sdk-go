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

package testutils

import (
	"testing"

	"github.com/go-openapi/strfmt"
)

// ParseDate parses a string to date and if parsing generates an error,
// it fails the given testing suite, otherwise it returns the
// strfmt.Datetime parsed value
func ParseDate(t *testing.T, date string) strfmt.DateTime {
	dt, err := strfmt.ParseDateTime(date)
	if err != nil {
		t.Fatal(err)
	}

	return dt
}
