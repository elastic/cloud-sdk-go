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

// Package planutil takes the functionality from the plan package and presents
// a more high-level API.
//
// Track a deployment's change via the Deployment ID with a blocking call which
// will wait until the deployment's resources pending plans have finished.
//
//  err := planutil.TrackChange(planutil.TrackChangeParams{
// 	TrackChangeParams: plan.TrackChangeParams{
// 		API:              &api.API{}, // A real API instance needs to be used.
// 		DeploymentID:     "2e9c997ff4d0bfc273da17f549e45e76",
// 		Config: plan.TrackFrequencyConfig{
// 			MaxRetries:    2, // # of API failures to accept. 2-4 recommended.
// 			PollFrequency: time.Second * 5, // 2-10s recommended.
// 		},
// 	},
// 	Format: "text", // "text", "json" or "" are allowed.
// 	Writer: os.Stdout,
//  )}
//  if err != nil {
//	return err
//  }
//
//
package planutil
