package main

import (
	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/workspace"
)

func CloneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clone",
		Short: "Clone all workspace repos",
		RunE: func(cmd *cobra.Command, args []string) error {
			return withWorkspace(func(config *workspace.Config, ws workspace.Workspace) error {
				filter := optionsToRepoFilter(cmd, config)

				return ws.Clone(filter)
			})
		},
	}

	return cmd
}
