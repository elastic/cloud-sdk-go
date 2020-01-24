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

package logger

import (
	"reflect"
	"testing"

	loggingdecorator "github.com/elastic/cloud-sdk-go/pkg/logging/decorator"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name              string
		expFormat         string
		expColoringScheme loggingdecorator.ColoringScheme
	}{
		{
			name:              "New standard output logger should return a properly constructed logger",
			expFormat:         defFmt,
			expColoringScheme: *loggingdecorator.DefaultColoringScheme(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sol := New()

			if sol.fmt != tt.expFormat {
				t.Errorf("New() logger format wanted %s lreturned %s", tt.expFormat, sol.fmt)
				return
			}
			if !reflect.DeepEqual(*sol.ColoringScheme, tt.expColoringScheme) {
				t.Errorf("New() logger coloring scheme wanted %v returned %v",
					tt.expColoringScheme, sol.ColoringScheme)
				return
			}
		})
	}
}

func TestNewWithArguments(t *testing.T) {
	var otherFormat = "other Format"
	var otherScheme = loggingdecorator.ColoringScheme{}

	tests := []struct {
		name              string
		expFormat         string
		expColoringScheme loggingdecorator.ColoringScheme
	}{
		{
			name:              "New standard output logger with arguments should return a properly constructed logger",
			expFormat:         otherFormat,
			expColoringScheme: otherScheme,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sol := New().WithColoringScheme(&otherScheme).WithFormat(otherFormat)

			if sol.fmt != tt.expFormat {
				t.Errorf("New() logger format wanted %s lreturned %s", tt.expFormat, sol.fmt)
				return
			}
			if !reflect.DeepEqual(*sol.ColoringScheme, tt.expColoringScheme) {
				t.Errorf("New() logger coloring scheme wanted %v returned %v",
					tt.expColoringScheme, sol.ColoringScheme)
				return
			}
		})
	}
}

func TestNewNullifyColoringScheme(t *testing.T) {
	tests := []struct {
		name              string
		expColoringScheme loggingdecorator.ColoringScheme
	}{
		{
			name:              "Standard output logger should return a properly constructed logger even when coloring scheme is nullified",
			expColoringScheme: *loggingdecorator.DefaultColoringScheme(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sol := New().WithColoringScheme(nil)
			if !reflect.DeepEqual(*sol.ColoringScheme, tt.expColoringScheme) {
				t.Errorf("New() logger coloring scheme wanted %v returned %v",
					tt.expColoringScheme, sol.ColoringScheme)
				return
			}
		})
	}
}
