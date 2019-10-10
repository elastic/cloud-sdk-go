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

package maps

import (
	"fmt"
)

// KeysToStringSlice accepts a map interface and returns its keys as a slice
// It supports the following map interfaces
// 1. map[string]string
// 2. map[string]int
// 1. map[string]bool
// 2. map[string]interface{}
func KeysToStringSlice(m interface{}) ([]string, error) {
	switch m := m.(type) {
	case map[string]string:
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		return keys, nil
	case map[string]int:
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		return keys, nil
	case map[string]bool:
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		return keys, nil
	case map[string]interface{}:
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		return keys, nil
	default:
		return nil, fmt.Errorf("unexpected key type %T", m)
	}
}

// Filter filters a string-key map values based on the given string
func Filter(m map[string]interface{}, filter string) map[string]interface{} {
	var filtered = make(map[string]interface{})
	for k, v := range m {
		if k[:len(filter)] == filter {
			filtered[k] = v
		}
	}
	return filtered
}

// HasKey returns true if a string-key map has a specific value, else false
func HasKey(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}
