package monorepo

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/viqueen/git-extensions/pkg/prompt"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Select(repo *git.Repo) error {
	views, err := listViews(repo)
	if err != nil {
		return fmt.Errorf("failed to list views: \n%w", err)
	}

	if len(views) == 0 {
		return nil
	}

	var selectedView string

	selectForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select view").
				Options(prompt.AsOptions(views)...).
				Value(&selectedView),
		),
	)

	err = selectForm.Run()
	if err != nil {
		return fmt.Errorf("failed to run form: %w", err)
	}

	return Apply(repo, selectedView)
}
