package cloud

import (
	"time"

	"github.com/salignatmoandal/flux/config"
)

type AzureClient struct {
	subscriptionID string
	tenantID       string
}

func NewAzureClient(cfg *config.Config) *AzureClient {
	return &AzureClient{
		subscriptionID: cfg.Cloud.Azure.SubscriptionID,
		tenantID:       cfg.Cloud.Azure.TenantID,
	}
}

func (c *AzureClient) Initialize() error {
	// Initialisation du client Azure
	return nil
}

func (c *AzureClient) GetUnusedResources() ([]Resource, error) {
	// Implémentation Azure
	return nil, nil
}

func (c *AzureClient) GetCosts(start, end time.Time) (float64, error) {
	// Implémentation Azure
	return 0, nil
}
