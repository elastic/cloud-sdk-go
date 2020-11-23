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

func logMessageMatcher(message string, level LogLevel) func(message LogMessage) bool {
	return func(logMessage LogMessage) bool {
		return logMessage.Message == message &&
			Level(logMessage.Log.Level) == level
	}
}

// DebugLogMessageMatcher returns a matcher function to be used by a Mockery mock which expects a debug message.
func DebugLogMessageMatcher(message string) func(message LogMessage) bool {
	return logMessageMatcher(message, DEBUG)
}

// InfoLogMessageMatcher returns a matcher function to be used by a Mockery mock which expects an info message.
func InfoLogMessageMatcher(message string) func(message LogMessage) bool {
	return logMessageMatcher(message, INFO)
}

// WarnLogMessageMatcher returns a matcher function to be used by a Mockery mock which expects a warning message.
func WarnLogMessageMatcher(message string) func(message LogMessage) bool {
	return logMessageMatcher(message, WARNING)
}

// ErrorLogMessageMatcher returns a matcher function to be used by a Mockery mock which expects an error message.
func ErrorLogMessageMatcher(message string) func(message LogMessage) bool {
	return logMessageMatcher(message, ERROR)
}
