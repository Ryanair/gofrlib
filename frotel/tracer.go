package frotel

import (
	"context"
	"fmt"
	"github.com/Ryanair/gofrlib/log"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/detectors/aws/lambda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"os"
)

const (
	otelServiceEnv = "OTEL_NAME"
)

func NewProvider(ctx context.Context) (*trace.TracerProvider, error) {
	resourcesMerged, err := buildResources(ctx)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(trace.WithResource(resourcesMerged))
	if err != nil {
		log.Error("Error creating otel tracer provider: %v", err)
		return nil, err
	}

	otel.SetTracerProvider(tp)
	return tp, nil
}

func buildResources(ctx context.Context) (*resource.Resource, error) {
	detector := lambda.NewResourceDetector()
	resourcesDetected, err := detector.Detect(ctx)
	if err != nil {
		log.Error("Error detecting resources: %v", err)
		return nil, errors.Wrapf(err, "Error detecting resources")
	}

	resources, err := resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", getServiceName()),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Error("Error creating custom resources: %v", err)
		return nil, errors.Wrapf(err, "Error creating custom resources")
	}

	return resource.Merge(resources, resourcesDetected)
}

func getServiceName() string {
	if serviceName, defined := os.LookupEnv(otelServiceEnv); defined {
		return serviceName
	} else {
		return fmt.Sprintf("%s_UNDEFINED", otelServiceEnv)
	}
}
