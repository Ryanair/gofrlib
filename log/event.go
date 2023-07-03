package log

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"go.opentelemetry.io/otel/trace"
)

//Semantic conventions for messaging: https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/messaging.md

func SetUpSns(ctx context.Context, event events.SNSEvent) {
	SetupTraceIds(ctx)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemSns),
		semconv.MessagingBatchMessageCount(len(event.Records)),
		semconv.MessagingSourceName(event.Records[0].EventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, "sns",
			EventBody, ToString(event))
	}
}

func SetUpSnsRecord(ctx context.Context, event events.SNSEventRecord) {
	SetupTraceIds(ctx)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemSns),
		semconv.MessagingSourceName(event.EventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, event.EventSource,
			EventBody, ToString(event))
	}
}

func SetUpSqs(ctx context.Context, event events.SQSEvent) {
	SetupTraceIds(ctx)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemSqs),
		semconv.MessagingBatchMessageCount(len(event.Records)),
		semconv.MessagingSourceName(event.Records[0].EventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, "sqs",
			EventBody, ToString(event))
	}
}

func SetUpSqsRecord(ctx context.Context, event events.SQSMessage) {
	SetupTraceIds(ctx)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemSqs),
		semconv.MessagingSourceName(event.EventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, event.EventSource,
			EventBody, ToString(event))
	}
}

func SetUpDynamoEvent(ctx context.Context, event events.DynamoDBEvent) {
	SetupTraceIds(ctx)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemDynamoDbStreams),
		semconv.MessagingBatchMessageCount(len(event.Records)),
		semconv.MessagingSourceName(event.Records[0].EventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, event.Records[0].EventSource,
			EventBody, ToString(event))
	}
}

func SetUpDynamoRecord(ctx context.Context, event events.DynamoDBEventRecord) {
	SetupTraceIds(ctx)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemDynamoDbStreams),
		semconv.MessagingSourceName(event.EventSource),
		semconv.MessagingMessageID(event.EventID),
		MessagingSourceSystemDynamoDbStreamsMessageKey.String(ToString(event.Change.Keys)),
		semconv.MessagingMessagePayloadSizeBytes(int(event.Change.SizeBytes)),
		semconv.MessagingOperationKey.String(event.EventName),
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, event.EventSource,
			EventBody, ToString(event))
	}
}
