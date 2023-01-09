package log

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var blackListHeader = map[string]bool{
	"x-auth-token":  true,
	"authorization": true,
	"fr-token-sig":  true,
}

type headerItems []headerItem
type headerItem struct {
	name  string
	value []string
}

func (hi headerItems) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	for _, item := range hi {
		encoder.AppendObject(item)
	}

	return nil
}

func (hi headerItem) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddString("name", hi.name)
	encoder.AddString("value", strings.Join(hi.value, ","))

	return nil
}

func SetUpAPIRequest(ctx context.Context, request events.APIGatewayProxyRequest) {
	SetupTraceIds(ctx)
	ReportAPIRequest(request)
}

func ReportAPIRequest(request events.APIGatewayProxyRequest) {
	if IsDebugEnabled() {
		DebugW("Got request", BuildRequestLogTrackingFields(request)...)
	}
}

func ReportAPIRequestFailure(request events.APIGatewayProxyRequest) {
	WarnW("Got request", BuildRequestLogTrackingFields(request)...)
}

func BuildRequestLogTrackingFields(request events.APIGatewayProxyRequest) []interface{} {
	return []interface{}{
		zap.String("Body.context.origin.request.method", request.HTTPMethod),
		zap.String("Body.context.origin.request.url", request.Path),
		zap.String("Body.context.origin.request.route", request.Path),
		zap.String("Body.context.origin.request.query", buildQueryParam(request)),
		zap.Array("Body.context.origin.request.headers", buildHeaders(request)),
	}
}

func buildHeaders(request events.APIGatewayProxyRequest) headerItems {
	var headerItems []headerItem

	for key, value := range request.MultiValueHeaders {
		if _, exists := blackListHeader[strings.ToLower(key)]; !exists {
			headerItems = append(headerItems, headerItem{name: key, value: value})
		}
	}
	for key, value := range request.Headers {
		if _, exists := blackListHeader[strings.ToLower(key)]; !exists {
			headerItems = append(headerItems, headerItem{name: key, value: []string{value}})
		}
	}

	return headerItems
}

func buildQueryParam(request events.APIGatewayProxyRequest) string {
	var params []string
	for key, value := range request.QueryStringParameters {
		params = append(params, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(params, "&")
}

func WithApigwV2Request(ctx context.Context, event events.APIGatewayV2HTTPRequest) {
	SetupTraceIds(ctx)
	log = log.
		With(EventContextMethod, event.RequestContext.HTTP.Method).
		With(EventContextUrl, event.RequestContext.DomainName).
		With(EventContextRoute, event.RequestContext.RouteKey).
		With(EventContextQuery, event.RawQueryString).
		With(EventContextAgent, event.RequestContext.HTTP.UserAgent)
}
