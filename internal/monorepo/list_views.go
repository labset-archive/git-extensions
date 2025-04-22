package monorepo

import (
	"fmt"
	"os"
	"strings"

	"github.com/viqueen/git-extensions/pkg/git"
)

func listViews(repo *git.Repo) ([]string, error) {
	viewsDir, err := ensureViewsFolder(repo)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure views folder: \n%w", err)
	}

	entries, err := os.ReadDir(viewsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: \n%w", viewsDir, err)
	}

	views := make([]string, 0, len(entries))
	views = append(views, "[all]")

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		viewName := strings.TrimSuffix(entry.Name(), ".json")
		views = append(views, viewName)
	}

	return views, nil
}
