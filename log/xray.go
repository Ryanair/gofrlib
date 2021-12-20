package log

import (
	"context"
	"fmt"
	"github.com/aws/aws-xray-sdk-go/header"
	"github.com/aws/aws-xray-sdk-go/strategy/ctxmissing"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
)

type xRayLogger struct {
}

func (x *xRayLogger) Log(level xraylog.LogLevel, msg fmt.Stringer) {
	switch level {
	case xraylog.LogLevelWarn:
		log.Warn(msg.String())
	case xraylog.LogLevelError:
		log.Error(msg.String())
	}
}

func setUpXRay() {
	if err := xray.Configure(xray.Config{ContextMissingStrategy: &ctxmissing.DefaultIgnoreErrorStrategy{}}); err != nil {
		log.Error("unable to configure xray: %+v", err)
	}
	setupXRayLogger()
}

func setupXRayLogger() {
	xray.SetLogger(&xRayLogger{})
}

func getTraceHeaderFromContext(ctx context.Context) *header.Header {
	var traceHeader string

	if traceHeaderValue := ctx.Value(xray.LambdaTraceHeaderKey); traceHeaderValue != nil {
		traceHeader = traceHeaderValue.(string)
		return header.FromString(traceHeader)
	}
	return nil
}
