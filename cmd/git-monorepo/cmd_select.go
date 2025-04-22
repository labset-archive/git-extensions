package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/monorepo"
	"github.com/viqueen/git-extensions/pkg/git"
)

func SelectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "select",
		Short: "Select a view",
		Long:  `Select a view`,
		RunE: func(cmd *cobra.Command, args []string) error {
			loader := git.NewLoader()
			repo, err := loader.Load()
			if err != nil {
				return fmt.Errorf("failed to load git repository: \n%w", err)
			}

			return monorepo.Select(repo)
		},
	}

	return cmd
}
