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

package loggingdecorator

import "github.com/fatih/color"

// ColoringScheme describes the colors used by the standard output logger
type ColoringScheme struct {
	Trace color.Attribute
	Debug color.Attribute
	Info  color.Attribute
	Warn  color.Attribute
	Error color.Attribute
}

// DefaultColoringScheme returns the default standard output logger coloring scheme
func DefaultColoringScheme() *ColoringScheme {
	return &ColoringScheme{
		Trace: color.FgWhite,
		Debug: color.FgWhite,
		Info:  color.FgWhite,
		Warn:  color.FgYellow,
		Error: color.FgRed,
	}
}
