package files_test

import (
	"fmt"
	"os"
	"path/filepath"
	"snekcheck/internal/files"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	t.Parallel()
	t.Run("NewPath()", func(t *testing.T) {
		t.Parallel()
		t.Run("creates a path", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"",
				"longstring",
				strings.Join([]string{"parent", "child"}, string(os.PathSeparator)),
				strings.Join([]string{"dir", "dir", "file.txt"}, string(os.PathSeparator)),
			}
			for _, input := range testCases {
				assert.NotNil(t, files.NewPath(input))
			}
		})
	})
	t.Run("Base()", func(t *testing.T) {
		t.Parallel()
		t.Run("returns the last element", func(t *testing.T) {
			t.Parallel()
			testCases := []files.Path{
				{"dir", "dir", "file.txt"},
				{"longstring"},
				{"dir", "README.md"},
			}
			for _, input := range testCases {
				assert.Equal(t, input[len(input)-1], input.Base())
			}
		})
		t.Run("panics if the path is empty", func(t *testing.T) {
			t.Parallel()
			assert.Panics(t, func() { files.Path{}.Base() })
		})
	})
	t.Run("Parent()", func(t *testing.T) {
		t.Parallel()
		t.Run("returns every element except the last element", func(t *testing.T) {
			t.Parallel()
			testCases := []files.Path{
				{"dir", "dir", "file.txt"},
				{"longstring"},
				{"dir", "README.md"},
			}
			for _, input := range testCases {
				assert.Equal(t, input[:len(input)-1], input.Parent())
			}
		})
		t.Run("panics if the path is empty", func(t *testing.T) {
			t.Parallel()
			assert.Panics(t, func() { files.Path{}.Parent() })
		})
	})
	t.Run("String()", func(t *testing.T) {
		t.Parallel()
		t.Run("returns the full path", func(t *testing.T) {
			t.Parallel()
			testCases := []files.Path{
				{"dir", "dir", "file.txt"},
				{"longstring"},
				{"dir", "README.md"},
			}
			for _, input := range testCases {
				assert.Equal(t, filepath.Join(input...), input.String())
			}
		})
		t.Run("formats correctly as a string", func(t *testing.T) {
			t.Parallel()
			testCases := []files.Path{
				{"dir", "dir", "file.txt"},
				{"longstring"},
				{"dir", "README.md"},
			}
			for _, input := range testCases {
				assert.Equal(t, filepath.Join(input...), fmt.Sprintf("%v", input))
			}
		})
	})
}
