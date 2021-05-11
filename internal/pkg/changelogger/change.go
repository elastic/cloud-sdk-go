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

package changelogger

type Changes []Change

type Change struct {
	// Required change category or change type, for a list of all supported
	// change types see matcher.go.
	Category string

	// Required change description.
	Description string

	// Optional title, if set, it will use the title.
	Title string

	// Optional reference for the change source, it must be a valid URL.
	// If not set, it will use the title of the note and resolve to the local
	// repository.
	Ref string
}

func (cs Changes) ByCategory(cat string) Changes {
	var result Changes
	for _, change := range cs {
		if change.Category == cat {
			result = append(result, change)
		}
	}

	return result
}

func (c Change) TitleOrRef() string {
	if c.Title != "" {
		return c.Title
	}
	return c.Ref
}

// Sort interface

func (cs Changes) Len() int           { return len(cs) }
func (cs Changes) Swap(i, j int)      { cs[i], cs[j] = cs[j], cs[i] }
func (cs Changes) Less(i, j int) bool { return cs[i].Ref < cs[j].Ref }
