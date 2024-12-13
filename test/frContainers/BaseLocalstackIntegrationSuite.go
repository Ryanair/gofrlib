package frContainers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
)

type BaseLocalstackIntegrationSuite struct {
	suite.Suite

	awsEndpoint      string
	container        testcontainers.Container
	provideContainer func() (testcontainers.Container, string)
}

func NewBaseLocalstackIntegrationSuite(provideContainer func() (testcontainers.Container, string)) BaseLocalstackIntegrationSuite {
	return BaseLocalstackIntegrationSuite{provideContainer: provideContainer}
}

func (this *BaseLocalstackIntegrationSuite) GetLocalstackConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background(), func(options *config.LoadOptions) error {
		options.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           this.awsEndpoint,
				SigningRegion: "eu-west-1",
			}, nil
		})

		return nil
	})
	if err != nil {
		panic(err)
	}

	return cfg
}

func (this *BaseLocalstackIntegrationSuite) SetupSuite() {
	container, serviceUrl := this.provideContainer()
	this.container = container
	this.awsEndpoint = serviceUrl
	fmt.Printf("Connecting to AWS using: %v\n", serviceUrl)
}

func (this *BaseLocalstackIntegrationSuite) TearDownSuite() {
	fmt.Println("Cleaning all")
	err := this.container.Terminate(context.Background())
	if err != nil {
		fmt.Printf("could not terminate: %v", err)
	}
}
