package workspace

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
)

func Init() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	handler := NewConfigHandler(wd)

	var workspacesDir string

	var workspacesDefault string

	var githubUsername string

	var githubToken string

	configForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Workspaces Directory").
				Placeholder("Enter the workspaces directory").
				Value(&workspacesDir),
			huh.NewInput().
				Title("Default Workspace").
				Placeholder("Enter the default workspace").
				Value(&workspacesDefault),
			huh.NewInput().
				Title("GitHub Username").
				Placeholder("Enter your GitHub username").
				Value(&githubUsername),
			huh.NewInput().
				Title("GitHub Token").
				Placeholder("Enter your GitHub token").
				EchoMode(huh.EchoModePassword).
				Value(&githubToken),
		),
	)

	err = configForm.Run()
	if err != nil {
		return fmt.Errorf("failed to run config form: %w", err)
	}

	workspacesDir, err = safeDir(workspacesDir)
	if err != nil {
		return fmt.Errorf("failed to resolve workspaces directory: %w", err)
	}

	config := &Config{
		Dir:            workspacesDir,
		Default:        workspacesDefault,
		GithubUsername: githubUsername,
		GithubToken:    githubToken,
	}

	err = handler.Save(config)

	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

func safeDir(dir string) (string, error) {
	if strings.HasPrefix(dir, "~/") || dir == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		return filepath.Join(homeDir, strings.TrimPrefix(dir, "~/")), nil
	}
	return dir, nil
}
