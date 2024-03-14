package frotel

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

// AddToCurrentSpan OpenTelemetry instructions https://opentelemetry.io/docs/instrumentation/go/manual/
func AddToCurrentSpan(ctx context.Context, kv ...attribute.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(kv...)
}

func SetStatus(ctx context.Context, code codes.Code, description string) {
	span := trace.SpanFromContext(ctx)
	span.SetStatus(code, description)
}

func RecordError(ctx context.Context, err error) {
	span := trace.SpanFromContext(ctx)
	span.RecordError(err)
}

func InstrumentSpan[T interface{}](ctx context.Context, spanName string, consumer func(ctx context.Context) T) T {
	if tracer == nil {
		tracer = otel.GetTracerProvider().Tracer("fr-otel-tracer")
	}
	spanCtx, span := tracer.Start(ctx, spanName)
	defer span.End()

	return consumer(spanCtx)
}

func InstrumentSpanWithError(ctx context.Context, spanName string, consumer func(ctx context.Context) error) error {
	if tracer == nil {
		tracer = otel.GetTracerProvider().Tracer("fr-otel-tracer")
	}
	spanCtx, span := tracer.Start(ctx, spanName)
	defer span.End()

	err := consumer(spanCtx)
	instrumentError(err, span)

	return err
}

func InstrumentSpanWithErr[T interface{}](ctx context.Context, spanName string, consumer func(ctx context.Context) (T, error)) (T, error) {
	if tracer == nil {
		tracer = otel.GetTracerProvider().Tracer("fr-otel-tracer")
	}
	spanCtx, span := tracer.Start(ctx, spanName)
	defer span.End()

	result, err := consumer(spanCtx)
	instrumentError(err, span)

	return result, err
}

func instrumentError(err error, span trace.Span) {
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
	}
}
