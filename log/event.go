package log

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

func SetUpSns(ctx context.Context, event events.SNSEvent) {
	SetupTraceIds(ctx)
	DebugW("Got event",
		EventSource, "sns",
		EventBody, ToString(event))
}

func SetUpSnsRecord(ctx context.Context, event events.SNSEventRecord) {
	SetupTraceIds(ctx)
	DebugW("Got event",
		EventSource, event.EventSource,
		EventBody, ToString(event))
}

func SetUpSqs(ctx context.Context, event events.SQSEvent) {
	SetupTraceIds(ctx)
	DebugW("Got event",
		EventSource, "sqs",
		EventBody, ToString(event))
}

func SetUpSqsRecord(ctx context.Context, event events.SQSMessage) {
	SetupTraceIds(ctx)
	DebugW("Got event",
		EventSource, event.EventSource,
		EventBody, ToString(event))
}
