package cloud

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
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
	// Implémentation pour AWS
	resources := []Resource{
		{
			ID:           "i-123456",
			Type:         "ec2",
			Provider:     "aws",
			LastUsed:     time.Now().Add(-24 * time.Hour),
			CostPerHour:  0.5,
			UsagePercent: 15.0,
		},
	}
	return resources, nil
}

func (c *AWSClient) GetCosts(start, end time.Time) (float64, error) {
	// Utiliser AWS Cost Explorer API
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(start.Format("2006-01-02")),
			End:   aws.String(end.Format("2006-01-02")),
		},
		Granularity: aws.String("DAILY"),
		Metrics:     []string{"UnblendedCost"},
	}

	result, err := c.client.GetCostAndUsage(context.Background(), input)
	if err != nil {
		return 0, err
	}

	var totalCost float64
	for _, data := range result.ResultsByTime {
		cost, _ := strconv.ParseFloat(*data.Total["UnblendedCost"].Amount, 64)
		totalCost += cost
	}

	return totalCost, nil
}
