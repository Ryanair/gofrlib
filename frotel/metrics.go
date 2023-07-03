package frotel

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

func deltaTemporalitySelector(metric.InstrumentKind) metricdata.Temporality {
	return metricdata.DeltaTemporality
}

func NewMetricProvider(ctx context.Context) (*metric.MeterProvider, error) {
	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithTemporalitySelector(deltaTemporalitySelector),
	)
	if err != nil {
		return nil, err
	}

	resourcesMerged, err := buildResources(ctx)
	if err != nil {
		return nil, err
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exporter)),
		metric.WithResource(resourcesMerged),
	)
	otel.SetMeterProvider(provider)

	return provider, nil
}
