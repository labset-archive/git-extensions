package branches

import (
	"fmt"
	"strings"

	"github.com/viqueen/git-extensions/pkg/prompt"

	"github.com/charmbracelet/huh"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Recent(repo *git.Repo) error {
	branches, err := listRecentBranches(repo)
	if err != nil {
		return fmt.Errorf("failed to list branches: %w", err)
	}

	var selectedBranch string

	recentForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a branch to checkout").
				Options(prompt.AsOptions(branches)...).
				Value(&selectedBranch),
		),
	)

	err = recentForm.Run()
	if err != nil {
		return fmt.Errorf("failed to run form: %w", err)
	}

	parts := strings.Split(selectedBranch, " ")
	_, err = repo.Executor.Git("checkout", parts[1])

	if err != nil {
		return fmt.Errorf("failed to checkout branch: %w", err)
	}

	return nil
}
