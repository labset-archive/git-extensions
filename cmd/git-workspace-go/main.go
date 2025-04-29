package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "git-workspace-go",
	}

	rootCmd.AddCommand(InitCmd())
	rootCmd.AddCommand(AddCmd())

	rootCmd.PersistentFlags().StringP("workspace", "w", "", "target workspace")
	rootCmd.PersistentFlags().StringP("owner", "o", "", "target owner")
	rootCmd.PersistentFlags().StringP("name", "n", "", "target name")
	rootCmd.PersistentFlags().StringP("keyword", "k", "", "target keyword")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
