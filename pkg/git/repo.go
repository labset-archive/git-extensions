package git

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type Repo struct {
	Executor ExecCommand
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

	var out bytes.Buffer

	executor := NewExecCommand(workingDirectory, &out)
	err = executor.Run("rev-parse", "--is-inside-work-tree")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrorNotAGitRepository, err)
	}

	if strings.TrimSpace(out.String()) != "true" {
		return nil, fmt.Errorf("%s: %w", ErrorNotAGitRepository, err)
	}

	return &Repo{Executor: executor}, nil
}
