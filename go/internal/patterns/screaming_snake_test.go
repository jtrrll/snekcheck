package patterns_test

import (
	"snekcheck/internal/patterns"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkScreamingSnakeCase(b *testing.B) {
	b.Run("IsScreamingSnakeCase()", func(b *testing.B) {
		for range b.N {
			patterns.IsScreamingSnakeCase("Bench mark")
		}
	})
	b.Run("ToScreamingSnakeCase()", func(b *testing.B) {
		for range b.N {
			patterns.ToScreamingSnakeCase("Bench mark")
		}
	})
}

func FuzzScreamingSnakeCase(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		output := patterns.ToScreamingSnakeCase(input)
		assert.True(t, patterns.IsScreamingSnakeCase(output))
		if patterns.IsScreamingSnakeCase(input) {
			assert.Equal(t, input, output)
		}
	})
}

func TestScreamingSnakeCase(t *testing.T) {
	t.Parallel()
	t.Run("IsScreamingSnakeCase()", func(t *testing.T) {
		t.Parallel()
		t.Run("identifies valid screaming snake case", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"",
				"_",
				"SNAKE",
				"_SNAKE_CASE_",
				"012_345",
				"FILE1",
			}
			for _, input := range testCases {
				assert.True(t, patterns.IsScreamingSnakeCase(input))
			}
		})
		t.Run("identifies invalid screaming snake case", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"Snake",
				"snake case 123",
				"snake-case",
				"snake_case",
			}
			for _, input := range testCases {
				assert.False(t, patterns.IsScreamingSnakeCase(input))
			}
		})
	})
	t.Run("ToScreamingSnakeCase()", func(t *testing.T) {
		t.Parallel()
		t.Run("does not change valid screaming snake case", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"",
				"_",
				"SNAKE",
				"SNAKE_CASE_123",
				"_DO_NOT_CHANGE_THIS_PLEASE_",
			}
			for _, input := range testCases {
				require.True(t, patterns.IsScreamingSnakeCase(input))
				assert.Equal(t, input, patterns.ToScreamingSnakeCase(input))
			}
		})
		t.Run("converts invalid screaming snake case to valid screaming snake case", func(t *testing.T) {
			t.Parallel()
			testCases := []struct {
				input  string
				output string
			}{
				{input: "lol#$", output: "LOL"},
				{input: "snake Case", output: "SNAKE_CASE"},
				{input: " SNake   caSE ", output: "_SNAKE_CASE_"},
			}
			for _, tc := range testCases {
				require.False(t, patterns.IsScreamingSnakeCase(tc.input))
				require.True(t, patterns.IsScreamingSnakeCase(tc.output))
				actual := patterns.ToScreamingSnakeCase(tc.input)
				assert.Equal(t, tc.output, actual)
				assert.True(t, patterns.IsScreamingSnakeCase(actual))
			}
		})
	})
}
