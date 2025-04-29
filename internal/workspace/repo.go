package workspace

import "fmt"

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
