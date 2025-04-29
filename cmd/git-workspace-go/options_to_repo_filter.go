package main

import (
	"github.com/spf13/cobra"
	"github.com/viqueen/git-extensions/internal/workspace"
	"strings"
)

func optionsToRepoFilter(cmd *cobra.Command, config *workspace.Config) workspace.RepoFilter {
	wsOption, _ := cmd.Flags().GetString("workspace")
	ownerOption, _ := cmd.Flags().GetString("owner")
	nameOption, _ := cmd.Flags().GetString("name")
	keywordOption, _ := cmd.Flags().GetString("keyword")

	if wsOption == "" {
		wsOption = config.Default
	}

	return func(repo workspace.Repo) bool {
		hasWorkspace := repo.Workspace == wsOption
		hasOwner := ownerOption == "" || repo.Owner == ownerOption
		hasName := nameOption == "" || repo.Name == nameOption
		hasKeyword := keywordOption == "" || strings.Contains(repo.Name, keywordOption) || strings.Contains(repo.Owner, keywordOption)
		return hasWorkspace && hasOwner && hasName && hasKeyword
	}
}
