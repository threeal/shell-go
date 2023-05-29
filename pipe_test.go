package shell

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// requirePipeOutput redirects the standard output and error streams to a pipe,
// allowing capturing of their content for testing purposes.
// It returns two functions: read and restore.
//
// The read function asynchronously reads the content from the pipe and returns it as a string.
//
// The restore function restores the original standard output and error streams.
func requirePipeOutput(t *testing.T) (read func() string, restore func()) {
	// Save the original standard output and error streams.
	oldStdout := os.Stdout
	oldStderr := os.Stderr

	// Replace the standard output and error streams with the pipe.
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	os.Stderr = w

	// Function that asynchronously reads the content from the pipe and returns it as a string.
	read = func() string {
		out := make(chan string)
		// Copy buffer asynchronously to prevent blocking.
		go func() {
			var b strings.Builder
			_, err := io.Copy(&b, r)
			require.NoError(t, err)
			out <- b.String()
		}()
		w.Close()
		return <-out
	}

	// Function that restores the original standard output and error streams.
	restore = func() {
		os.Stdout = oldStdout
		os.Stderr = oldStderr
	}

	return read, restore
}

func TestRequirePipeOutput(t *testing.T) {
	read, restore := requirePipeOutput(t)
	defer restore()
	fmt.Println("some log")
	fmt.Fprintln(os.Stderr, "some error")
	fmt.Println("some other log")
	require.Equal(t, "some log\nsome error\nsome other log\n", read())
}
