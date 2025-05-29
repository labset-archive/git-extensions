package workspace

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/viqueen/git-extensions/pkg/collections"
	"github.com/viqueen/git-extensions/pkg/registry"

	"github.com/viqueen/git-extensions/pkg/prompt"
)

func (w *workspace) Clone(filter RepoFilter) error {
	items, err := w.repos.ListItems(asRegistryFilter(filter))
	if err != nil {
		return fmt.Errorf("failed to list repos: %w", err)
	}

	repos := collections.Map(items, func(item registry.Item[Repo]) Repo {
		return item.Value
	})

	if len(repos) == 0 {
		log.Printf("[git-workspace] No repositories found matching the filter")

		return nil
	}

	execClone := func(repo Repo) error {
		repoLocation := filepath.Join(w.config.Dir, repo.Workspace, repo.Name)

		_, cloneErr := w.execCommand.Git("clone", repo.UrlConnection, repoLocation)
		if cloneErr != nil {
			return fmt.Errorf(
				"failed to clone %s/%s/%s: %w",
				repo.Workspace,
				repo.Owner,
				repo.Name,
				cloneErr,
			)
		}

		return nil
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
