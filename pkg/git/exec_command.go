package git

import (
	"fmt"
	"os/exec"
	"strings"
)

type ExecCommand interface {
	Run(name string, args ...string) (string, error)
}

type execCommand struct {
	wd string
}

// NewExecCommand creates a new instance of execCommand.
func NewExecCommand(wd string) ExecCommand {
	return &execCommand{
		wd: wd,
	}
}

const (
	ErrorFailedToRunCommand = "failed to run command"
)

func (e *execCommand) Run(name string, args ...string) (string, error) {
	cmdArgs := []string{name}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("git", cmdArgs...) //nolint:gosec
	cmd.Dir = e.wd

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%s: %w", ErrorFailedToRunCommand, err)
	}

	trimmed := strings.TrimSpace(string(output))

	return trimmed, nil
}
