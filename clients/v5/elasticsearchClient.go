package clients

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/edoardo849/apex-aws-signer"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"net/http"
	"time"
)

// EsConfig represents data needed by Elasticsearch driver to create Client.
type EsConfig struct {
	Address string
	Scheme  string
	Region  string
}

// EsClient interface defines simplified Elasticsearch client.
type EsClient interface {
	Search(index string, query elastic.Query, timeoutSec int) (*elastic.SearchResult, error)
	ExecuteBulk(requests []elastic.BulkableRequest, timeoutSec int) (*elastic.BulkResponse, error)
	DeleteByQuery(index string, query elastic.Query, timeoutSec int, shardWait string, waitForComplete bool) (*elastic.BulkIndexByScrollResponse, error)
	Close()
}

// AwsEsClient wrapper for Elasticsearch client. Created for our convenience.
type AwsEsClient struct {
	config *EsConfig
	client *elastic.Client
}

func (c *AwsEsClient) Search(index string, query elastic.Query, timeoutSec int) (*elastic.SearchResult, error) {
	service := elastic.NewSearchService(c.client).Query(query).Index(index)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	return service.Do(ctx)
}

// ExecuteBulk executes BulkableRequest.
func (c *AwsEsClient) ExecuteBulk(requests []elastic.BulkableRequest, timeoutSec int) (*elastic.BulkResponse, error) {
	service := elastic.NewBulkService(c.client).Add(requests...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	return service.Do(ctx)
}

// DeleteByQuery deletes all documents matching given query.
func (c *AwsEsClient) DeleteByQuery(index string, query elastic.Query, timeoutSec int, shardWait string, waitForComplete bool) (*elastic.BulkIndexByScrollResponse, error) {
	service := elastic.NewDeleteByQueryService(c.client).Query(query).Index(index)

	service.WaitForCompletion(waitForComplete)
	service.Timeout(shardWait)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	return service.Do(ctx)
}

// DeleteByQueryAndRouting deletes all documents matching given query for given routingKey.
func (c *AwsEsClient) DeleteByQueryAndRouting(index string, query elastic.Query, routingKey string, timeoutSec int) (*elastic.BulkIndexByScrollResponse, error) {
	service := elastic.NewDeleteByQueryService(c.client).
		Index(index).
		Routing(routingKey).
		Query(query)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	return service.Do(ctx)
}

// Close closes Elasticsearch client which is used internally by EsClient.
func (c *AwsEsClient) Close() {
	c.client.Stop()
}

// New is a factory method which creates new Elasticsearch client instance.
func New(config *EsConfig) (*AwsEsClient, error) {

	session, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
	})

	if err != nil {
		log.Printf("Error while creating new AWS session. Err: %v", err)
		return nil, err
	}

	httpClient := &http.Client{
		Transport: signer.NewTransport(session, elasticsearchservice.ServiceName),
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.Address),
		elastic.SetScheme(config.Scheme),
		elastic.SetHttpClient(httpClient),
	)

	if err != nil {
		log.Printf("Error while creating Elasticsearch client. Err: %v", err)
		return nil, err
	}

	return &AwsEsClient{
		config: config,
		client: client,
	}, nil
}
