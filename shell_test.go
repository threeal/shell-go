package shell

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Run("RunSuccessfully", func(t *testing.T) {
		read, restore := requirePipeOutput(t)
		defer restore()
		err := Run("go", "version")
		require.NoError(t, err)
		require.Regexp(t, "^go version go\\S+ \\S+\n$", read())
	})

	t.Run("RunWithError", func(t *testing.T) {
		read, restore := requirePipeOutput(t)
		defer restore()
		err := Run("go", "invalid")
		require.ErrorContains(t, err, "exit status")
		require.Regexp(t, "^go invalid: unknown command\nRun 'go help' for usage.\n$", read())
	})
}

func TestRunSilently(t *testing.T) {
	t.Run("RunSuccessfully", func(t *testing.T) {
		read, restore := requirePipeOutput(t)
		defer restore()
		err := RunSilently("go", "version")
		require.NoError(t, err)
		require.Empty(t, read())
	})

	t.Run("RunWithError", func(t *testing.T) {
		read, restore := requirePipeOutput(t)
		defer restore()
		err := RunSilently("go", "invalid")
		require.ErrorContains(t, err, "exit status")
		require.Empty(t, read())
	})
}

func TestOutputSilently(t *testing.T) {
	t.Run("RunSuccessfully", func(t *testing.T) {
		read, restore := requirePipeOutput(t)
		defer restore()
		output, err := OutputSilently("go", "version")
		require.NoError(t, err)
		require.Regexp(t, "^go version go\\S+ \\S+\n$", output)
		require.Empty(t, read())
	})

	t.Run("RunWithError", func(t *testing.T) {
		read, restore := requirePipeOutput(t)
		defer restore()
		output, err := OutputSilently("go", "invalid")
		require.ErrorContains(t, err, "exit status")
		require.Regexp(t, "^go invalid: unknown command\nRun 'go help' for usage.\n$", output)
		require.Empty(t, read())
	})
}
