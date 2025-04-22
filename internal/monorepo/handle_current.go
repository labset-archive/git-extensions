package monorepo

import (
	"fmt"

	"github.com/viqueen/git-extensions/pkg/git"
)

func Current(_ *git.Repo) error {
	if err := gitCmd("config", "--local", "monorepo.view"); err != nil {
		return fmt.Errorf("failed to set git config: \n%w", err)
	}

	return nil
}
