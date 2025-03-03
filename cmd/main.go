package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/salignatmoandal/flux/config"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
)

func main() {
	// Création de la commande racine
	rootCmd := &cobra.Command{
		Use:   "flux",
		Short: "Flux - Optimiseur de coûts ML dans le cloud",
		Long: `Flux est un outil d'analyse et d'optimisation des coûts 
			   pour les workloads d'apprentissage automatique dans AWS, GCP et Azure.
			   
			   Il permet de :
			   - Détecter les ressources inutilisées
			   - Planifier intelligemment les entraînements ML
			   - Monitorer les coûts en temps réel`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Chargement de la configuration
			cfg, err := config.LoadConfig()
			if err != nil {
				log.Fatalf("Erreur lors du chargement de la configuration: %v", err)
			}

			if verbose {
				fmt.Printf("Configuration chargée: %+v\n", cfg)
			}
		},
	}

	// Flags globaux
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "chemin du fichier de configuration")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "mode verbeux")

	// Ajout des sous-commandes
	rootCmd.AddCommand(
		cmd.NewAnalyzeCmd(),
		cmd.NewMonitorCmd(),
		cmd.NewScheduleCmd(),
		cmd.NewReportCmd(),
	)

	// Exécution de la commande
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		// Utilisation du fichier de configuration spécifié
		os.Setenv("FLUX_CONFIG", cfgFile)
	}
}
