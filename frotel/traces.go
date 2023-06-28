package frotel

import (
	"context"
	"github.com/Ryanair/gofrlib/log"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewTraceProvider(ctx context.Context) (*trace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	if err != nil {
		log.Error("Error creating exporter: %v", err)
		return nil, errors.Wrapf(err, "Error creating exporter")
	}

	resourcesMerged, err := buildResources(ctx)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resourcesMerged),
	)
	if err != nil {
		log.Error("Error creating otel tracer provider: %v", err)
		return nil, err
	}

	otel.SetTracerProvider(tp)
	return tp, nil
}
