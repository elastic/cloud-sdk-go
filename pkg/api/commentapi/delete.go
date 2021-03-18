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
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// Delete delete the comment of the given combination of comment ID, version, resource type and resource ID.
func Delete(params DeleteParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	_, err := params.V1API.Comments.DeleteComment(
		comments.NewDeleteCommentParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithVersion(ec.String(params.Version)).
			WithCommentID(params.CommentID).
			WithResourceType(params.ResourceType).
			WithResourceID(params.ResourceID),
		params.AuthWriter)

	if err != nil {
		return apierror.Wrap(err)
	}
	return nil
}
