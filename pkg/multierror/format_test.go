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

package multierror

import (
	"errors"
	"io/ioutil"
	"path"
	"reflect"
	"runtime"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_elasticsearch"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestWithFormat(t *testing.T) {
	type args struct {
		err    error
		format string
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: `non "json" format simply returns the error`,
			args: args{
				err:    errors.New("some error"),
				format: "some format",
			},
			err: errors.New("some error"),
		},
		{
			name: `"json" format on not a multierror returns the error`,
			args: args{
				err:    errors.New("some error"),
				format: "json",
			},
			err: errors.New("some error"),
		},
		{
			name: `non "json" format simply returns the error`,
			args: args{
				err:    NewPrefixed("some prefix"),
				format: "some format",
			},
			err: NewPrefixed("some prefix"),
		},
		{
			name: `"json" format on a multierror returns the error with the format func set`,
			args: args{
				err:    &Prefixed{Prefix: "some prefix"},
				format: "json",
			},
			err: &Prefixed{Prefix: "some prefix"},
		},
		{
			name: `"json" format on a multierror returns the error with the format func set`,
			args: args{
				err:    new(multierror.Error),
				format: "json",
			},
			err: new(multierror.Error),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WithFormat(tt.args.err, tt.args.format)

			// Sadly, in Go it's not possible to compare two functions, instead we're relying on the
			// declared function name to ensure that what we expect has been set, has been.
			if tt.args.format == "json" {
				// Prefixed asseertion.
				var merr *Prefixed
				if errors.As(err, &merr) {
					funcName := path.Base(runtime.FuncForPC(reflect.ValueOf(merr.FormatFunc).Pointer()).Name())
					if funcName != "multierror.JSONFormatFunc" {
						t.Errorf("WithFormat() error = %s, wantErr multierror.JSONFormatFunc", funcName)
					}
				}

				// multierror.Error assertion.
				var hashiErr *multierror.Error
				if errors.As(err, &hashiErr) {
					funcName := path.Base(runtime.FuncForPC(reflect.ValueOf(hashiErr.ErrorFormat).Pointer()).Name())
					if funcName != "multierror.JSONFormatFunc" {
						t.Errorf("WithFormat() error = %s, wantErr multierror.JSONFormatFunc", funcName)
					}
				}
			} else if !assert.Equal(t, err, tt.err) {
				t.Errorf("WithFormat() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func TestJSONFormatFunc(t *testing.T) {
	b, err := ioutil.ReadFile("./testdata/multierror.json")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		es []error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "marshal",
			args: args{es: []error{
				errors.New("some"),
				errors.New("some 2"),
				errors.New("some 3"),
				&clusters_elasticsearch.DeleteEsClusterRetryWith{
					Payload: &models.BasicFailedReply{
						Errors: []*models.BasicFailedReplyElement{
							{
								Code:    ec.String("clusters.cluster_plan_state_error"),
								Message: ec.String("There are running instances"),
							},
							{
								Code:    ec.String("auth.invalid_password"),
								Fields:  []string{"body.password"},
								Message: ec.String("request password doesn't match the user's password"),
							},
						},
					},
				},
			}},
			want: string(b),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JSONFormatFunc(tt.args.es); got != tt.want {
				t.Errorf("JSONFormatFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
