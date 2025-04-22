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

	views := make([]string, 0, len(entries))
	views = append(views, "default")

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		viewName := strings.TrimSuffix(entry.Name(), ".json")
		views = append(views, viewName)
	}

	// get current view
	// ignore error as it is not critical
	// see document for git config https://git-scm.com/docs/git-config
	currentView, _ := repo.Executor.Run("config", "get", "--local", "monorepo.view")

	if currentView == "" {
		currentView = "default"
	}

	for _, view := range views {
		if view == currentView {
			log.Printf("* \033[32m%s\033[0m\n", view)
		} else {
			log.Printf("  %s\n", view)
		}
	}

	return nil
}
