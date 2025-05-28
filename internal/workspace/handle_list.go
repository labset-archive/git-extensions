package workspace

import (
	"errors"
	"fmt"
	"log"

	"github.com/viqueen/git-extensions/pkg/collections"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/viqueen/git-extensions/pkg/prompt"
)

func List(ws Workspace, filer RepoFilter) error {
	items, err := ws.List(filer)
	if err != nil {
		return fmt.Errorf("failed to list repos: %w", err)
	}

	mapped := collections.Map(items, func(item Repo) DisplayRepo {
		return DisplayRepo{Name: item.Name, Owner: item.Owner, Workspace: item.Workspace}
	})

	tbl, err := prompt.NewTableComponent[DisplayRepo](mapped)
	if errors.Is(err, prompt.ErrEmptyItems) {
		log.Printf("empty workspace, nothing to list")

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
