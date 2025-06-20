package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	main "snekcheck/cmd/snekcheck"
)

func TestIsValid(t *testing.T) {
	t.Parallel()
	t.Run("identifies valid file names", func(t *testing.T) {
		t.Parallel()
		testCases := []string{
			"main.go",
			"flake.nix",
			"LICENSE",
			"README.md",
			"snake_test.go",
			"SPEC.TEST.txt",
		}
		for _, input := range testCases {
			assert.True(t, main.IsValid(input))
		}
	})
	t.Run("identifies invalid file names", func(t *testing.T) {
		t.Parallel()
		testCases := []string{
			"Snake",
			"snake case 123",
			"snake-case",
			"Readme.md",
			"snake.PNG",
			"snake.SNAKE.txt",
		}
		for _, input := range testCases {
			assert.False(t, main.IsValid(input))
		}
	})
}
