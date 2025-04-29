package branches

import (
	"fmt"
	"log"

	"github.com/viqueen/git-extensions/pkg/prompt"

	"github.com/charmbracelet/huh"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Squashed(repo *git.Repo, target string, forceDelete bool) error {
	branches, err := listSquashedBranches(repo, target)
	if err != nil {
		return fmt.Errorf("failed to list branches: %w", err)
	}

	if len(branches) == 0 {
		return nil
	}

	var selectedBranches []string

	squashedForm := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Select branches to delete").
				Options(prompt.AsOptions(branches)...).
				Value(&selectedBranches),
		),
	)

	err = squashedForm.Run()
	if err != nil {
		return fmt.Errorf("failed to run form: %w", err)
	}

	deleteFlag := "-d"
	if forceDelete {
		deleteFlag = "-D"
	}

	for _, selected := range selectedBranches {
		_, err = repo.Executor.Git("branch", deleteFlag, selected)
		if err != nil {
			log.Printf("failed to delete branch: %s, error: %v", selected, err)
		} else {
			log.Printf("deleted branch: %s", selected)
		}
	}

	return nil
}
