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

package output

import (
	"bytes"
	"io"
	"sync"
)

// Device masks an io.Writer with an intermediary buffer that receives
// the writes while the device is paused. Once the device is unpaused
// the masked io.Writer will receive the buffered writes.
type Device struct {
	writer io.Writer
	buf    *bytes.Buffer

	paused bool
	mu     sync.Mutex
}

// NewDevice instantiates a new Device from an io.Writer
func NewDevice(device io.Writer) *Device {
	return &Device{
		writer: device,
		buf:    new(bytes.Buffer),
	}
}

// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// If the device is paused, instead of writing to the actual device
// it writes to an intermediary buffer that will hold any writes
// while the device is paused. If the device is paused for too long
// and causes the
func (d *Device) Write(p []byte) (n int, err error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.paused {
		defer func() {
			// Recover from panic if the buffer becomes too large
			// nolint
			recover()
		}()
		return d.buf.Write(p)
	}
	return d.writer.Write(p)
}

// Pause pauses writes to the masked device and writes to an
// intermediary buffer.
func (d *Device) Pause() {
	defer d.mu.Unlock()
	d.mu.Lock()
	d.paused = true
}

// Resume copies the contents of the intermediary buffer to the device
// that is being masked.
func (d *Device) Resume() (n int64, err error) {
	defer func() {
		d.buf.Reset()
		d.mu.Unlock()
	}()
	d.mu.Lock()
	d.paused = false
	return io.Copy(d.writer, d.buf)
}

// Writer returns the io.Writer that Device is wrapping.
func (d *Device) Writer() io.Writer {
	return d.writer
}
