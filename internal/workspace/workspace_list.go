package workspace

import "fmt"

func (w *workspace) List(filter RepoFilter) ([]Repo, error) {
	repoItems, err := w.repos.ListItems(asRegistryFilter(filter))
	if err != nil {
		return nil, fmt.Errorf("failed to list repos: %w", err)
	}

	repos := make([]Repo, len(repoItems))
	for i, item := range repoItems {
		repos[i] = item.Value
	}

	return repos, nil
}
