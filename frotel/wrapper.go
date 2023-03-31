package frotel

import (
	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
)

func InstrumentHandler(handlerFunc interface{}) interface{} {
	return otellambda.InstrumentHandler(handlerFunc)
}

func Start(handlerFunc interface{}) {
	lambda.Start(InstrumentHandler(handlerFunc))
}
