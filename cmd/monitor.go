package cmd

import (
	"github.com/salignatmoandal/flux/internal/monitor"
	"github.com/spf13/cobra"
)

func NewMonitorCmd() *cobra.Command {
	var (
		interval string
		export   string
	)

	monitorCmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitorer les coûts en temps réel",
		RunE: func(cmd *cobra.Command, args []string) error {
			m := monitor.New(interval, export)
			return m.Start()
		},
	}

	monitorCmd.Flags().StringVarP(&interval, "interval", "i", "5m", "Intervalle de monitoring")
	monitorCmd.Flags().StringVarP(&export, "export", "e", "grafana", "Format d'export (grafana/prometheus)")

	return monitorCmd
}
