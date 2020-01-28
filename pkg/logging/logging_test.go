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
	"errors"
	"reflect"
	"testing"

	"github.com/fatih/color"

	loggingdecorator "github.com/elastic/cloud-sdk-go/pkg/logging/decorator"
)

func TestLogDispatcher_Add(t *testing.T) {
	tests := []struct {
		name   string
		logger *Mock
	}{
		{
			name:   "Add Logger should correctly add the logger to the dispatcher",
			logger: &Mock{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ld := MuxDispatcher{}
			ml := &Mock{}
			ld.Add(ml)

			if !contains(ld.loggers, ml) {
				t.Errorf("LoggerDispatcher.Add() failed to add logger. %v logger expected to be in slice %v", ml, ld.loggers)
				return
			}
		})
	}
}

func TestNewLogDispatcher(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "New log dispatcher should return a properly constructed dispatcher",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ld := NewLogDispatcher()

			if ld.level != INFO {
				t.Errorf("NewLogDispatcher() log Level wanted %d returned %d", INFO, ld.level)
				return
			}

			if len(ld.loggers) > 0 {
				t.Errorf("NewLogDispatcher() loggers slice should have zero(0) length but returned %d", len(ld.loggers))
				return
			}
		})
	}
}

func TestLogDispatcher_Dispatch(t *testing.T) {
	var msg = NewLogMessage().WithLog(NewLog().WithLevel(ERROR))

	tests := []struct {
		name           string
		wantedMessages []LogMessage
		msg            LogMessage
		err            error
		wantErr        bool
	}{
		{
			name:           "Log dispatcher should not log if message log Level if lower than dispatcher's log Level",
			msg:            NewLogMessage().WithLog(NewLog().WithLevel(TRACE)),
			wantedMessages: []LogMessage{},
			wantErr:        false,
			err:            nil,
		},
		{
			name:           "Log dispatcher should log if message log Level if same or higher than dispatcher's log Level",
			msg:            msg,
			wantedMessages: []LogMessage{msg},
			wantErr:        false,
			err:            nil,
		},
		{
			name:           "Log dispatcher should return err if one logger fails to log",
			msg:            msg,
			wantedMessages: []LogMessage{},
			wantErr:        true,
			err:            errors.New("logging error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ld := NewLogDispatcher().WithLogLevel(ERROR)

			logger := Mock{err: tt.err}
			ld.Add(&logger)
			err := ld.Dispatch(tt.msg)

			if (err != nil) != tt.wantErr {
				t.Errorf("Dispatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.wantedMessages) != len(logger.messagesLogged) {
				t.Errorf("Dispatch() number of logged messages wanted %d but returned %d", len(tt.wantedMessages), len(logger.messagesLogged))
				return
			}

			if len(tt.wantedMessages) > 0 && !reflect.DeepEqual(tt.wantedMessages[0], logger.messagesLogged[0]) {
				t.Errorf("Dispatch() logged message wanted %v but returned %v", tt.wantedMessages[0], logger.messagesLogged[0])
				return
			}
		})
	}
}

func TestLevel(t *testing.T) {
	tests := []struct {
		name        string
		level       string
		expLoglevel LogLevel
	}{
		{
			name:        "should return logging.TRACE",
			level:       trace,
			expLoglevel: TRACE,
		},
		{
			name:        "should return logging.DEBUG",
			level:       debug,
			expLoglevel: DEBUG,
		},
		{
			name:        "should return logging.INFO",
			level:       info,
			expLoglevel: INFO,
		},
		{
			name:        "should return logging.WARNING",
			level:       warn,
			expLoglevel: WARNING,
		},
		{
			name:        "should return logging.ERROR",
			level:       err,
			expLoglevel: ERROR,
		},
		{
			name:        "should return the default level",
			level:       "UNKNOWN",
			expLoglevel: INFO,
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

func TestLevelColor(t *testing.T) {
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
			logLevel: trace,
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the correct color for DEBUG log level ",
			expColor: color.FgYellow,
			logLevel: debug,
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the correct color for INFO log level ",
			expColor: color.FgHiBlue,
			logLevel: info,
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the correct color for WARN log level ",
			expColor: color.FgRed,
			logLevel: warn,
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the correct color for ERROR log level ",
			expColor: color.FgCyan,
			logLevel: err,
			scheme:   scheme,
		},
		{
			name:     "levelColor should return the default color for unknown log level ",
			expColor: color.FgWhite,
			logLevel: "UNKNOWN",
			scheme:   scheme,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logColor := LevelColor(tt.scheme, tt.logLevel)

			if logColor != tt.expColor {
				t.Errorf("levelColor() wanted %v returned %v", tt.expColor, logColor)
				return
			}
		})
	}
}

func TestDebugMessage(t *testing.T) {
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
			logMessage := DebugMessage(tt.msg)

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

func TestWarnMessage(t *testing.T) {
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
			logMessage := WarnMessage(tt.msg)

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

func TestErrorMessage(t *testing.T) {
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
			logMessage := ErrorMessage(tt.msg)

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

func TestTraceMessage(t *testing.T) {
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
			logMessage := TraceMessage(tt.msg)

			if logMessage.Message != tt.msg {
				t.Errorf("Error() logMessage.Message expected %s actual %s", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Error.Message != "" {
				t.Errorf("Error() logMessage.Message expected %s actual %s", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Log.Level != "TRACE" {
				t.Errorf("Error() logMessage.Log.Level expected %s actual %s", "TRACE", logMessage.Log.Level)
				return
			}
			if !reflect.DeepEqual(logMessage.Agent, agent) {
				t.Errorf("Error() logMessage.Agent expected %v actual %v", agent, logMessage.Agent)
				return
			}
		})
	}
}

func TestInfoMessage(t *testing.T) {
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
			logMessage := InfoMessage(tt.msg)

			if logMessage.Message != tt.msg {
				t.Errorf("Error() logMessage.Message expected %s actual %s", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Error.Message != "" {
				t.Errorf("Error() logMessage.Message expected [%s] actual [%s]", tt.msg, logMessage.Message)
				return
			}
			if logMessage.Log.Level != "INFO" {
				t.Errorf("Error() logMessage.Log.Level expected %s actual %s", "INFO", logMessage.Log.Level)
				return
			}
			if !reflect.DeepEqual(logMessage.Agent, agent) {
				t.Errorf("Error() logMessage.Agent expected %v actual %v", agent, logMessage.Agent)
				return
			}
		})
	}
}

// Contains checks if a logger slice contains an element
func contains(s []Logger, e Logger) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

var agent = NewAgent()
