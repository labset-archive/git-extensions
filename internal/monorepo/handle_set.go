package monorepo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Set(repo *git.Repo, name string, includes ...string) error {
	log.Printf("[git-monorepo] setting view %s", name)

	viewFile, err := ensureViewFileCreated(repo, name)
	if err != nil {
		return fmt.Errorf("failed to load view file: \n%w", err)
	}
	defer func(viewFile *os.File) {
		err = viewFile.Close()
		if err != nil {
			log.Fatalf("[git-monorepo] failed to close view file: %v", err)
		}
	}(viewFile)

	sort.Strings(includes)

	view := View{
		Name:     name,
		Includes: includes,
	}

	encoder := json.NewEncoder(viewFile)
	encoder.SetIndent("", "  ")

	if err = encoder.Encode(view); err != nil {
		return fmt.Errorf("failed to encode view to JSON: \n%w", err)
	}

	return nil
}
