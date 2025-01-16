package main_test

import (
	main "snekcheck/cmd/snekcheck"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Parallel()
	t.Run("Run()", func(t *testing.T) {
		t.Run("panics given an invalid file system", func(t *testing.T) {
			assert.Panics(t, func() { _ = main.Run(main.Config{Fs: nil}) })
		})
		t.Run("returns an error when given no paths", func(t *testing.T) {
			assert.NotNil(t, main.Run(main.Config{Fs: memfs.New(), Paths: nil}))
		})
	})
}
