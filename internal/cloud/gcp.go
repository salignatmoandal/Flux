package cloud

import (
	"time"

	"github.com/salignatmoandal/flux/config"
)

type GCPClient struct {
	projectID string
	credFile  string
}

func NewGCPClient(cfg *config.Config) *GCPClient {
	return &GCPClient{
		projectID: cfg.Cloud.GCP.ProjectID,
		credFile:  cfg.Cloud.GCP.CredentialFile,
	}
}

func (c *GCPClient) Initialize() error {
	// Initialisation du client GCP
	return nil
}

func (c *GCPClient) GetUnusedResources() ([]Resource, error) {
	// Implémentation GCP
	return nil, nil
}

func (c *GCPClient) GetCosts(start, end time.Time) (float64, error) {
	// Implémentation GCP
	return 0, nil
}
