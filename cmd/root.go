package cmd

import (
	"fmt"

	"github.com/salignatmoandal/flux/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
	cfg     *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "flux",
	Short: "Flux - Optimiseur de coûts ML dans le cloud",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.LoadConfig()
		if err != nil {
			return fmt.Errorf("erreur de configuration: %w", err)
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "fichier de configuration")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "mode verbeux")

	rootCmd.AddCommand(
		NewAnalyzeCmd(),
		NewMonitorCmd(),
	)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
}
