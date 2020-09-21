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

package pool_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/sync/pool"
)

type nilValidator struct {
	err error
}

func (v nilValidator) Validate() error { return v.err }

var output io.Writer = os.Stdout

func TestExamplePool(t *testing.T) {
	t.Run("testExample", func(t *testing.T) {
		var b = new(bytes.Buffer)
		output = b
		var want = new(bytes.Buffer)
		want.WriteString("stopped\n")
		want.WriteString("starting\n")
		want.WriteString("running\n")
		want.WriteString("finished\n")
		want.WriteString("stopped success\n")
		want.WriteString("leftovers: 0 leftovers errors: <nil>\n")

		ExamplePool()

		if b.String() != want.String() {
			t.Errorf("ExamplePool = %v, want %v", b.String(), want.String())
		}
	})

	t.Run("ExamplePool_failfast", func(t *testing.T) {
		var b = new(bytes.Buffer)
		output = b
		var want = new(bytes.Buffer)
		want.WriteString("stopped\n")
		want.WriteString("starting\n")
		want.WriteString("running\n")
		want.WriteString("execution errors: 1 error occurred:\n\t* first error\n\n\n")
		want.WriteString("stopped timeout\n")
		want.WriteString("stopped timeout\n")
		want.WriteString("leftovers: 3 leftovers errors: <nil>\n")

		ExamplePool_failfast()

		if b.String() != want.String() {
			t.Errorf("ExamplePool = %v, want %v", b.String(), want.String())
		}
	})
}

// This example shows how to create a new Pool and work with it.
func ExamplePool() {
	p, err := pool.NewPool(pool.Params{
		Size: 2,
		Run: func(params pool.Validator) error {
			<-time.After(time.Millisecond * 10)
			return params.Validate()
		},
		Timeout: pool.Timeout{
			Add:  time.Millisecond,
			Stop: time.Millisecond,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(output, pool.StatusText(p.Status()))

	// Start the pool so the workers start processing
	if err := p.Start(); err != nil {
		panic(err)
	}

	fmt.Fprintln(output, pool.StatusText(p.Status()))

	var work = []pool.Validator{
		new(nilValidator),
		new(nilValidator),
	}
	// Try to add work
	leftovers, err := p.Add(work...)
	if err != nil && err != pool.ErrAddOperationTimedOut {
		panic(err)
	}

	// Ensure there's no leftovers
	if len(leftovers) > 0 {
		for {
			leftovers, _ := p.Add(leftovers...)
			if len(leftovers) == 0 {
				break
			}
		}
	}

	<-time.After(time.Millisecond)
	fmt.Fprintln(output, pool.StatusText(p.Status()))

	// Wait until all of the work is consumed
	if err = p.Wait(); err != nil {
		fmt.Fprintln(output, "execution errors:", err.Error())
	}

	fmt.Fprintln(output, pool.StatusText(p.Status()))

	p.Stop()

	<-time.After(time.Millisecond)
	fmt.Fprintln(output, pool.StatusText(p.Status()))
	l, err := p.Leftovers()
	fmt.Fprintln(output, "leftovers:", len(l), "leftovers errors:", err)
}

// This example shows how to create a new Pool which stops processing work when
// an error is returned by a worker.
func ExamplePool_failfast() {
	p, err := pool.NewPool(pool.Params{
		Size: 2,
		Run: func(params pool.Validator) error {
			<-time.After(time.Millisecond * 10)
			return params.Validate()
		},
		Timeout: pool.Timeout{
			Add:  time.Millisecond,
			Stop: time.Millisecond,
		},
		// Setting FailFast will cause the pool to stop processing the queued
		// work and return the worker error when a worker returns with error.
		FailFast: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(output, pool.StatusText(p.Status()))

	// Start the pool so the workers start processing
	if err := p.Start(); err != nil {
		panic(err)
	}

	fmt.Fprintln(output, pool.StatusText(p.Status()))

	var work = []pool.Validator{
		new(nilValidator),
		&nilValidator{err: errors.New("first error")},
		new(nilValidator),
		new(nilValidator),
		&nilValidator{err: errors.New("last error")},
	}
	// Try to add work
	leftovers, err := p.Add(work...)
	if err != nil && err != pool.ErrAddOperationTimedOut {
		panic(err)
	}

	// Ensure there's no leftovers
	if len(leftovers) > 0 {
		for {
			leftovers, _ := p.Add(leftovers...)
			if len(leftovers) == 0 {
				break
			}
		}
	}

	<-time.After(time.Millisecond)
	fmt.Fprintln(output, pool.StatusText(p.Status()))

	// Wait until all of the work is consumed
	if err = p.Wait(); err != nil {
		fmt.Fprintln(output, "execution errors:", err.Error())
	}

	fmt.Fprintln(output, pool.StatusText(p.Status()))

	p.Stop()

	<-time.After(time.Millisecond)
	fmt.Fprintln(output, pool.StatusText(p.Status()))
	l, err := p.Leftovers()
	fmt.Fprintln(output, "leftovers:", len(l), "leftovers errors:", err)
}
