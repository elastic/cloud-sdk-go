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

package sync

import (
	"bytes"
	"sync"
)

// NewBuffer constructs a new concurrently safe byte buffer. Accepts multiple
// functions to be sent that will be called on each Write.
func NewBuffer(f ...func()) *Buffer {
	return &Buffer{buf: new(bytes.Buffer), f: f}
}

// Buffer is a partial bytes.Buffer implementation that complies both with
// the io.ReadWriter interface
type Buffer struct {
	buf *bytes.Buffer
	f   []func()

	mu sync.RWMutex
}

func (buf *Buffer) Write(p []byte) (n int, err error) {
	buf.mu.Lock()
	defer buf.mu.Unlock()
	if buf.f != nil {
		for _, cb := range buf.f {
			cb()
		}
	}
	return buf.buf.Write(p)
}

func (buf *Buffer) Read(p []byte) (n int, err error) {
	buf.mu.RLock()
	defer buf.mu.RUnlock()
	return buf.buf.Read(p)
}

func (buf *Buffer) String() string {
	buf.mu.RLock()
	defer buf.mu.RUnlock()
	return buf.buf.String()
}
