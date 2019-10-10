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
	"sync"
	"time"
)

// Worker is the structure that contains the configuration that a worker uses
// when it's spawned.
type Worker struct {
	// Work Queue where the worker obtains its work.
	Queue <-chan Validator
	// Stop channel is used to signal a worker to Stop processing items from
	// the queue.
	Stop <-chan struct{}
	// Stopped is the channel used to communicate back with the stopper to signal
	// that the worker has been successfully Stopped.
	Stopped chan<- bool
	// Finished channel is used to signal that the work has been completed. If
	// the worker happens to be stopping and the stopTimeout is hit before the
	// item has been processed by the run function the Finished signal is also
	// sent.
	Finished chan<- struct{}
	// error channel where the errors from the work will land.
	Errors chan<- error
	// Leftovers of any incompleted work items that the worker couldn't finish
	// before hitting the stop timeout.
	Leftovers chan<- Validator
	// Run is the function that the worker will Run on each work item received.
	Run RunFunc
	// controls the time.Duration to wait when a stop signal is received. If the
	// timeout is exceeded before the work is completed the current work will be
	// sent to the leftover queue.
	StopTimeout time.Duration
}

// StopParams is consumed by StopWorkers so a set of workers can be stopped.
type StopParams struct {
	// number of workers to stop
	Size int
	// Stop is used to send a signal to a worker to make it stop.
	Stop chan<- struct{}
	// Stopped is a signal given back by the worker that is being stopped
	// when it has stopped, either timing out or successfully.
	StoppedWithTimeout <-chan bool
}

// Work receives a Worker structure that contains the configuration to control
// the worker behaviour. It processes work using the worker.run function on
// worker.queue receive.
// When a stop signal is received it will wait the time.Duration defined by the
// stopTimeout and forcefully exit without waiting for the work to be completed.
func Work(worker Worker) {
	var timedout bool
	defer func() { worker.Stopped <- timedout }()

	var done = make(chan struct{})
	for {
		select {
		case task := <-worker.Queue:
			// This is necessary to allow the work to happen in the background
			// and still react to any sent to worker.stop and be able to clean
			// shutdown the worker even if the work is in flight.
			go func(w Validator, d chan<- struct{}) {
				if err := worker.Run(w); err != nil {
					worker.Errors <- err
				}
				d <- struct{}{}
			}(task, done)

			select {
			// Receives the done signal when the work is done.
			case <-done:
				worker.Finished <- struct{}{}

			// Handles the case where the worker is processing a work item and
			// and a stop signal is received while the work is in flight.
			case <-worker.Stop:
				// This gives one last chance to the worker to complete the
				// work before returning with a timeout and sending the item
				// to the leftover queue.
				select {
				case <-done:
					worker.Finished <- struct{}{}
				case <-time.After(worker.StopTimeout):
					worker.Leftovers <- task
					timedout = true
				}
				return
			}
		// Handle case where the worker is idle and the stop signal is received
		case <-worker.Stop:
			return
		}
	}
}

// StopWorkers stops all of the workers in parallel trying to honor their
// timeout settings. If the worker cannot be stopped before the params.timeout
// the function returns ErrStopOperationTimedOut.
func StopWorkers(params StopParams) error {
	var err = make(chan error, params.Size)
	var wg sync.WaitGroup
	wg.Add(params.Size)
	for index := 0; index < params.Size; index++ {
		go func() {
			defer wg.Done()
			params.Stop <- struct{}{}
			if v := <-params.StoppedWithTimeout; v {
				err <- ErrStopOperationTimedOut
			}
		}()
	}
	wg.Wait()
	close(err)

	return <-err
}
