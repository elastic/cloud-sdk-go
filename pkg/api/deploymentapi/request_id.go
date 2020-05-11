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

package deploymentapi

import "github.com/elastic/cloud-sdk-go/pkg/util/ec"

// RequestID creates a 64-character string when an empty s tring is provided,
// or returns the provided string. It is aimed to be used as a helper when
// creating a deployment to always provide a request ID.
func RequestID(s string) string {
	if s == "" {
		return ec.RandomResourceLength(64)
	}
	return s
}
