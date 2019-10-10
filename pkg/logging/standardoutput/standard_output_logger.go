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
	case "TRACE":
		return scheme.Trace
	case "DEBUG":
		return scheme.Debug
	case "INFO":
		return scheme.Info
	case "WARN":
		return scheme.Warn
	case "ERROR":
		return scheme.Error
	default:
		return color.FgWhite
	}
}
