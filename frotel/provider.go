package frotel

import (
	"context"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

type OtelProviders struct {
	TracerProvider *trace.TracerProvider
	MeterProvider  *metric.MeterProvider
}

func (op *OtelProviders) Shutdown(ctx context.Context) {
	op.MeterProvider.Shutdown(ctx)
	op.TracerProvider.Shutdown(ctx)
}

func NewProvider(ctx context.Context) (OtelProviders, error) {
	tp, err := NewTraceProvider(ctx)
	if err != nil {
		return OtelProviders{}, err
	}

	mp, err := NewMetricProvider(ctx)
	if err != nil {
		return OtelProviders{}, err
	}

	return OtelProviders{
		TracerProvider: tp,
		MeterProvider:  mp,
	}, nil
}

func DefaultProviders() OtelProviders {
	return OtelProviders{
		TracerProvider: trace.NewTracerProvider(),
		MeterProvider:  metric.NewMeterProvider(),
	}
}
