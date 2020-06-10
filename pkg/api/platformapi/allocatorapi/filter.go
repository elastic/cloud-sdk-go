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
	"strings"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// FilterByTag filters a list of allocators based on metadata tags and returns a new matched list
func FilterByTag(tags map[string]string, allocators []*models.AllocatorInfo) []*models.AllocatorInfo {
	var matchedAllocators []*models.AllocatorInfo

	if len(tags) > 0 {
		for _, a := range allocators {
			matches := 0
			for k, v := range tags {
				for _, m := range a.Metadata {
					if *m.Key == k && *m.Value == v {
						matches++
						break
					}
				}
			}
			if matches == len(tags) {
				matchedAllocators = append(matchedAllocators, a)
			}
		}
		return matchedAllocators
	}

	return allocators
}

// FilterConnectedOrWithInstances filters a list of allocators and returns only the connected ones or those who have more than one instance
func FilterConnectedOrWithInstances(allocators []*models.AllocatorInfo) []*models.AllocatorInfo {
	var matchedAllocators []*models.AllocatorInfo
	for _, a := range allocators {
		if *a.Status.Connected || len(a.Instances) > 0 {
			matchedAllocators = append(matchedAllocators, a)
		}
	}
	return matchedAllocators
}

func tagsToMap(filterArgs string) map[string]string {
	filterArgs = strings.ReplaceAll(filterArgs, "[", "")
	filterArgs = strings.ReplaceAll(filterArgs, "]", "")
	tags := strings.Split(filterArgs, ",")
	var tagsMap = make(map[string]string)

	for _, t := range tags {
		tag := strings.Split(t, ":")
		if len(tag) == 2 {
			tagKey := tag[0]
			tagValue := tag[1]

			tagsMap[tagKey] = tagValue
		}
	}

	return tagsMap
}
