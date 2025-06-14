package patterns_test

import (
	"snekcheck/internal/patterns"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkPosix(b *testing.B) {
	b.Run("IsPosixFileName()", func(b *testing.B) {
		for range b.N {
			patterns.IsPosixFileName("Bench mark")
		}
	})
	b.Run("ToPosixFileName()", func(b *testing.B) {
		for range b.N {
			patterns.ToPosixFileName("Bench mark")
		}
	})
}

func FuzzPosix(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		output := patterns.ToPosixFileName(input)
		assert.True(t, patterns.IsPosixFileName(output))
		if patterns.IsPosixFileName(input) {
			assert.Equal(t, input, output)
		}
	})
}

func TestPosix(t *testing.T) {
	t.Parallel()
	t.Run("IsPosixFileName()", func(t *testing.T) {
		t.Parallel()
		t.Run("identifies valid POSIX filenames", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"posix",
				"POSIX_FILE",
				"_POSIX__FILE_.md",
				"012_345",
				"FILE1.txt",
			}
			for _, input := range testCases {
				assert.True(t, patterns.IsPosixFileName(input))
			}
		})
		t.Run("identifies invalid POSIX filenames", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"-TEST",
				"lol@email",
				"invalid%",
				"file(12).pdf",
			}
			for _, input := range testCases {
				assert.False(t, patterns.IsPosixFileName(input))
			}
		})
	})
	t.Run("ToPosixFileName()", func(t *testing.T) {
		t.Parallel()
		t.Run("does not change valid POSIX filenames", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"POSIX",
				".POSIX_123_.md",
				"_DO_NOT_CHANGE_THIS_PLEASE___",
			}
			for _, input := range testCases {
				require.True(t, patterns.IsPosixFileName(input))
				assert.Equal(t, input, patterns.ToPosixFileName(input))
			}
		})
		t.Run("converts invalid POSIX filenames to valid POSIX filenames", func(t *testing.T) {
			t.Parallel()
			testCases := []struct {
				input  string
				output string
			}{
				{input: "lol#$", output: "lol"},
				{input: "spaced  name", output: "spaced_name"},
				{input: "__012 345.md", output: "__012_345.md"},
			}
			for _, tc := range testCases {
				require.False(t, patterns.IsPosixFileName(tc.input))
				require.True(t, patterns.IsPosixFileName(tc.output))
				actual := patterns.ToPosixFileName(tc.input)
				assert.Equal(t, tc.output, actual)
				assert.True(t, patterns.IsPosixFileName(actual))
			}
		})
	})
}
