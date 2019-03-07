package log

import (
	"context"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

var log *zap.SugaredLogger

//Creates logger with development config. Logs all from debug level
func init() {
	rawLogger, _ := zap.NewDevelopment()
	defer rawLogger.Sync()
	log = rawLogger.Sugar()
}

//Customizes logger to unify log format with ec2 application loggers
func Init(ctx context.Context) {

	logLevel := zapcore.DebugLevel
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		logLevel.Set(envLogLevel)
	}

	rawLogger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(logLevel),
		OutputPaths: []string{"stdout"},
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "timestamp",
			CallerKey:     "caller",
			MessageKey:    "message",
			LevelKey:      "level",
			StacktraceKey: "stack_trace",
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
		},
	}.Build()

	defer rawLogger.Sync()
	log = rawLogger.Sugar()

	context, _ := lambdacontext.FromContext(ctx)
	if context == nil || context.AwsRequestID == "" {
		log.Errorf("Empty context or missing AwsRequestID. Context: %v", context)
	} else {
		parts := strings.Split(context.InvokedFunctionArn, ":")
		application := parts[len(parts)-1]
		log = log.With("AwsRequestID", context.AwsRequestID).With("application", application)
	}
}

func Debug(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Info(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warn(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Error(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func HandleError(err error) error {
	if err != nil {
		Error("%v", err)
	}
	return err
}

func Metric(key string, duration time.Duration) {
	milliseconds := duration.Nanoseconds() / 1000000
	log.With(key, milliseconds).Debugf("%v took %vms", key, milliseconds)
}
