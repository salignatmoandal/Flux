package cloud

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/salignatmoandal/flux/config"
)

type AWSClient struct {
	client    *costexplorer.Client
	region    string
	accessKey string
	secretKey string
}

func NewAWSClient(cfg *config.Config) *AWSClient {
	return &AWSClient{
		region:    cfg.Cloud.AWS.Region,
		accessKey: cfg.Cloud.AWS.AccessKey,
		secretKey: cfg.Cloud.AWS.SecretKey,
	}
}

func (a *AWSClient) Initialize() error {
	// Initialisation du client AWS
	return nil
}

func (c *AWSClient) GetUnusedResources() ([]Resource, error) {
	// Implémentation AWS
	return nil, nil
}

func (c *AWSClient) GetCosts(start, end time.Time) (float64, error) {
	// Implémentation AWS
	return 0, nil
}
