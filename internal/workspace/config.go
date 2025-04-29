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

const (
	gitConfigDir            = "workspace.dir"
	gitConfigDefault        = "workspace.default"
	gitConfigGithubUsername = "workspace.github.username"
	gitConfigGithubToken    = "workspace.github.token"
)

func (c *configHandler) Load() (*Config, error) {
	workspacesDir, err := c.execCommand.Git("config", "--global", gitConfigDir)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %w", gitConfigDir, err)
	}

	workspacesDefault, err := c.execCommand.Git("config", "--global", gitConfigDefault)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %w", gitConfigDefault, err)
	}

	githubUsername, err := c.execCommand.Git("config", "--global", gitConfigGithubUsername)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %w", gitConfigGithubUsername, err)
	}

	githubToken, err := c.execCommand.Git("config", "--global", gitConfigGithubToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s: %w", gitConfigGithubToken, err)
	}

	return &Config{
		Dir:            workspacesDir,
		Default:        workspacesDefault,
		GithubUsername: githubUsername,
		GithubToken:    githubToken,
	}, nil
}

func (c *configHandler) Save(config *Config) error {
	_, err := c.execCommand.Git("config", "--global", gitConfigDir, config.Dir)
	if err != nil {
		return fmt.Errorf("failed to set %s: %w", gitConfigDir, err)
	}

	_, err = c.execCommand.Git("config", "--global", gitConfigDefault, config.Default)
	if err != nil {
		return fmt.Errorf("failed to set %s: %w", gitConfigDefault, err)
	}

	_, err = c.execCommand.Git(
		"config",
		"--global",
		gitConfigGithubUsername,
		config.GithubUsername,
	)
	if err != nil {
		return fmt.Errorf("failed to set %s: %w", gitConfigGithubUsername, err)
	}

	_, err = c.execCommand.Git("config", "--global", gitConfigGithubToken, config.GithubToken)
	if err != nil {
		return fmt.Errorf("failed to set %s: %w", gitConfigGithubToken, err)
	}

	return nil
}
