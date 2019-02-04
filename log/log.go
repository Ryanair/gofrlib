package log

import (
	"context"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	rawLogger, _ := zap.NewDevelopment()
	defer rawLogger.Sync()
	log = rawLogger.Sugar()
}

//Creates logger with development config. Logs all from debug level
func Init(ctx context.Context) {
	context, _ := lambdacontext.FromContext(ctx)
	if context == nil || context.AwsRequestID == "" {
		log.Errorf("Empty context or missing AwsRequestID. Context: %v", context)
	} else {
		log = log.With("context.AwsRequestID", context.AwsRequestID)
	}
}

func Debug(template string, args ...interface{}) {
	log.Debugf(template, args)
}

func Info(template string, args ...interface{}) {
	log.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warn(template string, args ...interface{}) {
	log.Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Error(template string, args ...interface{}) {
	log.Errorf(template, args)
}
