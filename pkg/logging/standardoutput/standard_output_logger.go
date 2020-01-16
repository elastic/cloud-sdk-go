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

package standardoutputlogger

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/elastic/cloud-sdk-go/pkg/logging"
	loggingdecorator "github.com/elastic/cloud-sdk-go/pkg/logging/decorator"
)

const (
	// TRACE represents the trace logging level
	TRACE = "TRACE"
	// DEBUG represents the debug logging level
	DEBUG = "DEBUG"
	// INFO represents the info logging level
	INFO = "INFO"
	// WARN represents the warn logging level
	WARN = "WARN"
	// ERROR represents the error logging level
	ERROR = "ERROR"
)

var (
	// DefaultColoringScheme format of log message
	defFmt = "[%s] [%s:%s] %s\n"
)

// StandardOutputLogger is the implementation of standard output logging
type StandardOutputLogger struct {
	// StandardOutputLogger format of log message and time
	fmt string
	*loggingdecorator.ColoringScheme
}

// Log writes the given log message to the standard output using the logger defined format
func (logger StandardOutputLogger) Log(msg logging.LogMessage) error {
	color.Set(levelColor(*logger.ColoringScheme, msg.Log.Level))
	fmt.Printf(logger.fmt, msg.Timestamp, msg.Agent.Name, msg.Log.Level, msg.Message)
	color.Unset()
	return nil
}

// WithFormat sets the logger format and returns the logger itself
func (logger *StandardOutputLogger) WithFormat(format string) *StandardOutputLogger {
	logger.fmt = format
	return logger
}

// WithColoringScheme sets the coloring scheme and returns the logger itself
func (logger *StandardOutputLogger) WithColoringScheme(scheme *loggingdecorator.ColoringScheme) *StandardOutputLogger {
	if scheme == nil {
		logger.ColoringScheme = loggingdecorator.DefaultColoringScheme()
	} else {
		logger.ColoringScheme = scheme
	}
	return logger
}

// New properly creates a new standard output logger initializing its internal state with default values
func New() *StandardOutputLogger {
	return &StandardOutputLogger{
		fmt:            defFmt,
		ColoringScheme: loggingdecorator.DefaultColoringScheme(),
	}
}

func levelColor(scheme loggingdecorator.ColoringScheme, logLevel string) color.Attribute {
	switch strings.ToUpper(logLevel) {
	case TRACE:
		return scheme.Trace
	case DEBUG:
		return scheme.Debug
	case INFO:
		return scheme.Info
	case WARN:
		return scheme.Warn
	case ERROR:
		return scheme.Error
	default:
		return color.FgWhite
	}
}

// StandardLogger is a struct to encapsulate a logging Dispatcher to allow us to add more methods
type StandardLogger struct {
	logging.Dispatcher
}

// NewStandardLogger creates a new standard Logger with required dispatcher properly configured
func NewStandardLogger() logging.Dispatcher {
	return StandardLogger{
		Dispatcher: NewDispatcher(),
	}
}

// NewDispatcher creates a new Logger dispatcher with required loggers properly configured
func NewDispatcher() logging.Dispatcher {
	return logging.NewLogDispatcher(New())
}

// Warn creates a new WARNING log message and properly initializes all required fields
func Warn(msg string) logging.LogMessage {
	return newLogMessage(logging.WARNING).WithMessage(msg)
}

// Error creates a new ERROR log message and properly initializes all required fields
func Error(msg string) logging.LogMessage {
	return newLogMessage(logging.ERROR).WithMessage(msg).WithError(logging.NewError().WithMessage(msg))
}

// Info creates a new INFO log message and properly initializes all required fields
func Info(msg string) logging.LogMessage {
	return newLogMessage(logging.INFO).WithMessage(msg)
}

// Debug creates a new DEBUG log message and properly initializes all required fields
func Debug(msg string) logging.LogMessage {
	return newLogMessage(logging.DEBUG).WithMessage(msg)
}

// Trace creates a new TRACE log message and properly initializes all required fields
func Trace(msg string) logging.LogMessage {
	return newLogMessage(logging.TRACE).WithMessage(msg)
}

// Level returns the LogLevel based on the given level as string
// If the level is not found then the default (INFO) is returned
func Level(level string) logging.LogLevel {
	var logLevel logging.LogLevel
	switch level {
	case TRACE:
		logLevel = logging.TRACE
	case DEBUG:
		logLevel = logging.DEBUG
	case INFO:
		logLevel = logging.INFO
	case WARN:
		logLevel = logging.WARNING
	case ERROR:
		logLevel = logging.ERROR
	default:
		logLevel = logging.INFO
	}

	return logLevel
}

func newLogMessage(level logging.LogLevel) logging.LogMessage {
	return logging.NewLogMessage().
		WithAgent(logging.NewAgent()).
		WithLog(logging.NewLog().
			WithLevel(level))
}

// Log will log the given message
func (l StandardLogger) Log(message logging.LogMessage) {
	if l.Dispatcher == nil {
		return
	}
	if err := l.Dispatch(message); err != nil {
		fmt.Printf("unable to dispatch log message [%s]. error : %s\n", message.Message, err.Error())
	}
}

// LogIf will log the given message only if `shouldLog` argument is `true`
func (l StandardLogger) LogIf(message logging.LogMessage, shouldLog bool) {
	if shouldLog {
		l.Log(message)
	}
}
