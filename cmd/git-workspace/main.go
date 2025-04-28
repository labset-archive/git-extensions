package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "git-workspace",
	}

	rootCmd.AddCommand(InitCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
