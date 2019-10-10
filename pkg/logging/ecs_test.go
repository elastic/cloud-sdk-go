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
	"fmt"
	"net/http"
	"testing"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/elastic/cloud-sdk-go/pkg/util/slice"
)

func TestLogMessage_Marshall(t *testing.T) {
	var msg = NewLogMessage()
	msg.Agent.Name = "soteria"

	tests := []struct {
		name string
		msg  LogMessage
	}{
		{
			name: "Log Level string should return the correct description if log Level is ERROR",
			msg:  msg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := msg.Marshall()

			println(string(bytes))

			if err != nil {
				t.Errorf("Marshall() wanted no error returned %s instead", err.Error())
				return
			}
		})
	}
}

func TestLogLevelString(t *testing.T) {
	tests := []struct {
		name      string
		logLevel  LogLevel
		wantLevel string
	}{
		{
			name:      "Log Level string should return the correct description if log Level is ERROR",
			logLevel:  ERROR,
			wantLevel: "ERROR",
		},
		{
			name:      "Log Level string should return the correct description if log Level is WARNING",
			logLevel:  WARNING,
			wantLevel: "WARN",
		},
		{
			name:      "Log Level string should return the correct description if log Level is INFO",
			logLevel:  INFO,
			wantLevel: "INFO",
		},
		{
			name:      "Log Level string should return the correct description if log Level is DEBUG",
			logLevel:  DEBUG,
			wantLevel: "DEBUG",
		},
		{
			name:      "Log Level string should return the correct description if log Level is TRACE",
			logLevel:  TRACE,
			wantLevel: "TRACE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel := tt.logLevel.toString()

			if logLevel != tt.wantLevel {
				t.Errorf("LogMessage.logLevelString() wanted %s returned %s instead", tt.wantLevel, logLevel)
				return
			}
		})
	}
}

func TestNewAgent(t *testing.T) {
	var (
		name = "agent"
		id   = "id"
		eid  = "eid"
		ver  = "version"
	)

	tests := []struct {
		name string
	}{
		{
			name: "Agent construction should properly set fields",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := NewAgent().WithEphemeralID(eid).WithID(id).WithName(name).WithVersion(ver)
			err := &multierror.Error{}
			if agent.EphemeralID != eid {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted EphemeralID %s but returned %s\n", eid, agent.EphemeralID)))
			}
			if agent.ID != id {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted ID %s but returned %s\n", id, agent.ID)))
			}
			if agent.Version != ver {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted Version %s but returned %s\n", ver, agent.Version)))
			}
			if agent.Name != name {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted Name %s but returned %s\n", name, agent.Name)))
			}

			if len(err.Errors) > 0 {
				t.Error(err.Error())
				return
			}
		})
	}
}

func TestNewError(t *testing.T) {
	var (
		msg  = "msg"
		id   = "id"
		code = "code"
	)

	tests := []struct {
		name string
	}{
		{
			name: "Error construction should properly set fields",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewError().WithID(id).WithCode(code).WithMessage(msg)
			err := &multierror.Error{}
			if e.Code != code {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted Code %s but returned %s\n", code, e.Code)))
			}
			if e.ID != id {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted ID %s but returned %s\n", id, e.ID)))
			}
			if e.Message != msg {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted Message %s but returned %s\n", msg, e.Message)))
			}

			if len(err.Errors) > 0 {
				t.Error(err.Error())
				return
			}
		})
	}
}

func TestNewLog(t *testing.T) {
	var (
		line     = uint32(200)
		offset   = uint32(10)
		original = "original msg"
	)

	tests := []struct {
		name string
	}{
		{
			name: "Log construction should properly set fields",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := NewLog().WithLine(line).WithLevel(ERROR).WithOriginal(original).WithOffset(offset)
			err := &multierror.Error{}
			if log.Line != line {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted Line %d but returned %d\n", line, log.Line)))
			}
			if log.Offset != offset {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted Offset %d but returned %d\n", offset, log.Offset)))
			}
			if log.Original != original {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted Original %s but returned %s\n", original, log.Original)))
			}
			if log.loglevel != ERROR {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted logLevel %d but returned %d\n", ERROR, log.loglevel)))
			}
			if log.Level != "ERROR" {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted Level ERROR but returned %s\n", log.Level)))
			}

			if len(err.Errors) > 0 {
				t.Error(err.Error())
				return
			}
		})
	}
}

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name         string
		req          Request
		wantedMethod string
	}{
		{
			name:         "Request construction should properly set default fields",
			req:          NewRequest(),
			wantedMethod: http.MethodGet,
		},
		{
			name:         "Request construction should properly set fields",
			req:          NewRequest().WithMethod(http.MethodDelete),
			wantedMethod: http.MethodDelete,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &multierror.Error{}
			if tt.req.Method != tt.wantedMethod {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted request method %v but returned %v\n", tt.wantedMethod, tt.req.Method)))
			}

			if len(err.Errors) > 0 {
				t.Error(err.Error())
				return
			}
		})
	}
}

func TestNewResponse(t *testing.T) {
	tests := []struct {
		name             string
		resp             Response
		wantedStatusCode uint16
		wantedBody       string
	}{
		{
			name:             "Response construction should properly set fields",
			resp:             NewResponse().WithStatusCode(http.StatusAccepted).WithBody("Response"),
			wantedStatusCode: http.StatusAccepted,
			wantedBody:       "Response",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &multierror.Error{}
			if tt.resp.StatusCode != tt.wantedStatusCode {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted response status code %d but returned %d\n", tt.wantedStatusCode, tt.resp.StatusCode)))
			}

			if tt.resp.Body != tt.wantedBody {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted responsebody %v but returned %v\n", tt.wantedBody, tt.resp.Body)))
			}

			if len(err.Errors) > 0 {
				t.Error(err.Error())
				return
			}
		})
	}
}

func TestNewHttp(t *testing.T) {
	var req = NewRequest().WithMethod(http.MethodPost)
	var resp = NewResponse().WithStatusCode(http.StatusAccepted)
	var version = "version"

	tests := []struct {
		name                string
		http                HTTP
		wantedRequestMethod string
		wantedResponseSc    uint16
		wantedVersion       string
	}{
		{
			name:                "HTTP construction should properly set objects",
			http:                NewHTTP().WithRequest(req).WithResponse(resp).WithVersion(version),
			wantedRequestMethod: http.MethodPost,
			wantedResponseSc:    http.StatusAccepted,
			wantedVersion:       version,
		},
		{
			name:                "HTTP construction should create the default object",
			http:                NewHTTP(),
			wantedRequestMethod: http.MethodGet,
			wantedResponseSc:    0,
			wantedVersion:       "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &multierror.Error{}
			if tt.wantedRequestMethod != tt.http.Request.Method {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted request method %v but returned %v\n", tt.wantedRequestMethod, tt.http.Request.Method)))
			}

			if tt.wantedResponseSc != tt.http.Response.StatusCode {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted response status code %d but returned %d\n", tt.wantedResponseSc, tt.http.Response.StatusCode)))
			}

			if tt.wantedVersion != tt.http.Version {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted http version %v but returned %v\n", tt.wantedVersion, tt.http.Version)))
			}

			if len(err.Errors) > 0 {
				t.Error(err.Error())
				return
			}
		})
	}
}

func TestNewLogMessage(t *testing.T) {
	var message = "message"
	var req = NewRequest().WithMethod(http.MethodPost)
	var resp = NewResponse().WithStatusCode(http.StatusAccepted)
	var agent = NewAgent().WithID("soteria")
	var err = NewError().WithMessage(message)
	var log = NewLog().WithLine(uint32(100))
	var httpInfo = NewHTTP().WithResponse(resp).WithRequest(req)
	var ecs = NewECS()
	var tags = []string{"tags"}
	var labels = map[string]string{
		"labelA": "valueA",
		"labelB": "valueB",
	}

	tests := []struct {
		name          string
		logMsg        LogMessage
		wantedAgent   Agent
		wantedError   Error
		wantedLog     Log
		wantedHTTP    HTTP
		wantedMessage string
		wantedTags    []string
		wantedLabels  map[string]string
		wantedECS     ECS
	}{
		{
			name: "log message construction should properly set objects and fields",
			logMsg: NewLogMessage().
				WithHTTP(httpInfo).
				WithAgent(agent).
				WithError(err).
				WithLog(log).
				WithMessage(message).
				WithTags(tags).
				WithLabels(labels),
			wantedAgent:   agent,
			wantedError:   err,
			wantedHTTP:    httpInfo,
			wantedLog:     log,
			wantedMessage: message,
			wantedTags:    tags,
			wantedLabels:  labels,
			wantedECS:     ecs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &multierror.Error{}
			if tt.wantedAgent != tt.logMsg.Agent {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted agent %v but returned %v\n", tt.wantedAgent, tt.logMsg.Agent)))
			}
			if tt.wantedError != tt.logMsg.Error {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted error %v but returned %v\n", tt.wantedError, tt.logMsg.Error)))
			}
			if tt.wantedLog != tt.logMsg.Log {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted log %v but returned %v\n", tt.wantedLog, tt.logMsg.Log)))
			}
			if tt.wantedHTTP != tt.logMsg.HTTP {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted http %v but returned %v\n", tt.wantedHTTP, tt.logMsg.HTTP)))
			}
			if tt.wantedECS != tt.logMsg.ECS {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted ECS %v but returned %v\n", tt.wantedHTTP, tt.logMsg.HTTP)))
			}
			if slice.ContainsAll(tt.logMsg.Tags, tt.wantedTags) && len(tt.wantedTags) == len(tt.logMsg.Labels) {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted tags %v but returned %v\n", tt.wantedTags, tt.logMsg.Tags)))
			}
			for k, v := range tt.wantedLabels {
				actualValue, ok := tt.logMsg.Labels[k]
				if !ok || actualValue != v {
					err = multierror.Append(err, errors.New(fmt.Sprintf("wanted label %v with value %v but got %v\n", k, v, actualValue)))
				}
			}
			if tt.wantedMessage != tt.logMsg.Message {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted tags %v but returned %v\n", tt.wantedMessage, tt.logMsg.Message)))
			}
			if len(err.Errors) > 0 {
				t.Error(err.Error())
				return
			}
		})
	}
}

func TestNewECS(t *testing.T) {
	tests := []struct {
		name          string
		wantedVersion string
	}{
		{
			name:          "ECS construction should properly set objects and fields",
			wantedVersion: "0.1.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newECS := NewECS()
			err := &multierror.Error{}
			if tt.wantedVersion != newECS.Version {
				err = multierror.Append(err, errors.New(fmt.Sprintf("wanted ECS version %v but got %v\n", tt.wantedVersion, newECS.Version)))
			}
			if len(err.Errors) > 0 {
				t.Error(err.Error())
				return
			}
		})
	}
}
