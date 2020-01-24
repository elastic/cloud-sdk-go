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
	"strings"
	"sync"

	"github.com/fatih/color"
	multierror "github.com/hashicorp/go-multierror"

	loggingdecorator "github.com/elastic/cloud-sdk-go/pkg/logging/decorator"
)

// LogLevel defines the logging level as untyped int
type LogLevel int

const (
	// ERROR defines the error log level as int value of 1
	ERROR LogLevel = iota + 1
	// WARNING defines the warning log level as int value of 2
	WARNING
	// INFO defines the info log level as int value of 3
	INFO
	// DEBUG defines the debug log level as int value of 4
	DEBUG
	// TRACE defines the trace log level as int value of 5
	TRACE

	warn  = "WARNING"
	trace = "TRACE"
	info  = "INFO"
	err   = "ERROR"
	debug = "DEBUG"
)

// Logger is the interface describing a loggers's behavior
type Logger interface {
	Log(msg LogMessage) error
}

// Dispatcher is the common Log dispatcher interface
type Dispatcher interface {
	Dispatch(msg LogMessage) error
	Add(loggers ...Logger)
	WithLogLevel(level LogLevel) Dispatcher
}

// MuxDispatcher holds is a concurrent-safe collection of loggers
type MuxDispatcher struct {
	lock    sync.Mutex
	loggers []Logger
	level   LogLevel
}

// NewLogDispatcher properly creates a new log dispatcher initializing the mutex and loggers slice
func NewLogDispatcher(loggers ...Logger) *MuxDispatcher {
	disp := &MuxDispatcher{}
	disp.level = INFO
	disp.loggers = make([]Logger, 0)
	disp.Add(loggers...)

	return disp
}

// WithLogLevel sets the log dispatcher level and returns the dispatcher
func (ld *MuxDispatcher) WithLogLevel(level LogLevel) Dispatcher {
	ld.level = level
	return ld
}

// Dispatch iterates over dispatcher loggers and logs the given message for each one of them
func (ld *MuxDispatcher) Dispatch(msg LogMessage) error {
	var merr = new(multierror.Error)
	for _, logger := range ld.loggers {
		if ld.level >= msg.Log.loglevel {
			if err := logger.Log(msg); err != nil {
				merr = multierror.Append(merr, err)
			}
		}
	}
	return merr.ErrorOrNil()
}

// Add adds aa logger to the dispatcher
func (ld *MuxDispatcher) Add(loggers ...Logger) {
	ld.lock.Lock()
	ld.loggers = append(ld.loggers, loggers...)
	ld.lock.Unlock()
}

// LevelColor returns the color attribute based on the log level and the coloring scheme
func LevelColor(scheme loggingdecorator.ColoringScheme, logLevel string) color.Attribute {
	switch strings.ToUpper(logLevel) {
	case trace:
		return scheme.Trace
	case debug:
		return scheme.Debug
	case info:
		return scheme.Info
	case warn:
		return scheme.Warn
	case err:
		return scheme.Error
	default:
		return color.FgWhite
	}
}

// WarnMessage creates a new WARNING log message and properly initializes all required fields
func WarnMessage(msg string) LogMessage {
	return newLogMessage(WARNING).WithMessage(msg)
}

// ErrorMessage creates a new ERROR log message and properly initializes all required fields
func ErrorMessage(msg string) LogMessage {
	return newLogMessage(ERROR).WithMessage(msg).WithError(NewError().WithMessage(msg))
}

// InfoMessage creates a new INFO log message and properly initializes all required fields
func InfoMessage(msg string) LogMessage {
	return newLogMessage(INFO).WithMessage(msg)
}

// DebugMessage creates a new DEBUG log message and properly initializes all required fields
func DebugMessage(msg string) LogMessage {
	return newLogMessage(DEBUG).WithMessage(msg)
}

// TraceMessage creates a new TRACE log message and properly initializes all required fields
func TraceMessage(msg string) LogMessage {
	return newLogMessage(TRACE).WithMessage(msg)
}

func newLogMessage(level LogLevel) LogMessage {
	return NewLogMessage().
		WithAgent(NewAgent()).
		WithLog(NewLog().
			WithLevel(level))
}

// Level returns the LogLevel based on the given level as string
// If the level is not found then the default (INFO) is returned
func Level(level string) LogLevel {
	var logLevel LogLevel
	switch level {
	case trace:
		logLevel = TRACE
	case debug:
		logLevel = DEBUG
	case info:
		logLevel = INFO
	case warn:
		logLevel = WARNING
	case err:
		logLevel = ERROR
	default:
		logLevel = INFO
	}

	return logLevel
}
