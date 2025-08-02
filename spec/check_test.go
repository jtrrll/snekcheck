package e2e_test

import (
	"os"
	"testing"

	e2e "snekcheck-e2e"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	t.Run("it passes with one valid file", func(t *testing.T) {
		e2e.ResetTestDir()
		path := e2e.CreateFile(e2e.TestDir, e2e.File{
			Name:     e2e.ValidChars(10),
			Children: []e2e.File{},
		})

		exitCode, _, _ := e2e.RunExecutable(path)
		assert.Equal(t, 0, exitCode)
	})

	t.Run("it fails with one invalid file", func(t *testing.T) {
		e2e.ResetTestDir()
		path := e2e.CreateFile(os.TempDir(), e2e.File{
			Name:     e2e.ValidChars(5) + " " + e2e.ValidChars(5),
			Children: []e2e.File{},
		})

		exitCode, _, _ := e2e.RunExecutable(path)
		assert.NotEqual(t, 0, exitCode)
	})
}
