/*
`snekcheck` recursively lints all provided file paths to ensure all filenames are snake_case.

Usage:

	snekcheck <flag> ... <path> ...
*/
package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"snekcheck/internal/files"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/spf13/pflag"
)

var (
	// A colorful logger.
	// TODO: Use a better view solution.
	logger = configureLogger()
)

// The snekcheck CLI builds a runtime configuration for the core process.
// Will exit with a non-zero exit code upon failure.
func main() {
	// Initialize filesystem.
	rootFs := osfs.New("/")

	// Parse CLI flags.
	cli := pflag.NewFlagSet("flags", pflag.ExitOnError)
	cli.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprint(os.Stderr, cli.FlagUsages())
	}
	depth := cli.UintP("depth", "d", math.MaxUint8, "The number of levels to descend into a directory")
	fix := cli.BoolP("fix", "f", false, "Whether to correct invalid filenames")
	help := cli.BoolP("help", "h", false, "Print usage help")
	if cli.Parse(os.Args) != nil {
		panic("failed to parse command line flags and arguments")
	}

	if *help {
		cli.Usage()
		exit(0)
	}

	paths := cli.Args()[1:]
	absPaths := make([]files.Path, len(paths))
	for i, path := range paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}
		absPaths[i] = files.NewPath(absPath)
	}
	if len(absPaths) == 0 {
		logger.Error("no valid files or directories specified")
		exit(1)
	}

	success, err := Run(Config{
		Fs:    rootFs,
		Paths: absPaths,
		Depth: *depth,
		Fix:   *fix,
	})
	if err != nil {
		logger.Error(err)
		exit(1)
	}
	if success {
		exit(0)
	} else {
		exit(1)
	}
}

// Configures the CLI logger.
func configureLogger() (logger *log.Logger) {
	logger = log.New(os.Stderr)
	styles := log.DefaultStyles()
	styles.Keys["INVALID"] = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#f44747"))
	styles.Values["INVALID"] = lipgloss.NewStyle()
	styles.Keys["VALID"] = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6a9955"))
	styles.Values["VALID"] = lipgloss.NewStyle()
	styles.Keys["FIXED"] = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#dcdcaa"))
	styles.Values["FIXED"] = lipgloss.NewStyle()
	logger.SetStyles(styles)
	return
}

// Terminates the current program with the given status code.
// Panics if the exit code is not in the range [0, 125].
func exit(code uint8) {
	if code > 125 {
		panic(fmt.Errorf("invalid exit code: %d", code))
	}
	os.Exit(int(code))
}
