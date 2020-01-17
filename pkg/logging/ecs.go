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
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-openapi/strfmt"
)

// LogMessage is the top-level structure
// The base set contains all fields which are on the top level. These fields are common across all types of events.
type LogMessage struct {
	Timestamp strfmt.DateTime   `json:"@timestamp,omitempty"`
	Tags      []string          `json:"tags,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
	Message   string            `json:"message,omitempty"`
	Agent     Agent             `json:"agent,omitempty"`
	Error     Err               `json:"error,omitempty"`
	HTTP      HTTP              `json:"http,omitempty"`
	Log       Log               `json:"log,omitempty"`
	ECS       ECS               `json:"ecs,omitempty"`
}

// NewLogMessage creates a new LogMessage and properly initializes all required fields
func NewLogMessage() LogMessage {
	msg := LogMessage{}
	msg.Agent = NewAgent()
	msg.Error = NewError()
	msg.HTTP = NewHTTP()
	msg.Log = NewLog()
	msg.ECS = NewECS()
	msg.Timestamp = strfmt.DateTime(time.Now())
	return msg
}

// WithTags sets the tags field and returns the message itself
func (msg LogMessage) WithTags(tags []string) LogMessage {
	msg.Tags = tags
	return msg
}

// WithLabels sets the labels field and returns the message itself
func (msg LogMessage) WithLabels(labels map[string]string) LogMessage {
	msg.Labels = labels
	return msg
}

// WithMessage sets the message field and returns the message itself
func (msg LogMessage) WithMessage(message string) LogMessage {
	msg.Message = message
	return msg
}

// WithAgent sets the agent object field and returns the message itself
func (msg LogMessage) WithAgent(agent Agent) LogMessage {
	msg.Agent = agent
	return msg
}

// WithError sets the agent object field and returns the message itself
func (msg LogMessage) WithError(err Err) LogMessage {
	msg.Error = err
	return msg
}

// WithHTTP sets the http object field and returns the message itself
func (msg LogMessage) WithHTTP(httpInfo HTTP) LogMessage {
	msg.HTTP = httpInfo
	return msg
}

// WithLog sets the log object field and returns the message itself
func (msg LogMessage) WithLog(log Log) LogMessage {
	msg.Log = log
	return msg
}

// Agent is the structure of an agent
type Agent struct {
	Version     string `json:"version,omitempty"`
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
	EphemeralID string `json:"ephemeral_id,omitempty"`
}

// NewAgent creates a new Agent and properly initializes all required fields
func NewAgent() Agent {
	agent := Agent{}
	return agent
}

// WithVersion sets the version field and returns the agent itself
func (agent Agent) WithVersion(version string) Agent {
	agent.Version = version
	return agent
}

// WithName sets the name field and returns the agent itself
func (agent Agent) WithName(name string) Agent {
	agent.Name = name
	return agent
}

// WithID sets the id field and returns the agent itself
func (agent Agent) WithID(id string) Agent {
	agent.ID = id
	return agent
}

// WithEphemeralID sets the ephemeral id field and returns the agent itself
func (agent Agent) WithEphemeralID(id string) Agent {
	agent.EphemeralID = id
	return agent
}

// Err is the structure of an error
// These fields can represent errors of any kind. Use them for errors that happen while fetching events or in cases where the event itself contains an error.
type Err struct {
	Message string `json:"message,omitempty"`
	ID      string `json:"id,omitempty"`
	Code    string `json:"code,omitempty"`
}

// NewError creates a new Error and properly initializes all required fields
func NewError() Err {
	err := Err{}
	return err
}

// WithMessage sets the message field and returns the error itself
func (err Err) WithMessage(msg string) Err {
	err.Message = msg
	return err
}

// WithID sets the ID field and returns the error itself
func (err Err) WithID(id string) Err {
	err.ID = id
	return err
}

// WithCode sets the Code field and returns the error itself
func (err Err) WithCode(code string) Err {
	err.Code = code
	return err
}

// HTTP is the structure of an http request and response
// Fields related to HTTP requests and responses.
type HTTP struct {
	Version  string   `json:"version,omitempty"`
	Response Response `json:"response,omitempty"`
	Request  Request  `json:"request,omitempty"`
}

// NewHTTP creates a new HTTP and properly initializes all required fields
func NewHTTP() HTTP {
	httpInfo := HTTP{}
	httpInfo.Response = NewResponse()
	httpInfo.Request = NewRequest()
	return httpInfo
}

// WithResponse sets the response object field and returns the http itself
func (httpInfo HTTP) WithResponse(resp Response) HTTP {
	httpInfo.Response = resp
	return httpInfo
}

// WithRequest sets the request object field and returns the http itself
func (httpInfo HTTP) WithRequest(req Request) HTTP {
	httpInfo.Request = req
	return httpInfo
}

// WithVersion sets the version field and returns the http itself
func (httpInfo HTTP) WithVersion(version string) HTTP {
	httpInfo.Version = version
	return httpInfo
}

// Request is the structure of an http request
type Request struct {
	Method string `json:"method,omitempty"`
}

// NewRequest creates a new Request and properly initializes all required fields
func NewRequest() Request {
	request := Request{}
	request.Method = http.MethodGet
	return request
}

// WithMethod sets the request method field and returns the request itself
func (req Request) WithMethod(method string) Request {
	req.Method = method
	return req
}

// Response is the structure of an http response
type Response struct {
	Body       string `json:"method,omitempty"`
	StatusCode uint16 `json:"status_code,omitempty"`
}

// NewResponse creates a new Response and properly initializes all required fields
func NewResponse() Response {
	resp := Response{}
	return resp
}

// WithBody sets the response body field and returns the request itself
func (resp Response) WithBody(body string) Response {
	resp.Body = body
	return resp
}

// WithStatusCode sets the response status code field and returns the request itself
func (resp Response) WithStatusCode(sc uint16) Response {
	resp.StatusCode = sc
	return resp
}

// Log is the structure of a Log event
type Log struct {
	Level    string `json:"level,omitempty"`
	Original string `json:"original,omitempty"`
	loglevel LogLevel
	Line     uint32 `json:"line,omitempty"`
	Offset   uint32 `json:"offset,omitempty"`
}

// NewLog creates a new Log and properly initializes all required fields
func NewLog() Log {
	log := Log{}
	return log
}

// WithLevel sets the log level field and returns the log itself
func (log Log) WithLevel(level LogLevel) Log {
	log.Level = level.toString()
	log.loglevel = level
	return log
}

// WithLine sets the log line field and returns the log itself
func (log Log) WithLine(line uint32) Log {
	log.Line = line
	return log
}

// WithOriginal sets the log original message field and returns the log itself
func (log Log) WithOriginal(original string) Log {
	log.Original = original
	return log
}

// WithOffset sets the log offset field and returns the log itself
func (log Log) WithOffset(offset uint32) Log {
	log.Offset = offset
	return log
}

// ECS is the structure of Meta-information specific to ECS
type ECS struct {
	Version string `json:"version,omitempty"`
}

// NewECS creates a new ECS and properly initializes all required fields
func NewECS() ECS {
	return ECS{
		Version: "0.1.0",
	}
}

// Returns the log Level as string
func (level LogLevel) toString() string {
	logLevels := [...]string{
		"ERROR",
		"WARN",
		"INFO",
		"DEBUG",
		"TRACE",
	}
	return logLevels[level-1]
}

// Marshall returns the json representation of a LogMessage
func (msg *LogMessage) Marshall() ([]byte, error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return b, nil
}
