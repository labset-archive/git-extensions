package workspace

import (
	"fmt"

	"github.com/viqueen/git-extensions/pkg/registry"
)

func (w *workspace) Add(repo Repo) error {
	repoItem := registry.Item[Repo]{ID: repo.Id(), Value: repo}

	_, err := w.repos.CreateOrUpdate(repoItem)
	if err != nil {
		return fmt.Errorf("failed to add repo: %w", err)
	}

	return nil
}
