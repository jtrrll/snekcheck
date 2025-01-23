package main_test

import (
	"math"
	"os"
	main "snekcheck/cmd/snekcheck"
	"snekcheck/internal/files"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RunTestSuite struct {
	suite.Suite
	Fs   billy.Filesystem
	Root files.Path
}

func TestRun(t *testing.T) {
	suite.Run(t, new(RunTestSuite))
}

func (suite *RunTestSuite) SetupSubTest() {
	suite.Fs = memfs.New()
	suite.Root = files.Path{suite.Fs.Root(), "tmp"}
	_ = suite.Fs.MkdirAll(suite.Root.String(), os.ModeDir)
}

func (suite *RunTestSuite) TestRun() {
	suite.Run("Run()", func() {
		suite.Run("panics given an invalid file system", func() {
			assert.Panics(suite.T(), func() { _ = main.Run(main.Config{Fs: nil}) })
		})
		suite.Run("returns an error when given no paths", func() {
			assert.NotNil(suite.T(), main.Run(main.Config{Fs: memfs.New(), Paths: nil}))
		})
		suite.Run("check", func() {
			suite.Run("succeeds with an empty directory", func() {
				config := main.Config{Depth: math.MaxUint, Fs: suite.Fs, Fix: false, Paths: []files.Path{suite.Root}, Verbose: false}
				assert.Nil(suite.T(), main.Run(config))
			})
			suite.Run("succeeds with a single valid file", func() {
				validFilename := append(suite.Root, "valid")
				require.True(suite.T(), main.IsValid(validFilename.Base()))
				_, createErr := suite.Fs.Create(validFilename.String())
				require.Nil(suite.T(), createErr)

				config := main.Config{Depth: math.MaxUint, Fs: suite.Fs, Fix: false, Paths: []files.Path{suite.Root}, Verbose: false}
				assert.Nil(suite.T(), main.Run(config))
			})
			suite.Run("fails with a single invalid file", func() {
				invalidFilename := append(suite.Root, "InVaLiD")
				require.False(suite.T(), main.IsValid(invalidFilename.Base()))
				_, createErr := suite.Fs.Create(invalidFilename.String())
				require.Nil(suite.T(), createErr)

				config := main.Config{Depth: math.MaxUint, Fs: suite.Fs, Fix: false, Paths: []files.Path{suite.Root}, Verbose: false}
				assert.NotNil(suite.T(), main.Run(config))
			})
		})
		suite.Run("fix", func() {
			suite.Run("succeeds with an empty directory", func() {
				config := main.Config{Depth: math.MaxUint, Fs: suite.Fs, Fix: true, Paths: []files.Path{suite.Root}, Verbose: false}
				assert.Nil(suite.T(), main.Run(config))
			})
			suite.Run("succeeds with a single valid file", func() {
				validFilename := append(suite.Root, "valid")
				require.True(suite.T(), main.IsValid(validFilename.Base()))
				_, createErr := suite.Fs.Create(validFilename.String())
				require.Nil(suite.T(), createErr)

				config := main.Config{Depth: math.MaxUint, Fs: suite.Fs, Fix: true, Paths: []files.Path{suite.Root}, Verbose: false}
				assert.Nil(suite.T(), main.Run(config))
			})
			suite.Run("succeeds with a single invalid file", func() {
				invalidFilename := append(suite.Root, "InVaLiD")
				require.False(suite.T(), main.IsValid(invalidFilename.Base()))
				_, createErr := suite.Fs.Create(invalidFilename.String())
				require.Nil(suite.T(), createErr)

				config := main.Config{Depth: math.MaxUint, Fs: suite.Fs, Fix: true, Paths: []files.Path{suite.Root}, Verbose: false}
				assert.Nil(suite.T(), main.Run(config))
			})
		})
	})
}
