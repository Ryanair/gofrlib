package frotel

import (
	"context"
	"github.com/Ryanair/gofrlib/log"
	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InstrumentHandler(tp *trace.TracerProvider, handlerFunc interface{}) interface{} {
	return otellambda.InstrumentHandler(handlerFunc,
		otellambda.WithTracerProvider(tp),
		otellambda.WithFlusher(tp),
		otellambda.WithPropagator(propagation.TraceContext{}))
}

func Start(handlerFunc interface{}) {
	ctx := context.Background()
	otelProviders, err := NewProvider(ctx)
	if err != nil {
		log.Error("creating tracing provider failed", err)
		otelProviders = DefaultProviders()
	}
	defer otelProviders.Shutdown(ctx)

	lambda.Start(InstrumentHandler(otelProviders.TracerProvider, handlerFunc))
}
