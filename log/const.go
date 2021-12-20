package log

const (
	TraceId       = "TraceId"
	CorrelationId = "CorrelationId"
	SpanId        = "SpanId"
	TraceFlags    = "TraceFlags"

	Timestamp = "Timestamp"
	Level     = "SeverityText"

	Message    = "Body.message"
	StackTrace = "Body.stacktrace"

	Logger       = "Resource.logger"
	Application  = "Resource.application"
	Project      = "Resource.project"
	ProjectGroup = "Resource.projectGroup"
	Version      = "Resource.version"

	EventSource = "Body.origin.event.eventSource"
	EventBody   = "Body.origin.event.eventBody"
)
