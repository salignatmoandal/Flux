package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/salignatmoandal/flux/config"
	"github.com/salignatmoandal/flux/internal/cloud"
)

type Monitor struct {
	config   *config.Config
	interval time.Duration
	export   string
	clients  map[string]cloud.Client
	ctx      context.Context
	cancel   context.CancelFunc
}

func New(cfg *config.Config, interval, export string) *Monitor {
	d, _ := time.ParseDuration(interval)
	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:   cfg,
		interval: d,
		export:   export,
		clients:  make(map[string]cloud.Client),
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (m *Monitor) Start() error {
	if err := m.initClients(); err != nil {
		return err
	}

	ticker := time.NewTicker(m.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := m.collectMetrics(); err != nil {
				fmt.Printf("Erreur de collecte: %v\n", err)
			}
		case <-m.ctx.Done():
			return nil
		}
	}
}

func (m *Monitor) Stop() {
	m.cancel()
}

func (m *Monitor) initClients() error {
	// AWS
	awsClient := cloud.NewAWSClient(m.config)
	if err := awsClient.Initialize(); err != nil {
		return fmt.Errorf("error during AWS initialization: %w", err)
	}
	m.clients["aws"] = awsClient

	// GCP
	gcpClient := cloud.NewGCPClient(m.config)
	if err := gcpClient.Initialize(); err != nil {
		return fmt.Errorf("error during GCP initialization: %w", err)
	}
	m.clients["gcp"] = gcpClient

	// Azure
	azureClient := cloud.NewAzureClient(m.config)
	if err := azureClient.Initialize(); err != nil {
		return fmt.Errorf("error during Azure initialization: %w", err)
	}
	m.clients["azure"] = azureClient

	return nil
}

func (m *Monitor) collectMetrics() error {
	now := time.Now()
	start := now.Add(-m.interval)

	for provider, client := range m.clients {
		costs, err := client.GetCosts(start, now)
		if err != nil {
			return fmt.Errorf("error during collection for %s: %w", provider, err)
		}

		// Selon le format d'export choisi
		switch m.export {
		case "grafana":
			// TODO: Exporter vers Grafana
			fmt.Printf("Costs %s: %.2f€\n", provider, costs)
		case "prometheus":
			// TODO: Exporter vers Prometheus
			fmt.Printf("Costs %s: %.2f€\n", provider, costs)
		default:
			fmt.Printf("Costs %s: %.2f€\n", provider, costs)
		}
	}

	return nil
}
