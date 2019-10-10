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

package input

import (
	"fmt"
	"io"
	"sync"

	"github.com/elastic/cloud-sdk-go/pkg/output"
)

// Scanner masks an io.Reader and an output.Device so that the output.Device
// can be paused before the Scanners input can be parsed.
type Scanner struct {
	reader io.Reader
	writer *output.Device

	mu sync.Mutex
}

// NewScanner instantiates a new Scanner from an io.Reader and output.Device.
func NewScanner(reader io.Reader, device *output.Device) *Scanner {
	return &Scanner{
		reader: reader,
		writer: device,
	}
}

// Scan scans the input that Scanner is wrapping first printing a statement to
// the output.Device so that the user can understand what it has to type. Only
// one concurrent scan is allowed to happen at any time in the same Scanner.
func (s *Scanner) Scan(text string) string {
	s.mu.Lock()
	s.writer.Pause()
	defer func() {
		// nolint
		s.writer.Resume()
		s.mu.Unlock()
	}()

	fmt.Fprint(s.writer.Writer(), text)

	var result string
	var err error
	_, err = fmt.Fscanln(s.reader, &result)
	for err != nil && result == "" {
		if err.Error() == "unexpected newline" {
			fmt.Fprint(s.writer.Writer(), text)
		}
		_, err = fmt.Fscanln(s.reader, &result)
	}

	return result
}
