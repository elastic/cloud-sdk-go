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
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/client"
	"github.com/elastic/cloud-sdk-go/pkg/client/authentication"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/sync"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestNewUserLogin(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want *UserLogin
		err  error
	}{
		{
			name: "fails on empty username and password",
			err: &multierror.Error{Errors: []error{
				errors.New("auth: Username must not be empty"),
				errors.New("auth: Password must not be empty"),
			}},
		},
		{
			name: "fails on empty username",
			args: args{password: "some"},
			err: &multierror.Error{Errors: []error{
				errors.New("auth: Username must not be empty"),
			}},
		},
		{
			name: "fails on empty password",
			args: args{username: "some"},
			err: &multierror.Error{Errors: []error{
				errors.New("auth: Password must not be empty"),
			}},
		},
		{
			name: "builds UserLogin with default holder",
			args: args{username: "some", password: "somepass"},
			want: &UserLogin{
				Username: "some",
				Password: "somepass",
				Holder:   new(GenericHolder),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserLogin(tt.args.username, tt.args.password)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewUserLogin() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLogin_Login(t *testing.T) {
	type fields struct {
		Username string
		Password string
		Holder   TokenHandler
	}
	type args struct {
		rc *client.Rest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantToken string
		err       error
	}{
		{
			name: "fails due to empty client",
			fields: fields{
				Username: "invalid",
				Password: "invalid",
				Holder:   new(GenericHolder),
			},
			err: errors.New("auth: login client cannot be empty"),
		},
		{
			name: "fails due to API error",
			fields: fields{
				Username: "invalid",
				Password: "invalid",
				Holder:   new(GenericHolder),
			},
			args: args{rc: newMock(mock.Response{Response: http.Response{
				Body:       mock.NewStructBody(models.BasicFailedReply{}),
				StatusCode: 401,
			}})},
			err: fmt.Errorf(
				"failed to login with user/password: %s",
				&authentication.LoginUnauthorized{Payload: new(models.BasicFailedReply)},
			),
		},
		{
			name: "succeeds",
			fields: fields{
				Username: "valid",
				Password: "even more valid",
				Holder:   new(GenericHolder),
			},
			args: args{rc: newMock(mock.Response{Response: http.Response{
				Body: mock.NewStructBody(models.TokenResponse{
					Token: ec.String("some token!!!"),
				}),
				StatusCode: 200,
			}})},
			wantToken: "some token!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ul := &UserLogin{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Holder:   tt.fields.Holder,
			}
			if err := ul.Login(tt.args.rc); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("UserLogin.Login() error = %v, wantErr %v", err, tt.err)
			}
			if ul.Holder != nil {
				if token := ul.Holder.Token(); token != tt.wantToken {
					t.Errorf("UserLogin.Login() token = %v, want %v", token, tt.wantToken)
				}
			}
		})
	}
}

func TestUserLogin_AuthRequest(t *testing.T) {
	type fields struct {
		Username string
		Password string
		Holder   TokenHandler
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *http.Request
	}{
		{
			name:   "auths the request",
			fields: fields{Holder: &GenericHolder{token: "some"}},
			args: args{req: &http.Request{
				Header: make(http.Header),
			}},
			want: &http.Request{
				Header: http.Header{
					"Authorization": []string{"Bearer some"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ul := &UserLogin{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Holder:   tt.fields.Holder,
			}
			if got := ul.AuthRequest(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserLogin.AuthRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLogin_RefreshTokenOnce(t *testing.T) {
	type fields struct {
		Username string
		Password string
		Holder   TokenHandler
	}
	type args struct {
		rc *client.Rest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantToken string
		err       error
	}{
		{
			name: "fails due to empty client",
			fields: fields{
				Username: "invalid",
				Password: "invalid",
				Holder:   new(GenericHolder),
			},
			err: errors.New("auth: login client cannot be empty"),
		},
		{
			name: "fails due to API error",
			fields: fields{
				Username: "invalid",
				Password: "invalid",
				Holder:   new(GenericHolder),
			},
			args: args{rc: newMock(mock.Response{Response: http.Response{
				Body:       mock.NewStructBody(models.BasicFailedReply{}),
				StatusCode: 401,
			}})},
			err: fmt.Errorf(
				"auth: failed to refresh the loaded token: %s",
				&authentication.RefreshTokenUnauthorized{Payload: new(models.BasicFailedReply)},
			),
		},
		{
			name: "succeeds",
			fields: fields{
				Username: "valid",
				Password: "even more valid",
				Holder:   &GenericHolder{token: "Once a valid token"},
			},
			args: args{rc: newMock(mock.Response{Response: http.Response{
				Body: mock.NewStructBody(models.TokenResponse{
					Token: ec.String("some token!!!"),
				}),
				StatusCode: 200,
			}})},
			wantToken: "some token!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ul := &UserLogin{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Holder:   tt.fields.Holder,
			}
			if err := ul.RefreshTokenOnce(tt.args.rc); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("UserLogin.RefreshTokenOnce() error = %v, wantErr %v", err, tt.err)
			}
			if ul.Holder != nil {
				if token := ul.Holder.Token(); token != tt.wantToken {
					t.Errorf("UserLogin.RefreshTokenOnce() token = %v, want %v", token, tt.wantToken)
				}
			}
		})
	}
}

func TestUserLogin_RefreshToken(t *testing.T) {
	type fields struct {
		Username string
		Password string
		Holder   TokenHandler
	}
	type args struct {
		params RefreshTokenParams
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantToken       string
		wantErrorDevice string
		err             error
	}{
		{
			name: "returns an error on invalid params",
			err: &multierror.Error{Errors: []error{
				errors.New("refresh token: ErrorDevice cannot be nil"),
				errors.New("refresh token: rest client cannot be nil"),
			}},
		},
		{
			name: "Refresh token",
			fields: fields{
				Holder: new(GenericHolder),
			},
			wantToken: "sometoken",
			args: args{params: RefreshTokenParams{
				Frequency:        time.Millisecond * 10,
				ErrorDevice:      sync.NewBuffer(),
				InterruptChannel: make(chan os.Signal, 1),
				// Add 3 responses as the ceiling for the multiplier. all are the same just guarding against errors.
				Client: newMock(mock.Response{Response: http.Response{
					Body:       mock.NewStructBody(models.TokenResponse{Token: ec.String("sometoken")}),
					StatusCode: 200,
				}}, mock.Response{Response: http.Response{
					Body:       mock.NewStructBody(models.TokenResponse{Token: ec.String("sometoken")}),
					StatusCode: 200,
				}}, mock.Response{Response: http.Response{
					Body:       mock.NewStructBody(models.TokenResponse{Token: ec.String("sometoken")}),
					StatusCode: 200,
				}}),
			}},
		},
		{
			name: "Refresh returns error the first time",
			fields: fields{
				Holder: new(GenericHolder),
			},
			wantToken: "sometoken",
			args: args{params: RefreshTokenParams{
				Frequency:        time.Millisecond * 10,
				ErrorDevice:      sync.NewBuffer(),
				InterruptChannel: make(chan os.Signal, 1),
				// Add 3 responses as the ceiling for the multiplier. all are the same just guarding against errors.
				Client: newMock(mock.Response{Response: http.Response{
					Body:       mock.NewStructBody(new(models.BasicFailedReply)),
					StatusCode: 401,
				}}, mock.Response{Response: http.Response{
					Body:       mock.NewStructBody(models.TokenResponse{Token: ec.String("sometoken")}),
					StatusCode: 200,
				}}, mock.Response{Response: http.Response{
					Body:       mock.NewStructBody(models.TokenResponse{Token: ec.String("sometoken")}),
					StatusCode: 200,
				}}),
			}},
			wantErrorDevice: fmt.Sprintf(
				"auth: failed to refresh the loaded token: %s\n",
				&authentication.RefreshTokenUnauthorized{Payload: new(models.BasicFailedReply)},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ul := &UserLogin{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Holder:   tt.fields.Holder,
			}

			if tt.args.params.InterruptChannel == nil {
				tt.args.params.InterruptChannel = make(chan os.Signal, 1)
			}

			if err := ul.RefreshToken(tt.args.params); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("UserLogin.RefreshToken() error = %v, wantErr %v", err, tt.err)
				return
			}

			<-time.After(tt.args.params.Frequency * 3)
			tt.args.params.InterruptChannel <- os.Interrupt

			if ul.Holder != nil {
				if token := ul.Holder.Token(); token != tt.wantToken {
					t.Errorf("UserLogin.RefreshToken() token = %v, want %v", token, tt.wantToken)
				}
			}

			if tt.args.params.ErrorDevice != nil {
				if ed, ok := tt.args.params.ErrorDevice.(*sync.Buffer); ok {
					if edContents := ed.String(); edContents != tt.wantErrorDevice {
						t.Errorf("UserLogin.RefreshToken() ErrorDevice = %v, want %v", edContents, tt.wantErrorDevice)
					}
				}
			}
		})
	}
}
