package monorepo

import "fmt"

func Current() error {
	if err := runCmd("git", "config", "--local", "monorepo.view"); err != nil {
		return fmt.Errorf("failed to set git config: \n%w", err)
	}
	return nil
}
