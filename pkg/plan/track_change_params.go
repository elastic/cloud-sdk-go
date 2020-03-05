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

package plan

import (
	"errors"
	"time"

	multierror "github.com/hashicorp/go-multierror"

	"github.com/elastic/cloud-sdk-go/pkg/api"
)

// TrackChangeParams is consumed by TrackChange. It can be used to track
// a whole deployment's resource changes or a specific resource change.
// DeploymentID and ResourceID with Kind cannot be used at the same time.
type TrackChangeParams struct {
	*api.API

	// DeploymentID to track its resource changes. Incompatible with ResourceID
	// and Kind.
	DeploymentID string

	// ResourceID to track (Formerly Cluster ID). Incompatible with DeploymentID
	ResourceID string

	// Resource kind to track. Incompatible with DeploymentID
	Kind string

	// IgnoreDownstream if set, will skip sending updates for any workload
	// plan changes other than the specified Kind. Only takes effect when Kind
	// and ResourceID is set.
	IgnoreDownstream bool

	// Tracking settings
	Config TrackFrequencyConfig
}

// TrackFrequencyConfig controls how the TrackChange function polls the API.
type TrackFrequencyConfig struct {
	// PollFrequency is the duration to use to poll the API for new changes
	// on the pending plan. The recommended setting is from 2 to 30 seconds.
	PollFrequency time.Duration

	// If set to > 1, allows up to that number of errors coming from the API.
	// It controls how many API errors can be tolerated. Or how many times
	// the polling has to come back with no changes in order to consider the
	// plan change finished.
	MaxRetries int
}

// Validate ensures the parameters are usable by the consuming function.
func (params TrackChangeParams) Validate() error {
	var merr = new(multierror.Error)
	var emptyDeploymentID = params.DeploymentID == ""
	var emptyResourceID = params.ResourceID == ""
	var emptyKind = params.Kind == ""

	if params.API == nil {
		merr = multierror.Append(merr, errors.New("plan track change: API cannot be nil"))
	}

	if emptyDeploymentID && emptyResourceID {
		merr = multierror.Append(merr,
			errors.New("plan track change: one of DeploymentID or ResourceID must be specified"),
		)
	}

	if !emptyDeploymentID && !emptyResourceID {
		merr = multierror.Append(merr,
			errors.New("plan track change: cannot specify both DeploymentID and ResourceID"),
		)
	}

	if emptyDeploymentID && emptyKind {
		merr = multierror.Append(merr,
			errors.New("plan track change: Kind cannot be empty"),
		)
	}

	merr = multierror.Append(merr, params.Config.Validate())

	return merr.ErrorOrNil()
}

// Validate ensures the parameters are usable by the consuming function.
func (params *TrackFrequencyConfig) Validate() error {
	var merr = new(multierror.Error)
	if params.MaxRetries <= 0 {
		merr = multierror.Append(merr, errors.New("plan track change: max retries must be at least 1"))
	}

	if params.PollFrequency.Nanoseconds() < 1 {
		merr = multierror.Append(merr, errors.New("plan track change: poll frequency must be at least 1 nanosecond"))
	}
	return merr.ErrorOrNil()
}

func (params *TrackFrequencyConfig) fillDefaults() {
	if params.MaxRetries <= 0 {
		params.MaxRetries = 1
	}

	if params.PollFrequency.Nanoseconds() < 1 {
		params.PollFrequency = time.Nanosecond * 1
	}
}
