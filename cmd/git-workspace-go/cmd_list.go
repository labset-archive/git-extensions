package main

import (
	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/workspace"
)

func ListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all workspace repos",
		RunE: func(cmd *cobra.Command, args []string) error {
			return withWorkspace(func(config *workspace.Config, ws workspace.Workspace) error {
				filter := optionsToRepoFilter(cmd, config)

				return workspace.List(ws, filter)
			})
		},
	}

	return cmd
}
