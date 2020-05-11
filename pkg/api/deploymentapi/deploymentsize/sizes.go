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

package deploymentsize

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const minsize = 512

// Parse converts the stringified size notation to an int32 Megabyte value.
// The minimum size allowed is 512 Megabytes.
func Parse(strSize string) (int32, error) {
	re := regexp.MustCompile(`(?m)(.*\w)(g|m)`)
	matches := re.FindStringSubmatch(strings.ToLower(strSize))
	if len(matches) < 2 {
		fmt.Println(matches, "length", len(matches))
		return 0, fmt.Errorf(`failed to convert "%s" to <size><g|m>`, strSize)
	}

	// Pops the first item out when the first match (all groups) are contained
	// within the full string.
	if strings.Contains(strSize, matches[0]) {
		matches = matches[1:]
	}

	var size int32
	sizeBody, sizeUnit := matches[0], matches[1]

	rawSize, err := strconv.ParseFloat(sizeBody, 32)
	if err != nil {
		return 0, err
	}

	switch sizeUnit {
	case "g":
		size = int32(rawSize * 1024)
	case "m":
		size = int32(rawSize)
	}

	if size < minsize {
		return 0, fmt.Errorf(`size %d is invalid, minimum size is %d`, size, minsize)
	}

	return size, nil
}
