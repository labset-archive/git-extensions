package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/monorepo"
	"github.com/viqueen/git-extensions/pkg/git"
)

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all monorepo views",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			loader := git.NewLoader()
			repo, err := loader.Load()
			if err != nil {
				return fmt.Errorf("failed to load git repository: \n%w", err)
			}

			return monorepo.List(repo)
		},
	}
}
