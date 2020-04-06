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
	"errors"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/spf13/cobra"
)

func TestIncompatibleFlags(t *testing.T) {
	cmdWithSliceFlag := &cobra.Command{
		Use: "something",
		Run: func(cmd *cobra.Command, args []string) {},
	}
	cmdWithSliceFlag.Flags().StringSlice("instance", []string{"1", "2", "3"}, "instance")
	cmdWithSliceFlag.Flags().Bool("all", false, "all")
	cmdWithSliceFlag.ParseFlags([]string{})

	cmdWithSliceFlagChanged := &cobra.Command{
		Use: "something",
		Run: func(cmd *cobra.Command, args []string) {},
	}
	cmdWithSliceFlagChanged.Flags().StringSlice("instance", []string{"1", "2", "3"}, "instance")
	cmdWithSliceFlagChanged.Flags().Bool("all", false, "all")
	cmdWithSliceFlagChanged.ParseFlags([]string{"--all", "--instance=1"})

	type args struct {
		cmd    *cobra.Command
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "returns no error when no flag is specified",
			args: args{
				cmd:    cmdWithSliceFlag,
				first:  "instance",
				second: "all",
			},
		},
		{
			name: "returns an error when both flags are specified",
			args: args{
				cmd:    cmdWithSliceFlagChanged,
				first:  "all",
				second: "instance",
			},
			err: errors.New(`incompatible flags "--all" and "--instance" specified, "--instance" will be ignored"`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IncompatibleFlags(tt.args.cmd, tt.args.first, tt.args.second); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("IncompatibleFlags() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func TestDecodeFlags(t *testing.T) {
	type nested struct {
		Prop string `mapstructure:"some-prop"`
	}
	// nolint
	type cfg struct {
		// Standard types
		SomeString   string        `mapstructure:"some-string"`
		SomeBool     bool          `mapstructure:"some-bool"`
		SomeInt      int           `mapstructure:"some-int"`
		SomeInt32    int32         `mapstructure:"some-int32"`
		SomeInt64    int64         `mapstructure:"some-int64"`
		SomeUint     uint          `mapstructure:"some-uint"`
		SomeUint32   uint32        `mapstructure:"some-uint32"`
		SomeUint64   uint64        `mapstructure:"some-uint64"`
		SomeFloat32  float32       `mapstructure:"some-float32"`
		SomeFloat64  float64       `mapstructure:"some-float64"`
		SomeDuration time.Duration `mapstructure:"some-duration"`
		SomeIP       net.IP        `mapstructure:"some-ip"`
		SomeIPMask   net.IPMask    `mapstructure:"some-ip-mask"`
		SomeIPNet    net.IPNet     `mapstructure:"some-ip-net"`

		nested `mapstructure:",squash"`

		// Slices
		SomeStringSlice   []string        `mapstructure:"some-string-slice"`
		SomeStringArray   []string        `mapstructure:"some-string-array"`
		SomeBoolSlice     []bool          `mapstructure:"some-bool-slice"`
		SomeIntSlice      []int           `mapstructure:"some-int-slice"`
		SomeUintSlice     []uint          `mapstructure:"some-uint-slice"`
		SomeInt32Slice    []int32         `mapstructure:"some-int32-slice"`
		SomeInt64Slice    []int64         `mapstructure:"some-int64-slice"`
		SomeFloat32Slice  []float32       `mapstructure:"some-float32-slice"`
		SomeFloat64Slice  []float64       `mapstructure:"some-float64-slice"`
		SomeDurationSlice []time.Duration `mapstructure:"some-duration-slice"`
		SomeIPSlice       []net.IP        `mapstructure:"some-ip-slice"`
	}

	cobraCommand := &cobra.Command{}

	// Generic types
	cobraCommand.Flags().String("some-string", "defaultVal", "some desc")
	cobraCommand.Flag("some-string").Value.Set("my value")

	cobraCommand.Flags().String("some-prop", "nestedPropVal", "some desc")

	cobraCommand.Flags().Bool("some-bool", false, "some desc")
	cobraCommand.Flag("some-bool").Value.Set("true")

	cobraCommand.Flags().Int("some-int", 100, "some desc")
	cobraCommand.Flags().Int32("some-int32", 2000, "some desc")
	cobraCommand.Flags().Int64("some-int64", 20000, "some desc")

	cobraCommand.Flags().Uint("some-uint", 100, "some desc")
	cobraCommand.Flags().Uint32("some-uint32", 2000, "some desc")
	cobraCommand.Flags().Uint64("some-uint64", 20000, "some desc")

	cobraCommand.Flags().Float32("some-float32", 100.1, "some desc")
	cobraCommand.Flags().Float64("some-float64", 100.2, "some desc")

	cobraCommand.Flags().Duration("some-duration", time.Second, "some desc")

	ip, ipNet, err := net.ParseCIDR("192.168.1.1/24")
	if err != nil {
		t.Fatal(err)
	}
	cobraCommand.Flags().IP("some-ip", ip, "some desc")
	cobraCommand.Flags().IPMask("some-ip-mask", ipNet.Mask, "some desc")
	cobraCommand.Flags().IPNet("some-ip-net", *ipNet, "some desc")

	// Slice types

	cobraCommand.Flags().StringSlice("some-string-slice", []string{"someval"}, "some desc")
	cobraCommand.Flags().StringArray("some-string-array", []string{"someval-array"}, "some desc")

	cobraCommand.Flags().BoolSlice("some-bool-slice", []bool{false, true}, "some desc")

	cobraCommand.Flags().IntSlice("some-int-slice", []int{100, 200}, "some desc")
	cobraCommand.Flags().Int32Slice("some-int32-slice", []int32{100, 200}, "some desc")
	cobraCommand.Flags().Int64Slice("some-int64-slice", []int64{100, 200}, "some desc")

	cobraCommand.Flags().UintSlice("some-uint-slice", []uint{100, 200}, "some desc")

	cobraCommand.Flags().Float32Slice("some-float32-slice", []float32{100.1, 100.2}, "some desc")
	cobraCommand.Flags().Float64Slice("some-float64-slice", []float64{100.1, 100.2, 100.3}, "some desc")

	cobraCommand.Flags().DurationSlice("some-duration-slice", []time.Duration{time.Second, time.Hour}, "some desc")

	ipSlice := []net.IP{net.ParseIP("192.168.1.1"), net.ParseIP("192.168.1.2")}
	cobraCommand.Flags().IPSlice("some-ip-slice", ipSlice, "some desc")

	type args struct {
		cmd *cobra.Command
		dst interface{}
	}
	tests := []struct {
		name       string
		args       args
		err        error
		wantStruct interface{}
	}{
		{
			name: "Parses all available types",
			args: args{cmd: cobraCommand, dst: &cfg{}},
			wantStruct: &cfg{
				nested: nested{
					Prop: "nestedPropVal",
				},
				SomeString:        "my value",
				SomeBool:          true,
				SomeInt:           100,
				SomeInt32:         2000,
				SomeInt64:         20000,
				SomeUint:          100,
				SomeUint32:        2000,
				SomeUint64:        20000,
				SomeFloat32:       100.1,
				SomeFloat64:       100.2,
				SomeDuration:      time.Second,
				SomeIP:            ip,
				SomeIPMask:        ipNet.Mask,
				SomeIPNet:         *ipNet,
				SomeIntSlice:      []int{100, 200},
				SomeInt32Slice:    []int32{100, 200},
				SomeInt64Slice:    []int64{100, 200},
				SomeUintSlice:     []uint{100, 200},
				SomeFloat32Slice:  []float32{100.1, 100.2},
				SomeFloat64Slice:  []float64{100.1, 100.2, 100.3},
				SomeDurationSlice: []time.Duration{time.Second, time.Hour},
				SomeBoolSlice:     []bool{false, true},
				SomeStringSlice:   []string{"someval"},
				SomeStringArray:   []string{"someval-array"},
				SomeIPSlice:       ipSlice,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecodeFlags(tt.args.cmd, tt.args.dst); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("DecodeFlags() error = %v, wantErr %v", err, tt.err)
			}

			if !reflect.DeepEqual(tt.args.dst, tt.wantStruct) {
				t.Errorf("DecodeFlags() values = \n%+v, wantStruct \n%+v", tt.args.dst, tt.wantStruct)
			}
		})
	}
}
