package monorepo

import (
	"fmt"
	"log"

	"github.com/viqueen/git-extensions/pkg/git"
)

func List(repo *git.Repo) error {
	views, err := listViews(repo)
	if err != nil {
		return fmt.Errorf("failed to list views: \n%w", err)
	}

	// get current view
	// ignore error as it is not critical
	// see document for git config https://git-scm.com/docs/git-config
	currentView, _ := repo.Executor.Git("config", "get", "--local", "monorepo.view")

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
