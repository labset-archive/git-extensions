package git

import (
	"fmt"
	"os"
	"os/exec"
)

type ExecCommand interface {
	Run(name string, args ...string) error
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

func (e *execCommand) Run(name string, args ...string) error {
	cmdArgs := []string{name}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("git", cmdArgs...) //nolint:gosec
	cmd.Dir = e.wd
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	return nil
}
