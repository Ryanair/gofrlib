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
	eventSource := retrieveSNSEventSource(event)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemSns),
		semconv.MessagingBatchMessageCount(len(event.Records)),
		semconv.MessagingSourceName(eventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, eventSource,
			EventBody, ToString(event))
	}
}

func SetUpSnsRecord(ctx context.Context, event events.SNSEventRecord) {
	SetupTraceIds(ctx)
	eventSource := retrieveTopic(event)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemSns),
		semconv.MessagingSourceName(eventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, eventSource,
			EventBody, ToString(event))
	}
}

func SetUpSqs(ctx context.Context, event events.SQSEvent) {
	SetupTraceIds(ctx)
	eventSource := retrieveSQSEventSource(event)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemSqs),
		semconv.MessagingBatchMessageCount(len(event.Records)),
		semconv.MessagingSourceName(eventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, eventSource,
			EventBody, ToString(event))
	}
}

func SetUpSqsRecord(ctx context.Context, event events.SQSMessage) {
	SetupTraceIds(ctx)
	eventSource := retrieveQueueArn(event)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemSqs),
		semconv.MessagingSourceName(eventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, eventSource,
			EventBody, ToString(event))
	}
}

func SetUpDynamoEvent(ctx context.Context, event events.DynamoDBEvent) {
	SetupTraceIds(ctx)
	eventSource := retrieveDynamoEventSource(event)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemDynamoDbStreams),
		semconv.MessagingBatchMessageCount(len(event.Records)),
		semconv.MessagingSourceName(eventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, eventSource,
			EventBody, ToString(event))
	}
}

func SetUpDynamoRecord(ctx context.Context, event events.DynamoDBEventRecord) {
	SetupTraceIds(ctx)
	eventSource := retrieveStreamArn(event)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemDynamoDbStreams),
		semconv.MessagingSourceName(eventSource),
		semconv.MessagingMessageID(event.EventID),
		MessagingSourceSystemDynamoDbStreamsMessageKey.String(ToString(event.Change.Keys)),
		semconv.MessagingMessagePayloadSizeBytes(int(event.Change.SizeBytes)),
		semconv.MessagingOperationKey.String(event.EventName),
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, eventSource,
			EventBody, ToString(event))
	}
}

func SetUpKinesisEvent(ctx context.Context, event events.KinesisEvent) {
	SetupTraceIds(ctx)
	eventSource := retrieveKinesisEventSource(event)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemKinesis),
		semconv.MessagingBatchMessageCount(len(event.Records)),
		semconv.MessagingSourceName(eventSource),
		semconv.MessagingOperationProcess,
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, eventSource,
			EventBody, ToString(event))
	}
}

func SetUpKinesisRecord(ctx context.Context, event events.KinesisEventRecord) {
	SetupTraceIds(ctx)
	eventSource := retrieveKinesisArn(event)
	trace.SpanFromContext(ctx).SetAttributes(
		semconv.MessagingSystem(MessagingSourceSystemKinesis),
		semconv.MessagingSourceName(eventSource),
		semconv.MessagingMessageID(event.EventID),
		semconv.MessagingMessageID(event.Kinesis.SequenceNumber),
		MessagingMessageShard.String(event.Kinesis.PartitionKey),
		semconv.MessagingMessagePayloadSizeBytes(len(event.Kinesis.Data)),
		semconv.MessagingOperationKey.String(event.EventName),
	)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, eventSource,
			EventBody, ToString(event))
	}
}

func retrieveSNSEventSource(event events.SNSEvent) string {
	if len(event.Records) == 0 {
		return "missing SNS topic Arn"
	}
	return retrieveTopic(event.Records[0])
}

func retrieveTopic(event events.SNSEventRecord) string {
	return event.SNS.TopicArn
}

func retrieveSQSEventSource(event events.SQSEvent) string {
	if len(event.Records) == 0 {
		return "missing SQS queue Arn"
	}
	return retrieveQueueArn(event.Records[0])
}

func retrieveQueueArn(event events.SQSMessage) string {
	return event.EventSourceARN
}

func retrieveDynamoEventSource(event events.DynamoDBEvent) string {
	if len(event.Records) == 0 {
		return "missing DynamoDB Stream Arn"
	}
	return retrieveStreamArn(event.Records[0])
}

func retrieveStreamArn(event events.DynamoDBEventRecord) string {
	return event.EventSourceArn
}

func retrieveKinesisEventSource(event events.KinesisEvent) string {
	if len(event.Records) == 0 {
		return "missing Kinesis Stream Arn"
	}
	return retrieveKinesisArn(event.Records[0])
}

func retrieveKinesisArn(event events.KinesisEventRecord) string {
	return event.EventSourceArn
}
