package monorepo

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/viqueen/git-extensions/pkg/git"
)

func Select(repo *git.Repo) error {
	views, err := listViews(repo)
	if err != nil {
		return fmt.Errorf("failed to list views: \n%w", err)
	}

	viewPrompt := promptui.Select{
		Label: "select view",
		Items: views,
	}

	_, selectedView, err := viewPrompt.Run()
	if err != nil {
		return fmt.Errorf("failed to select view: \n%w", err)
	}

	return Apply(repo, selectedView)
}
