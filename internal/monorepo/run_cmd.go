package monorepo

import (
	"fmt"
	"os"
	"os/exec"
)

func gitCmd(name string, args ...string) error {
	cmdArgs := []string{name}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("git", cmdArgs...) //nolint:gosec
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run command %s %v: %w", name, args, err)
	}

	return nil
}
