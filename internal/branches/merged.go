package branches

import (
	"fmt"
	"log"

	"github.com/viqueen/git-extensions/pkg/prompt"

	"github.com/charmbracelet/huh"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Merged(repo *git.Repo, target string) error {
	branches, err := listMergedBranches(repo, target)
	if err != nil {
		return fmt.Errorf("failed to list branches: %w", err)
	}

	if len(branches) == 0 {
		return nil
	}

	var selectedBranches []string

	mergedForm := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select branches to delete").
				Options(prompt.AsOptions(branches)...).
				Value(&selectedBranches)),
	)

	err = mergedForm.Run()
	if err != nil {
		return fmt.Errorf("failed to run form: %w", err)
	}

	for _, selected := range selectedBranches {
		_, err = repo.Executor.Run("branch", "-d", selected)
		if err != nil {
			log.Printf("failed to delete branch: %s, error: %v", selected, err)
		} else {
			log.Printf("deleted branch: %s", selected)
		}
	}

	return nil
}
