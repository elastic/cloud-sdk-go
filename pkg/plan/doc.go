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

// Package plan provides an API to interact with a deployment's pending plan.
// It's mainly structured in two primitives: TrackChanges and Stream.
//
// There are a couple of ways to track a deployment's change, by Deployment ID
// or by Resource ID and Kind (elasticsearch, kibana, apm, appsearch, etc).
//
// Although the plan package is ready for external consumption, if your goal is
// track and stream the changes to a device in either Text or JSON, please take
// a look at github.com/elastic/cloud-sdk-go/pkg/plan/planutil.
//
//  channel, err := plan.TrackChange(plan.TrackChangeParams{
// 	API:              &api.API{}, // A real API instance needs to be used.
// 	DeploymentID:     "2e9c997ff4d0bfc273da17f549e45e76",
//  // ResourceID:    "6779ce55fc0646309ef812d007bb2526",
//  // Kind:          "elasticsearch",
// 	Config: plan.TrackFrequencyConfig{
// 		MaxRetries:    2, // # of API failures to accept. 2-4 recommended.
// 		PollFrequency: time.Second * 5, // 2-10s recommended.
// 	},
//  })
//  if err != nil {
//	return err
//  }
//
//  // Alternatively, plan.StreamJSON(channel, os.Stdout, false) can be used to
//  // print JSON formatted updates to an io.Writer.
//  if err := plan.Stream(channel, os.Stdout); err != nik {
//	 return err
//  }
//
// Legacy Documentation
//
// The plan.Track function has been marked as deprecated and will be removed in
// a future version, please refrain from using it or migrate to a new version
// before the code is removed. See below.
//
//	channel, err := plan.Track(plan.TrackParams{
//		API:           params.API,
//		ID:            params.ID,
//		Kind:          params.Kind,
//		PollFrequency: time.Second,
//		MaxRetries:    4,
//	})
//	if err != nil {
//		return err
//	}
//
//	plan.Stream(channel, params.Output)
//
package plan
