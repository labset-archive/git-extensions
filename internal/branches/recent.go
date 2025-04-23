package branches

import (
	"fmt"
	"strings"

	"github.com/viqueen/git-extensions/pkg/prompt"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Recent(repo *git.Repo) error {
	branches, err := listRecentBranches(repo)
	if err != nil {
		return fmt.Errorf("failed to list branches: %w", err)
	}

	recentPrompt := prompt.Select{
		Message: "select branch to checkout",
		Options: branches,
	}

	output, err := recentPrompt.Prompt()
	if err != nil {
		return fmt.Errorf("failed to select branch: %w", err)
	}

	parts := strings.Split(output.Selected, " ")
	_, err = repo.Executor.Run("checkout", parts[1])

	if err != nil {
		return fmt.Errorf("failed to checkout branch: %w", err)
	}

	return nil
}
