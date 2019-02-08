package log

import (
	"context"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"testing"
)

type ExampleContext struct {
	context.Context
}

func (lc ExampleContext) Value(key interface{}) interface{} {
	return &lambdacontext.LambdaContext{
		AwsRequestID: "123",
	}
}

func TestInit(t *testing.T) {
	ctx := ExampleContext{}
	Init(ctx)
	log.Info("test")

}
