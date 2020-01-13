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
	"bytes"
	"os"
	"reflect"
	"syscall"
	"testing"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

var testTimeout = Timeout{
	Add:  time.Millisecond,
	Stop: time.Millisecond,
}

type nilValidator struct{}

func (v nilValidator) Validate() error {
	return nil
}

func TestNewPool(t *testing.T) {
	type args struct {
		params Params
	}

	success := new(mockRun).Run

	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Creating a pool succeeds",
			args: args{
				params: Params{
					Size:    2,
					Run:     success,
					Timeout: DefaultTimeout,
				},
			},
		},
		{
			name: "Creating a pool fails due to empty timeouts",
			args: args{
				params: Params{
					Size: 2,
					Run:  success,
				},
			},
			err: &multierror.Error{
				Errors: []error{
					errPoolStopTimeoutCannotBeZero,
					errPoolAddTimeoutCannotBeZero,
				},
			},
		},
		{
			name: "Creating a pool fails due to 0 size",
			args: args{
				params: Params{
					Size:    0,
					Run:     success,
					Timeout: DefaultTimeout,
				},
			},
			err: &multierror.Error{
				Errors: []error{
					errPoolSizeCannotBeZero,
				},
			},
		},
		{
			name: "Creating a pool fails due to empty Run",
			args: args{
				params: Params{
					Size:    2,
					Timeout: DefaultTimeout,
				},
			},
			err: &multierror.Error{
				Errors: []error{
					errPoolRunFuncCannotBeNil,
				},
			},
		},
		{
			name: "Creating a pool fails due to empty Run",
			args: args{
				params: Params{},
			},
			err: &multierror.Error{
				Errors: []error{
					errPoolSizeCannotBeZero,
					errPoolStopTimeoutCannotBeZero,
					errPoolAddTimeoutCannotBeZero,
					errPoolRunFuncCannotBeNil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := NewPool(tt.args.params); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewPool() error = %v, wantErr %v", err, tt.err)
				return
			}
		})
	}
}

func TestPoolAdd(t *testing.T) {
	type fields struct {
		size      uint16
		run       RunFunc
		queue     chan Validator
		signals   Signals
		state     *State
		leftovers chan Validator
		errors    chan error
		timeouts  Timeout
	}
	type args struct {
		work             []Validator
		limitConsumption uint
		consumeQueue     bool
		consumeSignal    bool
	}
	type want struct {
		leftovers      []Validator
		queueContents  []Validator
		signalContents []struct{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
		err    error
	}{
		{
			// To avoid test flakiness, the behaviour of this test is as follows:
			// The queue has 3 buffered spots and 2 work items are sent. We're simulating
			// that 2 items are consumed immediately via `limitConsumption` to 2.
			name: "Add succeeds",
			fields: fields{
				queue: make(chan Validator, 3),
				signals: Signals{
					Added: make(chan struct{}),
				},
				state: &State{
					Status: Counter{value: StartedStatus},
				},
				timeouts: testTimeout,
			},
			args: args{
				work:             generateWork(2),
				consumeQueue:     true,
				consumeSignal:    true,
				limitConsumption: 2,
			},
			want: want{
				queueContents: generateWork(2),
				signalContents: []struct{}{
					{},
					{},
				},
			},
		},
		{
			// The behaviour of this test is as follows, 6 work items are sent,
			// The queue has 1 buffered spot and we're simulating that 2 items
			// are consumed immediately via `limitConsumption` to 2.
			// that leaves 3 items sent to the queue, 1 in the buffer, 2 consumed
			// by the simulated workeds and 3 as leftovers.
			name: "Add fails when the queue is full",
			fields: fields{
				queue: make(chan Validator, 1),
				signals: Signals{
					Added: make(chan struct{}),
				},
				state: &State{
					Status: Counter{value: StartedStatus},
				},
				timeouts: testTimeout,
			},
			args: args{
				consumeSignal:    true,
				consumeQueue:     true,
				limitConsumption: 2,
				work:             generateWork(6),
			},
			want: want{
				leftovers:     generateWork(3),
				queueContents: generateWork(2),
				signalContents: []struct{}{
					{},
					{},
					{},
				},
			},
			err: ErrAddOperationTimedOut,
		},
		{
			name: "Add fails when the pool is being stopped",
			fields: fields{
				state: &State{
					Status: Counter{value: StoppingStatus},
				},
			},
			args: args{
				work: generateWork(3),
			},
			want: want{
				leftovers: generateWork(3),
			},
			err: ErrCannotAddWorkToStoppingPool,
		},
		{
			name: "Add fails when the pool is stopped by timeout",
			fields: fields{
				state: &State{
					Status: Counter{value: StoppedTimeout},
				},
			},
			args: args{
				work: generateWork(4),
			},
			want: want{
				leftovers: generateWork(4),
			},
			err: ErrCannotAddWorkToStoppingPool,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var fakeConsumer = &fakeWorkConsumer{
				limit: tt.args.limitConsumption,
				q:     tt.fields.queue,
				done:  make(chan struct{}),
			}
			var fakeSignalConsumer = &fakeSignalConsumer{
				q:    tt.fields.signals.Added,
				done: make(chan struct{}),
			}

			if tt.args.consumeQueue {
				go fakeConsumer.Consume()
				defer fakeConsumer.Stop()
			}

			if tt.args.consumeSignal {
				go fakeSignalConsumer.Consume()
				defer fakeSignalConsumer.Stop()
			}

			p := Pool{
				size:      tt.fields.size,
				run:       tt.fields.run,
				queue:     tt.fields.queue,
				signals:   tt.fields.signals,
				state:     tt.fields.state,
				leftovers: tt.fields.leftovers,
				errors:    tt.fields.errors,
				timeouts:  tt.fields.timeouts,
			}

			got, err := p.Add(tt.args.work...)
			if !reflect.DeepEqual(got, tt.want.leftovers) {
				t.Errorf("Pool.Add() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Pool.Add() error = %v, want %v", err, tt.err)
			}

			// Leave time for the goroutines to consume the channels
			<-time.After(time.Millisecond)

			var queueContents = fakeConsumer.Contents()
			var signalContents = fakeSignalConsumer.Contents()

			if !reflect.DeepEqual(queueContents, tt.want.queueContents) {
				t.Errorf("Pool.queue = %v, want %v", queueContents, tt.want.queueContents)
			}

			if !reflect.DeepEqual(signalContents, tt.want.signalContents) {
				t.Errorf("Pool.signals.added = %v, want %v", signalContents, tt.want.signalContents)
			}
		})
	}
}

func TestPoolStart(t *testing.T) {
	type fields struct {
		size      uint16
		run       RunFunc
		queue     chan Validator
		signals   Signals
		state     *State
		leftovers chan Validator
		errors    chan error
		timeouts  Timeout
	}
	type want struct {
		status uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   want
		err    error
	}{
		{
			name: "Start succeeds and leaves state as StartedStatus",
			fields: fields{
				size: 1,
				signals: Signals{
					Stop:        make(chan struct{}),
					Stopped:     make(chan bool),
					StopMonitor: make(chan struct{}),
				},
				state: new(State),
			},
			want: want{
				status: StartingStatus,
			},
			err: nil,
		},
		{
			name: "Start fails when the status is started",
			fields: fields{
				state: &State{
					Status: Counter{value: StartedStatus},
				},
			},
			want: want{
				status: StartedStatus,
			},
			err: ErrAlreadyStarted,
		},
		{
			name: "Start fails when the status is IdleStatus",
			fields: fields{
				state: &State{
					Status: Counter{value: IdleStatus},
				},
			},
			want: want{
				status: IdleStatus,
			},
			err: ErrAlreadyStarted,
		},
		{
			name: "Start fails when the status is FinishedStatus",
			fields: fields{
				state: &State{
					Status: Counter{value: FinishedStatus},
				},
			},
			want: want{
				status: FinishedStatus,
			},
			err: ErrAlreadyStarted,
		},
		{
			name: "Start fails when the status is StoppingStatus",
			fields: fields{
				state: &State{
					Status: Counter{value: StoppingStatus},
				},
			},
			want: want{
				status: StoppingStatus,
			},
			err: ErrAlreadyStarted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.signals.StopMonitor != nil {
				defer func() { tt.fields.signals.StopMonitor <- struct{}{} }()
			}
			p := &Pool{
				size:      tt.fields.size,
				run:       tt.fields.run,
				queue:     tt.fields.queue,
				signals:   tt.fields.signals,
				state:     tt.fields.state,
				leftovers: tt.fields.leftovers,
				errors:    tt.fields.errors,
				timeouts:  tt.fields.timeouts,
			}
			if err := p.Start(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Pool.Start() error = %v, want %v", err, tt.err)
			}

			if p.Status() != tt.want.status {
				t.Errorf("Pool.status = %v, want %v", StatusText(p.Status()), StatusText(tt.want.status))
			}
		})
	}
}

func generateWork(n uint) []Validator {
	var work = make([]Validator, n)
	for index := 0; index < int(n); index++ {
		work[index] = new(nilValidator)
	}
	return work
}

func TestPoolStopWithContext(t *testing.T) {
	var delayedMockRun = &mockRun{duration: time.Second * 30}

	type fields struct {
		size      uint16
		run       RunFunc
		queue     chan Validator
		signals   Signals
		state     *State
		leftovers chan Validator
		errors    chan error
		timeouts  Timeout
	}
	type args struct {
		work []Validator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
		status uint32
	}{
		{
			name: "Stop succeeds stopping on empty queue without hitting the timeout",
			fields: fields{
				size:      2,
				queue:     make(chan Validator, 10),
				leftovers: make(chan Validator, 12),
				timeouts: Timeout{
					Stop: time.Millisecond * 50,
				},
				signals: Signals{
					Stop:        make(chan struct{}),
					Stopped:     make(chan bool),
					Added:       make(chan struct{}),
					Finish:      make(chan struct{}),
					StopMonitor: make(chan struct{}),
				},
				state: new(State),
				run:   new(mockRun).Run,
			},
			err:    nil,
			status: StoppedSuccess,
		},
		{
			name: "Stop succeeds stopping on full queue",
			fields: fields{
				size:      100,
				queue:     make(chan Validator, 1000),
				leftovers: make(chan Validator, 1100),
				timeouts: Timeout{
					Stop: time.Millisecond * 500,
					Add:  time.Millisecond * 50,
				},
				signals: Signals{
					Stop:        make(chan struct{}),
					Stopped:     make(chan bool),
					Added:       make(chan struct{}),
					Finish:      make(chan struct{}),
					StopMonitor: make(chan struct{}),
				},
				state: new(State),
				run:   new(mockRun).Run,
			},
			args: args{
				work: generateWork(10000),
			},
			err:    nil,
			status: StoppedSuccess,
		},
		{
			name: "Stop returns ErrStopOperationTimedOut when the workers stop with timeout",
			fields: fields{
				size:      10,
				queue:     make(chan Validator, 100),
				leftovers: make(chan Validator, 110),
				timeouts: Timeout{
					Stop: time.Millisecond * 500,
					Add:  time.Millisecond * 50,
				},
				signals: Signals{
					Stop:        make(chan struct{}),
					Stopped:     make(chan bool),
					Added:       make(chan struct{}),
					Finish:      make(chan struct{}),
					StopMonitor: make(chan struct{}),
				},
				state: new(State),
				run:   delayedMockRun.Run,
			},
			args: args{
				work: generateWork(110),
			},
			err:    ErrStopOperationTimedOut,
			status: StoppedTimeout,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup consumers of the queues where the worker sends
			var finishedConsumer = &fakeSignalConsumer{
				q:    tt.fields.signals.Finish,
				done: make(chan struct{}),
			}
			var addedConsumer = &fakeSignalConsumer{
				q:    tt.fields.signals.Added,
				done: make(chan struct{}),
			}

			go finishedConsumer.Consume()
			go addedConsumer.Consume()

			p := &Pool{
				size:      tt.fields.size,
				run:       tt.fields.run,
				queue:     tt.fields.queue,
				signals:   tt.fields.signals,
				state:     tt.fields.state,
				leftovers: tt.fields.leftovers,
				errors:    tt.fields.errors,
				timeouts:  tt.fields.timeouts,
			}

			if err := p.Start(); err != nil {
				t.Fatalf("Pool.Start() error = %v, can't test Stop() with a stopped pool", err)
			}

			p.Add(tt.args.work...)

			defer finishedConsumer.Stop()
			defer addedConsumer.Stop()

			<-time.After(time.Millisecond * 10)
			if err := p.Stop(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Pool.Stop() error = %v, wantErr %v", err, tt.err)
			}

			if p.Status() != tt.status {
				t.Errorf("Pool.Stop().status = %v want %v", StatusText(p.Status()), StatusText(tt.status))
			}

			<-time.After(time.Millisecond * 2)
			leftover, err := p.Leftovers()
			if err != nil {
				t.Fatal("UNEXPECTED ERR", err)
			}

			t.Log("==== SUMMARY ====")
			t.Log("Consumed:", len(finishedConsumer.Contents()))
			t.Log("Leftover:", len(leftover))
		})
	}
}

func TestPoolStopStatus(t *testing.T) {
	type fields struct {
		state *State
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name: "Stop fails on a pool that has never run",
			fields: fields{
				state: new(State),
			},
			err: ErrAlreadyStopped,
		},
		{
			name: "Stop fails on a pool that has been already stopped",
			fields: fields{
				state: &State{
					Status: Counter{value: StoppedSuccess},
				},
			},
			err: ErrAlreadyStopped,
		},
		{
			name: "Stop fails on a pool that has been already stopped (Timeout)",
			fields: fields{
				state: &State{
					Status: Counter{value: StoppedTimeout},
				},
			},
			err: ErrAlreadyStopped,
		},
		{
			name: "Stop fails on a pool that has been already stopped (Timeout)",
			fields: fields{
				state: &State{
					Status: Counter{value: StoppedTimeout},
				},
			},
			err: ErrAlreadyStopped,
		},
		{
			name: "Stop fails on a pool that is already being stopped",
			fields: fields{
				state: &State{
					Status: Counter{value: StoppingStatus},
				},
			},
			err: ErrAlreadyStopping,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				state: tt.fields.state,
			}
			go p.monitor(nil)
			if err := p.Stop(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Pool.Stop() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func TestStatusText(t *testing.T) {
	type args struct {
		status uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "stopped status",
			args: args{status: 0},
			want: "stopped",
		},
		{
			name: "unknown status on missing map entry",
			args: args{status: 99},
			want: "unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusText(tt.args.status); got != tt.want {
				t.Errorf("StatusText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoolWait(t *testing.T) {
	var errs = []error{
		errors.New("AN ERROR"),
		errors.New("ANOTHER ERROR"),
		errors.New("YET ANOTHER ERROR"),
	}
	type fields struct {
		size      uint16
		run       RunFunc
		queue     chan Validator
		leftovers chan Validator
		errors    chan error
		signals   Signals
		state     *State
		timeouts  Timeout
		writer    *bytes.Buffer
	}
	type args struct {
		interrupt chan os.Signal
		wait      time.Duration
		errs      []error
		signals   []os.Signal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		err    error
		output string
	}{
		{
			name: "Returns err on stopped pool",
			fields: fields{
				errors: make(chan error),
				state:  new(State),
			},
			err: ErrCannotWaitOnStoppedPool,
		},
		{
			name: "Returns nil when the status is FinishedStatus",
			fields: fields{
				errors: make(chan error),
				state: &State{
					Status: Counter{value: FinishedStatus},
				},
			},
			err: nil,
		},
		{
			name: "Returns nil when the status is StoppedTimeout",
			fields: fields{
				errors: make(chan error),
				state: &State{
					Status: Counter{value: StoppedTimeout},
				},
			},
			err: nil,
		},
		{
			name: "Returns the errors received by the err channel",
			fields: fields{
				errors: make(chan error),
				state: &State{
					Status: Counter{value: StoppedSuccess},
					Errors: new(Errors),
				},
			},
			args: args{
				wait: time.Millisecond * 50,
				errs: errs,
			},
			err: &multierror.Error{
				Errors: errs,
			},
		},
		{
			name: `Stop the pool and retrieve the errors caught if "os.Interrupt" is sent`,
			fields: fields{
				errors: make(chan error),
				queue:  make(chan Validator),
				state: &State{
					Status: Counter{value: IdleStatus},
					Errors: new(Errors),
				},
				writer: new(bytes.Buffer),
			},
			args: args{
				wait:      time.Millisecond * 50,
				errs:      errs,
				interrupt: make(chan os.Signal),
				signals: []os.Signal{
					os.Interrupt,
				},
			},
			err: &multierror.Error{
				Errors: errs,
			},
			output: "pool: received interrupt, stopping pool...\n",
		},
		{
			name: `Stop the pool and retrieve the errors caught if "os.Kill" is sent`,
			fields: fields{
				errors: make(chan error),
				queue:  make(chan Validator),
				state: &State{
					Status: Counter{value: IdleStatus},
					Errors: new(Errors),
				},
			},
			args: args{
				wait:      time.Millisecond * 50,
				errs:      errs,
				interrupt: make(chan os.Signal),
				signals: []os.Signal{
					os.Kill,
				},
			},
			err: &multierror.Error{
				Errors: errs,
			},
		},
		{
			name: `Stop the pool and retrieve the errors caught if an "syscall.SIGTERM" is sent`,
			fields: fields{
				errors: make(chan error),
				queue:  make(chan Validator),
				state: &State{
					Status: Counter{value: IdleStatus},
					Errors: new(Errors),
				},
			},
			args: args{
				wait:      time.Millisecond * 50,
				errs:      errs,
				interrupt: make(chan os.Signal),
				signals: []os.Signal{
					syscall.SIGTERM,
				},
			},
			err: &multierror.Error{
				Errors: errs,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				size:      tt.fields.size,
				run:       tt.fields.run,
				queue:     tt.fields.queue,
				leftovers: tt.fields.leftovers,
				errors:    tt.fields.errors,
				signals:   tt.fields.signals,
				state:     tt.fields.state,
				timeouts:  tt.fields.timeouts,
			}

			if tt.fields.writer != nil {
				p.writer = tt.fields.writer
			}

			if tt.args.interrupt != nil {
				go func() {
					<-time.After(time.Millisecond)
					for _, s := range tt.args.signals {
						tt.args.interrupt <- s
					}
				}()
			}

			go p.monitor(tt.args.interrupt)

			for _, err := range tt.args.errs {
				tt.fields.errors <- err
			}

			<-time.After(tt.args.wait)

			if err := p.Wait(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Pool.Wait() error = %v, wantErr %v", err, tt.err)
			}

			if tt.fields.writer != nil {
				if out := tt.fields.writer.String(); out != tt.output {
					t.Errorf("Pool.Wait() output = %v, want %v", out, tt.output)
				}
			}
		})
	}
}

func TestPoolupdateState(t *testing.T) {
	type fields struct {
		size      uint16
		run       RunFunc
		queue     chan Validator
		leftovers chan Validator
		errors    chan error
		signals   Signals
		state     *State
		timeouts  Timeout
	}
	tests := []struct {
		name   string
		fields fields
		want   uint32
	}{
		{
			name: "Pool stopping leaves the status untouched",
			fields: fields{
				state: &State{
					Status: Counter{value: StoppingStatus},
				},
			},
			want: StoppingStatus,
		},
		{
			name: "Pool Started with 0 items queued moves the status to idle",
			fields: fields{
				state: &State{
					Status: Counter{value: StartedStatus},
				},
			},
			want: IdleStatus,
		},
		{
			name: "Pool Started with 1 item queued and processed moves the status to finished",
			fields: fields{
				state: &State{
					Queued:    Counter{value: 1},
					Processed: Counter{value: 1},
					Status:    Counter{value: StartedStatus},
				},
			},
			want: FinishedStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				size:      tt.fields.size,
				run:       tt.fields.run,
				queue:     tt.fields.queue,
				leftovers: tt.fields.leftovers,
				errors:    tt.fields.errors,
				signals:   tt.fields.signals,
				state:     tt.fields.state,
				timeouts:  tt.fields.timeouts,
			}
			p.updateState()

			if p.Status() != tt.want {
				t.Errorf("Pool.updateState().Status() = %v, want %v", p.Status(), tt.want)
			}
		})
	}
}

func TestPoolLeftovers(t *testing.T) {
	type fields struct {
		size      uint16
		run       RunFunc
		queue     chan Validator
		leftovers chan Validator
		errors    chan error
		signals   Signals
		state     *State
		timeouts  Timeout
	}
	type args struct {
		leftovers []Validator
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Validator
		err    error
	}{
		{
			name:   "Obtaining leftovers fails when pool is StartedStatus returns an error",
			fields: fields{state: &State{Status: Counter{value: StartedStatus}}},
			err:    ErrCannotGetLeftovers,
		},
		{
			name:   "Obtaining leftovers fails when pool is StoppedStatus returns an error",
			fields: fields{state: &State{Status: Counter{value: StoppedStatus}}},
			err:    ErrCannotGetLeftovers,
		},
		{
			name:   "Obtaining leftovers fails when pool is StartingStatus returns an error",
			fields: fields{state: &State{Status: Counter{value: StartingStatus}}},
			err:    ErrCannotGetLeftovers,
		},
		{
			name:   "Obtaining leftovers fails when pool is IdleStatus returns an error",
			fields: fields{state: &State{Status: Counter{value: IdleStatus}}},
			err:    ErrCannotGetLeftovers,
		},
		{
			name:   "Obtaining leftovers fails when pool is FinishedStatus returns an error",
			fields: fields{state: &State{Status: Counter{value: FinishedStatus}}},
			err:    ErrCannotGetLeftovers,
		},
		{
			name:   "Obtaining leftovers fails when pool is StoppingStatus returns an error",
			fields: fields{state: &State{Status: Counter{value: StoppingStatus}}},
			err:    ErrCannotGetLeftovers,
		},
		{
			name: "Obtaining leftovers fails when pool is StoppedTimeout returns the leftovers",
			fields: fields{
				state:     &State{Status: Counter{value: StoppedTimeout}},
				leftovers: make(chan Validator, 10),
			},
			args: args{leftovers: generateWork(10)},
			want: generateWork(10),
		},
		{
			name: "Obtaining leftovers fails when pool is StoppedSuccess returns the leftovers",
			fields: fields{
				state:     &State{Status: Counter{value: StoppedSuccess}},
				leftovers: make(chan Validator, 20),
			},
			args: args{leftovers: generateWork(20)},
			want: generateWork(20),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				size:      tt.fields.size,
				run:       tt.fields.run,
				queue:     tt.fields.queue,
				leftovers: tt.fields.leftovers,
				errors:    tt.fields.errors,
				signals:   tt.fields.signals,
				state:     tt.fields.state,
				timeouts:  tt.fields.timeouts,
			}

			for _, l := range tt.args.leftovers {
				tt.fields.leftovers <- l
			}

			got, err := p.Leftovers()
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Pool.Leftovers() error = %v, wantErr %v", err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pool.Leftovers() = %v, want %v", got, tt.want)
			}
		})
	}
}
