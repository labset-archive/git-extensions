package main

import (
	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/workspace"
)

func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize git workspaces",
		RunE: func(cmd *cobra.Command, args []string) error {
			return workspace.Init()
		},
	}
}
