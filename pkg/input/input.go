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

package input

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

// NewFileOrReader tries to open a file if a non empty filename is passed as a
// parameter, otherwise a reader passed as the first argument is to be read and
// copied in a buffer that is returned. Any error along the way is returned.
func NewFileOrReader(reader io.Reader, filename string) (io.ReadCloser, error) {
	if filename != "" {
		return os.Open(filename)
	}

	if reader == nil {
		return nil, errors.New("cannot read from nil reader")
	}

	var buf = new(bytes.Buffer)
	if _, err := io.Copy(buf, reader); err != nil {
		return nil, err
	}
	return ioutil.NopCloser(buf), nil
}
