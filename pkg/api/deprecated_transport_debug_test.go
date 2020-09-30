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

package api

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDebugTransport(t *testing.T) {
	var debugBuffer = new(bytes.Buffer)
	var debugTransport = &DebugTransport{
		transport: http.DefaultTransport,
		output:    new(bytes.Buffer),
		count:     -1,
	}
	type args struct {
		transport http.RoundTripper
		obscure   bool
	}
	tests := []struct {
		name string
		args args
		want *DebugTransport
	}{
		{
			name: "Returns a new DebugTransport when there's an empty transport",
			args: args{},
			want: &DebugTransport{
				transport: http.DefaultTransport,
				output:    debugBuffer,
				count:     -1,
			},
		},
		{
			name: "Returns the same debugTransport when a DebugTransport is sent as a parameter",
			args: args{
				transport: debugTransport,
			},
			want: &DebugTransport{
				transport: http.DefaultTransport,
				output:    debugBuffer,
				count:     -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &bytes.Buffer{}
			if got := NewDebugTransport(tt.args.transport, o, tt.args.obscure); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
