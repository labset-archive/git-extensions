package prompt

import (
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Select struct {
	Message string
	Options []string
}

type SelectOutput struct {
	Selected string
	Index    int
}

type model struct {
	message  string
	options  []string
	cursor   int
	selected int
	quitting bool
	done     bool
}

func (s *Select) Prompt() (*SelectOutput, error) {
	m := model{
		message: s.Message,
		options: s.Options,
		cursor:  0,
	}

	p := tea.NewProgram(m)

	result, err := p.Run()
	if err != nil {
		return nil, err
	}

	finalModel := result.(model)
	if finalModel.quitting {
		return nil, errors.New("prompt cancelled")
	}

	return &SelectOutput{
		Selected: finalModel.options[finalModel.selected],
		Index:    finalModel.selected,
	}, nil
}

// Init is the Bubble Tea initialization function.
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles user input.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch key := msg.(type) {
	case tea.KeyMsg:
		switch key.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true

			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "enter":
			m.selected = m.cursor
			m.done = true

			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the UI.
func (m model) View() string {
	if m.done {
		return ""
	}

	s := "\n" + m.message + "\n\n"

	for i, opt := range m.options {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("%s %s\n", cursor, opt)
	}

	s += "\n(↑/↓ to move, Enter to select, q to quit)"

	return s
}
