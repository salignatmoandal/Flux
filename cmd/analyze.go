package cmd

import (
	"fmt"

	"github.com/salignatmoandal/flux/internal/analyzer"
	"github.com/spf13/cobra"
)

func NewAnalyzeCmd() *cobra.Command {
	var (
		provider string
		dryRun   bool
	)

	analyzeCmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze unused cloud resources",
		Long: `Detect unused or underutilized cloud resources 
               like GPUs, VMs and Kubernetes clusters.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			a := analyzer.NewResourceAnalyzer()
			if err := a.Initialize(cfg); err != nil {
				return fmt.Errorf("erreur d'initialisation: %w", err)
			}

			results, err := a.AnalyzeResources(provider, dryRun)
			if err != nil {
				return fmt.Errorf("erreur pendant l'analyse: %w", err)
			}

			// Affichage des résultats
			for _, r := range results {
				fmt.Printf("Resource %s: Cost/hour=%.2f€, Usage=%.1f%%\n",
					r.ID, r.CostPerHour, r.UsagePercent)
			}

			return nil
		},
	}

	// Flags spécifiques à la commande analyze
	analyzeCmd.Flags().StringVarP(&provider, "provider", "p", "all",
		"Cloud provider to analyze (aws, gcp, azure, all)")
	analyzeCmd.Flags().BoolVar(&dryRun, "dry-run", false,
		"Run in simulation mode")

	return analyzeCmd
}
