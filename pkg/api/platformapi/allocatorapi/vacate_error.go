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

package allocatorapi

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
)

const (
	resourcePrefixErrFmt  = "allocator %s: resource id [%s][%s]: %s"
	allocatorPrefixErrFmt = "allocator %s: %s"
)

// VacateError wraps an error and enriches it with some fields to be able to
// have a consistent output (resourcePrefixErrFmt). Or by being serialized to
// JSON by using `multierror.WithFormat`.
type VacateError struct {
	AllocatorID string `json:"allocator_id,omitempty"`
	ResourceID  string `json:"resource_id,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Ctx         string `json:"context,omitempty"`
	Err         error  `json:"error,omitempty"`
}

func (e VacateError) Error() string {
	var err = e.Err.Error()

	if e.Ctx != "" {
		err = fmt.Sprintf("%s: %s", e.Ctx, err)
	}

	if e.Kind == "" {
		e.Kind = "unknown"
	}

	if e.AllocatorID != "" && e.ResourceID == "" && e.Kind == "" {
		return fmt.Sprintf(allocatorPrefixErrFmt, e.AllocatorID, err)
	}

	return fmt.Sprintf(resourcePrefixErrFmt,
		e.AllocatorID, e.ResourceID, e.Kind, err,
	)
}

// MarshalJSON Implements the json.Marshaler interface to be able to modify the
// error and format it in JSON.w
func (e VacateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		AllocatorID string `json:"allocator_id,omitempty"`
		ResourceID  string `json:"resource_id,omitempty"`
		Kind        string `json:"kind,omitempty"`
		Ctx         string `json:"context,omitempty"`
		Err         error  `json:"error,omitempty"`
	}{
		AllocatorID: e.AllocatorID,
		ResourceID:  e.ResourceID,
		Kind:        e.Kind,
		Ctx:         e.Ctx,
		Err:         apierror.NewJSONError(e.Err),
	})
}
