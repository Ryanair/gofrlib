package otel

import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"

func InstrumentHandler(handlerFunc interface{}) interface{} {
	return otellambda.InstrumentHandler(handlerFunc)
}
