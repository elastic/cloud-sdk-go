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

package noteapi

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/elastic/cloud-sdk-go/pkg/util/testutils"
)

func TestUpdate(t *testing.T) {
	var messageDateTime = testutils.ParseDate(t, "2018-04-13T07:11:54.999Z")
	const simpleNoteModified = `{
		"id": "1",
		"message": "a modified message",
		"user_id": "root",
		"timestamp": "2018-04-13T07:11:54.999Z",
		"version": 2
	}`

	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Note
		wantErr string
	}{
		{
			name: "Update note succeeds",
			args: args{params: UpdateParams{
				UserID:  "root",
				Message: "a modified message",
				NoteID:  "1",
				Params: Params{
					Region: "us-east-1",
					ID:     "e3dac8bf3dc64c528c295a94d0f19a77",
					API: api.NewMock(mock.Response{Response: http.Response{
						Body:       mock.NewStringBody(simpleNoteModified),
						StatusCode: 200,
					}}),
				},
			}},
			want: &models.Note{
				ID:        "1",
				Message:   ec.String("a modified message"),
				UserID:    "root",
				Timestamp: messageDateTime,
			},
		},
		{
			name: "Update note fails due to api error",
			args: args{params: UpdateParams{
				UserID:  "epifanio",
				Message: "a modified message",
				NoteID:  "1",
				Params: Params{
					Region: "us-east-1",
					ID:     "a2c4f423c1014941b75a48292264dd25",
					API:    api.NewMock(mock.SampleInternalError()),
				},
			}},
			wantErr: mock.MultierrorInternalError.Error(),
		},
		{
			name: "Update note fails due to empty params",
			wantErr: multierror.NewPrefixed("deployment note update",
				errors.New("user id cannot be empty"),
				errors.New("note comment cannot be empty"),
				errors.New("note id cannot be empty"),
				multierror.NewPrefixed("deployment note",
					errors.New("api reference is required for the operation"),
					errors.New(`id "" is invalid`),
					errors.New("region not specified and is required for this operation"),
				),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.wantErr) {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
