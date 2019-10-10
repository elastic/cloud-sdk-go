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

package datetime

import (
	"testing"
	"time"
)

func TestISO8601(t *testing.T) {
	// -05:00 hours from UTC
	zone := time.FixedZone("myzone", -5*60*60)

	tests := []struct {
		name    string
		args    string
		want    time.Time
		wantErr bool
	}{
		{
			name: "It should properly parse a date",
			args: "2018-01-01",
			want: time.Date(2018, 01, 01, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "It should properly parse a date+time-hours",
			args: "2018-01-02T08Z",
			want: time.
				Date(2018, 01, 02, 8, 0, 0, 0, time.UTC),
		},
		{
			name: "It should properly parse a date+time-min",
			args: "2018-01-02T08:30Z",
			want: time.
				Date(2018, 01, 02, 8, 30, 0, 0, time.UTC),
		},
		{
			name: "It should properly parse a date+time-sec",
			args: "2018-01-02T08:30:22Z",
			want: time.
				Date(2018, 01, 02, 8, 30, 22, 0, time.UTC),
		},
		{
			name: "It should properly parse a date+time-nanosec",
			args: "2018-01-02T08:30:22.944526000Z",
			want: time.Date(2018, 01, 02, 8, 30, 22, 944526000, time.UTC),
		},
		{
			name: "It should properly parse a date+time w/ timezone",
			args: "2018-01-03T08:00:00-05:00",
			want: time.Date(2018, 01, 03, 8, 0, 0, 0, zone),
		},
		{
			name:    "It should error on a bad date",
			args:    "certainly not a date",
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ISO8601(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ISO8601() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("ISO8601() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidDuration(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"MinusInf", true},
		{"PlusInf", true},
		{"-Inf", true},
		{"+Inf", true},
		{"Inf", true},
		{"  7 day  ", true},
		{"7 days", true},
		{"7 µs", true},
		{"12h", true},
		{"1.5 hours", true},
		{"1.5897987 ms", true},
		{"1.5. ms", false},
		{"1. ms", false},
		{"ms", false},
		{"randomstuff", false},
		{"1", false},
		{"12h garbageonend", false},
		{"garbageonbegin 12h garbageonend", false},
		{"garbageonbegin 12h", false},
		{"  ", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run("TestValidDuration", func(t *testing.T) {
			if got := ValidDuration(tt.args); got != tt.want {
				t.Errorf("ValidDuration() = args = \"%v\" got = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestValidSlowQueryThreshold(t *testing.T) {
	tests := []struct {
		args string
		want bool
	}{
		{"2ms", true},
		{"3μs", true},
		{"4ns", true},
		{"1s", true},
		{"2 ms", false},
		{"3 μs", false},
		{"4 ns", false},
		{"1 s", false},
		{"a100 s", false},
		{"2m", false},
		{"2μ", false},
		{"2n", false},
		{"2", false},
		{" 7 day  ", false},
		{"7 days", false},
		{"1 h  ", false},
		{"3hours", false},
		{"1.5. ms", false},
		{"1. ms", false},
		{"ms", false},
		{"", false},
		{"randomstuff", false},
		{"12h garbageonend", false},
		{"garbageonbegin 12h garbageonend", false},
		{"garbageonbegin 12h", false},
	}
	for _, tt := range tests {
		t.Run("TestdValidSlowQueryThreshold", func(t *testing.T) {
			if got := ValidSlowQueryThreshold(tt.args); got != tt.want {
				t.Errorf("TestdValidSlowQueryThreshold() = args = \"%v\" got = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
