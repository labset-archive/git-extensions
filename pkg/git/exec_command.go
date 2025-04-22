package git

import (
	"fmt"
	"io"
	"os/exec"
)

type ExecCommand interface {
	Run(name string, args ...string) error
}

type execCommand struct {
	wd  string
	out io.Writer
}

// NewExecCommand creates a new instance of execCommand.
func NewExecCommand(wd string, out io.Writer) ExecCommand {
	return &execCommand{
		wd:  wd,
		out: out,
	}
}

const (
	ErrorFailedToRunCommand = "failed to run command"
)

func (e *execCommand) Run(name string, args ...string) error {
	cmdArgs := []string{name}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("git", cmdArgs...) //nolint:gosec
	cmd.Dir = e.wd
	cmd.Stdout = e.out

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %w", ErrorFailedToRunCommand, err)
	}

	return nil
}
