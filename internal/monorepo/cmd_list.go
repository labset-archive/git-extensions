package monorepo

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func List() error {
	viewsDir, err := ensureViewsFolder()
	if err != nil {
		return fmt.Errorf("failed to ensure views folder: %w", err)
	}

	entries, err := os.ReadDir(viewsDir)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", viewsDir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		viewName := strings.Trim(entry.Name(), ".json")
		log.Printf("[git-monorepo] - %s", viewName)
	}

	return nil
}
