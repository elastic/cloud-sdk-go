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

package logging

import (
	"sync"
)

// MockDispatcher is a mock log dispatcher to be used in tests
type MockDispatcher struct {
	logger *Mock
}

// Mock is a mock logger to be used in tests
type Mock struct {
	lock           sync.Mutex
	messagesLogged []LogMessage
	err            error
}

// WithLogLevel sets the log dispatcher level and returns the dispatcher
func (ld *MockDispatcher) WithLogLevel(level LogLevel) Dispatcher {
	return ld
}

// Dispatch iterates over dispatcher loggers and logs the given message for each one of them
func (ld *MockDispatcher) Dispatch(msg LogMessage) error {
	return ld.logger.Log(msg)
}

// Add is a dummy implementation for mocking
func (ld *MockDispatcher) Add(loggers ...Logger) {
}

// GetDispatchedMessages returns a slice of all dispatched messages
func (ld *MockDispatcher) GetDispatchedMessages() []string {
	dms := make([]string, 0)
	for i := range ld.logger.messagesLogged {
		dms = append(dms, ld.logger.messagesLogged[i].Message)
	}
	return dms
}

// Log mocks logging used the mock logger implementation
func (logger *Mock) Log(msg LogMessage) error {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	if logger.err != nil {
		return logger.err
	}
	logger.messagesLogged = append(logger.messagesLogged, msg)
	return nil
}

// GetLoggedMessages returns a slice of all logged messages
func (logger *Mock) GetLoggedMessages() []LogMessage {
	return logger.messagesLogged
}

// NewMockDispatcher creates a new mock log dispatcher to be used in testing
func NewMockDispatcher() *MockDispatcher {
	return &MockDispatcher{
		logger: &Mock{},
	}
}
