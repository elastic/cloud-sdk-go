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

package mock

import (
	"io"
	"net/http"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// APIError can be used to create a multiple error response.
type APIError struct {
	Code, Message string
	Fields        []string
}

// NewErrorBody creates a replica of a body representing an EC API error.
func NewErrorBody(errs ...APIError) io.ReadCloser {
	var replies = make([]*models.BasicFailedReplyElement, 0, len(errs))
	for _, e := range errs {
		replies = append(replies, &models.BasicFailedReplyElement{
			Code: &e.Code, Message: &e.Message, Fields: e.Fields,
		})
	}
	return NewStructBody(models.BasicFailedReply{
		Errors: replies,
	})
}

// NewErrorResponse creates a replica of a responnse representing an EC API
// error.
func NewErrorResponse(code int, errs ...APIError) Response {
	return Response{Response: http.Response{
		StatusCode: code,
		Status:     http.StatusText(code),
		Body:       NewErrorBody(errs...),
	}}
}
