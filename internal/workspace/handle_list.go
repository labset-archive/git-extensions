package workspace

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/viqueen/git-extensions/pkg/prompt"
)

func List(ws Workspace, filer RepoFilter) error {
	items, err := ws.List(filer)
	if err != nil {
		return fmt.Errorf("failed to list repos: %w", err)
	}

	tbl, err := prompt.NewTableComponent[Repo](items)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	program := tea.NewProgram(tbl)
	_, err = program.Run()
	if err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}

	return nil
}
