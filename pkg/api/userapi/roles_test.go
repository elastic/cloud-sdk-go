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

package userapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

const platformViewerRole = "ece_platform_viewer"

func TestValidateRoles(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		err  string
	}{
		{
			name: "validate should return an error when ece_platform_admin is used along other roles",
			arg:  []string{platformAdminRole, platformViewerRole},
			err: multierror.NewPrefixed("invalid user params",
				errors.New("ece_platform_admin cannot be used in conjunction with other roles"),
			).Error(),
		},
		{
			name: "validate should return an error when ece_platform_admin is used along other roles",
			arg:  []string{deploymentsManagerRole, deploymentsViewerRole},
			err: multierror.NewPrefixed("invalid user params",
				errors.New("only one of ece_deployment_manager or ece_deployment_viewer can be chosen"),
			).Error(),
		},
		{
			name: "validate should pass if all params are properly set",
			arg:  []string{platformViewerRole, deploymentsManagerRole},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRoles(tt.arg)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
		})
	}
}

func TestHasBothDeploymentRoles(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want bool
	}{
		{
			name: "should return true when both deployment roles are present",
			arg:  []string{deploymentsManagerRole, deploymentsViewerRole},
			want: true,
		},
		{
			name: "should return false if both deployment roles are not present",
			arg:  []string{deploymentsManagerRole, platformViewerRole, platformAdminRole},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasBothDeploymentRoles(tt.arg)

			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
