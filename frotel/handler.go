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

func initOtelProviders(ctx context.Context) *OtelProviders {
	otelProviders, err := NewProvider(ctx)
	if err != nil {
		log.Error("creating tracing provider failed", err)
		otelProviders = DefaultProviders()
	}

	return &otelProviders
}

func Start(handlerFunc interface{}) {
	ctx := context.Background()
	otelProviders := initOtelProviders(ctx)
	defer otelProviders.Shutdown(ctx)

	lambda.Start(InstrumentHandler(otelProviders.TracerProvider, handlerFunc))
}

func StartWithTerminationHook(handlerFunc interface{}, hook func()) {
	ctx := context.Background()
	otelProviders := initOtelProviders(ctx)
	defer otelProviders.Shutdown(ctx)

	lambda.StartWithOptions(InstrumentHandler(otelProviders.TracerProvider, handlerFunc),
		lambda.WithEnableSIGTERM(hook))
}
