package log

import (
	"context"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
	"strings"
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


	//rawLogger, _ := zap.Config{
	//
	//
	//	//Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
	//	//Development: false,
	//	//Sampling: &zap.SamplingConfig{
	//	//	Initial:    100,
	//	//	Thereafter: 100,
	//	//},
	//	//Encoding:         "json",
	//	//EncoderConfig:    NewProductionEncoderConfig(),
	//	//OutputPaths:      []string{"stderr"},
	//	//ErrorOutputPaths: []string{"stderr"},
	//	Encoding:    "json",
	//	Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
	//	OutputPaths: []string{"stdout"},
	//	EncoderConfig: zapcore.EncoderConfig{
	//		MessageKey: "message",  // <--
	//	},
	//}.Build()


	defer rawLogger.Sync()
	log = rawLogger.Sugar()

	context, _ := lambdacontext.FromContext(ctx)
	if context == nil || context.AwsRequestID == "" {
		log.Errorf("Empty context or missing AwsRequestID. Context: %v", context)
	} else {
		parts := strings.Split(context.InvokedFunctionArn, ":")
		application := parts[len(parts) - 1]
		log = log.With("context.AwsRequestID", context.AwsRequestID).With("application", application)
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
