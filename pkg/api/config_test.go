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
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/auth"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name   string
		fields Config
		err    error
	}{
		{
			name: "Validate succeeds",
			fields: Config{
				Client:     new(http.Client),
				Host:       "https://localhost",
				AuthWriter: auth.APIKey("dummy"),
			},
		},
		{
			name:   "Validate fails due on empty config",
			fields: Config{},
			err: multierror.NewPrefixed("invalid api config",
				errors.New("client cannot be empty"),
				errEmptyAuthWriter,
				errors.New("host cannot be empty"),
			),
		},
		{
			name: "Validate fails due to missing Authenticator",
			fields: Config{
				Client: new(http.Client),
			},
			err: multierror.NewPrefixed("invalid api config",
				errEmptyAuthWriter,
				errors.New("host cannot be empty"),
			),
		},
		{
			name: "Validate fails due to verbose set but device empty",
			fields: Config{
				Client:          new(http.Client),
				VerboseSettings: VerboseSettings{Verbose: true},
			},
			err: multierror.NewPrefixed("invalid api config",
				errEmptyAuthWriter,
				errors.New("host cannot be empty"),
				multierror.NewPrefixed("invalid verbose settings",
					errors.New("output device cannot be empty when verbose is enabled"),
				),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.fields
			if err := c.Validate(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
