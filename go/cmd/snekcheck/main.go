/*
`snekcheck` recursively lints all provided file paths to ensure all filenames are snake_case.

Usage:

	snekcheck <flag> ... <path> ...
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jtrrll/snekcheck/internal/cli"
	"github.com/jtrrll/snekcheck/internal/files"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/spf13/pflag"
)

// The snekcheck CLI builds a runtime configuration for the core process.
// Will exit with a non-zero exit code upon failure.
func main() {
	rootFs := osfs.New("/")

	flagSet := pflag.NewFlagSet("flags", pflag.ExitOnError)
	if flagSet == nil {
		panic("failed to initialize flag set")
	}

	flagSet.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			"Usage:\n  snekcheck <flag> ... <path> ...\n\n%s",
			flagSet.FlagUsages(),
		)
	}

	fix := flagSet.BoolP("fix", "f", false, "Whether to correct invalid filenames")
	help := flagSet.BoolP("help", "h", false, "Print usage help")
	verbose := flagSet.BoolP("verbose", "v", false, "Whether to print filenames")

	if flagSet.Parse(os.Args) != nil {
		panic("failed to parse command line flags and arguments")
	}

	paths := flagSet.Args()[1:]
	if *help || len(paths) == 0 {
		flagSet.Usage()
		cli.Exit(nil)
	}

	absPaths := make([]files.Path, len(paths))
	for i, path := range paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}

		absPaths[i] = files.NewPath(absPath)
	}

	err := Run(Config{
		Fs:      rootFs,
		Paths:   absPaths,
		Fix:     *fix,
		Verbose: *verbose,
	})
	cli.Exit(err)
}
