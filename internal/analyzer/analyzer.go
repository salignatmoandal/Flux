package analyzer

import (
	"fmt"

	"github.com/salignatmoandal/flux/config"
	"github.com/salignatmoandal/flux/internal/cloud"
)

type Analyzer struct {
	config   *config.Config
	provider string
	clients  map[string]cloud.Client
}

func NewResourceAnalyzer() *Analyzer {
	return &Analyzer{
		clients: make(map[string]cloud.Client),
	}
}

func (a *Analyzer) Initialize(cfg *config.Config) error {
	a.config = cfg
	return nil
}

func (a *Analyzer) initClients() error {
	switch a.provider {
	case "aws":
		client := cloud.NewAWSClient(a.config)
		if err := client.Initialize(); err != nil {
			return fmt.Errorf("erreur AWS: %w", err)
		}
		a.clients["aws"] = client
	case "gcp":
		client := cloud.NewGCPClient(a.config)
		if err := client.Initialize(); err != nil {
			return fmt.Errorf("erreur GCP: %w", err)
		}
		a.clients["gcp"] = client
	case "azure":
		client := cloud.NewAzureClient(a.config)
		if err := client.Initialize(); err != nil {
			return fmt.Errorf("erreur Azure: %w", err)
		}
		a.clients["azure"] = client
	default:
		return fmt.Errorf("provider non supporté: %s", a.provider)
	}
	return nil
}

func (a *Analyzer) AnalyzeResources(provider string, dryRun bool) ([]cloud.Resource, error) {
	a.provider = provider

	if err := a.initClients(); err != nil {
		return nil, err
	}

	var resources []cloud.Resource
	for _, client := range a.clients {
		unused, err := client.GetUnusedResources()
		if err != nil {
			continue
		}
		resources = append(resources, unused...)
	}

	return resources, nil
}
