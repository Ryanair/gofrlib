package dynamodbutils

import (
	"fmt"
	"github.com/Ryanair/gofrlib/log"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ToAttributeValue(value events.DynamoDBAttributeValue) dynamodb.AttributeValue {
	switch value.DataType() {
	case events.DataTypeBinary:
		return dynamodb.AttributeValue{B: value.Binary()}
	case events.DataTypeBinarySet:
		return dynamodb.AttributeValue{BS: value.BinarySet()}
	case events.DataTypeBoolean:
		b := value.Boolean()
		return dynamodb.AttributeValue{BOOL: &b}
	case events.DataTypeNumber:
		n := value.Number()
		return dynamodb.AttributeValue{N: &n}
	case events.DataTypeNumberSet:
		return dynamodb.AttributeValue{NS: value.NumberSet()}
	case events.DataTypeString:
		s := value.String()
		return dynamodb.AttributeValue{S: &s}
	case events.DataTypeStringSet:
		return dynamodb.AttributeValue{NS: value.StringSet()}
	case events.DataTypeNull:
		n := value.IsNull()
		return dynamodb.AttributeValue{NULL: &n}
	case events.DataTypeList:
		l := make([]dynamodb.AttributeValue, 0)
		for _, i := range value.List() {
			l = append(l, ToAttributeValue(i))
		}
		return dynamodb.AttributeValue{L: l}
	case events.DataTypeMap:
		m := make(map[string]dynamodb.AttributeValue, 0)
		for k, i := range value.Map() {
			m[k] = ToAttributeValue(i)
		}
		return dynamodb.AttributeValue{M: m}
	}
	msg := fmt.Sprintf("Couldn't map value %v", value)
	log.Error("%v", msg)
	panic(msg)
}

func ToAttributeMap(eventMap map[string]events.DynamoDBAttributeValue) map[string]dynamodb.AttributeValue {
	m := make(map[string]dynamodb.AttributeValue, 0)
	for k, i := range eventMap {
		m[k] = ToAttributeValue(i)
	}
	return m
}
