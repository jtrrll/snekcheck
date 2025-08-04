package e2e_test

import (
	"path/filepath"
	"strings"
	"testing"

	e2e "github.com/jtrrll/snekcheck-e2e"

	"github.com/stretchr/testify/assert"
)

func TestFix(t *testing.T) {
	t.Run("it does not rename one file that matches the default pattern", func(t *testing.T) {
		e2e.ResetTestDir()
		path := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(10))

		exitCode, _, _ := e2e.RunExecutable("--fix", path)
		assert.Equal(t, 0, exitCode)
		assert.FileExists(t, path)
	})

	t.Run("it renames one file that does not match the default pattern", func(t *testing.T) {
		e2e.ResetTestDir()
		path := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(5)+" "+e2e.ValidChars(5))

		exitCode, _, _ := e2e.RunExecutable("--fix", path)
		assert.Equal(t, 0, exitCode)
		assert.NoFileExists(t, path)
		assert.FileExists(t, strings.ReplaceAll(path, " ", "_"))
	})

	t.Run("it renames only the file that does not match the default pattern", func(t *testing.T) {
		e2e.ResetTestDir()
		path1 := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(10))
		path2 := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(5)+" "+e2e.ValidChars(5))
		path3 := e2e.CreateFile(e2e.TestDir, e2e.ValidChars(10))

		exitCode, _, _ := e2e.RunExecutable("--fix", path1, path2, path3)
		assert.Equal(t, 0, exitCode)
		assert.FileExists(t, path1)
		assert.NoFileExists(t, path2)
		assert.FileExists(t, strings.ReplaceAll(path2, " ", "_"))
		assert.FileExists(t, path3)
	})

	t.Run("it condenses separators into one underscore", func(t *testing.T) {
		e2e.ResetTestDir()
		path := e2e.CreateFile(e2e.TestDir, "01 - Song.mp3")

		exitCode, _, _ := e2e.RunExecutable("--fix", path)
		assert.Equal(t, 0, exitCode)
		assert.NoFileExists(t, path)
		assert.FileExists(t, filepath.Join(e2e.TestDir, "01_song.mp3"))
	})

	t.Run("it renames a non-matching directory of non-matching files without failing", func(t *testing.T) {
		e2e.ResetTestDir()
		dir := e2e.CreateDirectory(e2e.TestDir, e2e.ValidChars(5)+" "+e2e.ValidChars(5))
		path := e2e.CreateFile(dir, e2e.ValidChars(5)+" "+e2e.ValidChars(5))

		exitCode, _, _ := e2e.RunExecutable("--fix", path, dir)
		assert.Equal(t, 0, exitCode)
		assert.NoDirExists(t, dir)
		assert.DirExists(t, strings.ReplaceAll(dir, " ", "_"))
		assert.NoFileExists(t, path)
		assert.FileExists(t, strings.ReplaceAll(path, " ", "_"))
	})
}
