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
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

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

	wantVerbose := mock.NewClient()
	verboseTransport := CustomTransport{
		rt:      mock.NewRoundTripper(),
		writer:  output.NewDevice(new(bytes.Buffer)),
		verbose: true,
		agent:   DefaultUserAgent,
		backoff: defaultBackoff,
	}
	wantVerbose.Transport = &verboseTransport

	wantRetries := mock.NewClient()

	retryTransport := verboseTransport
	retryTransport.retries = 5
	wantRetries.Transport = &retryTransport

	clientWithTimeout := mock.NewClient()
	clientWithTimeout.Timeout = time.Hour

	type args struct {
		c Config
	}
	tests := []struct {
		name       string
		args       args
		wantClient *http.Client
		err        error
	}{
		{
			name: "fails due to empty parameters",
			err: multierror.NewPrefixed("invalid api config",
				errors.New("client cannot be empty"),
				errEmptyAuthWriter,
				errors.New("apikey is the only valid authentication mechanism when targeting the Elasticsearch Service"),
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
			wantClient: mock.NewClient(),
		},
		{
			name: "succeeds with verbose",
			args: args{c: Config{
				Host:       ESSEndpoint,
				AuthWriter: dummyKey,
				VerboseSettings: VerboseSettings{
					Verbose: true,
					Device:  output.NewDevice(new(bytes.Buffer)),
				},
				Client: mock.NewClient(),
			}},
			wantClient: wantVerbose,
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
			wantClient: wantVerbose,
		},
		{
			name: "succeeds with retries",
			args: args{c: Config{
				Host:       ESSEndpoint,
				AuthWriter: dummyKey,
				Retries:    5,
				VerboseSettings: VerboseSettings{
					Verbose: true,
					Device:  output.NewDevice(new(bytes.Buffer)),
				},
				Client: mock.NewClient(),
			}},
			wantClient: wantRetries,
		},
		{
			name: "succeeds with and removes the http.Client.Timeout",
			args: args{c: Config{
				Host:       ESSEndpoint,
				AuthWriter: dummyKey,
				Retries:    5,
				VerboseSettings: VerboseSettings{
					Verbose: true,
					Device:  output.NewDevice(new(bytes.Buffer)),
				},
				Client: clientWithTimeout,
			}},
			wantClient: wantRetries,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewAPI(tt.args.c)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}

			assert.Equal(t, tt.wantClient, tt.args.c.Client)
		})
	}
}
