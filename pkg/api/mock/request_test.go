package mock

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/stretchr/testify/assert"
)

func TestAssertRequest(t *testing.T) {
	type args struct {
		want *RequestAssertion
		req  *http.Request
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "expects a body but finds none",
			args: args{
				want: &RequestAssertion{
					Body: NewStringBody(`{"some field":1}`),
				},
				req: &http.Request{},
			},
			err: multierror.NewPrefixed("request assertion",
				errors.New(`got body , want {"some field":1}`),
			),
		},
		{
			name: "matches all fields",
			args: args{
				want: &RequestAssertion{
					Body: NewStringBody(`{"some field":1}`),
					Header: map[string][]string{
						"Authorization": {"Apikey Someapikey"},
						"Content-Type":  {"application/json"},
					},
					Method: "POST",
					Host:   "somehost",
					Path:   "/somepath/somesubpath",
				},
				req: &http.Request{
					Header: map[string][]string{
						"Authorization": {"Apikey Someapikey"},
						"Content-Type":  {"application/json"},
					},
					Body: NewStringBody(`{"some field":1}`),
					URL: &url.URL{
						Path: "/somepath/somesubpath",
						Host: "somehost",
					},
					Host:   "somehost",
					Method: "POST",
				},
			},
		},
		{
			name: "matches no fields, returns an error",
			args: args{
				want: &RequestAssertion{
					Body: NewStringBody(`{"some field":2}`),
					Header: map[string][]string{
						"Content-Type": {"application/json"},
					},
					Method: "GET",
					Host:   "someotherhost",
					Path:   "/someotherpath/somesubpath",
				},
				req: &http.Request{
					Header: map[string][]string{
						"Authorization": {"Apikey Someapikey"},
						"Content-Type":  {"application/json"},
					},
					Body: NewStringBody(`{"some field":1}`),
					URL: &url.URL{
						Path: "/somepath/somesubpath",
						Host: "somehost",
					},
					Host:   "somehost",
					Method: "POST",
				},
			},
			err: multierror.NewPrefixed("request assertion",
				errors.New(`got body {"some field":1}, want {"some field":2}`),
				errors.New(`headers do not match: map[Content-Type:[application/json]] != map[Authorization:[Apikey Someapikey] Content-Type:[application/json]]`),
				errors.New(`methods do not match: GET != POST`),
				errors.New(`paths do not match: /someotherpath/somesubpath != /somepath/somesubpath`),
				errors.New(`hosts do not match: someotherhost != somehost`),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AssertRequest(tt.args.want, tt.args.req)
			var errString string
			if tt.err != nil {
				errString = tt.err.Error()
			}
			if err != nil || errString != "" {
				assert.EqualError(t, err, errString)
				return
			}

			assert.NoError(t, err)
		})
	}
}
