package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "git-monorepo",
	}

	rootCmd.AddCommand(ApplyCmd())
	rootCmd.AddCommand(SetCmd())
	rootCmd.AddCommand(ListCmd())
	rootCmd.AddCommand(SelectCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
