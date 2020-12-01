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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/auth"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestLoginUser(t *testing.T) {
	var apiKeyAPI = NewMock()

	var userLoginError = NewMock(mock.Response{
		Response: http.Response{
			StatusCode: 401,
			Body: mock.NewStructBody(&models.BasicFailedReply{Errors: []*models.BasicFailedReplyElement{
				{Code: ec.String("a code"), Message: ec.String("message")},
			}}),
		},
	})
	fail, err := auth.NewUserLogin("user", "pass")
	if err != nil {
		t.Fatal(err)
	}
	userLoginError.AuthWriter = fail

	var userLoginSuccess = NewMock(mock.New200Response(mock.NewStructBody(models.TokenResponse{
		Token: ec.String("sometoken"),
	})))
	success, err := auth.NewUserLogin("user", "pass")
	if err != nil {
		t.Fatal(err)
	}
	userLoginSuccess.AuthWriter = success
	type args struct {
		instance *API
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		err        string
	}{
		{
			name: "skips logging in when the AuthWriter isn't *auth.UserLogin",
			args: args{instance: apiKeyAPI},
		},
		{
			name: "fails logging in",
			args: args{instance: userLoginError},
			err: multierror.NewPrefixed("failed to login with user/password",
				errors.New("api error: a code: message"),
			).Error(),
		},
		{
			name: "succeeds logging in",
			args: args{instance: userLoginSuccess},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := LoginUser(tt.args.instance, writer); err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("LoginUser() error = %v, wantErr %v", err, tt.err)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("LoginUser() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
