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
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client"
	"github.com/elastic/cloud-sdk-go/pkg/client/authentication"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var (
	defaultRefreshTickerTime = time.Minute * 1
	errLoginClientEmpty      = errors.New("auth: login client cannot be empty")
)

// UserLogin uses a user's username and password to login against the Login
// API Endpoint. Doing so obtains a JWT token which is then persisted in the
// token field, guarded by a mutex.
// This is a form of user authentication, but API Keys are still the preferred
// authentication mechanism.
type UserLogin struct {
	Username, Password string
	Holder             TokenHandler
}

// NewUserLogin creates a UserLogin from a username and password. It does not
// automatically login against the API until Auth() is called.
func NewUserLogin(username, password string) (*UserLogin, error) {
	userLogin := UserLogin{Username: username, Password: password}
	if err := userLogin.Validate(); err != nil {
		return nil, err
	}

	if userLogin.Holder == nil {
		userLogin.Holder = new(GenericHolder)
	}

	return &userLogin, nil
}

// Validate ensures the validity of the data container.
func (t *UserLogin) Validate() error {
	var merr = multierror.NewPrefixed("auth")
	if t.Username == "" {
		merr = merr.Append(errors.New("username must not be empty"))
	}
	if t.Password == "" {
		merr = merr.Append(errors.New("password must not be empty"))
	}
	return merr.ErrorOrNil()
}

// RefreshTokenParams is used to refresh a bearer token, which is necessary
// before its validity expires.
type RefreshTokenParams struct {
	Client           *client.Rest
	Frequency        time.Duration
	ErrorDevice      io.Writer
	InterruptChannel chan os.Signal
}

// Validate ensures that the parameters are valid.
func (params *RefreshTokenParams) Validate() error {
	params.fillValues()
	var merr = multierror.NewPrefixed("auth")
	if params.ErrorDevice == nil {
		merr = merr.Append(errors.New("errorDevice cannot be nil"))
	}

	if params.Client == nil {
		merr = merr.Append(errors.New("rest client cannot be nil"))
	}

	return merr.ErrorOrNil()
}

// fillValues sets the default values for the structure.
func (params *RefreshTokenParams) fillValues() {
	if params.Frequency.Nanoseconds() == 0 {
		params.Frequency = defaultRefreshTickerTime
	}
}

// Login calls the authentication/login endpoint with a username and password
// persisting the returned token.
func (t *UserLogin) Login(c *client.Rest) error {
	if c == nil {
		return errLoginClientEmpty
	}

	res, err := c.Authentication.Login(authentication.NewLoginParams().
		WithBody(&models.LoginRequest{
			Username: ec.String(t.Username),
			Password: ec.String(t.Password),
		}),
		nil,
	)
	if err != nil {
		return multierror.NewPrefixed("failed to login with user/password", apierror.Unwrap(err))
	}

	return t.Holder.Update(*res.Payload.Token)
}

// AuthenticateRequest authenticates a runtime.ClientRequest. Implements the
// runtime.ClientAuthInfoWriter interface using the JWT Bearer token.
func (t *UserLogin) AuthenticateRequest(c runtime.ClientRequest, r strfmt.Registry) error {
	return httptransport.BearerToken(t.Holder.Token()).AuthenticateRequest(c, r)
}

// AuthRequest adds the Authorization header to an http.Request
func (t *UserLogin) AuthRequest(req *http.Request) *http.Request {
	req.Header.Add("Authorization", "Bearer "+t.Holder.Token())
	return req
}

// RefreshToken creates a goroutine which will run in the background refreshing
// the token every Frequency. It does not refresh the token until the first
// period has passed.
func (t *UserLogin) RefreshToken(params RefreshTokenParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	if params.InterruptChannel == nil {
		params.InterruptChannel = make(chan os.Signal, 1)
	}
	signal.Notify(params.InterruptChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		ticker := time.NewTicker(params.Frequency)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := t.RefreshTokenOnce(params.Client); err != nil {
					fmt.Fprintln(params.ErrorDevice, err)
					continue
				}
			case <-params.InterruptChannel:
				return
			}
		}
	}()

	return nil
}

// RefreshTokenOnce refreshRefreshTokenOncees the current JWT token once.
func (t *UserLogin) RefreshTokenOnce(c *client.Rest) error {
	if c == nil {
		return errLoginClientEmpty
	}

	res, err := c.Authentication.RefreshToken(
		authentication.NewRefreshTokenParams(), t,
	)
	if err != nil {
		return multierror.NewPrefixed("failed to refresh the loaded token", apierror.Unwrap(err))
	}

	return t.Holder.Update(*res.Payload.Token)
}
