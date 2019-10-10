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
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/output"
)

func TestScannerScan(t *testing.T) {
	type fields struct {
		reader io.Reader
		writer *output.Device
	}
	type args struct {
		text string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult string
	}{
		{
			name: "scans text",
			fields: fields{
				reader: strings.NewReader("something\n"),
				writer: output.NewDevice(new(bytes.Buffer)),
			},
			wantResult: "something",
		},
		{
			name: "obtains the first line of text",
			fields: fields{
				reader: strings.NewReader("something\nsomethingelse"),
				writer: output.NewDevice(new(bytes.Buffer)),
			},
			wantResult: "something",
		},
		{
			name: "obtains the first word of text",
			fields: fields{
				reader: strings.NewReader("something something wut\nsomething"),
				writer: output.NewDevice(new(bytes.Buffer)),
			},
			wantResult: "something",
		},
		{
			name: "ignores the empty lines",
			fields: fields{
				reader: strings.NewReader("\n\n\nigotthistext"),
				writer: output.NewDevice(new(bytes.Buffer)),
			},
			wantResult: "igotthistext",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scanner{
				reader: tt.fields.reader,
				writer: tt.fields.writer,
			}
			if gotResult := s.Scan(tt.args.text); gotResult != tt.wantResult {
				t.Errorf("Scanner.Scan() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestScannerFull(t *testing.T) {
	// SIMULATES STDOUT; WRAPPED IN output.NewDevice
	var out = &ConcBuf{buf: new(bytes.Buffer)}

	// THIS IS THE DEVICE WHERE WE WRITE
	var outputDevice = output.NewDevice(out)

	// We wrap it like this to avoid races in reads & writes.
	var inputDevice = &Stdin{
		in:  &ConcBuf{buf: new(bytes.Buffer)},
		out: out,
	}

	var scanner = NewScanner(inputDevice, outputDevice)
	var want = `
Yes or no [y/N]: y
somehing
somehing
somehing
Yes or no [y/N]: 
Yes or no [y/N]: n
output is paused
we'll see this lines after the yes/no question
`[1:]

	// string channel used to see when the Scan call is finished.
	c := make(chan string)
	// Signal channel used to wait until Scan is scheduled.
	s := make(chan bool)

	go func() {
		s <- true
		c <- scanner.Scan("Yes or no [y/N]: ")
	}()
	<-s
	<-time.After(time.Millisecond)
	fmt.Fprintln(outputDevice, "somehing")
	fmt.Fprintln(outputDevice, "somehing")
	fmt.Fprintln(outputDevice, "somehing")
	fmt.Fprintln(inputDevice, "y")
	<-c

	go func() {
		s <- true
		c <- scanner.Scan("Yes or no [y/N]: ")
	}()
	<-s

	<-time.After(time.Millisecond)
	fmt.Fprintln(inputDevice, "")

	fmt.Fprintln(outputDevice, "output is paused")
	fmt.Fprintln(outputDevice, "we'll see this lines after the yes/no question")

	<-time.After(time.Millisecond * 2)
	fmt.Fprintln(inputDevice, "n")
	<-c

	if out.String() != want {
		t.Errorf("Output = %v, want %v", out, want)
	}
}

// TSBuf simulates the standard input, any write that it receives will
// be automatically written to the out device.
type Stdin struct {
	in  io.ReadWriter
	out io.Writer
}

func (in *Stdin) Write(p []byte) (n int, err error) {
	in.out.Write(p)
	return in.in.Write(p)
}

func (in *Stdin) Read(p []byte) (n int, err error) {
	return in.in.Read(p)
}

type ConcBuf struct {
	buf *bytes.Buffer
	mu  sync.RWMutex
}

func (buf *ConcBuf) Write(p []byte) (n int, err error) {
	buf.mu.Lock()
	defer buf.mu.Unlock()
	return buf.buf.Write(p)
}

func (buf *ConcBuf) Read(p []byte) (n int, err error) {
	buf.mu.RLock()
	defer buf.mu.RUnlock()
	return buf.buf.Read(p)
}

func (buf *ConcBuf) String() string {
	buf.mu.RLock()
	defer buf.mu.RUnlock()
	return buf.buf.String()
}
