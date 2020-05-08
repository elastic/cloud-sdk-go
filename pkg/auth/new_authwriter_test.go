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

package auth

import (
	"errors"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestNewAuthWriter(t *testing.T) {
	var apikey = APIKey("someapikey")

	type args struct {
		c Config
	}
	tests := []struct {
		name string
		args args
		want Writer
		err  error
	}{
		{
			name: "when APIKey is set returns an apikey AuthWriter",
			args: args{c: Config{
				APIKey: "someapikey",
			}},
			want: &apikey,
		},
		{
			name: "when Username and Password is set returns an UserLogin",
			args: args{c: Config{
				Username: "myuser", Password: "my very secret password",
			}},
			want: &UserLogin{
				Username: "myuser", Password: "my very secret password",
				Holder: new(GenericHolder),
			},
		},
		{
			name: "when Username is set but password is empty returns an error",
			args: args{c: Config{
				Username: "myuser",
			}},
			err: multierror.NewPrefixed("auth", errors.New("password must not be empty")),
		},
		{
			name: "when Password is set but username is empty returns an error",
			args: args{c: Config{
				Password: "my very secret password",
			}},
			err: multierror.NewPrefixed("auth", errors.New("username must not be empty")),
		},
		{
			name: "when the credentials are empty returns an error",
			args: args{c: Config{}},
			err: multierror.NewPrefixed("authwriter",
				errors.New("one of apikey or username and password must be specified"),
			),
		},
		{
			name: "when all the credentials are set returns an error",
			args: args{c: Config{APIKey: "a", Username: "b", Password: "c"}},
			err: multierror.NewPrefixed("authwriter",
				errors.New("only one of of apikey or username and password can be specified"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAuthWriter(tt.args.c)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewAuthWriter() error = %v, wantErr %v", err, tt.err)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthWriter() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
