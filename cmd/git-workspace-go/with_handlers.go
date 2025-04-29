package main

import (
	"fmt"

	"github.com/viqueen/git-extensions/internal/workspace"
	"github.com/viqueen/git-extensions/pkg/git"
	"github.com/viqueen/git-extensions/pkg/registry"
)

type WithConfig func(config *workspace.Config) error

func withConfig(handler WithConfig) error {
	configHandler := workspace.NewConfigHandler("")

	config, err := configHandler.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	return handler(config)
}

type WithRegistry func(registry registry.Registry[workspace.Repo]) error

func withRegistry(handler WithRegistry) error {
	return withConfig(func(config *workspace.Config) error {
		workspaceDir := config.Dir

		repos, err := registry.NewBadgerRegistry[workspace.Repo](workspaceDir)
		if err != nil {
			return fmt.Errorf("failed to create registry: %w", err)
		}

		return handler(repos)
	})
}

type WithWorkspace func(workspace workspace.Workspace) error

func withWorkspace(handler WithWorkspace) error {
	return withRegistry(func(repos registry.Registry[workspace.Repo]) error {
		execCommand := git.NewExecCommand("")
		ws := workspace.NewWorkspace(repos, execCommand)

		return handler(ws)
	})
}
