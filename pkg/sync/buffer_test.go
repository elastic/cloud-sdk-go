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
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestBufferWrite(t *testing.T) {
	var validationChan = make(chan struct{})
	type fields struct {
		buf *bytes.Buffer
		f   []func()
	}
	type args struct {
		p         []byte
		emptyChan func()
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "Write succeeds",
			fields: fields{
				buf: new(bytes.Buffer),
			},
			args: args{
				p: []byte("something to write there"),
			},
			wantN: len([]byte("something to write there")),
		},
		{
			name: "Write succeeds (with callback)",
			fields: fields{
				buf: new(bytes.Buffer),
				f: []func(){
					func() { validationChan <- struct{}{} },
				},
			},
			args: args{
				p:         []byte("something to write there"),
				emptyChan: func() { <-validationChan },
			},
			wantN: len([]byte("something to write there")),
		},
		{
			name: "Write succeeds (with callbacks)",
			fields: fields{
				buf: new(bytes.Buffer),
				f: []func(){
					func() { validationChan <- struct{}{} },
					func() { validationChan <- struct{}{} },
				},
			},
			args: args{
				p: []byte("something to write there"),
				emptyChan: func() {
					<-validationChan
					<-validationChan
				},
			},
			wantN: len([]byte("something to write there")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &Buffer{
				buf: tt.fields.buf,
				f:   tt.fields.f,
			}

			if tt.args.emptyChan != nil {
				go tt.args.emptyChan()
			}

			gotN, err := buf.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buffer.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Buffer.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestBufferRead(t *testing.T) {
	type fields struct {
		buf *bytes.Buffer
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "Read succeeds",
			fields: fields{
				buf: func() *bytes.Buffer {
					buf := new(bytes.Buffer)
					buf.WriteString("something will be read here")
					return buf
				}(),
			},
			args: args{
				p: []byte("something will be read here"),
			},
			wantN: len([]byte("something will be read here")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &Buffer{
				buf: tt.fields.buf,
			}
			gotN, err := buf.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Buffer.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Buffer.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestBufferString(t *testing.T) {
	type fields struct {
		buf *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "String returns the contents of the buffer",
			fields: fields{
				buf: func() *bytes.Buffer {
					buf := new(bytes.Buffer)
					buf.WriteString("something will be read here")
					return buf
				}(),
			},
			want: "something will be read here",
		},
		{
			name: "String returns the contents of the buffer",
			fields: fields{
				buf: func() *bytes.Buffer {
					buf := new(bytes.Buffer)
					buf.WriteString("something will be also read here")
					return buf
				}(),
			},
			want: "something will be also read here",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &Buffer{
				buf: tt.fields.buf,
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("Buffer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBufferConcurrentAccess(t *testing.T) {
	type fields struct {
		buf *bytes.Buffer
	}
	type args struct {
		N uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Does not panic with 2 goroutines",
			fields: fields{buf: new(bytes.Buffer)},
			args:   args{N: 2},
		},
		{
			name:   "Does not panic with 10 goroutines",
			fields: fields{buf: new(bytes.Buffer)},
			args:   args{N: 10},
		},
		{
			name:   "Does not panic with 100 goroutines",
			fields: fields{buf: new(bytes.Buffer)},
			args:   args{N: 100},
		},
		{
			name:   "Does not panic with 1000 goroutines",
			fields: fields{buf: new(bytes.Buffer)},
			args:   args{N: 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &Buffer{
				buf: tt.fields.buf,
			}

			var wg sync.WaitGroup
			wg.Add(int(tt.args.N))
			for index := 0; uint(index) < tt.args.N; index++ {
				// Pseudorandom is fine here.
				ruint := time.Duration(
					rand.Intn(int(time.Millisecond)),
				)
				go time.AfterFunc(ruint, func() {
					buf.Write([]byte(
						strconv.Itoa(int(ruint))),
					)
					wg.Done()
				})
			}

			wg.Wait()
		})
	}
}
