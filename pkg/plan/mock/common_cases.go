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

package planmock

// type LogConfig struct {
// 	N             int
// 	Elasticsearch string
// 	Kibana        string
// 	Apm           string
// 	Appsearch     string
// }

// func NewPending(cfg LogConfig) *models.DeploymentGetResponse {
// 	if cfg.N == 0 {
// 		cfg.N = 3
// 	}

// 	var config GenerateConfig
// 	if cfg.Elasticsearch != "" {
// 		config.Elasticsearch = append(config.Elasticsearch, GeneratedResourceConfig{
// 			PendingLog: NewPendingLog(cfg.N),
// 		})
// 	}
// 	if cfg.Kibana != "" {
// 		config.Kibana = append(config.Kibana, GeneratedResourceConfig{
// 			PendingLog: NewPendingLog(cfg.N),
// 		})
// 	}
// 	if cfg.Apm != "" {
// 		config.Apm = append(config.Apm, GeneratedResourceConfig{
// 			PendingLog: NewPendingLog(cfg.N),
// 		})
// 	}
// 	if cfg.Appsearch != "" {
// 		config.Appsearch = append(config.Appsearch, GeneratedResourceConfig{
// 			PendingLog: NewPendingLog(cfg.N),
// 		})
// 	}
// 	return Generate(config)
// }

// func NewFinished(cfg LogConfig) *models.DeploymentGetResponse {
// 	var config GenerateConfig
// 	if cfg.Elasticsearch != "" {
// 		config.Elasticsearch = append(config.Elasticsearch, GeneratedResourceConfig{
// 			CurrentLog: NewFinishedLog(4),
// 		})
// 	}
// 	if cfg.Kibana != "" {
// 		config.Kibana = append(config.Kibana, GeneratedResourceConfig{
// 			CurrentLog: NewFinishedLog(4),
// 		})
// 	}
// 	if cfg.Apm != "" {
// 		config.Apm = append(config.Apm, GeneratedResourceConfig{
// 			CurrentLog: NewFinishedLog(4),
// 		})
// 	}
// 	if cfg.Appsearch != "" {
// 		config.Appsearch = append(config.Appsearch, GeneratedResourceConfig{
// 			CurrentLog: NewFinishedLog(4),
// 		})
// 	}
// 	return Generate(config)
// }

// func NewPendingLog(count int) []*models.ClusterPlanStepInfo {
// 	var logs = make([]*models.ClusterPlanStepInfo, count)
// 	for i := 0; i < count; i++ {
// 		if i+1 == count {
// 			logs = append(logs, NewPlanStep(fmt.Sprint("step-", i+1), "pending"))
// 			continue
// 		}
// 		logs = append(logs, NewPlanStep(fmt.Sprint("step-", i+1), "success"))
// 	}
// 	return logs
// }

// func NewFinishedLog(count int) []*models.ClusterPlanStepInfo {
// 	var logs = make([]*models.ClusterPlanStepInfo, count)
// 	for i := 0; i < count; i++ {
// 		if i+1 == count {
// 			logs = append(logs, NewPlanStep(fmt.Sprint("plan-completed", i+1), "success"))
// 			continue
// 		}
// 		logs = append(logs, NewPlanStep(fmt.Sprint("step-", i+1), "success"))
// 	}
// 	return logs
// }
