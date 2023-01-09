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

	EventSource        = "Body.origin.event.eventSource"
	EventBody          = "Body.origin.event.eventBody"
	EventContextSource = "Body.context.origin.event.eventSource"
	EventContextBody   = "Body.context.origin.event.eventBody"
	EventContextParams = "Body.context.origin.event.eventParams"

	EventContextMethod = "Body.context.origin.request.method"
	EventContextUrl    = "Body.context.origin.request.url"
	EventContextRoute  = "Body.context.origin.request.route"
	EventContextQuery  = "Body.context.origin.request.query"
	EventContextAgent  = "Body.context.origin.request.userAgent"
)
