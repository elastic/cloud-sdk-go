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

package apierror

import (
	"fmt"
	"strings"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// newBasicFailedReplyMultierror returns a multierror from a BasicFailedReply.
func newBasicFailedReplyMultierror(prefix string, r *models.BasicFailedReply) *multierror.Prefixed {
	merr := multierror.NewPrefixed(prefix)
	if r == nil {
		return merr
	}

	for _, e := range r.Errors {
		merr = merr.Append(newBasicFailedReply(e))
	}

	return merr
}

func newBasicFailedReply(elem *models.BasicFailedReplyElement) error {
	var code, message = "unknown", "unknown"
	var fields string

	if elem.Code != nil {
		code = *elem.Code
	}

	if elem.Message != nil {
		message = *elem.Message
	}

	if elem.Fields != nil {
		fields = strings.Join(elem.Fields, ", ")
	}

	if fields != "" {
		return fmt.Errorf("%s: %s (%s)", code, message, fields)
	}

	return fmt.Errorf("%s: %s", code, message)
}
