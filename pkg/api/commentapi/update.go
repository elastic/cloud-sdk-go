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

package commentapi

import (
	"context"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/comments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// Update creates a comment to the given resource and returns a reference ID.
func Update(params UpdateParams) (string, error) {
	if err := params.Validate(); err != nil {
		return "", err
	}

	res, err := params.V1API.Comments.UpdateComment(
		comments.NewUpdateCommentParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithCommentID(params.CommentID).
			WithVersion(ec.String(params.Version)).
			WithBody(&models.CommentUpdateRequest{Message: ec.String(params.Message)}).
			WithResourceType(params.ResourceType).
			WithResourceID(params.ResourceID),
		params.AuthWriter)

	if err != nil {
		return "", apierror.Wrap(err)
	}
	return *res.GetPayload().ID, nil
}
