package main

import (
	"log"

	"github.com/viqueen/git-extensions/internal/monorepo"

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

	applyCmd := &cobra.Command{
		Use:   "apply",
		Short: "Apply a monorepo view",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return monorepo.Apply(args[0])
		},
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new monorepo view",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return monorepo.Create(args[0], args[1:]...)
		},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all monorepo views",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return monorepo.List()
		},
	}

	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)

	return rootCmd
}
