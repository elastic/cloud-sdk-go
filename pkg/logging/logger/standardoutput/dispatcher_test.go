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

	"github.com/elastic/cloud-sdk-go/pkg/logging"
)

func TestNewStandardOutputDispatcher(t *testing.T) {
	tests := []struct {
		name     string
		expected logging.Dispatcher
	}{
		{
			name:     "New standard output logger should return a properly constructed logger",
			expected: StandardOutputDispatcher{Dispatcher: logging.NewLogDispatcher(New())},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewStandardOutputDispatcher()

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("NewStandardOutputDispatcher() wanted %v returned %v",
					tt.expected, actual)
				return
			}
		})
	}
}

func TestStandardOutputDispatcher_Log(t *testing.T) {
	tests := []struct {
		name    string
		logger  StandardOutputDispatcher
		message string
	}{
		{
			name:    "should log message",
			logger:  StandardOutputDispatcher{Dispatcher: logging.NewMockDispatcher()},
			message: "test message",
		},
		{
			name:   "should not log message",
			logger: StandardOutputDispatcher{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Log(logging.InfoMessage(tt.message))
			logging.AssertLoggedMessages(t, tt.logger.Dispatcher, []string{tt.message})
		})
	}
}

func TestStandardOutputDispatcher_LogIf(t *testing.T) {
	tests := []struct {
		name        string
		logger      StandardOutputDispatcher
		message     string
		shouldLog   bool
		expMessages []string
	}{
		{
			name:        "should log message",
			logger:      StandardOutputDispatcher{Dispatcher: logging.NewMockDispatcher()},
			message:     "test message",
			shouldLog:   true,
			expMessages: []string{"test message"},
		},
		{
			name:    "should not log message",
			logger:  StandardOutputDispatcher{Dispatcher: logging.NewMockDispatcher()},
			message: "test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.LogIf(logging.InfoMessage(tt.message), tt.shouldLog)
			logging.AssertLoggedMessages(t, tt.logger.Dispatcher, tt.expMessages)
		})
	}
}
