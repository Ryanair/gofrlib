package mapper

import (
	"github.com/aws/aws-lambda-go/events"
	"strings"
)

func ToApiGatewayRequest(request events.ALBTargetGroupRequest) events.APIGatewayProxyRequest {
	return events.APIGatewayProxyRequest{
		HTTPMethod: request.HTTPMethod,
		Path:       request.Path,
		PathParameters: map[string]string{
			"proxy": request.Path,
		},
		QueryStringParameters:           request.QueryStringParameters,
		MultiValueQueryStringParameters: request.MultiValueQueryStringParameters,
		Headers:                         request.Headers,
		MultiValueHeaders:               request.MultiValueHeaders,
		RequestContext:                  mapContext(request),
		Body:                            request.Body,
		IsBase64Encoded:                 request.IsBase64Encoded,
		StageVariables:                  map[string]string{},
	}
}

func ToAlbResponse(response events.APIGatewayProxyResponse) events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		StatusCode:        response.StatusCode,
		Headers:           buildHeaders(response),
		MultiValueHeaders: response.MultiValueHeaders,
		Body:              response.Body,
		IsBase64Encoded:   response.IsBase64Encoded,
	}

}

func mapContext(request events.ALBTargetGroupRequest) events.APIGatewayProxyRequestContext {
	return events.APIGatewayProxyRequestContext{
		HTTPMethod:   request.HTTPMethod,
		ResourcePath: request.Path,
	}
}

func buildHeaders(response events.APIGatewayProxyResponse) map[string]string {
	if response.Headers != nil {
		return response.Headers
	}

	// According to: https://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2
	var headers = make(map[string]string)
	for key, value := range response.MultiValueHeaders {
		headers[key] = strings.Join(value, ",")
	}
	return headers
}
