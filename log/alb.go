package log

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
	"strings"
)

func SetUpALBApiRequest(ctx context.Context, req events.ALBTargetGroupRequest) {
	SetupTraceIds(ctx)
	ReportALBApiRequest(req)
}

func ReportALBApiRequest(req events.ALBTargetGroupRequest) {
	if IsDebugEnabled() {
		DebugW("Got request", buildRequestLogTrackingFields(req)...)
	}
}

func buildRequestLogTrackingFields(request events.ALBTargetGroupRequest) []interface{} {
	return []interface{}{
		zap.String("Body.context.origin.request.method", request.HTTPMethod),
		zap.String("Body.context.origin.request.url", request.Path),
		zap.String("Body.context.origin.request.route", request.Path),
		zap.String("Body.context.origin.request.query", buildAlbQueryParam(request)),
		zap.Array("Body.context.origin.request.headers", buildAlbHeaders(request)),
	}
}

func buildAlbHeaders(request events.ALBTargetGroupRequest) headerItems {
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

func buildAlbQueryParam(request events.ALBTargetGroupRequest) string {
	var params []string
	for key, value := range request.QueryStringParameters {
		params = append(params, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(params, "&")
}
