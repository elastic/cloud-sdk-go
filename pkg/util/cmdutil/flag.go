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

package cmdutil

import (
	"fmt"
	"net"
	"reflect"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// IncompatibleFlags checks if both flags have been specified, and if so
// returns an error.
func IncompatibleFlags(cmd *cobra.Command, first, second string) error {
	if cmd.Flag(first).Changed && cmd.Flag(second).Changed {
		return fmt.Errorf(
			`incompatible flags "--%s" and "--%s" specified, "--%s" will be ignored"`,
			first, second, second,
		)
	}
	return nil
}

// DecodeFlags decodes the set flags of a cobra.Command and unpacks all the
// values to the specified pointer of the passed "output" structure.
func DecodeFlags(cmd *cobra.Command, output interface{}) error {
	var flagMap = make(map[string]interface{})
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		flagMap[f.Name] = parseValue(f.Value)
	})

	return mapstructure.WeakDecode(flagMap, output)
}

func parseValue(val pflag.Value) interface{} {
	// All types which encapsulate the pflag slice type have a GetSlice
	// method, which obtains the slice as []string. Combined with the
	// mapstructure call to WeakDecode, it'll decode the slice correctly.
	flagType := val.Type()
	isSlice := strings.HasSuffix(flagType, "Slice")
	isArray := strings.HasSuffix(flagType, "Array")
	if isSlice || isArray {
		// Reflects the ValueOf the pflag.Value and obtains "GetSlice" functio
		// as reflect.Value type, which is then called with no arguments.
		getSlice := reflect.ValueOf(val).MethodByName("GetSlice")

		// Since Call returns a slice of reflect.Value, and the signature of
		// "GetSlice" has a single return, the 0 index is accessed to obtain
		// the actual result of the function call and is transformed into the
		// underlying type through the Interface() call.
		result := getSlice.Call(nil)[0].Interface()

		// Since WeakDecode is used, the types are loosely matched with the
		// flag type. Duration would be returned as a []int, and ipSlice as
		// a slice of []byte since those are the underlying types.
		if flagType == "durationSlice" {
			return parseDurationSlice(result)
		}
		if flagType == "ipSlice" {
			return parseIPSlice(result)
		}

		return result
	}

	return val
}

func parseDurationSlice(i interface{}) []time.Duration {
	var durationSlice = make([]time.Duration, 0)
	for _, d := range i.([]string) {
		dVal, _ := time.ParseDuration(d)
		durationSlice = append(durationSlice, dVal)
	}
	return durationSlice
}

func parseIPSlice(i interface{}) []net.IP {
	var ipSlice = make([]net.IP, 0)
	for _, ip := range i.([]string) {
		ipSlice = append(ipSlice, net.ParseIP(ip))
	}
	return ipSlice
}
