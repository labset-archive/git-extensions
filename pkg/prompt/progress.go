package prompt

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ProgressModel is a generic progress model for displaying progress of operations.
type ProgressModel[T any] struct {
	items       []T
	progress    progress.Model
	currentItem int
	done        bool
	err         error
	execFunc    func(item T) error
	itemLabel   func(item T) string
}

// NewProgressModel creates a new progress model.
func NewProgressModel[T any](
	items []T,
	execFunc func(item T) error,
	itemLabel func(item T) string,
) *ProgressModel[T] {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)

	return &ProgressModel[T]{
		items:     items,
		progress:  p,
		execFunc:  execFunc,
		itemLabel: itemLabel,
	}
}

// Init initializes the progress model.
func (m *ProgressModel[T]) Init() tea.Cmd {
	return tickCmd()
}

// Update updates the progress model.
func (m *ProgressModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tickMsg:
		if m.currentItem >= len(m.items) {
			m.done = true

			return m, tea.Quit
		}

		item := m.items[m.currentItem]
		err := m.execFunc(item)

		if err != nil {
			m.err = err

			return m, tea.Quit
		}

		m.currentItem++
		percent := float64(m.currentItem) / float64(len(m.items))
		cmd := m.progress.SetPercent(percent)

		return m, tea.Batch(tickCmd(), cmd)
	}

	progressModel, cmd := m.progress.Update(msg)
	m.progress = progressModel.(progress.Model)

	return m, cmd
}

// View renders the progress model.
func (m *ProgressModel[T]) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	if m.done {
		return "Operation completed successfully!\n"
	}

	itemInfo := ""

	if m.currentItem < len(m.items) {
		item := m.items[m.currentItem]
		itemInfo = m.itemLabel(item)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		itemInfo,
		m.progress.View(),
		fmt.Sprintf("%d/%d items processed", m.currentItem, len(m.items)),
	) + "\n"
}

// RunProgress runs the progress model.
func RunProgress[T any](
	items []T,
	execFunc func(item T) error,
	itemLabel func(item T) string,
) error {
	if len(items) == 0 {
		return errors.New("no items to process")
	}

	model := NewProgressModel(items, execFunc, itemLabel)

	if _, err := tea.NewProgram(model).Run(); err != nil {
		return fmt.Errorf("failed to run progress program: %w", err)
	}

	return nil
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
