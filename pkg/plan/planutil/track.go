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

package planutil

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/elastic/cloud-sdk-go/pkg/plan"
	"github.com/elastic/cloud-sdk-go/pkg/util/slice"

	"github.com/hashicorp/go-multierror"
)

var validFormats = []string{"json", "text"}

// TrackChangeParams is consumed by TrackChange.
type TrackChangeParams struct {
	plan.TrackChangeParams
	Writer io.Writer
	Format string
}

// Validate ensures the parameters are usable by the consuming function.
// Note this doesn't validate plan.TrackChangeParams as it's already validated
// when plan.TrackChange is called.
func (params TrackChangeParams) Validate() error {
	var merr = new(multierror.Error)

	if params.Format != "" && params.Writer == nil {
		merr = multierror.Append(merr,
			errors.New("planutil track change: writer needs to be specified when format is not empty"),
		)
	}

	if !slice.HasString(validFormats, params.Format) && params.Format != "" {
		merr = multierror.Append(merr,
			fmt.Errorf("planutil track change: invalid format \"%s\"", params.Format),
		)
	}

	return merr.ErrorOrNil()
}

// TrackChange combines the plan.TrackChange and plan.Stream with configurable
// format and writer ouptuts.
func TrackChange(params TrackChangeParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	channel, err := plan.TrackChange(params.TrackChangeParams)
	if err != nil {
		return err
	}

	if params.Format == "json" {
		return plan.StreamJSON(channel, params.Writer, false)
	}

	if params.Format == "text" {
		return plan.Stream(channel, params.Writer)
	}

	return plan.StreamJSON(channel, ioutil.Discard, false)
}
