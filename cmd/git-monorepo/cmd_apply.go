package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/monorepo"
	"github.com/viqueen/git-extensions/pkg/git"
)

func ApplyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "apply <name>",
		Short: "Apply a monorepo view",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			loader := git.NewLoader()
			repo, err := loader.Load()
			if err != nil {
				return fmt.Errorf("failed to load git repository: \n%w", err)
			}

			return monorepo.Apply(repo, args[0])
		},
	}
}
