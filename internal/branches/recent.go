package branches

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Recent(repo *git.Repo) error {
	branches, err := listRecentBranches(repo)
	if err != nil {
		return fmt.Errorf("failed to list branches: %w", err)
	}

	recentPrompt := promptui.Select{
		Label: "select branch to checkout",
		Items: branches,
	}

	_, selectedBranch, err := recentPrompt.Run()
	if err != nil {
		return fmt.Errorf("failed to select branch: %w", err)
	}

	parts := strings.Split(selectedBranch, " ")
	_, err = repo.Executor.Run("checkout", parts[1])

	if err != nil {
		return fmt.Errorf("failed to checkout branch: %w", err)
	}

	return nil
}
