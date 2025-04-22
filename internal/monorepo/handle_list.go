package monorepo

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/viqueen/git-extensions/pkg/git"
)

func List(repo *git.Repo) error {
	viewsDir, err := ensureViewsFolder(repo)
	if err != nil {
		return fmt.Errorf("failed to ensure views folder: \n%w", err)
	}

	entries, err := os.ReadDir(viewsDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: \n%w", viewsDir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		viewName := strings.TrimSuffix(entry.Name(), ".json")
		log.Printf("[git-monorepo] - %s", viewName)
	}

	return nil
}
