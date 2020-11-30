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

import "strings"

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

// StringSlice wraps a string slice to provide methods on top of it
type StringSlice []string

// ToMap iterates over the slice and slices each element into substrings separated by sep. Then it adds a map key value
// pair where the key is the first substring and value is the second substring
//
// If the sep is not found in the slice element then it's ignored
//
// It returns the produced map after iterating all slice elements
func (ss StringSlice) ToMap(sep string) map[string]string {
	newMap := make(map[string]string)

	for _, s := range ss {
		sliced := strings.Split(s, sep)
		if len(sliced) == 2 {
			newMap[sliced[0]] = sliced[1]
		}
	}

	return newMap
}

// IsEmpty returns true if a string slice is empty or contains empty string
// values, otherwise returns false.
func (ss StringSlice) IsEmpty() bool {
	return IsEmpty(ss)
}

// Contains returns true if the given string value is found in the provided
// slice, otherwise returns false.
func (ss StringSlice) Contains(s string) bool {
	return HasString(ss, s)
}

// ToCommaSeparatedString returns the elements of the string slice as a single string where all values are separated
// by the comma.
func (ss StringSlice) ToCommaSeparatedString() string {
	return strings.Join(ss, ",")
}
