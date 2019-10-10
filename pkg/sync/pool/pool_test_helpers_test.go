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

type mockRun struct {
	processed []Validator
	err       error
	duration  time.Duration
	mu        sync.RWMutex
}

func (m *mockRun) Run(params Validator) error {
	<-time.After(m.duration)
	m.mu.Lock()
	defer m.mu.Unlock()
	m.processed = append(m.processed, params)
	return m.err
}

func (m *mockRun) Contents() []Validator {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.processed
}

// fakeWorkConsumer is used as a fake consumer
type fakeWorkConsumer struct {
	list  []Validator
	limit uint
	mu    sync.RWMutex
	q     chan Validator
	done  chan struct{}
}

func (l *fakeWorkConsumer) Consume() {
	for {
		select {
		case item := <-l.q:
			if !l.Add(item) {
				l.q = nil
			}
		case <-l.done:
			return
		}
	}
}

func (l *fakeWorkConsumer) Stop() {
	l.done <- struct{}{}
}

func (l *fakeWorkConsumer) Add(i ...Validator) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list = append(l.list, i...)

	return uint(len(l.list)) != l.limit
}

func (l *fakeWorkConsumer) Contents() []Validator {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list
}

func (l *fakeWorkConsumer) Len() uint {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return uint(len(l.list))
}

type fakeSignalConsumer struct {
	list []struct{}
	mu   sync.RWMutex
	q    chan struct{}
	done chan struct{}
}

func (l *fakeSignalConsumer) Consume() {
	for {
		select {
		case item := <-l.q:
			l.Add(item)
		case <-l.done:
			return
		}
	}
}

func (l *fakeSignalConsumer) Stop() {
	l.done <- struct{}{}
}

func (l *fakeSignalConsumer) Add(i ...struct{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list = append(l.list, i...)
}

func (l *fakeSignalConsumer) Contents() []struct{} {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list
}

type fakeErrorConsumer struct {
	list []error
	mu   sync.RWMutex
	q    chan error
	done chan struct{}
}

func (l *fakeErrorConsumer) Consume() {
	for {
		select {
		case item := <-l.q:
			l.Add(item)
		case <-l.done:
			return
		}
	}
}

func (l *fakeErrorConsumer) Stop() {
	l.done <- struct{}{}
}

func (l *fakeErrorConsumer) Add(i ...error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.list = append(l.list, i...)
}

func (l *fakeErrorConsumer) Contents() []error {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.list
}
