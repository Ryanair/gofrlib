package frotel

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"go.opentelemetry.io/otel/propagation"
)

const (
	traceparentHeader = "traceparent"
	tracestateHeader  = "tracestate"
)

type TracingContext struct {
	TraceParent string `json:"traceparent"`
	TraceState  string `json:"tracestate"`
}

func RetrieveContext(ctx context.Context) TracingContext {
	carrier := make(map[string]string)
	propagation.TraceContext{}.Inject(ctx, propagation.MapCarrier(carrier))

	return TracingContext{
		TraceParent: carrier[traceparentHeader],
		TraceState:  carrier[tracestateHeader],
	}
}

func (receiver TracingContext) ExtractSpanContext(ctx context.Context) context.Context {
	return receiver.BuildSpanContext(ctx, receiver.TraceParent, receiver.TraceState)
}

func (receiver TracingContext) BuildSpanContext(ctx context.Context, traceParent, traceState string) context.Context {
	return propagation.TraceContext{}.
		Extract(ctx, propagation.MapCarrier(map[string]string{
			traceparentHeader: traceParent,
			tracestateHeader:  traceState,
		}))
}

func (receiver TracingContext) BuildMessageAttributes() map[string]types.MessageAttributeValue {
	attributes := make(map[string]types.MessageAttributeValue)
	if receiver.TraceParent != "" {
		attributes[traceparentHeader] = types.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: &receiver.TraceParent,
		}
	}
	if receiver.TraceState != "" {
		attributes[tracestateHeader] = types.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: &receiver.TraceState,
		}
	}
	return attributes
}

func ExtractContextFromMessageAttributes(ctx context.Context, event events.SQSMessage) context.Context {
	attributes := make(map[string]string)
	if val, ok := event.MessageAttributes[traceparentHeader]; ok {
		attributes[traceparentHeader] = *val.StringValue
	}
	if val, ok := event.MessageAttributes[tracestateHeader]; ok {
		attributes[tracestateHeader] = *val.StringValue
	}

	if len(attributes) == 0 {
		return ctx
	}

	return propagation.TraceContext{}.Extract(ctx, propagation.MapCarrier(attributes))
}
