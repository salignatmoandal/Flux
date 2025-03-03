package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flux",
	Short: "Flux - Optimize your ML costs in the cloud",
	Long: `Flux is an optimization tool for ML workloads 
           in AWS, GCP and Azure.`,
}

func Execute() error {
	return rootCmd.Execute()
}
