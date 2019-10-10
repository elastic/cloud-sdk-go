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
	"errors"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"syscall"
	"testing"
)

// nolint
func createTempFile(t *testing.T, name string, contents ...[]byte) (*os.File, func()) {
	f, err := ioutil.TempFile("", name)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range contents {
		if _, err := f.Write(c); err != nil {
			t.Fatal(err)
		}
	}
	return f, func() { os.RemoveAll(f.Name()) }
}

func TestNewFileOrReader(t *testing.T) {
	successFile, cleanSuccess := createTempFile(t, "successFile", []byte("something"))
	defer cleanSuccess()

	emptyReader := new(bytes.Buffer)

	type args struct {
		reader   io.Reader
		filename string
	}
	tests := []struct {
		name string
		args args
		want []byte
		err  error
	}{
		{
			name: "Read contents from file name",
			args: args{
				filename: successFile.Name(),
			},
			want: []byte("something"),
		},
		{
			name: "Read contents from reader",
			args: args{
				reader: strings.NewReader("something inside"),
			},
			want: []byte("something inside"),
		},
		{
			name: "Reads empty reader",
			args: args{reader: emptyReader},
			want: []byte(""),
		},
		{
			name: "Fails to read unexisting file",
			args: args{
				filename: "somethingUnexisting",
			},
			err: &os.PathError{
				Op:   "open",
				Path: "somethingUnexisting",
				Err:  syscall.Errno(0x2),
			},
			want: []byte(""),
		},
		{
			name: "Fails to read nil reader",
			args: args{},
			err:  errors.New("cannot read from nil reader"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFileOrReader(tt.args.reader, tt.args.filename)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewFileOrReader() error = %v, wantErr %v", err, tt.err)
				return
			}

			var contents []byte
			if got != nil {
				contents, _ = ioutil.ReadAll(got)
			}

			if !reflect.DeepEqual(contents, tt.want) {
				t.Errorf("NewFileOrReader() = %v, want %v", contents, tt.want)
			}
		})
	}
}
