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

import (
	"fmt"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/elastic/cloud-sdk-go/pkg/util/slice"
)

// TrackChange iterates over a deployment's resources pending plans, sending
// updates to the returned channel in the form of TrackResponse every frequency
// configured in the parameter's TrackFrequencyConfig.
// When all of deployment's resources pending plans have finished, the channel
// is automatically closed by the goroutine that this function launches. It is
// possible to iterate with a for loop and assume that the loop will exit after
// all of the updates have been sent and the channel has been closed.
// If a ResourceID and Kind are set instead of the DeploymentID, a reverse
// lookup will be performed in order to find the DeploymentID and be able to
// track the pending plan.
func TrackChange(params TrackChangeParams) (<-chan TrackResponse, error) {
	params.Config.fillDefaults()
	if err := params.Validate(); err != nil {
		return nil, err
	}

	deploymentID, err := getDeploymentID(params)
	if err != nil {
		return nil, err
	}
	params.DeploymentID = deploymentID

	var out = make(chan TrackResponse)
	go trackChange(params, out, time.NewTicker(params.Config.PollFrequency))

	return out, nil
}

func trackChange(params TrackChangeParams, c chan<- TrackResponse, ticker *time.Ticker) {
	// Close the channel before the function returns. This particularly
	// important so that clients consuming this channel can use it in
	// a for loop and assume that when the foor loop ends, the change is
	// complete.
	defer close(c)

	// retries is used as a simple counter which is incremented every time an
	// error occurs, or when the returned pending plan slice is 0.
	var retries int

	// changedResources is a list of resource IDs which have been seen to have
	// a pending plan. It's used to filter out any resources which weren't
	// part of the last plan change.
	var changedResources []string
	for range ticker.C {
		// After the retries number is higher or equal to MaxRetries, the plan
		// changed is considered complete. In which case, the current plan or
		// the last plan in the plan history is checked to obtain the last plan
		// step log, and decode any errors which might've happened or mark the
		// plan as succeeded.
		if retries >= params.Config.MaxRetries {
			var checkRetries int
			checkCurrentStatus(params, c, changedResources, checkRetries)
			return
		}

		res, err := params.V1API.Deployments.GetDeployment(
			deployments.NewGetDeploymentParams().
				WithDeploymentID(params.DeploymentID).
				WithShowPlanLogs(ec.Bool(true)).
				WithShowPlans(ec.Bool(true)),
			params.AuthWriter,
		)
		if err != nil {
			retries++
			continue
		}

		var plans = buildTrackResponse(res.Payload.Resources, false)
		if len(plans) == 0 {
			retries++
		}

		for _, p := range plans {
			changedResources = append(changedResources, p.ID)
			p.DeploymentID = *res.Payload.ID
			ignoreChange := params.Kind != p.Kind && params.IgnoreDownstream
			if ignoreChange {
				continue
			}
			c <- p
		}
	}
}

// getDeploymentID ensures that a DeploymentID is found, if the DeploymentID
// has already been set in the parameters, it simply returns that ID, otherwise
// performs a deployment search to obtain the Deployment ID from a resource ID
// and Kind.
func getDeploymentID(params TrackChangeParams) (string, error) {
	if params.DeploymentID != "" {
		return params.DeploymentID, nil
	}

	res, err := params.V1API.Deployments.SearchDeployments(
		deployments.NewSearchDeploymentsParams().
			WithBody(NewReverseLookupQuery(params.ResourceID, params.Kind)),
		params.AuthWriter,
	)
	if err != nil {
		return "", api.UnwrapError(err)
	}

	if len(res.Payload.Deployments) > 0 {
		return *res.Payload.Deployments[0].ID, nil
	}

	return "", fmt.Errorf(
		"plan track change: couldn't find a deployment containing Kind %s with ID %s",
		params.Kind, params.ResourceID,
	)
}

// checkCurrentStatus is run after the deployment's resources pending plans
// have finished. It's necessary for a couple of reasons:
//   1. Catching any errors which might've happen in the pending plan but
//      weren't caught because the plan finished in between polling periods.
//   2. Posting the end result of the resource back to the channel.
// Additionally, changedResources is sent as a parameter to filter out any of
// the deployment's resources which weren't involved in the plan change.
func checkCurrentStatus(params TrackChangeParams, c chan<- TrackResponse, changedResources []string, retries int) {
	res, err := params.V1API.Deployments.GetDeployment(
		deployments.NewGetDeploymentParams().
			WithDeploymentID(params.DeploymentID).
			WithShowPlanLogs(ec.Bool(true)).
			WithShowPlans(ec.Bool(true)).
			// Necessary for deployments which failed on creation
			WithShowPlanHistory(ec.Bool(true)),
		params.AuthWriter,
	)
	if err != nil {
		// retry the API call again until params.Config.MaxRetries is reached.
		if retries < params.Config.MaxRetries {
			retries++
			checkCurrentStatus(params, c, changedResources, retries)
		}
		return
	}

	for _, trackResponse := range buildTrackResponse(res.Payload.Resources, true) {
		if slice.HasString(changedResources, trackResponse.ID) {
			ignoreChange := params.Kind != trackResponse.Kind && params.IgnoreDownstream
			if ignoreChange {
				continue
			}
			trackResponse.DeploymentID = *res.Payload.ID
			c <- trackResponse
		}
	}
}
