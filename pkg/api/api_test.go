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
	"errors"
	"net/url"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/auth"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/output"
)

func TestNewAPI(t *testing.T) {
	dummyKey, err := auth.NewAPIKey("dummy")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		c Config
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "fails due to empty parameters",
			err: multierror.NewPrefixed("invalid api config",
				errors.New("client cannot be empty"),
				errEmptyAuthWriter,
				errESSInvalidAuth,
			),
		},
		{
			name: "fails with invalid url",
			args: args{c: Config{
				Host: "very.much.invalid",
				VerboseSettings: VerboseSettings{
					Verbose: true,
					Device:  output.NewDevice(new(bytes.Buffer)),
				},
				Client: mock.NewClient(),
			}},
			err: multierror.NewPrefixed("invalid api config",
				errEmptyAuthWriter,
				&url.Error{Op: "parse", URL: "very.much.invalid/", Err: errors.New("invalid URI for request")},
			),
		},
		{
			name: "succeeds with region",
			args: args{c: Config{
				Host:       "https://cloud.elastic.co",
				Region:     "us-east-1",
				AuthWriter: dummyKey,
				VerboseSettings: VerboseSettings{
					Verbose: true,
					Device:  output.NewDevice(new(bytes.Buffer)),
				},
				Client: mock.NewClient(),
			}},
		},
		{
			name: "succeeds without region",
			args: args{c: Config{
				Host:       ESSEndpoint,
				AuthWriter: dummyKey,
				VerboseSettings: VerboseSettings{
					Verbose: true,
					Device:  output.NewDevice(new(bytes.Buffer)),
				},
				Client: mock.NewClient(),
			}},
		},
		{
			name: "succeeds without host",
			args: args{c: Config{
				AuthWriter: dummyKey,
				VerboseSettings: VerboseSettings{
					Verbose: true,
					Device:  output.NewDevice(new(bytes.Buffer)),
				},
				Client: mock.NewClient(),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewAPI(tt.args.c); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewAPI() error = %v, wantErr %v", err, tt.err)
				return
			}
		})
	}
}
