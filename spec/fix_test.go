package e2e_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	e2e "snekcheck-e2e"

	"github.com/stretchr/testify/assert"
)

func TestFix(t *testing.T) {
	t.Run("it does not rename one valid file", func(t *testing.T) {
		e2e.ResetTestDir()
		path := e2e.CreateFile(e2e.TestDir, e2e.File{
			Name:     e2e.ValidChars(10),
			Children: []e2e.File{},
		})

		exitCode, _, _ := e2e.RunExecutable("--fix", path)
		assert.Equal(t, 0, exitCode)
		assert.FileExists(t, path)
	})

	t.Run("it renames one invalid file", func(t *testing.T) {
		e2e.ResetTestDir()
		path := e2e.CreateFile(os.TempDir(), e2e.File{
			Name:     e2e.ValidChars(5) + " " + e2e.ValidChars(5),
			Children: []e2e.File{},
		})

		exitCode, _, _ := e2e.RunExecutable("--fix", path)
		assert.Equal(t, 0, exitCode)
		assert.NoFileExists(t, path)
		assert.FileExists(t, strings.ReplaceAll(path, " ", "_"))
	})

	t.Run("it condenses separators into one underscore", func(t *testing.T) {
		e2e.ResetTestDir()
		path := e2e.CreateFile(os.TempDir(), e2e.File{
			Name:     "01 - Song.mp3",
			Children: []e2e.File{},
		})

		exitCode, _, _ := e2e.RunExecutable("--fix", path)
		assert.Equal(t, 0, exitCode)
		assert.NoFileExists(t, path)
		assert.FileExists(t, filepath.Join(os.TempDir(), "01_song.mp3"))
	})
}
