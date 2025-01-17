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

//go:generate gowrap gen -p go.temporal.io/server/common/persistence -i QueueV2 -t gowrap_template -o queue_v2_gen.go -l ""

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	_sourcePersistence "go.temporal.io/server/common/persistence"
)

// telemetryQueueV2 implements QueueV2 interface instrumented with OpenTelemetry.
type telemetryQueueV2 struct {
	_sourcePersistence.QueueV2
	tracer trace.Tracer
}

// newTelemetryQueueV2 returns telemetryQueueV2.
func newTelemetryQueueV2(base _sourcePersistence.QueueV2, tracer trace.Tracer) telemetryQueueV2 {
	return telemetryQueueV2{
		QueueV2: base,
		tracer:  tracer,
	}
}

// CreateQueue wraps QueueV2.CreateQueue.
func (d telemetryQueueV2) CreateQueue(ctx context.Context, request *_sourcePersistence.InternalCreateQueueRequest) (ip1 *_sourcePersistence.InternalCreateQueueResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.QueueV2/CreateQueue")
	defer span.End()

	ip1, err = d.QueueV2.CreateQueue(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// EnqueueMessage wraps QueueV2.EnqueueMessage.
func (d telemetryQueueV2) EnqueueMessage(ctx context.Context, request *_sourcePersistence.InternalEnqueueMessageRequest) (ip1 *_sourcePersistence.InternalEnqueueMessageResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.QueueV2/EnqueueMessage")
	defer span.End()

	ip1, err = d.QueueV2.EnqueueMessage(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// ListQueues wraps QueueV2.ListQueues.
func (d telemetryQueueV2) ListQueues(ctx context.Context, request *_sourcePersistence.InternalListQueuesRequest) (ip1 *_sourcePersistence.InternalListQueuesResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.QueueV2/ListQueues")
	defer span.End()

	ip1, err = d.QueueV2.ListQueues(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// RangeDeleteMessages wraps QueueV2.RangeDeleteMessages.
func (d telemetryQueueV2) RangeDeleteMessages(ctx context.Context, request *_sourcePersistence.InternalRangeDeleteMessagesRequest) (ip1 *_sourcePersistence.InternalRangeDeleteMessagesResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.QueueV2/RangeDeleteMessages")
	defer span.End()

	ip1, err = d.QueueV2.RangeDeleteMessages(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}

// ReadMessages wraps QueueV2.ReadMessages.
func (d telemetryQueueV2) ReadMessages(ctx context.Context, request *_sourcePersistence.InternalReadMessagesRequest) (ip1 *_sourcePersistence.InternalReadMessagesResponse, err error) {
	ctx, span := d.tracer.Start(ctx, "persistence.QueueV2/ReadMessages")
	defer span.End()

	ip1, err = d.QueueV2.ReadMessages(ctx, request)
	if err != nil {
		span.RecordError(err)
	}

	return
}
