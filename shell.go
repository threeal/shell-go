// Package shell provides functions for executing shell commands.
package shell

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

// Run executes a shell command, displaying its output to the current process's standard output.
// It returns an error if the command execution fails.
func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunSilently executes a shell command silently, without displaying its output.
// It returns an error if the command execution fails.
func RunSilently(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	return cmd.Run()
}

// Output executes a shell command, returning its output as a string
// and displaying its output to the current process's standard output.
// It returns an error if the command execution fails.
func Output(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var out strings.Builder
	cmd.Stdout = io.MultiWriter(&out, os.Stdout)
	cmd.Stderr = io.MultiWriter(&out, os.Stderr)
	err := cmd.Run()
	return out.String(), err
}

// OutputSilently executes a shell command silently and returns its output as a string.
// It returns an error if the command execution fails.
func OutputSilently(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var out strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}
