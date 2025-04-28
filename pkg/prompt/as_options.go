package prompt

import "github.com/charmbracelet/huh"

func AsOptions(values []string) []huh.Option[string] {
	options := make([]huh.Option[string], len(values))
	for i, value := range values {
		options[i] = huh.NewOption[string](value, value)
	}

	return options
}
