package dynamodbutils

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func UnmarshalImage[T any](image map[string]events.DynamoDBAttributeValue) (*T, error) {
	attr, err := ToAttributeMap(image)
	if err != nil {
		return nil, err
	}

	var targetObject T
	if err := attributevalue.UnmarshalMap(attr, &targetObject); err != nil {
		return nil, err
	}

	return &targetObject, nil
}
