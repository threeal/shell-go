// Package shell provides functions for executing shell commands.
package shell

import (
	"os"
	"os/exec"
)

// Run executes a shell command, displaying its output to the current process's standard output.
// Returns an error if the command execution fails.
func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunSilently executes a shell command silently, without displaying its output.
// Returns an error if the command execution fails.
func RunSilently(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}
