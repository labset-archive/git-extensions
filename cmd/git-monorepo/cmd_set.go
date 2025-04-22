package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/monorepo"
	"github.com/viqueen/git-extensions/pkg/git"
)

func SetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set <name> <paths...>",
		Short: "Set a new monorepo view",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			loader := git.NewLoader()
			repo, err := loader.Load()
			if err != nil {
				return fmt.Errorf("failed to load git repository: \n%w", err)
			}

			return monorepo.Set(repo, args[0], args[1:]...)
		},
	}
}
