package workspace

import (
	"errors"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/viqueen/git-extensions/pkg/collections"
	"github.com/viqueen/git-extensions/pkg/prompt"
	"github.com/viqueen/git-extensions/pkg/registry"
)

func (w *workspace) List(filter RepoFilter) error {
	items, err := w.repos.ListItems(asRegistryFilter(filter))
	if err != nil {
		return fmt.Errorf("failed to list repos: %w", err)
	}

	repos := collections.Map(items, func(item registry.Item[Repo]) DisplayRepo {
		return DisplayRepo{
			Name:      item.Value.Name,
			Owner:     item.Value.Owner,
			Workspace: item.Value.Workspace,
		}
	})

	tbl, err := prompt.NewTableComponent[DisplayRepo](repos)
	if errors.Is(err, prompt.ErrEmptyItems) {
		log.Printf("[git-workspace] No repositories found matching the filter")

		return nil
	} else if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	program := tea.NewProgram(tbl)

	_, err = program.Run()

	if err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}

	return nil
}
