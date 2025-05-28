package workspace

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/viqueen/git-extensions/pkg/prompt"
)

func (w *workspace) Clone(filter RepoFilter) error {
	repos, err := w.List(filter)
	if err != nil {
		return fmt.Errorf("failed to list repos: %w", err)
	}

	if len(repos) == 0 {
		return errors.New("no repositories to clone")
	}

	execClone := func(repo Repo) error {
		repoLocation := filepath.Join(w.config.Dir, repo.Workspace, repo.Name)
		_, cloneErr := w.execCommand.Git("clone", repo.UrlConnection, repoLocation)

		return cloneErr
	}

	labelFunc := func(repo Repo) string {
		return fmt.Sprintf("Cloning %s/%s/%s", repo.Workspace, repo.Owner, repo.Name)
	}

	err = prompt.RunProgress(repos, execClone, labelFunc)
	if err != nil {
		return fmt.Errorf("failed to run clone progress: %w", err)
	}

	return nil
}
