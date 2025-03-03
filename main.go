package main

import (
	"fmt"
	"os"

	"github.com/salignatmoandal/flux/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		os.Exit(1)
	}
}
