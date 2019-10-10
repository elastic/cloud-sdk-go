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

// Contains checks if a logger slice contains an element
func contains(s []Logger, e Logger) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}
