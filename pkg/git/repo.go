package git

import (
	"fmt"
	"os"
	"strings"
)

type Repo struct {
	WorkingDirectory string
	Executor         ExecCommand
}

type Loader interface {
	Load() (*Repo, error)
}

type loader struct {
}

func NewLoader() Loader {
	return &loader{}
}

const (
	ErrorFailedToGetWorkingDirectory = "failed to get working directory"
	ErrorNotAGitRepository           = "not a git repository"
)

func (l *loader) Load() (*Repo, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrorFailedToGetWorkingDirectory, err)
	}

	executor := NewExecCommand(workingDirectory)
	output, err := executor.Run("rev-parse", "--is-inside-work-tree")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrorNotAGitRepository, err)
	}

	if strings.TrimSpace(output) != "true" {
		return nil, fmt.Errorf("%s: %w", ErrorNotAGitRepository, err)
	}

	return &Repo{WorkingDirectory: workingDirectory, Executor: executor}, nil
}
