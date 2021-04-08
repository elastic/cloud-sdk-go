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

const minsize = 0

// ParseGb converts the stringified gigabyte size notation to an int32 Megabyte
// notation. The minimum size allowed is 0g with 0.5g increments.
func ParseGb(strSize string) (int32, error) {
	strSize = strings.ToLower(strSize)
	if strSize == "0" {
		return 0, nil
	}

	re := regexp.MustCompile(`(?m)(.*\w)(g)`)
	matches := re.FindStringSubmatch(strSize)
	if len(matches) < 2 {
		return 0, fmt.Errorf(`failed to convert "%s" to <size><g>`, strSize)
	}

	// Pops the first item out when the first match (all groups) are contained
	// within the full string.
	if strings.Contains(strSize, matches[0]) {
		matches = matches[1:]
	}

	rawSize, err := strconv.ParseFloat(matches[0], 32)
	if err != nil {
		return 0, err
	}

	var size = int32(rawSize * 1024)
	if size < minsize {
		return 0, fmt.Errorf(`size "%s" is invalid: minimum size is %.1fg`, strSize, float32(minsize)/1024)
	}

	if size%512 > 0 {
		return 0, fmt.Errorf(`size "%s" is invalid: only increments of 0.5g are permitted`, strSize)
	}

	return size, nil
}
