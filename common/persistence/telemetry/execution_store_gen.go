// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by gowrap. DO NOT EDIT.
// template: gowrap_template
// gowrap: http://github.com/hexdigest/gowrap

package telemetry

//go:generate gowrap gen -p go.temporal.io/server/common/persistence -i ExecutionStore -t gowrap_template -o execution_store_gen.go -l ""

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	_sourcePersistence "go.temporal.io/server/common/persistence"
)

// telemetryExecutionStore implements ExecutionStore interface instrumented with OpenTelemetry.
type telemetryExecutionStore struct {
	_sourcePersistence.ExecutionStore
	tracer trace.Tracer
}

// newTelemetryExecutionStore returns telemetryExecutionStore.
func newTelemetryExecutionStore(base _sourcePersistence.ExecutionStore, tracer trace.Tracer) telemetryExecutionStore {
	return telemetryExecutionStore{
		ExecutionStore: base,
		tracer:         tracer,
	}
}

// AddHistoryTasks wraps ExecutionStore.AddHistoryTasks.
func (d telemetryExecutionStore) AddHistoryTasks(ctx context.Context, request *_sourcePersistence.InternalAddHistoryTasksRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/AddHistoryTasks")
	defer span.End()

	err = d.ExecutionStore.AddHistoryTasks(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// AppendHistoryNodes wraps ExecutionStore.AppendHistoryNodes.
func (d telemetryExecutionStore) AppendHistoryNodes(ctx context.Context, request *_sourcePersistence.InternalAppendHistoryNodesRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/AppendHistoryNodes")
	defer span.End()

	err = d.ExecutionStore.AppendHistoryNodes(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// CompleteHistoryTask wraps ExecutionStore.CompleteHistoryTask.
func (d telemetryExecutionStore) CompleteHistoryTask(ctx context.Context, request *_sourcePersistence.CompleteHistoryTaskRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/CompleteHistoryTask")
	defer span.End()

	err = d.ExecutionStore.CompleteHistoryTask(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// ConflictResolveWorkflowExecution wraps ExecutionStore.ConflictResolveWorkflowExecution.
func (d telemetryExecutionStore) ConflictResolveWorkflowExecution(ctx context.Context, request *_sourcePersistence.InternalConflictResolveWorkflowExecutionRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/ConflictResolveWorkflowExecution")
	defer span.End()

	err = d.ExecutionStore.ConflictResolveWorkflowExecution(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// CreateWorkflowExecution wraps ExecutionStore.CreateWorkflowExecution.
func (d telemetryExecutionStore) CreateWorkflowExecution(ctx context.Context, request *_sourcePersistence.InternalCreateWorkflowExecutionRequest) (ip1 *_sourcePersistence.InternalCreateWorkflowExecutionResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/CreateWorkflowExecution")
	defer span.End()

	ip1, err = d.ExecutionStore.CreateWorkflowExecution(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// DeleteCurrentWorkflowExecution wraps ExecutionStore.DeleteCurrentWorkflowExecution.
func (d telemetryExecutionStore) DeleteCurrentWorkflowExecution(ctx context.Context, request *_sourcePersistence.DeleteCurrentWorkflowExecutionRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/DeleteCurrentWorkflowExecution")
	defer span.End()

	err = d.ExecutionStore.DeleteCurrentWorkflowExecution(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// DeleteHistoryBranch wraps ExecutionStore.DeleteHistoryBranch.
func (d telemetryExecutionStore) DeleteHistoryBranch(ctx context.Context, request *_sourcePersistence.InternalDeleteHistoryBranchRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/DeleteHistoryBranch")
	defer span.End()

	err = d.ExecutionStore.DeleteHistoryBranch(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// DeleteHistoryNodes wraps ExecutionStore.DeleteHistoryNodes.
func (d telemetryExecutionStore) DeleteHistoryNodes(ctx context.Context, request *_sourcePersistence.InternalDeleteHistoryNodesRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/DeleteHistoryNodes")
	defer span.End()

	err = d.ExecutionStore.DeleteHistoryNodes(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// DeleteReplicationTaskFromDLQ wraps ExecutionStore.DeleteReplicationTaskFromDLQ.
func (d telemetryExecutionStore) DeleteReplicationTaskFromDLQ(ctx context.Context, request *_sourcePersistence.DeleteReplicationTaskFromDLQRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/DeleteReplicationTaskFromDLQ")
	defer span.End()

	err = d.ExecutionStore.DeleteReplicationTaskFromDLQ(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// DeleteWorkflowExecution wraps ExecutionStore.DeleteWorkflowExecution.
func (d telemetryExecutionStore) DeleteWorkflowExecution(ctx context.Context, request *_sourcePersistence.DeleteWorkflowExecutionRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/DeleteWorkflowExecution")
	defer span.End()

	err = d.ExecutionStore.DeleteWorkflowExecution(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// ForkHistoryBranch wraps ExecutionStore.ForkHistoryBranch.
func (d telemetryExecutionStore) ForkHistoryBranch(ctx context.Context, request *_sourcePersistence.InternalForkHistoryBranchRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/ForkHistoryBranch")
	defer span.End()

	err = d.ExecutionStore.ForkHistoryBranch(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// GetAllHistoryTreeBranches wraps ExecutionStore.GetAllHistoryTreeBranches.
func (d telemetryExecutionStore) GetAllHistoryTreeBranches(ctx context.Context, request *_sourcePersistence.GetAllHistoryTreeBranchesRequest) (ip1 *_sourcePersistence.InternalGetAllHistoryTreeBranchesResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/GetAllHistoryTreeBranches")
	defer span.End()

	ip1, err = d.ExecutionStore.GetAllHistoryTreeBranches(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// GetCurrentExecution wraps ExecutionStore.GetCurrentExecution.
func (d telemetryExecutionStore) GetCurrentExecution(ctx context.Context, request *_sourcePersistence.GetCurrentExecutionRequest) (ip1 *_sourcePersistence.InternalGetCurrentExecutionResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/GetCurrentExecution")
	defer span.End()

	ip1, err = d.ExecutionStore.GetCurrentExecution(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// GetHistoryTasks wraps ExecutionStore.GetHistoryTasks.
func (d telemetryExecutionStore) GetHistoryTasks(ctx context.Context, request *_sourcePersistence.GetHistoryTasksRequest) (ip1 *_sourcePersistence.InternalGetHistoryTasksResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/GetHistoryTasks")
	defer span.End()

	ip1, err = d.ExecutionStore.GetHistoryTasks(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// GetHistoryTreeContainingBranch wraps ExecutionStore.GetHistoryTreeContainingBranch.
func (d telemetryExecutionStore) GetHistoryTreeContainingBranch(ctx context.Context, request *_sourcePersistence.InternalGetHistoryTreeContainingBranchRequest) (ip1 *_sourcePersistence.InternalGetHistoryTreeContainingBranchResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/GetHistoryTreeContainingBranch")
	defer span.End()

	ip1, err = d.ExecutionStore.GetHistoryTreeContainingBranch(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// GetReplicationTasksFromDLQ wraps ExecutionStore.GetReplicationTasksFromDLQ.
func (d telemetryExecutionStore) GetReplicationTasksFromDLQ(ctx context.Context, request *_sourcePersistence.GetReplicationTasksFromDLQRequest) (ip1 *_sourcePersistence.InternalGetReplicationTasksFromDLQResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/GetReplicationTasksFromDLQ")
	defer span.End()

	ip1, err = d.ExecutionStore.GetReplicationTasksFromDLQ(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// GetWorkflowExecution wraps ExecutionStore.GetWorkflowExecution.
func (d telemetryExecutionStore) GetWorkflowExecution(ctx context.Context, request *_sourcePersistence.GetWorkflowExecutionRequest) (ip1 *_sourcePersistence.InternalGetWorkflowExecutionResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/GetWorkflowExecution")
	defer span.End()

	ip1, err = d.ExecutionStore.GetWorkflowExecution(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// IsReplicationDLQEmpty wraps ExecutionStore.IsReplicationDLQEmpty.
func (d telemetryExecutionStore) IsReplicationDLQEmpty(ctx context.Context, request *_sourcePersistence.GetReplicationTasksFromDLQRequest) (b1 bool, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/IsReplicationDLQEmpty")
	defer span.End()

	b1, err = d.ExecutionStore.IsReplicationDLQEmpty(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// ListConcreteExecutions wraps ExecutionStore.ListConcreteExecutions.
func (d telemetryExecutionStore) ListConcreteExecutions(ctx context.Context, request *_sourcePersistence.ListConcreteExecutionsRequest) (ip1 *_sourcePersistence.InternalListConcreteExecutionsResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/ListConcreteExecutions")
	defer span.End()

	ip1, err = d.ExecutionStore.ListConcreteExecutions(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// PutReplicationTaskToDLQ wraps ExecutionStore.PutReplicationTaskToDLQ.
func (d telemetryExecutionStore) PutReplicationTaskToDLQ(ctx context.Context, request *_sourcePersistence.PutReplicationTaskToDLQRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/PutReplicationTaskToDLQ")
	defer span.End()

	err = d.ExecutionStore.PutReplicationTaskToDLQ(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// RangeCompleteHistoryTasks wraps ExecutionStore.RangeCompleteHistoryTasks.
func (d telemetryExecutionStore) RangeCompleteHistoryTasks(ctx context.Context, request *_sourcePersistence.RangeCompleteHistoryTasksRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/RangeCompleteHistoryTasks")
	defer span.End()

	err = d.ExecutionStore.RangeCompleteHistoryTasks(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// RangeDeleteReplicationTaskFromDLQ wraps ExecutionStore.RangeDeleteReplicationTaskFromDLQ.
func (d telemetryExecutionStore) RangeDeleteReplicationTaskFromDLQ(ctx context.Context, request *_sourcePersistence.RangeDeleteReplicationTaskFromDLQRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/RangeDeleteReplicationTaskFromDLQ")
	defer span.End()

	err = d.ExecutionStore.RangeDeleteReplicationTaskFromDLQ(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// ReadHistoryBranch wraps ExecutionStore.ReadHistoryBranch.
func (d telemetryExecutionStore) ReadHistoryBranch(ctx context.Context, request *_sourcePersistence.InternalReadHistoryBranchRequest) (ip1 *_sourcePersistence.InternalReadHistoryBranchResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/ReadHistoryBranch")
	defer span.End()

	ip1, err = d.ExecutionStore.ReadHistoryBranch(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// SetWorkflowExecution wraps ExecutionStore.SetWorkflowExecution.
func (d telemetryExecutionStore) SetWorkflowExecution(ctx context.Context, request *_sourcePersistence.InternalSetWorkflowExecutionRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/SetWorkflowExecution")
	defer span.End()

	err = d.ExecutionStore.SetWorkflowExecution(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// UpdateWorkflowExecution wraps ExecutionStore.UpdateWorkflowExecution.
func (d telemetryExecutionStore) UpdateWorkflowExecution(ctx context.Context, request *_sourcePersistence.InternalUpdateWorkflowExecutionRequest) (err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.ExecutionStore/UpdateWorkflowExecution")
	defer span.End()

	err = d.ExecutionStore.UpdateWorkflowExecution(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}
