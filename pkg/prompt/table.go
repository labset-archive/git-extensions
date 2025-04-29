package prompt

import (
	"errors"
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

type TableComponent[T any] struct {
	table table.Model
}

var (
	ErrEmptyItems  = errors.New("items cannot be empty")
	ErrInvalidType = errors.New("items must be a slice of structs")
)

func NewTableComponent[T any](items []T) (*TableComponent[T], error) {
	if len(items) == 0 {
		return nil, ErrEmptyItems
	}

	elementType := reflect.TypeFor[T]()
	if elementType.Kind() != reflect.Struct {
		return nil, ErrInvalidType
	}

	columns := make([]table.Column, elementType.NumField())

	for index := range elementType.NumField() {
		field := elementType.Field(index)
		columns[index] = table.Column{
			Title: field.Name,
			Width: 20,
		}
	}

	rows := make([]table.Row, len(items))

	for index, item := range items {
		row := make([]string, elementType.NumField())

		for fieldIndex := range elementType.NumField() {
			field := elementType.Field(fieldIndex)
			value := reflect.ValueOf(item).FieldByName(field.Name)
			row[fieldIndex] = fmt.Sprintf("%v", value.Interface())
		}

		rows[index] = row
	}

	tbl := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)
	tbl.SetStyles(defaultTableStyles())

	return &TableComponent[T]{table: tbl}, nil
}

func (c *TableComponent[T]) Init() tea.Cmd {
	return nil
}

func (c *TableComponent[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if ok {
		if key.String() == "ctrl+c" || key.String() == "ctrl+d" {
			return c, tea.Quit
		}
	}

	var cmd tea.Cmd
	c.table, cmd = c.table.Update(msg)

	return c, cmd
}

func (c *TableComponent[T]) View() string {
	return c.table.View()
}

func defaultTableStyles() table.Styles {
	styles := table.DefaultStyles()
	styles.Header = styles.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		Bold(true)
	styles.Selected = styles.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	return styles
}
