package prompt

import (
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MultiSelect struct {
	Message string
	Options []string
}

type MultiSelectOutput struct {
	Selected []string
	Indices  []int
}

type multiModel struct {
	message  string
	options  []string
	selected map[int]bool
	cursor   int
	done     bool
	quitting bool
}

func (m *MultiSelect) Prompt() (*MultiSelectOutput, error) {
	initial := multiModel{
		message:  m.Message,
		options:  m.Options,
		selected: make(map[int]bool),
		cursor:   0,
	}

	program := tea.NewProgram(initial)

	result, err := program.Run()
	if err != nil {
		return nil, err
	}

	finalModel := result.(multiModel)
	if finalModel.quitting {
		return nil, errors.New("prompt cancelled")
	}

	var selected []string

	var indices []int

	for i, ok := range finalModel.selected {
		if ok {
			selected = append(selected, finalModel.options[i])
			indices = append(indices, i)
		}
	}

	return &MultiSelectOutput{
		Selected: selected,
		Indices:  indices,
	}, nil
}

func (m multiModel) Init() tea.Cmd {
	return nil
}

func (m multiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case " ":
			m.selected[m.cursor] = !m.selected[m.cursor]
		case "enter":
			m.done = true

			return m, tea.Quit
		}
	}

	return m, nil
}

func (m multiModel) View() string {
	if m.done {
		return ""
	}

	s := "\n" + m.message + "\n\n"

	for i, opt := range m.options {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // current line
		}

		checked := " " // not selected
		if m.selected[i] {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, opt)
	}

	s += "\n(↑/↓ to move, Space to select, Enter to confirm, q to quit)"

	return s
}
