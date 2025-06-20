package e2e_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	e2e "snekcheck-e2e"
)

func TestHelp(t *testing.T) {
	t.Run("it shows the help message given --help", func(t *testing.T) {
		exitCode, stdout, stderr := e2e.RunExecutable("--help")
		assert.Equal(t, 0, exitCode)
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Usage")
	})
	t.Run("it shows the help message given -h", func(t *testing.T) {
		exitCode, stdout, stderr := e2e.RunExecutable("-h")
		assert.Equal(t, 0, exitCode)
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Usage")
	})
	t.Run("it shows the help message given no arguments", func(t *testing.T) {
		exitCode, stdout, stderr := e2e.RunExecutable()
		assert.Equal(t, 0, exitCode)
		assert.Empty(t, stdout)
		assert.Contains(t, stderr, "Usage")
	})
}
