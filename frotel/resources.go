package frotel

import (
	"context"
	"fmt"
	"github.com/Ryanair/gofrlib/log"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/detectors/aws/lambda"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv/v1.24.0"
	"os"
)

const (
	otelServiceVersion = "VERSION"
)

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
		log.Error("Error creating custom resources", err)
		return nil, errors.Wrapf(err, "Error creating custom resources")
	}

	return resources, nil
}

func buildMetricResources(ctx context.Context) (*resource.Resource, error) {
	resources, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithAttributes(
			semconv.ServiceVersionKey.String(getServiceVersion()),
			semconv.TelemetrySDKLanguageGo,
		),
	)
	if err != nil {
		log.Error("Error creating custom resources", err)
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
