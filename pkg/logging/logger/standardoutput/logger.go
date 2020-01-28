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

	"github.com/fatih/color"

	"github.com/elastic/cloud-sdk-go/pkg/logging"
	loggingdecorator "github.com/elastic/cloud-sdk-go/pkg/logging/decorator"
)

var (
	// Default format of log message
	defFmt = "[%s] [%s:%s] %s\n"
)

// StandardOutput is the implementation of standard output logging
type StandardOutput struct {
	format string
	*loggingdecorator.ColoringScheme
}

// Log writes the given log message to the standard output using the logger defined format
func (logger StandardOutput) Log(msg logging.LogMessage) error {
	color.Set(logging.LevelColor(*logger.ColoringScheme, msg.Log.Level))
	fmt.Printf(logger.format, msg.Timestamp, msg.Agent.Name, msg.Log.Level, msg.Message)
	color.Unset()
	return nil
}

// WithFormat sets the logger format and returns the logger itself
func (logger *StandardOutput) WithFormat(format string) *StandardOutput {
	logger.format = format
	return logger
}

// WithColoringScheme sets the coloring scheme and returns the logger itself
func (logger *StandardOutput) WithColoringScheme(scheme *loggingdecorator.ColoringScheme) *StandardOutput {
	if scheme == nil {
		logger.ColoringScheme = loggingdecorator.DefaultColoringScheme()
	} else {
		logger.ColoringScheme = scheme
	}
	return logger
}

// New properly creates a new standard output logger initializing its internal state with default values
func New() *StandardOutput {
	return &StandardOutput{
		format:         defFmt,
		ColoringScheme: loggingdecorator.DefaultColoringScheme(),
	}
}
