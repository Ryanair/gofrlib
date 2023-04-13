package frotel

import (
	"context"
	"fmt"
	"github.com/Ryanair/gofrlib/log"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/detectors/aws/lambda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"os"
)

const (
	otelServiceVersion = "VERSION"
)

func NewProvider(ctx context.Context) (*trace.TracerProvider, error) {
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

func buildResources(ctx context.Context) (*resource.Resource, error) {
	resources, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithDetectors(lambda.NewResourceDetector()),
		resource.WithAttributes(
			semconv.ServiceVersionKey.String(getServiceVersion()),
			semconv.TelemetrySDKLanguageGo,
		),
	)
	if err != nil {
		log.Error("Error creating custom resources: %v", err)
		return nil, errors.Wrapf(err, "Error creating custom resources")
	}

	return resources, nil
}

func getServiceVersion() string {
	if version, defined := os.LookupEnv(otelServiceVersion); defined {
		return version
	} else {
		return fmt.Sprintf("%s_UNDEFINED", otelServiceVersion)
	}
}
