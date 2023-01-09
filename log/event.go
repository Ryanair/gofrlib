package log

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func SetUpSns(ctx context.Context, event events.SNSEvent) {
	SetupTraceIds(ctx)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, "sns",
			EventBody, ToString(event))
	}
}

func SetUpSnsRecord(ctx context.Context, event events.SNSEventRecord) {
	SetupTraceIds(ctx)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, event.EventSource,
			EventBody, ToString(event))
	}
}

func SetUpSqs(ctx context.Context, event events.SQSEvent) {
	SetupTraceIds(ctx)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, "sqs",
			EventBody, ToString(event))
	}
}

func SetUpSqsRecord(ctx context.Context, event events.SQSMessage) {
	SetupTraceIds(ctx)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, event.EventSource,
			EventBody, ToString(event))
	}
}

func SetUpDynamoRecord(ctx context.Context, event events.DynamoDBEventRecord) {
	SetupTraceIds(ctx)
	if IsDebugEnabled() {
		DebugW("Got event",
			EventSource, event.EventSource,
			EventBody, ToString(event))
	}
}

func WithSQSEvent(ctx context.Context, event events.SQSMessage) {
	SetupTraceIds(ctx)
	raw, _ := json.Marshal(event.Attributes)
	log = log.
		With(EventContextSource, event.EventSource).
		With(EventContextBody, event.Body).
		With(EventContextParams, string(raw))
}

func WithEvent(ctx context.Context, source string, body string, params map[string]string) {
	SetupTraceIds(ctx)
	raw, _ := json.Marshal(params)
	log = log.
		With(EventContextSource, source).
		With(EventContextBody, body).
		With(EventContextParams, string(raw))
}
