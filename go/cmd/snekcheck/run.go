package main

import (
	"fmt"
	"slices"

	"github.com/jtrrll/snekcheck/internal/cli"
	"github.com/jtrrll/snekcheck/internal/files"
	"github.com/jtrrll/snekcheck/internal/patterns"

	"github.com/fatih/color"
	"github.com/go-git/go-billy/v5"
)

// Config is a runtime configuration for snekcheck.
type Config struct {
	Fs      billy.Filesystem
	Paths   []files.Path
	Fix     bool
	Verbose bool
}

// Run is the core snekcheck process.
func Run(config Config) cli.Error {
	if config.Fs == nil {
		panic("invalid filesystem")
	}
	if len(config.Paths) == 0 {
		return errNoPathsProvided
	}

	// Initialize results.
	validPaths := make([]files.Path, 0, len(config.Paths))
	invalidPaths := make([]files.Path, 0, len(config.Paths))
	renamedPaths := make([]struct {
		old files.Path
		new files.Path
	}, 0, len(config.Paths))

	// Define processing function.
	process := func(path files.Path) {
		if IsValid(path.Base()) {
			if config.Verbose {
				color.Green("%s\n", path)
			} else {
				fmt.Print(color.GreenString("."))
			}
			validPaths = append(validPaths, path)

			return
		}
		if !config.Fix {
			if config.Verbose {
				color.Red("%s\n", path)
			} else {
				fmt.Print(color.RedString("."))
			}
			invalidPaths = append(invalidPaths, path)

			return
		}
		var newPath files.Path
		newPath = append(append(newPath, path.Parent()...), patterns.ToSnakeCase(path.Base()))
		if config.Fs.Rename(path.String(), newPath.String()) != nil {
			panic(fmt.Errorf("unable to rename %s to %s", path.String(), newPath.String()))
		}
		if config.Verbose {
			color.Yellow("%s -> %s\n", path, newPath)
		} else {
			fmt.Print(color.YellowString("C"))
		}
		renamedPaths = append(renamedPaths, struct {
			old files.Path
			new files.Path
		}{old: path, new: newPath})
	}

	// Process paths.
	slices.SortFunc(config.Paths, func(a, b files.Path) int {
		return len(b) - len(a)
	})
	for _, path := range config.Paths {
		if path == nil {
			panic("invalid path")
		}
		process(path)
	}

	// Print results.
	if config.Verbose {
		fmt.Print("\n")
	} else {
		fmt.Print("\n\n")
	}
	if config.Fix {
		fmt.Printf(
			"%s valid filenames, %s filenames changed\n",
			color.GreenString("%d", len(validPaths)),
			color.YellowString("%d", len(renamedPaths)),
		)
	} else {
		fmt.Printf("%s valid filenames, %s invalid filenames\n", color.GreenString("%d", len(validPaths)), color.RedString("%d", len(invalidPaths)))
	}

	// Terminate.
	if len(invalidPaths) != 0 {
		return errInvalidFileNames
	}

	return nil
}
