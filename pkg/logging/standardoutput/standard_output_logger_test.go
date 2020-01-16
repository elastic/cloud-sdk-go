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
	"reflect"
	"testing"

	"github.com/fatih/color"

	"github.com/elastic/cloud-sdk-go/pkg/logging"
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
func Test_levelColor(t *testing.T) {
	scheme := loggingdecorator.ColoringScheme{
		Trace: color.FgWhite,
		Debug: color.FgYellow,
		Info:  color.FgHiBlue,
		Warn:  color.FgRed,
		Error: color.FgCyan,
	}

	tests := []struct {
		name     string
		expColor color.Attribute
		logLevel string
		scheme   loggingdecorator.ColoringScheme
	}{
		{
			name:     "levelColor should return the correct color for TRACE log level ",
			expColor: color.FgWhite,
			logLevel: "trace",
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the correct color for DEBUG log level ",
			expColor: color.FgYellow,
			logLevel: "debug",
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the correct color for INFO log level ",
			expColor: color.FgHiBlue,
			logLevel: "info",
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the correct color for WARN log level ",
			expColor: color.FgRed,
			logLevel: "warn",
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the correct color for ERROR log level ",
			expColor: color.FgCyan,
			logLevel: "error",
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the default color for unknown log level ",
			expColor: color.FgWhite,
			logLevel: "unknown",
			scheme:   scheme,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logColor := levelColor(tt.scheme, tt.logLevel)

			if logColor != tt.expColor {
				t.Errorf("levelColor() wanted %v returned %v", tt.expColor, logColor)
				return
			}
		})
	}
}

func TestDebug(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{
			name: "Log debug should properly create a debug message",
			msg:  "Log message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logMessage := Debug(tt.msg)

			if logMessage.Message != tt.msg {
				t.Errorf("Debug() logMessage.Message expected %s actual %s", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Log.Level != "DEBUG" {
				t.Errorf("Debug() logMessage.Log.Level expected %s actual %s", "DEBUG", logMessage.Log.Level)
				return
			}
			if !reflect.DeepEqual(logMessage.Agent, agent) {
				t.Errorf("Debug() logMessage.Agent expected %v actual %v", agent, logMessage.Agent)
				return
			}
		})
	}
}

func TestInfo(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{
			name: "Log info should properly create a info message",
			msg:  "Log message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logMessage := Info(tt.msg)

			if logMessage.Message != tt.msg {
				t.Errorf("Info() logMessage.Message expected %s actual %s", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Log.Level != "INFO" {
				t.Errorf("Info() logMessage.Log.Level expected %s actual %s", "INFO", logMessage.Log.Level)
				return
			}
			if !reflect.DeepEqual(logMessage.Agent, agent) {
				t.Errorf("Info() logMessage.Agent expected %v actual %v", agent, logMessage.Agent)
				return
			}
		})
	}
}

func TestWarn(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{
			name: "Log warn should properly create a warn message",
			msg:  "Log message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logMessage := Warn(tt.msg)

			if logMessage.Message != tt.msg {
				t.Errorf("Warn() logMessage.Message expected %s actual %s", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Log.Level != "WARN" {
				t.Errorf("Warn() logMessage.Log.Level expected %s actual %s", "WARN", logMessage.Log.Level)
				return
			}
			if !reflect.DeepEqual(logMessage.Agent, agent) {
				t.Errorf("Warn() logMessage.Agent expected %v actual %v", agent, logMessage.Agent)
				return
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name string
		msg  string
	}{
		{
			name: "Log error should properly create an error message",
			msg:  "Log message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logMessage := Error(tt.msg)

			if logMessage.Message != tt.msg {
				t.Errorf("Error() logMessage.Message expected %s actual %s", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Error.Message != tt.msg {
				t.Errorf("Error() logMessage.Message expected %s actual %s", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Log.Level != "ERROR" {
				t.Errorf("Error() logMessage.Log.Level expected %s actual %s", "ERROR", logMessage.Log.Level)
				return
			}
			if !reflect.DeepEqual(logMessage.Agent, agent) {
				t.Errorf("Error() logMessage.Agent expected %v actual %v", agent, logMessage.Agent)
				return
			}
		})
	}
}

func TestLevel(t *testing.T) {
	tests := []struct {
		name        string
		level       string
		expLoglevel logging.LogLevel
	}{
		{
			name:        "should return logging.TRACE",
			level:       TRACE,
			expLoglevel: logging.TRACE,
		},
		{
			name:        "should return logging.DEBUG",
			level:       DEBUG,
			expLoglevel: logging.DEBUG,
		},
		{
			name:        "should return logging.INFO",
			level:       INFO,
			expLoglevel: logging.INFO,
		},
		{
			name:        "should return logging.WARNING",
			level:       WARN,
			expLoglevel: logging.WARNING,
		},
		{
			name:        "should return logging.ERROR",
			level:       ERROR,
			expLoglevel: logging.ERROR,
		},
		{
			name:        "should return the default level",
			level:       "UNKNOWN",
			expLoglevel: logging.INFO,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expLoglevel != Level(tt.level) {
				t.Errorf("Level() expected %v but actual %s", tt.expLoglevel, tt.level)
				return
			}
		})
	}
}

func TestStandardLogger_Log(t *testing.T) {
	tests := []struct {
		name    string
		logger  StandardLogger
		message string
	}{
		{
			name:    "should log message",
			logger:  StandardLogger{Dispatcher: logging.NewMockDispatcher()},
			message: "test message",
		},
		{
			name:   "should not log message",
			logger: StandardLogger{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Log(Info(tt.message))
			logging.AssertLoggedMessages(t, tt.logger.Dispatcher, []string{tt.message})
		})
	}
}

func TestStandardLogger_LogIf(t *testing.T) {
	tests := []struct {
		name        string
		logger      StandardLogger
		message     string
		shouldLog   bool
		expMessages []string
	}{
		{
			name:        "should log message",
			logger:      StandardLogger{Dispatcher: logging.NewMockDispatcher()},
			message:     "test message",
			shouldLog:   true,
			expMessages: []string{"test message"},
		},
		{
			name:    "should not log message",
			logger:  StandardLogger{Dispatcher: logging.NewMockDispatcher()},
			message: "test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.LogIf(Info(tt.message), tt.shouldLog)
			logging.AssertLoggedMessages(t, tt.logger.Dispatcher, tt.expMessages)
		})
	}
}

var agent = logging.NewAgent()
