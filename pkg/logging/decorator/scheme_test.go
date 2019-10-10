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

package loggingdecorator

import (
	"testing"

	"github.com/fatih/color"
)

func TestDefaultColoringScheme(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "New standard output logger should return a properly constructed logger",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheme := DefaultColoringScheme()

			if scheme.Trace != color.FgWhite {
				t.Errorf("scheme trace color wanted %v returned %v", color.FgWhite, scheme.Debug)
				return
			}
			if scheme.Debug != color.FgWhite {
				t.Errorf("scheme debug color wanted %v returned %v", color.FgWhite, scheme.Debug)
				return
			}
			if scheme.Info != color.FgWhite {
				t.Errorf("scheme info color wanted %v returned %v", color.FgWhite, scheme.Debug)
				return
			}
			if scheme.Warn != color.FgYellow {
				t.Errorf("scheme warn color wanted %v returned %v", color.FgYellow, scheme.Debug)
				return
			}
			if scheme.Error != color.FgRed {
				t.Errorf("scheme debug color wanted %v returned %v", color.FgRed, scheme.Debug)
				return
			}
		})
	}
}
