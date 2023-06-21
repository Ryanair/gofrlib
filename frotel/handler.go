package frotel

import (
	"context"
	"github.com/Ryanair/gofrlib/log"
	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
	"os"
)

func InstrumentHandler(tp *trace.TracerProvider, handlerFunc interface{}) interface{} {
	return otellambda.InstrumentHandler(handlerFunc,
		otellambda.WithTracerProvider(tp),
		otellambda.WithFlusher(tp),
		otellambda.WithPropagator(propagation.TraceContext{}))
}

func Start(handlerFunc interface{}) {
	ctx := context.Background()
	tp, err := NewProvider(ctx)
	if err != nil {
		log.Error("creating tracing provider failed", err)
		tp = trace.NewTracerProvider()
	}
	defer tp.Shutdown(ctx)

	lambda.Start(InstrumentHandler(tp, WrapWithPanicHandler(handlerFunc)))
}

func WrapWithPanicHandler(handlerFunc interface{}) interface{} {
	return func() {
		//defer handlePanic()

		// invoke handlerFunc if it's function
		if handlerFunc, ok := handlerFunc.(func()); ok {
			handlerFunc()
		}
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		log.Error("Panic occured: %v", r)
		os.Exit(1)
	}
}
