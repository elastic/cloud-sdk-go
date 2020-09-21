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
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

const (
	// StoppedStatus is the default status for a pool. The pool is stopped
	// and no work is being processed.
	StoppedStatus = iota
	// StartingStatus represents the pool status when it's being started.
	// This means that the workers start processing the work incrementally.
	StartingStatus
	// StartedStatus all the workers have been started.
	StartedStatus
	// IdleStatus when all the workers have been started but no work is
	// being processed.
	IdleStatus
	// FinishedStatus all the workers are started and the queued work has
	// completely been processed.
	FinishedStatus
	// StoppingStatus represents the pool status when it's being stopped.
	// Some items might still be in flight.
	StoppingStatus
	// StoppedTimeout represents the pool status when has been stopped and
	// some of the workers have been forcefully stopped. This means that the
	// work that was being done by N workers was not finished, so the user that
	// consumes the Pool object might want to perform some checks or clean ups
	// to check which work wasn't completed.
	StoppedTimeout
	// StoppedSuccess represents the pool status when has been stopped without
	// hitting the stop timeout. The pool can still contain queued events that
	// have been moved to the leftover list.
	StoppedSuccess
)

var (
	// ErrAddOperationTimedOut is returned when the add timeout is exceeded
	ErrAddOperationTimedOut = errors.New("pool: failed adding work, queue full")
	// ErrStopOperationTimedOut is returned when the stop timeout is exceeded
	ErrStopOperationTimedOut = errors.New("pool: stop timeout exceeded")
	// ErrAlreadyStarted is returned when Start called on a non stopped pool
	ErrAlreadyStarted = errors.New("pool: cannot start a non stopped pool")
	// ErrAlreadyStopped is returned when Stop called on a stopped pool
	ErrAlreadyStopped = errors.New("pool: cannot stop an already stopped pool")
	// ErrAlreadyStopping is returned when Stop called on a stopping pool
	ErrAlreadyStopping = errors.New("pool: cannot stop a stopping pool")
	// ErrCannotAddWorkToStoppingPool is returned when work is added to a stopped
	// pool
	ErrCannotAddWorkToStoppingPool = errors.New("pool: cannot add work to stopping pool")
	// ErrCannotWaitOnStoppedPool is thrown by Wait() when the pool is stopped.
	ErrCannotWaitOnStoppedPool = errors.New("pool: cannot wait for workers to finish on a stopped pool")
	// ErrCannotGetLeftovers is returned when the pool is not in a stopped state.
	ErrCannotGetLeftovers = errors.New("pool: cannot get the work leftovers on a non stopped pool")

	failFastSetStopMsg = `pool: fail fast is set and received an error, stopping pool...`
)

var (
	defaultStatus = "unknown"
	statusMap     = map[uint32]string{
		StoppedStatus:  "stopped",
		StartingStatus: "starting",
		IdleStatus:     "idle",
		StartedStatus:  "running",
		FinishedStatus: "finished",
		StoppingStatus: "stopping",
		StoppedTimeout: "stopped timeout",
		StoppedSuccess: "stopped success",
	}
)

// Pool is a generic worker pool implementation that can be used to complete a
// series of tasks concurrently and obtain any errors that have been returned
// by the workers. The usage of the pool is quite simple by itself and relies
// on the constructor function NewPool().
type Pool struct {
	// Number of workers to create within the worker pool.
	size uint16

	// RunFunc that is used by each worker to process work.
	run RunFunc

	// internal queue used to feed the work to the worker pool.
	queue chan Validator

	// leftovers from a previously stopped pool, meaning work that did not
	// get processed and work that workers were processing and didn't complete.
	// The unfinished items will be ordered first, followed by the queue
	// contents that didn't get processed.
	leftovers chan Validator

	// error channel where all of the worker errors are received.
	errors chan error

	// signals is the structure that contains the sync signals that are
	// used to trigger changes in the Pool
	signals Signals

	// state contains the internal pool state.
	state *State

	// Pool timeouts for different
	timeouts Timeout

	// writer where any (log, info) messages will be sent.
	writer io.Writer

	// FailFast can be set to stop all the pool when any of the workers returns
	// with an error.
	failFast bool
}

// Signals contains all of the channels that are used to trigger different
// status changes in the Pool.
type Signals struct {
	// Stop a channel that is used to signal workers to to be Stop.
	Stop chan struct{}
	// Stopped is a channel that is used for backwards communication with the
	// stopper to verify that the worker has been Stopped.
	Stopped chan bool
	// Finish is a channel used by workers to signal that they've finished
	// processing a task.
	Finish chan struct{}
	// Added is a channel used by the pool to signal that a work item has been
	// pushed to the queue for processing.
	Added chan struct{}
	// StopMonitor is used to stop the monitoring goroutine that updates the
	// pool's internal state.
	StopMonitor chan struct{}
}

// State contains the pool State
type State struct {
	// Number of work items that have been added to the queue.
	Queued Counter
	// Number of work items that have been Processed by a worker.
	Processed Counter
	// Pool global Status
	Status Counter
	// monitoring is a condition that checks if the pool monitor is running.
	monitoring bool
	// Errors that have been returned by the worker.
	Errors *Errors
}

// Errors wraps a multierror.Error with a Mutex so that it can be safely used
// when accesses
type Errors struct {
	err *multierror.Error
	mu  sync.RWMutex
}

// Add appends a new error to the error list
func (e *Errors) Add(errs ...error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.err == nil {
		e.err = new(multierror.Error)
	}
	e.err = multierror.Append(e.err, errs...)
}

// Error returns an error or nil
func (e *Errors) Error() error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.err == nil {
		return nil
	}
	return e.err.ErrorOrNil()
}

// Counter represents a safe uint32 that can be used as a shared counter.
type Counter struct {
	value uint32
}

// Add increments the counter
func (c *Counter) Add(incr uint32) {
	atomic.AddUint32(&c.value, incr)
}

// Set overwrites the value of the counter in favour of the passed value.
func (c *Counter) Set(n uint32) {
	atomic.SwapUint32(&c.value, n)
}

// Get obtains the value of the counter
func (c *Counter) Get() uint32 {
	return atomic.LoadUint32(&c.value)
}

// NewPool initializes a new Pool from a set of parameters.
func NewPool(params Params) (*Pool, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// this buffer roughly means that each worker can have 128 items
	// in the queue
	var queueBuffer = params.Size * 128
	// Since there can be in flight items already picked up by the workers
	// the max leftover size should be the buffer + the number of workers.
	var leftoverBuffer = queueBuffer + params.Size

	var pool = Pool{
		size:      params.Size,
		run:       params.Run,
		timeouts:  params.Timeout,
		queue:     make(chan Validator, queueBuffer),
		leftovers: make(chan Validator, leftoverBuffer),
		signals: Signals{
			Stop:        make(chan struct{}),
			Stopped:     make(chan bool),
			Finish:      make(chan struct{}),
			Added:       make(chan struct{}),
			StopMonitor: make(chan struct{}),
		},
		state: &State{
			Errors:     new(Errors),
			monitoring: true,
		},
		errors:   make(chan error),
		writer:   params.Writer,
		failFast: params.FailFast,
	}

	go pool.monitor(nil)

	return &pool, nil
}

// Start starts the workers in the worker pool, and starts all of the internal
// goroutines that the pool relies in.
func (p *Pool) Start() error {
	if p.Status() > StoppedStatus && p.Status() < StoppedTimeout {
		return ErrAlreadyStarted
	}

	if !p.state.monitoring {
		go p.monitor(nil)
	}

	// If the pool was previously stopped, recreate the channels
	if p.Status() > StoppedTimeout {
		if len(p.queue) == 0 {
			p.queue = make(chan Validator, cap(p.queue))
		}

		if len(p.leftovers) == 0 {
			p.leftovers = make(chan Validator, cap(p.leftovers))
		}
	}

	p.setStatus(StartingStatus)
	for worker := 0; worker < int(p.size); worker++ {
		<-StartWorker(Worker{
			Queue:       p.queue,
			Stop:        p.signals.Stop,
			Stopped:     p.signals.Stopped,
			Finished:    p.signals.Finish,
			Errors:      p.errors,
			Leftovers:   p.leftovers,
			Run:         p.run,
			StopTimeout: p.timeouts.Stop,
		})
	}

	return nil
}

// StartWorker starts a worker in the background waiting for the goroutine to
// actually be schedules. It returns a channel that can be used to wait until
// the Goroutine has been run as in the code below:
//   wait := StartWorker(Worker{})
//   // This blocks execution
//   <-wait
func StartWorker(worker Worker) chan struct{} {
	var spawned = make(chan struct{})
	go func(c chan<- struct{}) {
		go Work(worker)
		c <- struct{}{}
	}(spawned)
	return spawned
}

// monitor receives signals from the Finish and Added signal channels and
// updates the counters accordingly. It also listens to any errors that might
// be received through the error channel and adds them to the error state.
// Also monitors an interrupt channel that stops the pool on interrupt, Kill
// and SIGTERM signals.
func (p *Pool) monitor(interrupt chan os.Signal) {
	if interrupt == nil {
		interrupt = make(chan os.Signal, 1)
	}
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	var interrupted bool
	for {
		select {
		case <-p.signals.Finish:
			p.state.Processed.Add(1)
		case <-p.signals.Added:
			p.state.Queued.Add(1)
		case err := <-p.errors:
			p.state.Errors.Add(err)
			if p.failFast {
				if p.writer != nil {
					fmt.Fprintln(p.writer, failFastSetStopMsg)
				}
				//nolint
				go p.Stop()
			}
		case <-interrupt:
			if interrupted {
				continue
			}
			interrupted = true
			if p.writer != nil {
				fmt.Fprintln(p.writer, "pool: received interrupt, stopping pool...")
			}
			//nolint
			go p.Stop()
		case <-p.signals.StopMonitor:
			p.state.monitoring = false
			break
		}
		p.updateState()
	}
}

// updateState updates the Pool state when the pool has already been started
// setting the state to Idle or Finished.
func (p *Pool) updateState() {
	if p.Status() >= StoppingStatus || p.Status() == StoppedStatus {
		return
	}

	var workersIdle = p.state.Processed.Get() == p.state.Queued.Get()
	var emptyQueue = p.state.Queued.Get() == 0
	var workersProcessing = p.state.Queued.Get() > p.state.Processed.Get()

	if workersProcessing {
		p.setStatus(StartedStatus)
	}

	if emptyQueue && workersIdle && p.Status() != IdleStatus {
		p.setStatus(IdleStatus)
	}

	if workersIdle && !emptyQueue && p.Status() != FinishedStatus {
		p.setStatus(FinishedStatus)
	}
}

// setStatus sets the pool to the specified status
func (p *Pool) setStatus(status int) {
	p.state.Status.Set(uint32(status))
}

// Wait will stop execution until the pool has finished processing all of the
// work that it had in the queue. It returns any errors that the workers might
// have returned.
func (p *Pool) Wait() error {
	if p.Status() == StoppedStatus {
		return ErrCannotWaitOnStoppedPool
	}

	for {
		<-time.After(time.Millisecond)
		if p.Status() == FinishedStatus || isStopped(p.Status()) {
			return p.Result()
		}
	}
}

// Result returns the results from the work that was done by the workers,
// namely returns any error in the multierror format.
func (p *Pool) Result() error {
	if p.state.Errors != nil {
		return p.state.Errors.Error()
	}
	return nil
}

// Leftovers obtains a list of unfinished work with the following order:
// First it returns any items that might have been in flight and did not
// complete before hitting the stop timeout.
// Following those, the items that did not get processed by a worker.
// This function can only be called after the pool has been stopped.
func (p *Pool) Leftovers() ([]Validator, error) {
	if p.state.Status.Get() < StoppedTimeout {
		return nil, ErrCannotGetLeftovers
	}

	// Close the channel before entering the loop, so the loop ranges
	// over the buffered items and exits when the channel has no more
	// items in it.
	close(p.leftovers)

	var leftovers = make([]Validator, 0, len(p.leftovers))
	for l := range p.leftovers {
		leftovers = append(leftovers, l)
	}

	return leftovers, nil
}

// isStopped determines if the status is stopped.
func isStopped(s uint32) bool {
	return s >= StoppedTimeout || s == StoppedStatus
}

// Stop attempts to gracefully shutdown the workers from the pool. If the stop
// timeout is reached, the work that was being processed by the worker is sent
// to the leftover queue as are any items that were not processed, returning
// ErrStopOperationTimedOut.
func (p *Pool) Stop() error {
	if isStopped(p.Status()) {
		return ErrAlreadyStopped
	}

	if p.Status() == StoppingStatus {
		return ErrAlreadyStopping
	}

	p.setStatus(StoppingStatus)
	err := StopWorkers(StopParams{
		Size:               int(p.size),
		Stop:               p.signals.Stop,
		StoppedWithTimeout: p.signals.Stopped,
	})

	close(p.queue)
	drain(p.queue, p.leftovers)

	if err != nil && err == ErrStopOperationTimedOut {
		p.setStatus(StoppedTimeout)
	}
	if err == nil {
		p.setStatus(StoppedSuccess)
	}

	p.signals.StopMonitor <- struct{}{}

	return err
}

// drain dumps the items from the first queue to the second queue, it assumes
// that the channel is already closed, since this function is only useful for
// buffered queues.
func drain(from <-chan Validator, to chan<- Validator) {
	for w := range from {
		to <- w
	}
}

// Status returns the numeric status of the pool
func (p *Pool) Status() uint32 {
	return p.state.Status.Get()
}

// StatusText obtains the current pool status as a string, for all available
// states see the statusMap which contains the mappings from int to string.
func StatusText(status uint32) string {
	if status, ok := statusMap[status]; ok {
		return status
	}

	return defaultStatus
}

// Add adds N amount of work to the pool's queue, timing out if the queue is
// full for more than the defined timeout.Add. If an error is returned it will
// be ErrAddingOperationTimedOut, meaning that the first parameter is the list
// of work that didn't get added, leaving any possible retries to add work to
// the user.
func (p *Pool) Add(work ...Validator) ([]Validator, error) {
	if p.Status() >= StoppingStatus {
		return work, ErrCannotAddWorkToStoppingPool
	}

	var err error
	var leftover []Validator
	for _, w := range work {
		select {
		case p.queue <- w:
			p.signals.Added <- struct{}{}
		case <-time.After(p.timeouts.Add):
			err = ErrAddOperationTimedOut
			leftover = append(leftover, w)
		}
	}

	return leftover, err
}
