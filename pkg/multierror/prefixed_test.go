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
	"reflect"
	"testing"

	"github.com/hashicorp/go-multierror"
)

func TestNewPrefixed(t *testing.T) {
	type args struct {
		prefix string
		errs   []error
	}
	tests := []struct {
		name string
		args args
		want *Prefixed
	}{
		{
			name: "New without explicit prefix",
			want: &Prefixed{
				Errors: make([]error, 0),
			},
		},
		{
			name: "New with prefix and errors",
			args: args{prefix: "some prefix here", errs: []error{
				errors.New("an error"),
				errors.New("another error"),
				errors.New("yet another error"),
			}},
			want: &Prefixed{
				Prefix: "some prefix here",
				Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
					errors.New("yet another error"),
				},
			},
		},
		{
			name: "New with prefix and errors and unpacking some other prefixed errors",
			args: args{prefix: "some prefix here", errs: []error{
				errors.New("an error"),
				errors.New("another error"),
				&Prefixed{Prefix: "a prefix", Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
				}},
			}},
			want: &Prefixed{
				Prefix: "some prefix here",
				Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
					errors.New("a prefix: an error"),
					errors.New("a prefix: another error"),
				},
			},
		},
		{
			name: "New with prefix and errors and unpacking some other prefixed errors and some multierrors",
			args: args{prefix: "some prefix here", errs: []error{
				errors.New("an error"),
				errors.New("another error"),
				&Prefixed{Prefix: "a prefix", Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
				}},
				&multierror.Error{Errors: []error{
					errors.New("multierror error"),
					errors.New("multierror error 2"),
				}},
				&Prefixed{Prefix: "some prefix here", Errors: []error{
					errors.New("unprefixed error 1"),
					errors.New("unprefixed error 2"),
				}},
			}},
			want: &Prefixed{
				Prefix: "some prefix here",
				Errors: []error{
					errors.New("an error"),
					errors.New("another error"),
					errors.New("a prefix: an error"),
					errors.New("a prefix: another error"),
					errors.New("multierror error"),
					errors.New("multierror error 2"),
					errors.New("unprefixed error 1"),
					errors.New("unprefixed error 2"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrefixed(tt.args.prefix, tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %+v, want %+v", *got, *tt.want)
			}
		})
	}
}

func TestPrefixed_Append(t *testing.T) {
	type fields struct {
		Prefix     string
		Errors     []error
		FormatFunc FormatFunc
	}
	type args struct {
		errs []error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Prefixed
	}{
		{
			name: "Adds some errors",
			fields: fields{Prefix: "some prefix", Errors: []error{
				errors.New("an error"),
			}},
			args: args{errs: []error{
				&Prefixed{Prefix: "prefix 1", Errors: []error{
					errors.New("a prefixed error"),
					errors.New("another error"),
				}},
				&Prefixed{Prefix: "prefix 2", Errors: []error{
					errors.New("a prefixed error"),
					errors.New("another error"),
				}},
				errors.New("a normal error"),
				// No prefix
				&Prefixed{Errors: []error{
					errors.New("a prefixed error"),
					errors.New("another error"),
				}},
			}},
			want: &Prefixed{
				Prefix: "some prefix",
				Errors: []error{
					errors.New("an error"),
					errors.New("prefix 1: a prefixed error"),
					errors.New("prefix 1: another error"),
					errors.New("prefix 2: a prefixed error"),
					errors.New("prefix 2: another error"),
					errors.New("a normal error"),
					errors.New("github.com/elastic/cloud-sdk-go/pkg/multierror.TestPrefixed_Append.func1: a prefixed error"),
					errors.New("github.com/elastic/cloud-sdk-go/pkg/multierror.TestPrefixed_Append.func1: another error"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prefixed{
				Prefix:     tt.fields.Prefix,
				Errors:     tt.fields.Errors,
				FormatFunc: tt.fields.FormatFunc,
			}
			if got := p.Append(tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prefixed.Append() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPrefixed_ErrorOrNil(t *testing.T) {
	type fields struct {
		Prefix     string
		Errors     []error
		FormatFunc FormatFunc
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name: "empty error returns nil",
			err:  nil,
		},
		{
			name: "a Prefixed error with multiple errors returns the error",
			fields: fields{Errors: []error{
				errors.New("some error"),
				errors.New("some other error"),
			}},
			err: &Prefixed{Errors: []error{
				errors.New("some error"),
				errors.New("some other error"),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prefixed{
				Prefix:     tt.fields.Prefix,
				Errors:     tt.fields.Errors,
				FormatFunc: tt.fields.FormatFunc,
			}
			if err := p.ErrorOrNil(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Prefixed.ErrorOrNil() error = %+v, wantErr %+v", err, tt.err)
			}
		})
	}
}

func TestPrefixed_Error(t *testing.T) {
	type fields struct {
		Prefix     string
		Errors     []error
		FormatFunc FormatFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty errors return empty string",
			want: "",
		},
		{
			name: "Empty FormatFunc uses the default one",
			fields: fields{
				Prefix: "prefix",
				Errors: []error{
					errors.New("some error"),
				},
			},
			want: "prefix: " + multierror.ListFormatFunc([]error{
				errors.New("some error"),
			}),
		},
		{
			name: "Empty FormatFunc uses the default one with more than 1 error",
			fields: fields{
				Prefix: "prefix",
				Errors: []error{
					errors.New("some error"),
					errors.New("another error"),
				},
			},
			want: "prefix: " + multierror.ListFormatFunc([]error{
				errors.New("some error"),
				errors.New("another error"),
			}),
		},
		{
			name: "With a custom FormatFunc ",
			fields: fields{
				Prefix: "prefix",
				FormatFunc: func(es []error) string {
					return "some bogus return"
				},
				Errors: []error{
					errors.New("some error"),
				},
			},
			want: "prefix: some bogus return",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Prefixed{
				Prefix:     tt.fields.Prefix,
				Errors:     tt.fields.Errors,
				FormatFunc: tt.fields.FormatFunc,
			}
			if got := p.Error(); got != tt.want {
				t.Errorf("Prefixed.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
