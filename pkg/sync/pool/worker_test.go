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
	"reflect"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestWork(t *testing.T) {
	var workerError = errors.New("ERROR")

	type config struct {
		queue       chan Validator
		stop        chan struct{}
		finished    chan struct{}
		stopped     chan bool
		errChan     chan error
		leftover    chan Validator
		stopTimeout time.Duration
		worker      *mockRun
	}
	type args struct {
		work   []Validator
		config config
	}
	type want struct {
		processed []Validator
		finished  []struct{}
		errored   []error
		leftovers []Validator
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Worker consumes the sent",
			args: args{
				config: config{
					queue:       make(chan Validator),
					stop:        make(chan struct{}),
					stopped:     make(chan bool, 5),
					finished:    make(chan struct{}),
					errChan:     make(chan error),
					leftover:    make(chan Validator),
					stopTimeout: time.Millisecond * 5,
					worker:      new(mockRun),
				},
				work: generateWork(5),
			},
			want: want{
				processed: generateWork(5),
				leftovers: nil,
				errored:   nil,
				finished: []struct{}{
					{},
					{},
					{},
					{},
					{},
				},
			},
		},
		{
			name: "Worker never finishes consuming work, gets stopped by timeout",
			args: args{
				config: config{
					queue:       make(chan Validator),
					stop:        make(chan struct{}),
					stopped:     make(chan bool, 1),
					finished:    make(chan struct{}),
					errChan:     make(chan error),
					leftover:    make(chan Validator),
					stopTimeout: time.Millisecond * 5,
					worker:      &mockRun{duration: time.Second * 10},
				},
				work: generateWork(1),
			},
			want: want{
				processed: nil,
				leftovers: []Validator{
					new(nilValidator),
				},
				errored:  nil,
				finished: nil,
			},
		},
		{
			name: "Worker is stopped while processing work but finishes work before timeout returns",
			args: args{
				config: config{
					queue:       make(chan Validator),
					stop:        make(chan struct{}),
					stopped:     make(chan bool, 10),
					finished:    make(chan struct{}),
					errChan:     make(chan error),
					leftover:    make(chan Validator),
					stopTimeout: time.Millisecond * 20,
					worker:      &mockRun{duration: time.Millisecond * 15},
				},
				work: generateWork(1),
			},
			want: want{
				processed: generateWork(1),
				leftovers: nil,
				errored:   nil,
				finished: []struct{}{
					{},
				},
			},
		},
		{
			name: "Worker consumes work which returns an error(s)",
			args: args{
				config: config{
					queue:       make(chan Validator),
					stop:        make(chan struct{}),
					stopped:     make(chan bool, 10),
					finished:    make(chan struct{}),
					errChan:     make(chan error),
					leftover:    make(chan Validator),
					stopTimeout: time.Millisecond * 5,
					worker:      &mockRun{err: workerError},
				},
				work: generateWork(3),
			},
			want: want{
				processed: generateWork(3),
				leftovers: nil,
				errored: []error{
					workerError,
					workerError,
					workerError,
				},
				finished: []struct{}{
					{},
					{},
					{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup consumers of the queues where the worker sends
			var leftoverConsumer = &fakeWorkConsumer{
				q:    tt.args.config.leftover,
				done: make(chan struct{}),
			}
			var finishedConsumer = &fakeSignalConsumer{
				q:    tt.args.config.finished,
				done: make(chan struct{}),
			}
			var errorConsumer = &fakeErrorConsumer{
				q:    tt.args.config.errChan,
				done: make(chan struct{}),
			}

			// Start consumers
			go leftoverConsumer.Consume()
			go finishedConsumer.Consume()
			go errorConsumer.Consume()

			// Send the actual work to the queue
			go func() {
				for _, item := range tt.args.work {
					t.Log("SEND WORK")
					tt.args.config.queue <- item
					t.Log("SENT WORK")
				}
			}()

			// Set up the stop signal so we can assert the results of the test
			go func() {
				<-time.After(time.Millisecond * 10)
				tt.args.config.stop <- struct{}{}
			}()

			// Start the worker
			Work(Worker{
				Queue:       tt.args.config.queue,
				Stop:        tt.args.config.stop,
				Stopped:     tt.args.config.stopped,
				Finished:    tt.args.config.finished,
				Errors:      tt.args.config.errChan,
				Leftovers:   tt.args.config.leftover,
				Run:         tt.args.config.worker.Run,
				StopTimeout: tt.args.config.stopTimeout,
			})

			// Give time for the workers to consume all of the signals
			<-time.After(time.Millisecond * 10)

			// Stop consumers
			defer leftoverConsumer.Stop()
			defer finishedConsumer.Stop()
			defer errorConsumer.Stop()

			// Assert result after the worker has exited
			processedWork := tt.args.config.worker.Contents()
			if !reflect.DeepEqual(processedWork, tt.want.processed) {
				t.Errorf("worker.processed = %v, want %v", processedWork, tt.want.processed)
			}

			receivedErrors := errorConsumer.Contents()
			if !reflect.DeepEqual(receivedErrors, tt.want.errored) {
				t.Errorf("Pool.errors = %v, want %v", receivedErrors, tt.want.errored)
			}

			finishedSignals := finishedConsumer.Contents()
			if !reflect.DeepEqual(finishedSignals, tt.want.finished) {
				t.Errorf("Pool.finished = %v, want %v", finishedSignals, tt.want.finished)
			}

			leftovers := leftoverConsumer.Contents()
			if !reflect.DeepEqual(leftovers, tt.want.leftovers) {
				t.Errorf("Pool.leftovers = %v, want %v", leftovers, tt.want.leftovers)
			}
		})
	}
}
