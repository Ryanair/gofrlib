# Ryanair common testing libraries

## Starting testcontainers utilities
| Function | Description |
| --- | --- |
| RunLocalstackDynamo | Runs localstack with dynamo |
| RunLocalstackSqs | Runs localstack with sqs |
| RunLocalstackServices | Runs localstack with any list of services |
| RunContainer | Runs any container |


## Base integration test suite
Simple utility for creating integration tests with testcontainers-go and localstack.

```go
package handler

import (
	"context"
	"fmt"
	"bitbucket.org/ryanair/gofrlib/test/frContainers"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"testing"
)

var (
	testTableName = "FR-LOCAL-DP-PRICE-TEST"
)

type HandlerIntegrationSuite struct {
	frContainers.BaseLocalstackIntegrationSuite
}

func TestHandlerIntegrationSuite(t *testing.T) {
	suite.Run(t, &HandlerIntegrationSuite{
		BaseLocalstackIntegrationSuite: frContainers.NewBaseLocalstackIntegrationSuite(func() (testcontainers.Container, string) {
			return frContainers.RunLocalstackDynamo(context.Background())
		}),
	})
}

func (suite *HandlerIntegrationSuite) buildDynamoClient() *dynamodb.Client {
	return dynamodb.NewFromConfig(suite.BaseLocalstackIntegrationSuite.GetLocalstackConfig())
}

func (suite *HandlerIntegrationSuite) SetupTest() {
	fmt.Println("Before test")
	_, err := suite.buildDynamoClient().CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName:   &testTableName,
		BillingMode: types.BillingModePayPerRequest,
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("sortKey"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("sortKey"),
				KeyType:       types.KeyTypeRange,
			},
		},
	})
	if err != nil {
		suite.TearDownSuite()
	}
}

func (suite *HandlerIntegrationSuite) TearDownTest() {
	fmt.Println("After test")
	_, err := suite.buildDynamoClient().DeleteTable(context.Background(), &dynamodb.DeleteTableInput{
		TableName: &testTableName,
	})
	if err != nil {
		fmt.Println(err)
		suite.TearDownSuite()
	}
}

func (suite *HandlerIntegrationSuite) Test_TableExists() {
	// WHEN
	client := suite.buildDynamoClient()
	tables, err := client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	suite.NoError(err)

	// THEN
	assert.Equal(suite.T(), 1, len(tables.TableNames))
}
```

## Customizations
localstack images can be overriden using TESTCONTAINERS_HUB_IMAGE_NAME_PREFIX env variable.