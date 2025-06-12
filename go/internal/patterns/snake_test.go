package patterns_test

import (
	"snekcheck/internal/patterns"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSnakeCase(b *testing.B) {
	b.Run("IsSnakeCase()", func(b *testing.B) {
		for range b.N {
			patterns.IsSnakeCase("Bench mark")
		}
	})
	b.Run("ToSnakeCase()", func(b *testing.B) {
		for range b.N {
			patterns.ToSnakeCase("Bench mark")
		}
	})
}

func FuzzSnakeCase(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		output := patterns.ToSnakeCase(input)
		assert.True(t, patterns.IsSnakeCase(output))
		if patterns.IsSnakeCase(input) {
			assert.Equal(t, input, output)
		}
	})
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()
	t.Run("IsSnakeCase()", func(t *testing.T) {
		t.Parallel()
		t.Run("identifies valid snake case", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"",
				"_",
				"snake",
				"_snake_case_",
				"012_345",
				"file1",
			}
			for _, input := range testCases {
				assert.True(t, patterns.IsSnakeCase(input))
			}
		})
		t.Run("identifies invalid snake case", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"Snake",
				"snake case 123",
				"snake-case",
				"SCREAMING_SNAKE_CASE",
			}
			for _, input := range testCases {
				assert.False(t, patterns.IsSnakeCase(input))
			}
		})
	})
	t.Run("ToSnakeCase()", func(t *testing.T) {
		t.Parallel()
		t.Run("does not change valid snake case", func(t *testing.T) {
			t.Parallel()
			testCases := []string{
				"",
				"_",
				"snake",
				"snake_case_123",
				"_do_not_change_this_please_",
			}
			for _, input := range testCases {
				require.True(t, patterns.IsSnakeCase(input))
				assert.Equal(t, input, patterns.ToSnakeCase(input))
			}
		})
		t.Run("converts invalid snake case to valid snake case", func(t *testing.T) {
			t.Parallel()
			testCases := []struct {
				input  string
				output string
			}{
				{input: "LOL.go", output: "lol.go"},
				{input: "snake Case", output: "snake_case"},
				{input: " SNake   caSE ", output: "_snake_case_"},
			}
			for _, tc := range testCases {
				require.False(t, patterns.IsSnakeCase(tc.input))
				require.True(t, patterns.IsSnakeCase(tc.output))
				actual := patterns.ToSnakeCase(tc.input)
				assert.Equal(t, tc.output, actual)
				assert.True(t, patterns.IsSnakeCase(actual))
			}
		})
	})
}
