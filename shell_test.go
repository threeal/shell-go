package shell

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Run("RunSuccessfully", func(t *testing.T) {
		err := Run("go", "version")
		require.NoError(t, err)
	})

	t.Run("RunWithError", func(t *testing.T) {
		err := Run("go", "invalid")
		require.ErrorContains(t, err, "exit status")
	})
}

func TestRunSilently(t *testing.T) {
	t.Run("RunSuccessfully", func(t *testing.T) {
		err := RunSilently("go", "version")
		require.NoError(t, err)
	})

	t.Run("RunWithError", func(t *testing.T) {
		err := RunSilently("go", "invalid")
		require.ErrorContains(t, err, "exit status")
	})
}
