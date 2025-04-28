package workspace

import (
	"fmt"

	"github.com/viqueen/git-extensions/pkg/git"
)

type Config struct {
	Dir            string
	Default        string
	GithubUsername string
	GithubToken    string
}

type ConfigHandler interface {
	Load() (*Config, error)
	Save(config *Config) error
}

type configHandler struct {
	execCommand git.ExecCommand
}

func NewConfigHandler(wd string) ConfigHandler {
	return &configHandler{
		execCommand: git.NewExecCommand(wd),
	}
}

func (c *configHandler) Load() (*Config, error) {
	workspacesDir, err := c.execCommand.Run("config", "--global", "workspaces.dir")
	if err != nil {
		return nil, fmt.Errorf("failed to get workspaces.dir: %w", err)
	}

	workspacesDefault, err := c.execCommand.Run("config", "--global", "workspaces.default")
	if err != nil {
		return nil, fmt.Errorf("failed to get workspaces.default: %w", err)
	}

	githubUsername, err := c.execCommand.Run("config", "--global", "workspaces.github.username")
	if err != nil {
		return nil, fmt.Errorf("failed to get workspaces.github.username: %w", err)
	}

	githubToken, err := c.execCommand.Run("config", "--global", "workspaces.github.token")
	if err != nil {
		return nil, fmt.Errorf("failed to get workspaces.github.token: %w", err)
	}

	return &Config{
		Dir:            workspacesDir,
		Default:        workspacesDefault,
		GithubUsername: githubUsername,
		GithubToken:    githubToken,
	}, nil
}

func (c *configHandler) Save(config *Config) error {
	_, err := c.execCommand.Run("config", "--global", "workspaces.dir", config.Dir)
	if err != nil {
		return fmt.Errorf("failed to set workspaces.dir: %w", err)
	}

	_, err = c.execCommand.Run("config", "--global", "workspaces.default", config.Default)
	if err != nil {
		return fmt.Errorf("failed to set workspaces.default: %w", err)
	}

	_, err = c.execCommand.Run(
		"config",
		"--global",
		"workspaces.github.username",
		config.GithubUsername,
	)
	if err != nil {
		return fmt.Errorf("failed to set workspaces.github.username: %w", err)
	}

	_, err = c.execCommand.Run("config", "--global", "workspaces.github.token", config.GithubToken)
	if err != nil {
		return fmt.Errorf("failed to set workspaces.github.token: %w", err)
	}

	return nil
}
