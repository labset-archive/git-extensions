package monorepo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Create(name string, includes ...string) error {
	log.Printf("[git-monorepo] creating view %s", name)

	viewFile, err := ensureViewFile(name)
	if err != nil {
		return fmt.Errorf("failed to create view file: %w", err)
	}
	defer func(viewFile *os.File) {
		err = viewFile.Close()
		if err != nil {
			log.Fatalf("[git-monorepo] failed to close view file: %v", err)
		}
	}(viewFile)

	view := View{
		Name:     name,
		Includes: includes,
	}

	encoder := json.NewEncoder(viewFile)
	encoder.SetIndent("", "  ")

	if err = encoder.Encode(view); err != nil {
		return fmt.Errorf("failed to encode view to JSON: %w", err)
	}

	log.Printf("[git-monorepo] created view %s", viewFile.Name())

	return nil
}
