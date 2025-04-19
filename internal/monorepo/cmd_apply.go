package monorepo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Apply(name string) error {
	viewFile, err := ensureViewFileExists(name)
	if err != nil {
		return fmt.Errorf("failed to ensure view file exists: \n%w", err)
	}

	defer func(viewFile *os.File) {
		err = viewFile.Close()
		if err != nil {
			log.Printf("[git-monorepo] failed to close view file: %v", err)
		}
	}(viewFile)

	log.Printf("[git-monorepo] applying view %s", name)

	var view View

	decoder := json.NewDecoder(viewFile)
	if err = decoder.Decode(&view); err != nil {
		return fmt.Errorf("failed to decode view file: \n%w", err)
	}

	if err = runCmd("git", "config", "--local", "monorepo.view", name); err != nil {
		return fmt.Errorf("failed to set git config: \n%w", err)
	}

	if err = runCmd("git", "sparse-checkout", "init", "--cone"); err != nil {
		return fmt.Errorf("failed to init sparse-checkout: \n%w", err)
	}

	sparseCheckoutSetArgs := []string{"sparse-checkout", "set"}
	sparseCheckoutSetArgs = append(sparseCheckoutSetArgs, view.Includes...)
	sparseCheckoutSetArgs = append(sparseCheckoutSetArgs, ".monorepo")

	if err = runCmd("git", sparseCheckoutSetArgs...); err != nil {
		return fmt.Errorf("failed to set sparse-checkout: \n%w", err)
	}

	return nil
}
