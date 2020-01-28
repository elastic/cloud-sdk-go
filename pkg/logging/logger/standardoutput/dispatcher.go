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
	"fmt"

	"github.com/elastic/cloud-sdk-go/pkg/logging"
)

// StandardOutputDispatcher is a struct to encapsulate a logging logging.Dispatcher to allow us to add more methods
type StandardOutputDispatcher struct {
	logging.Dispatcher
}

// NewStandardOutputDispatcher creates a new StandardOutputDispatcher by initializing the required fields
// The new dispatcher has a standard output logger defined
func NewStandardOutputDispatcher() StandardOutputDispatcher {
	return StandardOutputDispatcher{
		Dispatcher: logging.NewLogDispatcher(New()),
	}
}

// Log will log the given message
func (l StandardOutputDispatcher) Log(message logging.LogMessage) {
	if l.Dispatcher == nil {
		return
	}
	if err := l.Dispatch(message); err != nil {
		fmt.Printf("unable to dispatch log message [%s]. error : %s\n", message.Message, err.Error())
	}
}

// LogIf will log the given message only if `shouldLog` argument is `true`
func (l StandardOutputDispatcher) LogIf(message logging.LogMessage, shouldLog bool) {
	if shouldLog {
		l.Log(message)
	}
}

// Info logs an info message
func (l StandardOutputDispatcher) Info(msg string) {
	l.Log(logging.InfoMessage(msg))
}

// Warn logs a warn message
func (l StandardOutputDispatcher) Warn(msg string) {
	l.Log(logging.WarnMessage(msg))
}

// Error logs an error message
func (l StandardOutputDispatcher) Error(msg string) {
	l.Log(logging.ErrorMessage(msg))
}

// Debug logs a debug message
func (l StandardOutputDispatcher) Debug(msg string) {
	l.Log(logging.DebugMessage(msg))
}

// Trace logs a trace message
func (l StandardOutputDispatcher) Trace(msg string) {
	l.Log(logging.TraceMessage(msg))
}
