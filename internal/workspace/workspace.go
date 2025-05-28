package workspace

import (
	"github.com/viqueen/git-extensions/pkg/git"
	"github.com/viqueen/git-extensions/pkg/registry"
)

type Workspace interface {
	Add(repo Repo) error
	Remove(filter RepoFilter) error
	Clone(filter RepoFilter) error
	List(filter RepoFilter) ([]Repo, error)
}

type workspace struct {
	config      *Config
	repos       registry.Registry[Repo]
	execCommand git.ExecCommand
}

func NewWorkspace(
	config *Config,
	repos registry.Registry[Repo],
	execCommand git.ExecCommand,
) Workspace {
	return &workspace{
		config:      config,
		repos:       repos,
		execCommand: execCommand,
	}
}

func asRegistryFilter(filter RepoFilter) registry.Filter[Repo] {
	return func(item registry.Item[Repo]) bool {
		return filter(item.Value)
	}
}
