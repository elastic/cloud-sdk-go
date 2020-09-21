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

package pool

import (
	"io"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// DefaultTimeout is a set of default timeout settings that can be used to
// create a new Pool.
var DefaultTimeout = Timeout{
	Add:  time.Second * 30,
	Stop: time.Second * 30,
}

var (
	errPoolSizeCannotBeZero        = errors.New("pool: size cannot be 0")
	errPoolStopTimeoutCannotBeZero = errors.New("pool: stop timeout cannot be 0")
	errPoolAddTimeoutCannotBeZero  = errors.New("pool: add timeout cannot be 0")
	errPoolRunFuncCannotBeNil      = errors.New("pool: run function cannot be nil")
)

// Validator interface is consumed by the RunFunc.
type Validator interface {
	Validate() error
}

// RunFunc is the function that the pool will use as the worker.
type RunFunc func(params Validator) error

// Params is used to configure a Pool
type Params struct {
	// Size controls how many concurrent operations are running at the same
	// time.
	Size uint16
	// Run represents the actual function that will be run by each worker in
	// the worker pool.
	Run RunFunc
	// Timeout structure that controls the different timeout times.
	Timeout Timeout

	// Writer is the device where any (log, info) messages will be sent.
	Writer io.Writer

	// FailFast can be set to stop all the pool when any of the workers returns
	// with an error.
	FailFast bool
}

// Timeout is an object that encloses different Pool operation timeouts.
type Timeout struct {
	// Add timeout per Add operation, used to add items to the queue, the
	// timeout is evaluated per work item. Must be greater than 0s.
	Add time.Duration
	// Stop timeout that is used when stopping the workers, this timeout
	// is evaluated per worker, so the global timeout is * N workers. Must be
	// greater than 0s.
	Stop time.Duration
}

// Validate verifies that the parameters are valid and returns a multierror if
// any invalid parameters are found.
func (params Params) Validate() error {
	var merr = new(multierror.Error)

	if params.Size == 0 {
		merr = multierror.Append(merr, errPoolSizeCannotBeZero)
	}

	if params.Timeout.Stop.Seconds() == 0 {
		merr = multierror.Append(merr, errPoolStopTimeoutCannotBeZero)
	}

	if params.Timeout.Add.Seconds() == 0 {
		merr = multierror.Append(merr, errPoolAddTimeoutCannotBeZero)
	}

	if params.Run == nil {
		merr = multierror.Append(merr, errPoolRunFuncCannotBeNil)
	}

	return merr.ErrorOrNil()
}
