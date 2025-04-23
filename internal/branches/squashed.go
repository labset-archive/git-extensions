package branches

import (
	"fmt"
	"log"

	"github.com/viqueen/git-extensions/pkg/git"
	"github.com/viqueen/git-extensions/pkg/prompt"
)

func Squashed(repo *git.Repo, target string, forceDelete bool) error {
	branches, err := listSquashedBranches(repo, target)
	if err != nil {
		return fmt.Errorf("failed to list branches: %w", err)
	}

	if len(branches) == 0 {
		return nil
	}

	squashedPrompt := prompt.MultiSelect{
		Message: "select branch to delete",
		Options: branches,
	}

	output, err := squashedPrompt.Prompt()
	if err != nil {
		return fmt.Errorf("failed to select branches: %w", err)
	}

	deleteFlag := "-d"
	if forceDelete {
		deleteFlag = "-D"
	}

	for _, selected := range output.Selected {
		_, err = repo.Executor.Run("branch", deleteFlag, selected)
		if err != nil {
			log.Printf("failed to delete branch: %s, error: %v", selected, err)
		} else {
			log.Printf("deleted branch: %s", selected)
		}
	}

	return nil
}
