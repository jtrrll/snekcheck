package e2e_test

import (
	"testing"

	e2e "github.com/jtrrll/snekcheck-e2e"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	t.Run("it passes with one file that matches the default pattern", func(t *testing.T) {
		e2e.ResetTestDir()

		path := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(10))

		exitCode, _, _ := e2e.RunExecutable(path)
		assert.Equal(t, 0, exitCode)
	})

	t.Run("it fails with one file that does not match the default pattern", func(t *testing.T) {
		e2e.ResetTestDir()

		path := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(5)+" "+e2e.ValidChars(5))

		exitCode, _, _ := e2e.RunExecutable(path)
		assert.NotEqual(t, 0, exitCode)
	})

	t.Run("it fails with many files and one file that does not match the default pattern",
		func(t *testing.T) {
			e2e.ResetTestDir()

			path1 := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(10))
			path2 := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(5)+" "+e2e.ValidChars(5))
			path3 := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(10))

			exitCode, _, _ := e2e.RunExecutable(path1, path2, path3)
			assert.NotEqual(t, 0, exitCode)
		})
}
