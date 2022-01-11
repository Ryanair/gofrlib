package log_test

import (
	"context"
	"github.com/Ryanair/gofrlib/log"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/stretchr/testify/assert"
	"testing"
)

//doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestInit(t *testing.T) {
	config := log.NewConfiguration(
		"DEBUG",
		"TEST-APPLICATION",
		"TEST-PROJECT",
		"TEST-PROJECT-GROUP",
		"1.0.0",
		"testPrefix")
	log.Init(config)
	log.Debug("Debug msg: %v", "test-message")
	log.DebugW("DebugW msg with attribute string", "test-key-1", "test-value-1")
	log.DebugW("DebugW msg with attribute bool", "test-key-2", false)
	log.DebugW("DebugW msg with attribute int", "test-key-3", 123)

	log.With("test-key-4", "test-value-4")
	log.WithCustomAttr("CustomAttrKey1", "CustomAttr1Value")
	log.WithCustomAttr("CustomAttrKey2", true)
	log.WithCustomAttr("CustomAttrKey3", 123456)
	log.Info("Info msg with custom attributes")
}

//doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestInitShouldClearExistingContext(t *testing.T) {
	config := log.NewConfiguration(
		"DEBUG",
		"TEST-APPLICATION",
		"TEST-PROJECT",
		"TEST-PROJECT-GROUP",
		"1.0.0",
		"testPrefix")
	log.Init(config)
	log.With("test-key-1", "test-value-1")
	log.Debug("Debug msg with value in context")
	log.Init(config)
	log.Debug("Debug msg without value in context")
}

//doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestSkipLowerLogLevel(t *testing.T) {
	config := log.NewConfiguration(
		"INFO",
		"TEST-APPLICATION",
		"TEST-PROJECT",
		"TEST-PROJECT-GROUP",
		"1.0.0",
		"testPrefix")
	log.Init(config)
	log.Debug("Debug msg")
	log.Info("Info msg")
	log.Warn("Warn msg")
	log.Error("Error msg")
}

func TestLogLevelCheck(t *testing.T) {
	config := log.NewConfiguration(
		"WARN",
		"TEST-APPLICATION",
		"TEST-PROJECT",
		"TEST-PROJECT-GROUP",
		"1.0.0",
		"testPrefix")
	log.Init(config)
	assert.False(t, log.IsDebugEnabled())
	assert.False(t, log.IsInfoEnabled())
	assert.True(t, log.IsWarnEnabled())
}

//doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestLogEmptyVersion(t *testing.T) {
	config := log.NewConfiguration(
		"DEBUG",
		"TEST-APPLICATION",
		"TEST-PROJECT",
		"TEST-PROJECT-GROUP",
		"",
		"testPrefix")
	log.Init(config)
	log.Debug("Debug msg with value in context")
}

//doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestLogTraceIds(t *testing.T) {
	config := log.NewConfiguration(
		"DEBUG",
		"TEST-APPLICATION",
		"TEST-PROJECT",
		"TEST-PROJECT-GROUP",
		"",
		"testPrefix")
	log.Init(config)
	ctx := context.Background()
	ctx = context.WithValue(ctx, xray.LambdaTraceHeaderKey, "Sampled=1;Root=TraceIdValue;Parent=ParentIdValue")
	log.SetupTraceIds(ctx)
	log.Debug("Debug msg with value in context")
}
