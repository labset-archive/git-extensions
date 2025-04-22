package monorepo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Apply(repo *git.Repo, name string) error {
	if name == "[all]" {
		return applyDefaultView(repo)
	}

	return applyConfiguredView(repo, name)
}

func applyDefaultView(repo *git.Repo) error {
	// disable sparse checkout
	_, err := repo.Executor.Run("sparse-checkout", "disable")
	if err != nil {
		return fmt.Errorf("failed to disable sparse checkout: \n%w", err)
	}
	// remove git config
	_, err = repo.Executor.Run("config", "--local", "--remove-section", "monorepo")
	if err != nil {
		return fmt.Errorf("failed to remove git config: \n%w", err)
	}

	return nil
}

func applyConfiguredView(repo *git.Repo, name string) error {
	viewFile, err := ensureViewFileExists(repo, name)
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

	// set git config
	_, err = repo.Executor.Run("config", "--local", "monorepo.view", name)
	if err != nil {
		return fmt.Errorf("failed to set git config: \n%w", err)
	}

	// set sparse checkout
	_, err = repo.Executor.Run("sparse-checkout", "init", "--cone")
	if err != nil {
		return fmt.Errorf("failed to init sparse checkout: \n%w", err)
	}

	sparseCheckoutSetArgs := []string{"set", ".monorepo"}
	sparseCheckoutSetArgs = append(sparseCheckoutSetArgs, view.Includes...)

	_, err = repo.Executor.Run("sparse-checkout", sparseCheckoutSetArgs...)
	if err != nil {
		return fmt.Errorf("failed to set sparse-checkout: \n%w", err)
	}

	return nil
}
