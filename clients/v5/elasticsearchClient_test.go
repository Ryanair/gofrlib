package clients

import (
	"testing"
)

func TestAwsEsClientImplementsEsClient(t *testing.T) {
	var _ EsClient = &AwsEsClient{}
}

