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
	"regexp"
	"time"

	me "github.com/hashicorp/go-multierror"
)

const (
	// ISO8601Date defines an ISO 8601 Date format
	ISO8601Date = "2006-01-02"
	// ISO8601Hour defines an ISO 8601 Date format including the hour
	ISO8601Hour = "2006-01-02T15Z07:00"
	// ISO8601Minute defines an ISO 8601 Date format including the hour and the minute
	ISO8601Minute = "2006-01-02T15:04Z07:00"
	// ISO8601Second defines an ISO 8601 Date format including the hour, the minute and the second
	ISO8601Second = time.RFC3339
	// ISO8601Nano defines an ISO 8601 Date format including the hour, the minute, the second and the nanosecond
	ISO8601Nano = time.RFC3339Nano

	validDurationRegexStr =
	// Infinity
	`^\s*(([-+]|Plus|Minus)?Inf` +
		// number followed by optional whitespace
		`|\d+(.\d+)?\s*?` +
		// unit
		`(d|days?|h|hours?|min|minutes?` +
		`|s|sec|seconds?|ms|milli|milliseconds?` +
		`|µs|micro|microseconds?|ns|nano|nanoseconds?))\s*$`
)

var (
	// ISO8601Formats contains a slice of all declared ISO 8601 date formats
	ISO8601Formats = []string{
		ISO8601Date,
		ISO8601Hour,
		ISO8601Minute,
		ISO8601Second,
		ISO8601Nano,
	}

	validDurationRegex = regexp.MustCompile(validDurationRegexStr)
)

// ISO8601 is a utility function that tries all of the date and date+time formats
// of ISO8601
func ISO8601(s string) (time.Time, error) {
	errs := new(me.Error)

	for _, v := range ISO8601Formats {
		out, err := time.Parse(v, s)
		if err == nil {
			return out, nil
		}
		errs = me.Append(errs, err)
	}

	return time.Time{}, errs
}

// ValidDuration Validates the given string according to the Scala spec.
// https://github.com/scala/scala/blob/e3df10a0c06fbbdbaa77d654c9cc64495c4cac29/src/library/scala/concurrent/duration/Duration.scala#L52
func ValidDuration(s string) bool {
	return validDurationRegex.Match([]byte(s))
}

// ValidSlowQueryThreshold Validates the given string according to what ES threshold values are expected spec.
func ValidSlowQueryThreshold(s string) bool {
	// only errors if regex doesn't compile
	matched, _ := regexp.MatchString(
		`\d+(ms|ns|μs|s)`,
		s,
	)

	return matched
}
