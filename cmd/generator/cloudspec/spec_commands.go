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

package cloudspec

import (
	"fmt"
	"io"
	"sort"

	"github.com/go-openapi/spec"
)

const header = `# ECE %s Command Declaration

Total of %d commands found

## Command list

`

// GetCommands generates a markdown formatted list of the swagger spec'd paths.
func GetCommands(cloudSpec *spec.Swagger, w io.Writer, version string) {
	var commands = make([]string, 0, len(cloudSpec.Paths.Paths))
	// Iterate over the paths and parameters as well.
	// nolint rangeValCopy: each iteration copies 128 bytes (consider pointers or indexing)
	for path, pathItem := range cloudSpec.Paths.Paths {
		var command []string
		if pathItem.Head != nil {
			command = append(command, bulletCommand(pathItem.Head.ID, path, "HEAD"))
		}
		if pathItem.Get != nil {
			command = append(command, bulletCommand(pathItem.Get.ID, path, "GET"))
		}
		if pathItem.Post != nil {
			command = append(command, bulletCommand(pathItem.Post.ID, path, "POST"))
		}
		if pathItem.Put != nil {
			command = append(command, bulletCommand(pathItem.Put.ID, path, "PUT"))
		}
		if pathItem.Patch != nil {
			command = append(command, bulletCommand(pathItem.Patch.ID, path, "PATCH"))
		}
		if pathItem.Delete != nil {
			command = append(command, bulletCommand(pathItem.Delete.ID, path, "DELETE"))
		}

		commands = append(commands, command...)
	}

	sort.SliceStable(commands, func(i, j int) bool {
		return commands[i] < commands[j]
	})

	fmt.Fprintf(w, header, version, len(commands))

	for i := range commands {
		fmt.Fprintln(w, commands[i])
	}
}

func bulletCommand(id, path, method string) string {
	return fmt.Sprint("- `", id, "` (**", method, "** ", path, ")")
}
