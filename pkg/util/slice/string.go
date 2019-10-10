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

package slice

// HasString returns true if the given string value is found in the provided
// slice, otherwise returns false.
func HasString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// IsEmpty returns true if a string slice is empty or contains empty string
// values, otherwise returns false.
func IsEmpty(slice []string) bool {
	for _, i := range slice {
		if i != "" {
			return false
		}
	}
	return true
}

// ContainsAll return true if a string slice contains all elements of another
// string slice, otherwise returns false.
func ContainsAll(slice, containedSlice []string) bool {
	var all = true
	for _, c := range containedSlice {
		all = all && HasString(slice, c)
	}
	return all
}
