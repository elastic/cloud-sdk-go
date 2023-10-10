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

package allocatorapi

import (
	"context"
	"fmt"
	"strings"

	hashimultierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/plan"
	"github.com/elastic/cloud-sdk-go/pkg/plan/planutil"
	"github.com/elastic/cloud-sdk-go/pkg/sync/pool"
	"github.com/elastic/cloud-sdk-go/pkg/util"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/elastic/cloud-sdk-go/pkg/util/slice"
)

const (
	// PlanPendingMessage is used to discard
	PlanPendingMessage = "There is a plan still pending, cancel that or wait for it to complete before restarting"
)

// Vacate drains allocated resource instances away from the allocator list either to
// a specific allocator list or we let the constructor decide if that is empty.
// If resources is set, it will only move the instances that are part of those IDs.
// If kind is specified, it will only move the resources that match that kind.
// If none is specified it will add all of the resources in the allocator.
// The maximum concurrent moves is controlled by the Concurrency parameter.
func Vacate(params *VacateParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	var emptyTimeout pool.Timeout
	if params.PoolTimeout == emptyTimeout {
		params.PoolTimeout = pool.DefaultTimeout
	}

	p, err := pool.NewPool(pool.Params{
		Size:    params.Concurrency,
		Run:     VacateClusterInPool,
		Timeout: params.PoolTimeout,
		Writer:  params.Output,
	})
	if err != nil {
		return err
	}

	// Errors reported here are returned by a dry-run execution of the move api (validation-only flag is used)
	// and we don't want to stop the real vacate.
	// Instead we are returning the validateOnlyErr with the actual vacate validateOnlyErr at the end of the function
	leftovers, hasWork, validateOnlyErr := moveAllocators(params, p)

	if err := p.Start(); err != nil {
		return err
	}

	// If the queue was full prior to starting it, there might be some
	// leftovers, ranging until the leftovers are inexistent as the pool
	// clears out the work items.
	for len(leftovers) > 0 {
		leftovers, _ = p.Add(leftovers...)
	}

	var merr = multierror.NewPrefixed("vacate error")
	unpackMultierror(merr, validateOnlyErr)

	// Wait until all of the items have been processed and unpack the errors
	// from the pooled vacate calls tos the multierror.
	if err := waitVacateCompletion(p, hasWork); err != nil {
		unpackMultierror(merr, err)
	}

	return multierror.WithFormat(merr.ErrorOrNil(), params.OutputFormat)
}

// moveAllocators ranges over the list of provided allocators and moves the
// nodes off each allocator, finally, returns any leftovers from full pool
// queues, whether or not any work was added to the pool, and potential errors
// returned from API calls.
func moveAllocators(params *VacateParams, p *pool.Pool) ([]pool.Validator, bool, error) {
	var leftovers []pool.Validator
	var merr = multierror.NewPrefixed("vacate error")
	var hasWork bool
	for _, id := range params.Allocators {
		left, moved, err := moveNodes(id, params, p)
		merr = merr.Append(err)
		if len(left) > 0 {
			leftovers = append(leftovers, left...)
		}

		if moved {
			hasWork = true
		}
	}
	return leftovers, hasWork, merr.ErrorOrNil()
}

// moveNodes moves all of the nodes off the specified allocator
func moveNodes(id string, params *VacateParams, p *pool.Pool) ([]pool.Validator, bool, error) {
	var merr = multierror.NewPrefixed("vacate error")
	res, err := params.API.V1API.PlatformInfrastructure.MoveClusters(
		platform_infrastructure.NewMoveClustersParams().
			WithAllocatorID(id).
			WithMoveOnly(params.MoveOnly).
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithValidateOnly(ec.Bool(true)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, false, merr.Append(VacateError{
			AllocatorID: id,
			Err:         apierror.Wrap(err),
		})
	}

	if err := CheckVacateFailures(res.Payload.Failures, params.ClusterFilter, id); err != nil {
		// Errors already wrapped in VacateError
		merr = merr.Append(err)
	}

	work, hasWork := addAllocatorMovesToPool(addAllocatorMovesToPoolParams{
		ID:           id,
		Pool:         p,
		Moves:        res.Payload.Moves,
		VacateParams: params,
	})

	return work, hasWork, merr.ErrorOrNil()
}

// waitVacateCompletion waits for the pool to be finished if there's work
// items that were added and when all of the items have been processed
// stops the pool, and returns a multierror with any leftovers from the
// stopped pool.
func waitVacateCompletion(p *pool.Pool, hasWork bool) error {
	var merr = multierror.NewPrefixed("vacate error")
	if hasWork {
		if err := p.Wait(); err != nil {
			unpackMultierror(merr, err)
		}
	}

	if p.Status() < pool.StoppingStatus {
		// Stop the pool once we've finished all the work
		if err := p.Stop(); err != nil && err != pool.ErrStopOperationTimedOut {
			unpackMultierror(merr, err)
		}
	}

	leftovers, _ := p.Leftovers()
	for _, lover := range leftovers {
		if params, ok := lover.(*VacateClusterParams); ok {
			merr = merr.Append(VacateError{
				AllocatorID: params.ID,
				ResourceID:  params.ClusterID,
				Kind:        params.Kind,
				Err: apierror.JSONError{
					Message: "was either cancelled or not processed, follow up accordingly",
				},
			})
		}
	}

	return merr.ErrorOrNil()
}

// nolint since the linter says it's too complex, it's because of the for loop
// combination with ifs. It's better to have this grouped rather than scattered
// around.
func addAllocatorMovesToPool(params addAllocatorMovesToPoolParams) ([]pool.Validator, bool) {
	var leftovers []pool.Validator
	var vacates = make([]pool.Validator, 0)
	if params.Moves == nil {
		return leftovers, len(vacates) > 0
	}

	var filter = params.VacateParams.ClusterFilter
	var kindFilter = params.VacateParams.KindFilter
	for _, move := range params.Moves.ElasticsearchClusters {
		if len(filter) > 0 && !slice.HasString(filter, *move.ClusterID) {
			continue
		}

		var kind = util.Elasticsearch
		if kindFilter != "" && kind != kindFilter {
			break
		}
		vacates = append(vacates, newVacateClusterParams(params, *move.ClusterID, kind))
	}

	for _, move := range params.Moves.KibanaClusters {
		if len(filter) > 0 && !slice.HasString(filter, *move.ClusterID) {
			continue
		}

		var kind = util.Kibana
		if kindFilter != "" && kind != kindFilter {
			break
		}
		vacates = append(vacates, newVacateClusterParams(params, *move.ClusterID, kind))
	}

	for _, move := range params.Moves.ApmClusters {
		if len(filter) > 0 && !slice.HasString(filter, *move.ClusterID) {
			continue
		}

		var kind = util.Apm
		if kindFilter != "" && kind != kindFilter {
			break
		}
		vacates = append(vacates, newVacateClusterParams(params, *move.ClusterID, kind))
	}

	for _, move := range params.Moves.AppsearchClusters {
		if len(filter) > 0 && !slice.HasString(filter, *move.ClusterID) {
			continue
		}

		var kind = util.Appsearch
		if kindFilter != "" && kind != kindFilter {
			break
		}
		vacates = append(vacates, newVacateClusterParams(params, *move.ClusterID, kind))
	}

	for _, move := range params.Moves.EnterpriseSearchClusters {
		if len(filter) > 0 && !slice.HasString(filter, *move.ClusterID) {
			continue
		}

		var kind = util.EnterpriseSearch
		if kindFilter != "" && kind != kindFilter {
			break
		}
		vacates = append(vacates, newVacateClusterParams(params, *move.ClusterID, kind))
	}

	if leftover, _ := params.Pool.Add(vacates...); len(leftover) > 0 {
		leftovers = append(leftovers, leftover...)
	}

	return leftovers, len(vacates) > 0
}

func newVacateClusterParams(params addAllocatorMovesToPoolParams, id, kind string) *VacateClusterParams {
	clusterParams := VacateClusterParams{
		API:                 params.VacateParams.API,
		ID:                  params.ID,
		Kind:                kind,
		ClusterID:           id,
		Region:              params.VacateParams.Region,
		SkipTracking:        params.VacateParams.SkipTracking,
		ClusterFilter:       params.VacateParams.ClusterFilter,
		PreferredAllocators: params.VacateParams.PreferredAllocators,
		MaxPollRetries:      params.VacateParams.MaxPollRetries,
		TrackFrequency:      params.VacateParams.TrackFrequency,
		Output:              params.VacateParams.Output,
		OutputFormat:        params.VacateParams.OutputFormat,
		MoveOnly:            params.VacateParams.MoveOnly,
		PlanOverrides:       params.VacateParams.PlanOverrides,
	}

	if params.VacateParams.AllocatorDown != nil {
		clusterParams.AllocatorDown = params.VacateParams.AllocatorDown
	}

	return &clusterParams
}

// VacateClusterInPool vacates a resource from an allocator, complying
// with the pool.RunFunc signature.
func VacateClusterInPool(p pool.Validator) error {
	if p == nil {
		return errors.New("allocator vacate: params cannot be nil")
	}

	if params, ok := p.(*VacateClusterParams); ok {
		return VacateCluster(params)
	}

	return errors.New("allocator vacate: failed casting parameters to *VacateClusterParams")
}

// VacateCluster moves a resource node off an allocator.
func VacateCluster(params *VacateClusterParams) error {
	params, err := fillVacateClusterParams(params)
	if err != nil {
		return err
	}

	if err := moveClusterByType(params); err != nil {
		return multierror.WithFormat(err, params.OutputFormat)
	}

	if params.SkipTracking {
		return nil
	}

	return planutil.TrackChange(planutil.TrackChangeParams{
		TrackChangeParams: plan.TrackChangeParams{
			API:              params.API,
			ResourceID:       params.ClusterID,
			Kind:             params.Kind,
			IgnoreDownstream: true,
			Config: plan.TrackFrequencyConfig{
				PollFrequency: params.TrackFrequency,
				MaxRetries:    int(params.MaxPollRetries),
			},
		},
		Writer: params.Output,
		Format: params.OutputFormat,
	})
}

// fillVacateClusterParams validates the parameters and fills any missing
// properties that are set to a default if empty. Performs a Get on the
// allocator to discover the allocator health if AllocatorDown is nil.
func fillVacateClusterParams(params *VacateClusterParams) (*VacateClusterParams, error) {
	if params == nil {
		return nil, errors.New("allocator vacate: params cannot be nil")
	}

	if err := params.Validate(); err != nil {
		return nil, multierror.NewPrefixed(fmt.Sprintf(
			"allocator %s: resource id [%s][%s]",
			params.ID, params.ClusterID, params.Kind), err,
		)
	}

	if params.AllocatorDown == nil {
		alloc, err := Get(
			GetParams{API: params.API, ID: params.ID, Region: params.Region},
		)
		if err != nil {
			return nil, VacateError{
				AllocatorID: params.ID,
				ResourceID:  params.ClusterID,
				Kind:        params.Kind,
				Ctx:         "allocator health autodiscovery",
				Err:         err,
			}
		}
		if alloc.Status != nil {
			params.AllocatorDown = ec.Bool(!*alloc.Status.Connected || !*alloc.Status.Healthy)
		}
	}

	if params.MaxPollRetries == 0 {
		params.MaxPollRetries = util.DefaultRetries
	}

	if params.TrackFrequency.Nanoseconds() == 0 {
		params.TrackFrequency = util.DefaultPollFrequency
	}

	return params, nil
}

// newMoveClusterParams
func newMoveClusterParams(params *VacateClusterParams) (*platform_infrastructure.MoveClustersByTypeParams, error) {
	// By setting the ClusterID in the request body, the API will only return the matched cluster's plan information.
	// This greatly reduces the amount of work that the API has to perform to return the calculated plan.
	req := getVacateRequestByClusterID(params.ClusterID, params.Kind)

	res, err := params.API.V1API.PlatformInfrastructure.MoveClusters(
		platform_infrastructure.NewMoveClustersParams().
			WithAllocatorDown(params.AllocatorDown).
			WithMoveOnly(params.MoveOnly).
			WithAllocatorID(params.ID).
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithValidateOnly(ec.Bool(true)).
			WithBody(req),
		params.AuthWriter,
	)
	if err != nil {
		return nil, VacateError{
			AllocatorID: params.ID,
			ResourceID:  params.ClusterID,
			Kind:        params.Kind,
			Ctx:         "failed obtaining default vacate parameters",
			Err:         apierror.Wrap(err),
		}
	}

	req = ComputeVacateRequest(res.Payload.Moves,
		[]string{params.ClusterID},
		params.PreferredAllocators,
		params.PlanOverrides,
	)

	var moveParams = platform_infrastructure.NewMoveClustersByTypeParams().
		WithAllocatorID(params.ID).
		WithAllocatorDown(params.AllocatorDown).
		WithContext(api.WithRegion(context.Background(), params.Region)).
		WithBody(req)

	if len(req.ElasticsearchClusters) > 0 {
		moveParams.SetClusterType(util.Elasticsearch)
	}

	if len(req.KibanaClusters) > 0 {
		moveParams.SetClusterType(util.Kibana)
	}

	if len(req.ApmClusters) > 0 {
		moveParams.SetClusterType(util.Apm)
	}

	if len(req.AppsearchClusters) > 0 {
		moveParams.SetClusterType(util.Appsearch)
	}

	if len(req.EnterpriseSearchClusters) > 0 {
		moveParams.SetClusterType(util.EnterpriseSearch)
	}

	return moveParams, nil
}

// moveClusterByType moves a resource's node from its allocator
func moveClusterByType(params *VacateClusterParams) error {
	moveParams, err := newMoveClusterParams(params)
	if err != nil {
		return err
	}

	res, err := params.API.V1API.PlatformInfrastructure.MoveClustersByType(
		moveParams, params.AuthWriter,
	)

	if err != nil {
		return VacateError{
			AllocatorID: params.ID,
			ResourceID:  params.ClusterID,
			Kind:        params.Kind,
			Ctx:         "resource move API call error",
			Err:         apierror.Wrap(err),
		}
	}

	return CheckVacateFailures(res.Payload.Failures, params.ClusterFilter, params.ID)
}

// CheckVacateFailures iterates over the list of failures returning a multierror
// of VacateError if any of failures are found.
//
// nolint because of the complexity score here
func CheckVacateFailures(failures *models.MoveClustersDetails, filter []string, allocatorID string) error {
	if failures == nil {
		return nil
	}

	var merr = multierror.NewPrefixed("vacate error")
	const errMsgFmt = "%s (%s)"
	const errCtx = "failed vacating"
	for _, failure := range failures.ElasticsearchClusters {
		if len(filter) > 0 && !slice.HasString(filter, *failure.ClusterID) {
			continue
		}

		var ferr error
		if len(failure.Errors) > 0 {
			var err = failure.Errors[0]
			ferr = fmt.Errorf("%s (%s)", *err.Message, *err.Code)
		}
		if !strings.Contains(ferr.Error(), PlanPendingMessage) {
			merr = merr.Append(VacateError{
				AllocatorID: allocatorID,
				ResourceID:  *failure.ClusterID,
				Kind:        util.Elasticsearch,
				Ctx:         errCtx,
				Err:         ferr,
			})
		}
	}

	for _, failure := range failures.KibanaClusters {
		if len(filter) > 0 && !slice.HasString(filter, *failure.ClusterID) {
			continue
		}

		var ferr error
		if len(failure.Errors) > 0 {
			var err = failure.Errors[0]
			ferr = fmt.Errorf("%s (%s)", *err.Message, *err.Code)
		}

		if !strings.Contains(ferr.Error(), PlanPendingMessage) {
			merr = merr.Append(VacateError{
				AllocatorID: allocatorID,
				ResourceID:  *failure.ClusterID,
				Kind:        util.Kibana,
				Ctx:         errCtx,
				Err:         ferr,
			})
		}
	}

	for _, failure := range failures.ApmClusters {
		if len(filter) > 0 && !slice.HasString(filter, *failure.ClusterID) {
			continue
		}

		var ferr error
		if len(failure.Errors) > 0 {
			var err = failure.Errors[0]
			ferr = fmt.Errorf("%s (%s)", *err.Message, *err.Code)
		}

		if !strings.Contains(ferr.Error(), PlanPendingMessage) {
			merr = merr.Append(VacateError{
				AllocatorID: allocatorID,
				ResourceID:  *failure.ClusterID,
				Kind:        util.Apm,
				Ctx:         errCtx,
				Err:         ferr,
			})
		}
	}

	for _, failure := range failures.AppsearchClusters {
		if len(filter) > 0 && !slice.HasString(filter, *failure.ClusterID) {
			continue
		}

		var ferr error
		if len(failure.Errors) > 0 {
			var err = failure.Errors[0]
			ferr = fmt.Errorf("%s (%s)", *err.Message, *err.Code)
		}

		if !strings.Contains(ferr.Error(), PlanPendingMessage) {
			merr = merr.Append(VacateError{
				AllocatorID: allocatorID,
				ResourceID:  *failure.ClusterID,
				Kind:        util.Appsearch,
				Ctx:         errCtx,
				Err:         ferr,
			})
		}
	}

	for _, failure := range failures.EnterpriseSearchClusters {
		if len(filter) > 0 && !slice.HasString(filter, *failure.ClusterID) {
			continue
		}

		var ferr error
		if len(failure.Errors) > 0 {
			var err = failure.Errors[0]
			ferr = fmt.Errorf("%s (%s)", *err.Message, *err.Code)
		}

		if !strings.Contains(ferr.Error(), PlanPendingMessage) {
			merr = merr.Append(VacateError{
				AllocatorID: allocatorID,
				ResourceID:  *failure.ClusterID,
				Kind:        util.EnterpriseSearch,
				Ctx:         errCtx,
				Err:         ferr,
			})
		}
	}

	return merr.ErrorOrNil()
}

// getVacateRequestByClusterID makes models.MoveClusterRequest object which contains a cluster ID
// and the object will be set to body of an API call which will retrieve calculated plan data to
// be used to move a node.
func getVacateRequestByClusterID(clusterID, clusterType string) *models.MoveClustersRequest {
	var req models.MoveClustersRequest

	switch clusterType {
	case util.Elasticsearch:
		req.ElasticsearchClusters = append(req.ElasticsearchClusters,
			&models.MoveElasticsearchClusterConfiguration{
				ClusterIds: []string{clusterID},
			},
		)
	case util.Kibana:
		req.KibanaClusters = append(req.KibanaClusters,
			&models.MoveKibanaClusterConfiguration{
				ClusterIds: []string{clusterID},
			},
		)
	case util.Apm:
		req.ApmClusters = append(req.ApmClusters,
			&models.MoveApmClusterConfiguration{
				ClusterIds: []string{clusterID},
			},
		)
	case util.Appsearch:
		req.AppsearchClusters = append(req.AppsearchClusters,
			&models.MoveAppSearchConfiguration{
				ClusterIds: []string{clusterID},
			},
		)
	case util.EnterpriseSearch:
		req.EnterpriseSearchClusters = append(req.EnterpriseSearchClusters,
			&models.MoveEnterpriseSearchConfiguration{
				ClusterIds: []string{clusterID},
			},
		)
	}

	return &req
}

// ComputeVacateRequest filters the tentative resources that would be moved and
// filters those by ID if it's specified, also setting any preferred allocators
// if that is sent. Any resource plan overrides will be set in this function.
// nolint due to complexity
func ComputeVacateRequest(pr *models.MoveClustersDetails, resources, to []string, overrides PlanOverrides) *models.MoveClustersRequest {
	var req models.MoveClustersRequest
	for _, c := range pr.ElasticsearchClusters {
		if len(resources) > 0 && !slice.HasString(resources, *c.ClusterID) {
			continue
		}

		if overrides.SkipSnapshot != nil {
			c.CalculatedPlan.PlanConfiguration.SkipSnapshot = overrides.SkipSnapshot
		}

		if overrides.SkipDataMigration != nil {
			c.CalculatedPlan.PlanConfiguration.SkipDataMigration = overrides.SkipDataMigration
		}

		if overrides.OverrideFailsafe != nil {
			c.CalculatedPlan.PlanConfiguration.OverrideFailsafe = overrides.OverrideFailsafe
		}

		c.CalculatedPlan.PlanConfiguration.PreferredAllocators = to
		req.ElasticsearchClusters = append(req.ElasticsearchClusters,
			&models.MoveElasticsearchClusterConfiguration{
				ClusterIds:   []string{*c.ClusterID},
				PlanOverride: c.CalculatedPlan,
			},
		)
	}

	for _, c := range pr.KibanaClusters {
		if len(resources) > 0 && !slice.HasString(resources, *c.ClusterID) {
			continue
		}

		c.CalculatedPlan.PlanConfiguration.PreferredAllocators = to
		req.KibanaClusters = append(req.KibanaClusters,
			&models.MoveKibanaClusterConfiguration{
				ClusterIds:   []string{*c.ClusterID},
				PlanOverride: c.CalculatedPlan,
			},
		)
	}

	for _, c := range pr.ApmClusters {
		if len(resources) > 0 && !slice.HasString(resources, *c.ClusterID) {
			continue
		}

		c.CalculatedPlan.PlanConfiguration.PreferredAllocators = to
		req.ApmClusters = append(req.ApmClusters,
			&models.MoveApmClusterConfiguration{
				ClusterIds:   []string{*c.ClusterID},
				PlanOverride: c.CalculatedPlan,
			},
		)
	}

	for _, c := range pr.AppsearchClusters {
		if len(resources) > 0 && !slice.HasString(resources, *c.ClusterID) {
			continue
		}

		c.CalculatedPlan.PlanConfiguration.PreferredAllocators = to
		req.AppsearchClusters = append(req.AppsearchClusters,
			&models.MoveAppSearchConfiguration{
				ClusterIds:   []string{*c.ClusterID},
				PlanOverride: c.CalculatedPlan,
			},
		)
	}

	for _, c := range pr.EnterpriseSearchClusters {
		if len(resources) > 0 && !slice.HasString(resources, *c.ClusterID) {
			continue
		}

		c.CalculatedPlan.PlanConfiguration.PreferredAllocators = to
		req.EnterpriseSearchClusters = append(req.EnterpriseSearchClusters,
			&models.MoveEnterpriseSearchConfiguration{
				ClusterIds:   []string{*c.ClusterID},
				PlanOverride: c.CalculatedPlan,
			},
		)
	}

	return &req
}

// unpackMultierror transforms a appends the individual errors to a multierror.Prefixed.
func unpackMultierror(merr *multierror.Prefixed, err error) {
	var hashimerr *hashimultierror.Error
	if errors.As(err, &hashimerr) {
		for _, v := range hashimerr.Errors {
			_ = merr.Append(v)
		}
		return
	}

	var prefixed *multierror.Prefixed
	if errors.As(err, &prefixed) {
		merr.Errors = append(merr.Errors, prefixed.Errors...)
		return
	}

	_ = merr.Append(err)
}
