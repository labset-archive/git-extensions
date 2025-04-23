package prompt

import (
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
		return nil, fmt.Errorf("failed to run prompt: %w", err)
	}

	finalModel, ok := result.(multiModel)
	if !ok {
		return nil, ErrPromptEmpty
	}

	if finalModel.quitting {
		return nil, ErrPromptCancelled
	}

	var selected []string

	var indices []int

	for index, found := range finalModel.selected {
		if found {
			selected = append(selected, finalModel.options[index])
			indices = append(indices, index)
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
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

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

	return m, nil
}

func (m multiModel) View() string {
	if m.done {
		return ""
	}

	output := "\n" + m.message + "\n\n"

	for index, opt := range m.options {
		cursor := " " // no cursor
		if m.cursor == index {
			cursor = ">" // current line
		}

		checked := " " // not selected
		if m.selected[index] {
			checked = "x"
		}

		output += fmt.Sprintf("%s [%s] %s\n", cursor, checked, opt)
	}

	output += "\n(↑/↓ to move, Space to select, Enter to confirm, q to quit)"

	return output
}
