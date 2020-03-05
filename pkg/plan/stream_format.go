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

package plan

// Legacy cluster format
const (
	// legacyStreamFormat with green color
	legacyStreamFinishFormat = "\x1b[92;mCluster [%s][%s]: finished running all the plan steps\x1b[0m (Total plan duration: %s)\n"
	// legacyStreamFormatOnError with red color
	legacyStreamFinishErrFormat = "\x1b[91;1mCluster [%s][%s]: caught error: \"%s\"\x1b[0m (Total plan duration: %s)\n"

	// These formats are used when the plan has not yet finished.
	legacyStreamFormat    = "Cluster [%s][%s]: running step \"%s\" (Plan duration %s)...\n"
	legacyStreamErrFormat = "Cluster [%s][%s]: running step \"%s\" caught error: \"%s\" (Plan duration %s)...\n"
)

// Deployment format (Current).
const (
	// streamFinishFormat with green color
	streamFinishFormat = "\x1b[92;mDeployment [%s] - [%s][%s]: finished running all the plan steps\x1b[0m (Total plan duration: %s)\n"
	// streamFinishErrFormat with red color
	streamFinishErrFormat = "\x1b[91;1mDeployment [%s] - [%s][%s]: caught error: \"%s\"\x1b[0m (Total plan duration: %s)\n"

	// These formats are used when the plan has not yet finished.
	streamFormat    = "Deployment [%s] - [%s][%s]: running step \"%s\" (Plan duration %s)...\n"
	streamErrFormat = "Deployment [%s] - [%s][%s]: running step \"%s\" caught error: \"%s\" (Plan duration %s)...\n"
)
