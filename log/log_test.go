package log

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"testing"
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
	ctx := ExampleContext{}
	Init(ctx)
	log.Info("test")
	log.Error("error test", errors.New("bum"))
}