package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/branches"
	"github.com/viqueen/git-extensions/pkg/git"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "git-merged-go",
		RunE: func(cmd *cobra.Command, args []string) error {
			loader := git.NewLoader()
			repo, err := loader.Load()
			if err != nil {
				return fmt.Errorf("failed to load git repository: \n%w", err)
			}

			target, err := cmd.Flags().GetString("target")
			if err != nil {
				return fmt.Errorf("failed to get target flag: %w", err)
			}

			return branches.Merged(repo, target)
		},
	}

	rootCmd.PersistentFlags().StringP(
		"target", "t",
		"origin/main",
		"target branch to check for squashed branches",
	)

	rootCmd.PersistentFlags().BoolP(
		"force", "f",
		false,
		"force delete branches",
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
