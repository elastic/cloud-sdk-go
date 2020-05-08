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
	"io"

	"github.com/elastic/cloud-sdk-go/pkg/auth"
)

// LoginUser logs in a user when its AuthWriter is of type *auth.UserLogin.
// Additionally, calls the RefreshToken pethod in *auth.UserLogin launching a
// background Go routine which will keep the JWT token always valid.
func LoginUser(instance *API, writer io.Writer) error {
	aw, ok := instance.AuthWriter.(*auth.UserLogin)
	if !ok {
		return nil
	}

	if err := aw.Login(instance.V1API); err != nil {
		return err
	}

	return aw.RefreshToken(auth.RefreshTokenParams{
		Client:      instance.V1API,
		ErrorDevice: writer,
	})
}
