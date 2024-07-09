package log

import (
	"fmt"
	"go.opentelemetry.io/otel/attribute"
)

const (
	TraceId       = "TraceId"
	CorrelationId = "CorrelationId"
	SpanId        = "SpanId"
	TraceFlags    = "TraceFlags"

	Timestamp = "Timestamp"
	Level     = "SeverityText"

	Message    = "Body.message"
	StackTrace = "Body.stacktrace"

	Logger                 = "Resource.logger"
	Application            = "Resource.application"
	Project                = "Resource.project"
	Environment            = "@env"
	ProjectGroup           = "Resource.projectGroup"
	ResourceServiceName    = "Resource.service.name"
	ResourceServiceVersion = "Resource.service.version"
	Version                = "Resource.version"

	EventSource = "Body.origin.event.eventSource"
	EventBody   = "Body.origin.event.eventBody"

	MessagingSourceSystemSqs             = "sqs"
	MessagingSourceSystemSns             = "sns"
	MessagingSourceSystemDynamoDbStreams = "dyanamodbstreams"
	MessagingSourceSystemKinesis         = "kinesis"
)

var MessagingSourceSystemDynamoDbStreamsMessageKey = attribute.Key(fmt.Sprintf("messaging.%s.message.key", MessagingSourceSystemDynamoDbStreams))

var MessagingMessageShard = attribute.Key("messaging.message.shard")
