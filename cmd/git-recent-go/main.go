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
		Use: "git-recent",
		RunE: func(cmd *cobra.Command, args []string) error {
			loader := git.NewLoader()
			repo, err := loader.Load()
			if err != nil {
				return fmt.Errorf("failed to load git repository: \n%w", err)
			}

			return branches.Recent(repo)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}
