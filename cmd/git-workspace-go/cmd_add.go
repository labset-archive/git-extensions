package main

import (
	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/workspace"
)

func AddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <urlConnection>",
		Args:  cobra.ExactArgs(1),
		Short: "Add a workspace repo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return withWorkspace(func(ws workspace.Workspace) error {
				return nil
			})
		},
	}

	return cmd
}
