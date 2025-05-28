package workspace

import (
	"errors"
	"fmt"
	"regexp"
)

type Repo struct {
	UrlConnection string
	Host          string
	Workspace     string
	Owner         string
	Name          string
}

func (r *Repo) Id() string {
	return fmt.Sprintf("%s/%s/%s", r.Workspace, r.Owner, r.Name)
}

type DisplayRepo struct {
	Workspace string
	Owner     string
	Name      string
}

func (r *DisplayRepo) Id() string {
	return fmt.Sprintf("%s/%s/%s", r.Workspace, r.Owner, r.Name)
}

type RepoFilter func(repo Repo) bool

var (
	sshURLPattern = regexp.MustCompile(
		`^(ssh://)?([a-zA-Z\d]+)@(?P<host>[a-zA-Z\d.]+(:\d+)?)([/:])(?P<owner>[a-zA-Z\d-_]+)/(?P<name>[a-zA-Z\d-_.]+)\.git$`,
	)
	httpsURLPattern = regexp.MustCompile(
		`^(https://)(?P<host>[a-zA-Z\d.]+)/(?P<owner>[a-zA-Z\d-_]+)/(?P<name>[a-zA-Z\d-_.]+)\.git$`,
	)
)

var (
	ErrInvalidUrlConnection = errors.New("invalid url connection")
)

func RepoFromUrlConnection(urlConnection string, workspace string) (*Repo, error) {
	groups := extractNamedGroups(sshURLPattern, urlConnection)
	if groups == nil {
		groups = extractNamedGroups(httpsURLPattern, urlConnection)
	}

	if groups == nil {
		return nil, ErrInvalidUrlConnection
	}

	repo := &Repo{
		UrlConnection: urlConnection,
		Workspace:     workspace,
		Host:          groups["host"],
		Owner:         groups["owner"],
		Name:          groups["name"],
	}

	return repo, nil
}

func extractNamedGroups(expr *regexp.Regexp, text string) map[string]string {
	match := expr.FindStringSubmatch(text)
	if match == nil {
		return nil
	}

	result := make(map[string]string)

	for index, name := range expr.SubexpNames() {
		if index != 0 && name != "" {
			result[name] = match[index]
		}
	}

	return result
}
