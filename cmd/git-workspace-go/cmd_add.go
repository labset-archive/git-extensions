package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/workspace"
)

func AddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <urlConnection>",
		Args:  cobra.ExactArgs(1),
		Short: "Add a workspace repo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return withWorkspace(func(config *workspace.Config, ws workspace.Workspace) error {
				wsOption, _ := cmd.Flags().GetString("workspace")
				if wsOption == "" {
					wsOption = config.Default
				}
				urlConnection := args[0]
				repo, err := workspace.RepoFromUrlConnection(urlConnection, wsOption)
				if err != nil {
					return fmt.Errorf("failed to create repo from url connection: %w", err)
				}

				return ws.Add(*repo)
			})
		},
	}

	return cmd
}
