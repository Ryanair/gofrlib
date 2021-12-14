package dynamodbutils

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func ToAttributeValue(value events.DynamoDBAttributeValue) (types.AttributeValue, error) {
	switch value.DataType() {
	case events.DataTypeNull:
		return &types.AttributeValueMemberNULL{Value: value.IsNull()}, nil

	case events.DataTypeBoolean:
		return &types.AttributeValueMemberBOOL{Value: value.Boolean()}, nil

	case events.DataTypeBinary:
		return &types.AttributeValueMemberB{Value: value.Binary()}, nil

	case events.DataTypeBinarySet:
		bs := make([][]byte, len(value.BinarySet()))
		for i := 0; i < len(value.BinarySet()); i++ {
			bs[i] = append([]byte{}, value.BinarySet()[i]...)
		}
		return &types.AttributeValueMemberBS{Value: bs}, nil

	case events.DataTypeNumber:
		return &types.AttributeValueMemberN{Value: value.Number()}, nil

	case events.DataTypeNumberSet:
		return &types.AttributeValueMemberNS{Value: append([]string{}, value.NumberSet()...)}, nil

	case events.DataTypeString:
		return &types.AttributeValueMemberS{Value: value.String()}, nil

	case events.DataTypeStringSet:
		return &types.AttributeValueMemberSS{Value: append([]string{}, value.StringSet()...)}, nil

	case events.DataTypeList:
		values, err := ToDynamoList(value.List())
		if err != nil {
			return nil, err
		}
		return &types.AttributeValueMemberL{Value: values}, nil

	case events.DataTypeMap:
		values, err := ToAttributeMap(value.Map())
		if err != nil {
			return nil, err
		}
		return &types.AttributeValueMemberM{Value: values}, nil

	default:
		return nil, fmt.Errorf("unknown AttributeValue union member, %T", value)
	}
}

func ToDynamoList(from []events.DynamoDBAttributeValue) (to []types.AttributeValue, err error) {
	to = make([]types.AttributeValue, len(from))
	for i := 0; i < len(from); i++ {
		to[i], err = ToAttributeValue(from[i])
		if err != nil {
			return nil, err
		}
	}

	return to, nil
}

func ToAttributeMap(eventMap map[string]events.DynamoDBAttributeValue) (to map[string]types.AttributeValue, err error) {
	to = make(map[string]types.AttributeValue, len(eventMap))
	for field, value := range eventMap {
		to[field], err = ToAttributeValue(value)
		if err != nil {
			return nil, err
		}
	}

	return to, nil
}
