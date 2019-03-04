package log

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"os"
	"testing"
	"time"
)

type ExampleContext struct {
	context.Context
}

func (lc ExampleContext) Value(key interface{}) interface{} {
	return &lambdacontext.LambdaContext{
		AwsRequestID: "309a4277-c267-4a1d-abf5-1eaa7f2bbacf",
		InvokedFunctionArn: "arn:aws:lambda:eu-west-1:487943794540:function:FR-SANDBOX-OTA-APPS-HEALTH-OtaAppsHealthMailLambda-1DXP7KVTF5JG4",
	}
}

//doesn't assert anything because we have no method output, it's only to check if log format is valid
func TestInit(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	ctx := ExampleContext{}
	Init(ctx)
	Debug("debug level %v", 4444)
	Metric("metric", time.Second)
	Info("%v test %v", 123, 456)
	Error("error test %v", errors.New("bum"))
}