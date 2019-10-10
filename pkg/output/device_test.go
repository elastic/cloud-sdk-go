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
	"fmt"
	"os"
	"testing"
)

func TestDevice_Write(t *testing.T) {
	type fields struct {
		dev    *bytes.Buffer
		buf    *bytes.Buffer
		paused bool
	}
	type args struct {
		ps [][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		output  string
	}{
		{
			name: "Unpaused device writes directly to the io.Writer",
			fields: fields{
				dev: new(bytes.Buffer),
				buf: new(bytes.Buffer),
			},
			args: args{
				ps: [][]byte{
					[]byte("something\n"),
					[]byte("something else\n"),
					[]byte("something more\n"),
				},
			},
			output: "something\nsomething else\nsomething more\n",
		},
		{
			name: "Paused device writes directly to the buffer",
			fields: fields{
				dev:    new(bytes.Buffer),
				buf:    new(bytes.Buffer),
				paused: true,
			},
			args: args{
				ps: [][]byte{
					[]byte("something\n"),
					[]byte("something else\n"),
					[]byte("something more\n"),
				},
			},
			output: "something\nsomething else\nsomething more\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Device{
				writer: tt.fields.dev,
				buf:    tt.fields.buf,
				paused: tt.fields.paused,
			}
			for _, p := range tt.args.ps {
				if _, err := d.Write(p); (err != nil) != tt.wantErr {
					t.Errorf("Device.Write() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}

			var got = tt.fields.dev.String()
			if tt.fields.paused {
				got = tt.fields.buf.String()
			}

			if got != tt.output {
				t.Errorf("Device.Write() output = %v, wantErr %v", got, tt.output)
			}
		})
	}
}

func ExampleDevice() {
	var device = NewDevice(os.Stdout)

	device.Write([]byte("something\n"))
	device.Write([]byte("something unpaused\n"))

	device.Pause()

	device.Write([]byte("something paused\n"))
	device.Write([]byte("something paused 2\n"))
	device.Write([]byte("something paused 3\n"))

	fmt.Println("== paused device ==")

	device.Resume()

	// Output: something
	// something unpaused
	// == paused device ==
	// something paused
	// something paused 2
	// something paused 3
}

func Example() {
	var device = NewDevice(os.Stdout)

	fmt.Fprintln(device, "something")
	fmt.Fprintln(device, "something unpaused")

	device.Pause()

	fmt.Fprintln(device, "something paused")
	fmt.Fprintln(device, "something paused 2")
	fmt.Fprintln(device, "something paused 3")

	fmt.Println("== paused device ==")

	device.Resume()

	// Output: something
	// something unpaused
	// == paused device ==
	// something paused
	// something paused 2
	// something paused 3
}
