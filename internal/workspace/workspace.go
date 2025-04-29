package workspace

import (
	"fmt"

	"github.com/viqueen/git-extensions/pkg/git"
	"github.com/viqueen/git-extensions/pkg/registry"
)

type Repo struct {
	UrlConnection string
	Workspace     string
	Owner         string
	Name          string
}

func (r *Repo) Id() string {
	return fmt.Sprintf("%s/%s/%s", r.Workspace, r.Owner, r.Name)
}

type RepoFilter func(repo Repo) bool

type Workspace interface {
	Add(repo Repo) error
	Remove(filter RepoFilter) error
	Clone(filter RepoFilter) error
	List(filter RepoFilter) ([]Repo, error)
}

type workspace struct {
	repos       registry.Registry[Repo]
	execCommand git.ExecCommand
}

func NewWorkspace(repos registry.Registry[Repo], execCommand git.ExecCommand) Workspace {
	return &workspace{
		repos:       repos,
		execCommand: execCommand,
	}
}

func asRegistryFilter(filter RepoFilter) registry.Filter[Repo] {
	return func(item registry.Item[Repo]) bool {
		return filter(item.Value)
	}
}
