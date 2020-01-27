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

package elasticsearch

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/elastic/cloud-sdk-go/pkg/logging"
	logger "github.com/elastic/cloud-sdk-go/pkg/logging/logger/standardoutput"
)

const (
	logIndex = "/logs/_doc"
)

// Logger is the implementation of ES cluster logging
type Logger struct {
	host, user, pass string
}

// Log sends a logging message to an ES cluster
// If logger fails to deliver the message then logs an error message to the standard output logger
func (esLogger Logger) Log(msg logging.LogMessage) error {
	b, err := msg.Marshall()
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", esLogger.host, logIndex), bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.SetBasicAuth(esLogger.user, esLogger.pass)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		msg := logging.NewLogMessage().
			WithLog(
				logging.NewLog().WithLevel(logging.ERROR)).
			WithAgent(
				logging.NewAgent().WithName("cloud-sdk-go")).
			WithMessage(fmt.Sprintf(
				"cannot send logs to ES cluster %s. error code is %d\n detailed error %s",
				esLogger.host,
				resp.StatusCode,
				string(body)))

		// Do we want to check the error here?
		// nolint
		logger.NewStandardOutputDispatcher().Log(msg)
	}

	return nil
}

// WithHost sets the ES host and returns the Logger itself
func (esLogger Logger) WithHost(host string) Logger {
	esLogger.host = host
	return esLogger
}

// WithUser sets the ES user and returns the Logger itself
func (esLogger Logger) WithUser(user string) Logger {
	esLogger.user = user
	return esLogger
}

// WithPass sets the ES password and returns the Logger itself
func (esLogger Logger) WithPass(pass string) Logger {
	esLogger.pass = pass
	return esLogger
}

// NewLogger properly creates a new elasticsearch logger
func NewLogger() Logger {
	return Logger{}
}
