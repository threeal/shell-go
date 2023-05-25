// Package shell provides functions for executing shell commands.
package shell

import (
	"os"
	"os/exec"
)

// Run executes a shell command, displaying its outputs to the current process's standard outputs.
// Returns an error if the command execution fails.
func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
