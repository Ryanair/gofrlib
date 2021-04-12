package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger
var logConfig Configuration

type Configuration struct {
	logLevel               string
	application            string
	project                string
	projectGroup           string
	customAttributesPrefix string
}

func NewConfiguration(logLevel, application, project, projectGroup, customAttributesPrefix string) Configuration {
	return Configuration{
		logLevel:               logLevel,
		application:            application,
		project:                project,
		projectGroup:           projectGroup,
		customAttributesPrefix: customAttributesPrefix,
	}
}

//Customizes logger to unify log format with ec2 application loggers
func Init(config Configuration) {
	logConfig = config
	var logLevel zap.AtomicLevel
	if err := logLevel.UnmarshalText([]byte(config.logLevel)); err != nil {
		fmt.Printf("malformed log level: %+v\n", config.logLevel)
		logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	rawLogger, _ := zap.Config{
		Level:       logLevel,
		Development: false,
		Encoding:    "json",
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Timestamp",
			LevelKey:       "SeverityText",
			NameKey:        "logger",
			CallerKey:      "Resource.logger",
			MessageKey:     "Body.message",
			StacktraceKey:  "Body.stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		ErrorOutputPaths: []string{"stderr"},
		OutputPaths:      []string{"stderr"},
	}.Build()

	defer rawLogger.Sync()

	log = rawLogger.
		WithOptions(zap.AddCallerSkip(1)).
		With(zap.String("Resource.application", config.application)).
		With(zap.String("Resource.project", config.project)).
		With(zap.String("Resource.projectGroup", config.projectGroup)).
		Sugar()
}

func Flush() error {
	return log.Sync()
}

func Debug(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func DebugW(msg string, keysAndValues ...interface{}) {
	log.Debugw(msg, keysAndValues...)
}

func Info(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func InfoW(msg string, keysAndValues ...interface{}) {
	log.Infow(msg, keysAndValues...)
}

func Warn(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func WarnW(msg string, keysAndValues ...interface{}) {
	log.Warnw(msg, keysAndValues...)
}

func Error(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func ErrorW(msg string, keysAndValues ...interface{}) {
	log.Errorw(msg, keysAndValues...)
}

func With(args ...interface{}) {
	log = log.With(args...)
}

func WithCustomAttr(key string, value interface{}) {
	log = log.With(fmt.Sprintf("Body.%s.%s", logConfig.customAttributesPrefix, key), value)
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
