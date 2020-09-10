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
func Init(ctx context.Context, withArgs ...interface{}) {

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
			TimeKey:       "@timestamp",
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
		if app := os.Getenv("APP"); app != "" {
			application = app
		}
		log = log.With("AwsRequestID", context.AwsRequestID).With("application", application)
		if project := os.Getenv("PROJECT"); project != "" {
			log = log.With("project", project)
		}
	}
	log = log.With(withArgs...)
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

func With(args ...interface{}) {
	log.With(args...)
}

func HandleError(err error) error {
	if err != nil {
		Error("%v", err)
	}
	return err
}

func IsDebugEnabled() bool {
	return log.Desugar().Check(zapcore.DebugLevel, "") != nil
}

func IsInfoEnabled() bool {
	return log.Desugar().Check(zapcore.InfoLevel, "") != nil
}

func IsWarnEnabled() bool {
	return log.Desugar().Check(zapcore.WarnLevel, "") != nil
}

func MetricInt(key string, value int) {
	log.With(key, value).Debugf("%v value: %v", key, value)
}

func Metric(key string, duration time.Duration) {
	milliseconds := duration.Nanoseconds() / 1000000
	log.With(key, milliseconds).Debugf("%v took %vms", key, milliseconds)
}
