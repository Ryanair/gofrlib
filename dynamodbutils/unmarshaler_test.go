package dynamodbutils_test

import (
	"github.com/Ryanair/gofrlib/dynamodbutils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalImage(t *testing.T) {
	// GIVEN
	type B struct {
		Text   string
		Number int64
	}
	type A struct {
		Text   string
		Number int64
		Nested B
	}
	image := map[string]events.DynamoDBAttributeValue{
		"Text":   events.NewStringAttribute("Text1"),
		"Number": events.NewNumberAttribute("1"),
		"Nested": events.NewMapAttribute(
			map[string]events.DynamoDBAttributeValue{
				"Text":   events.NewStringAttribute("TextNested"),
				"Number": events.NewNumberAttribute("11"),
			},
		),
	}
	// WHEN
	res, err := dynamodbutils.UnmarshalImage[A](image)
	// THEN
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, A{
		Text:   "Text1",
		Number: 1,
		Nested: B{Text: "TextNested", Number: 11},
	}, *res)
}
