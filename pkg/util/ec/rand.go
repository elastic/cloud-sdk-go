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

package ec

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomResourceID generates a random string of 32 characters which emulates
// a real Elastic Cloud resource ID.
func RandomResourceID() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[seededRand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// RandomResourceLength generates a random string of n characters.
func RandomResourceLength(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[seededRand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
