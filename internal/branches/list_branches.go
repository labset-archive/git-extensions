package branches

import (
	"fmt"
	"strings"

	"github.com/viqueen/git-extensions/pkg/git"
)

func listRecentBranches(repo *git.Repo) ([]string, error) {
	output, err := repo.Executor.Run(
		"for-each-ref",
		"refs/heads/",
		"--sort=committerdate",
		"--format='%(committerdate:short) %(refname:short)'",
	)

	if err != nil {
		return nil, fmt.Errorf("failed to list branches: %w", err)
	}

	lines := strings.Split(output, "\n")
	branches := make([]string, 0, len(lines))

	for _, line := range lines {
		branches = append(branches, strings.Trim(line, "'"))
	}

	return branches, nil
}

func listSquashedBranches(repo *git.Repo, target string) ([]string, error) {
	output, err := repo.Executor.Run(
		"branch",
		"--no-contains",
		target,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to list branches: %w", err)
	}

	lines := strings.Split(output, "\n")
	branches := make([]string, 0, len(lines))

	for _, line := range lines {
		branches = append(branches, strings.TrimSpace(line))
	}

	return branches, nil
}

func listMergedBranches(repo *git.Repo, target string) ([]string, error) {
	output, err := repo.Executor.Run(
		"branch",
		"--merged",
		target,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to list branches: %w", err)
	}

	lines := strings.Split(output, "\n")
	branches := make([]string, 0, len(lines))

	for _, line := range lines {
		branches = append(branches, strings.TrimSpace(line))
	}

	return branches, nil
}
