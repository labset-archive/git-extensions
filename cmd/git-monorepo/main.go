package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	cmd := command()
	if err := cmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}

func command() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "git-monorepo",
	}

	rootCmd.AddCommand(ApplyCmd())
	rootCmd.AddCommand(CreateCmd())
	rootCmd.AddCommand(ListCmd())

	return rootCmd
}
