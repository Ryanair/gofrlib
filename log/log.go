package log

import (
	"context"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

//Creates logger with development config. Logs all from debug level
func init() {
	rawLogger, _ := zap.NewDevelopment()
	defer rawLogger.Sync()
	log = rawLogger.Sugar()
}

//Switch into production mode (JSON format) and appends AwsRequestID to all log messages
func Init(ctx context.Context) {
	rawLogger, _ := zap.NewProduction()
	defer rawLogger.Sync()
	log = rawLogger.Sugar()

	context, _ := lambdacontext.FromContext(ctx)
	if context == nil || context.AwsRequestID == "" {
		log.Errorf("Empty context or missing AwsRequestID. Context: %v", context)
	} else {
		log = log.With("context.AwsRequestID", context.AwsRequestID)
	}
}

func Info(template string, args ...interface{}) {
	log.Infof(template, args)
}

func Warn(template string, args ...interface{}) {
	log.Warnf(template, args)
}

func Error(template string, args ...interface{}) {
	log.Errorf(template, args)
}
