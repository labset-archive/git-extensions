package workspace

import "fmt"

func (w *workspace) Remove(filter RepoFilter) error {
	repoItems, err := w.repos.ListItems(asRegistryFilter(filter))
	if err != nil {
		return fmt.Errorf("failed to list repos: %w", err)
	}

	for _, item := range repoItems {
		_, err = w.repos.Delete(item.ID)
		if err != nil {
			return fmt.Errorf("failed to remove repo: %w", err)
		}
	}

	return nil
}
