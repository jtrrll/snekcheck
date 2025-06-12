// Provides utilities for running end-to-end tests for snekcheck.
package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Runs the snekcheck executable with the provided arguments.
// Panics if the executable is not found.
func RunExecutable(args ...string) (int, string, string) {
	binaryPath := filepath.Join(os.Getenv("PROJECT_ROOT"), "result", "bin", "snekcheck")

	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command(binaryPath, args...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	stdout := stdoutBuf.String()
	stderr := stderrBuf.String()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode(), stdout, stderr
		}
		panic(fmt.Errorf("failed to run %s: %w", binaryPath, err))
	}

	return 0, stdout, stderr
}
