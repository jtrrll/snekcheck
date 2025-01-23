package main

import (
	"fmt"
	"slices"

	"snekcheck/internal/cli"
	"snekcheck/internal/files"
	"snekcheck/internal/patterns"
	"snekcheck/internal/tree"

	"github.com/fatih/color"
	"github.com/go-git/go-billy/v5"
)

// A runtime configuration for snekcheck.
type Config struct {
	Fs      billy.Filesystem
	Paths   []files.Path
	Depth   uint
	Fix     bool
	Verbose bool
}

// The core snekcheck process.
func Run(config Config) cli.Error {
	if config.Fs == nil {
		panic("invalid filesystem")
	}
	if len(config.Paths) == 0 {
		return noPathsProvidedErr
	}

	// Build file tree.
	fileTree := tree.NewUniqueTree[string]()
	slices.SortFunc(config.Paths, func(a, b files.Path) int {
		return len(a) - len(b)
	})
	for _, path := range config.Paths {
		if addPathWithChildren(fileTree, config.Fs, path, config.Depth) != nil {
			return failedToBuildFileTreeErr
		}
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
	seenPaths := make(map[string]struct{}, len(config.Paths))
	for _, startPath := range config.Paths {
		if startPath == nil {
			panic("invalid path")
		}
		if _, seen := seenPaths[startPath.String()]; seen {
			continue
		}
		seenPaths[startPath.String()] = struct{}{}

		startNode := fileTree.FindByPath(startPath)
		if startNode == nil {
			panic("failed to find start node")
		}

		process(startPath)
		for path := range startNode.IterPreOrder() {
			fullPath := append(startPath, path...)
			seenPaths[fullPath.String()] = struct{}{}
			process(fullPath)
		}
	}

	// Print results.
	if config.Verbose {
		fmt.Print("\n")
	} else {
		fmt.Print("\n\n")
	}
	if config.Fix {
		fmt.Printf("%s valid filenames, %s filenames changed\n", color.GreenString("%d", len(validPaths)), color.YellowString("%d", len(renamedPaths)))
	} else {
		fmt.Printf("%s valid filenames, %s invalid filenames\n", color.GreenString("%d", len(validPaths)), color.RedString("%d", len(invalidPaths)))
	}

	// Terminate.
	if len(invalidPaths) != 0 {
		return invalidFileNamesErr
	}
	return nil
}

// Recursively adds matching child paths to a file tree, up to a maximum depth.
func addPathWithChildren(fileTree tree.UniqueNode[string], fs billy.Filesystem, path files.Path, maxDepth uint) error {
	if fileTree == nil {
		panic("invalid file tree")
	}
	if fs == nil {
		panic("invalid filesystem")
	}

	gitIgnore := loadGlobalGitIgnore(fs)
	match := func(path files.Path, isDir bool) bool {
		return !gitIgnore.Match(path, isDir)
	}

	var recurse func(path files.Path, depth uint) error
	recurse = func(path files.Path, depth uint) error {
		fileInfo, statErr := fs.Stat(path.String())
		if statErr != nil {
			return statErr
		}

		if fileInfo.IsDir() {
			gitIgnore = append(gitIgnore, parseGitIgnorePatterns(fs, path)...)
		}

		if !match(path, fileInfo.IsDir()) {
			return nil
		}
		fileTree.AddPath(path)
		if depth >= maxDepth {
			return nil
		}

		entries, readErr := fs.ReadDir(path.String())
		if readErr != nil {
			entries = nil
		}
		for _, entry := range entries {
			err := recurse(append(path, entry.Name()), depth+1)
			if err != nil {
				return err
			}
		}
		return nil
	}

	return recurse(path, 0)
}

// Parses gitignore patterns in a single directory
func parseGitIgnorePatterns(fs billy.Filesystem, path files.Path) files.GitIgnore {
	patterns, ignoreErr := files.ParseGitIgnore(fs, path)
	if ignoreErr != nil {
		patterns = nil
	}
	return patterns
}

// Parses the list of global gitignore patterns.
// Produces an empty list of patterns upon failure.
func loadGlobalGitIgnore(fs billy.Filesystem) files.GitIgnore {
	globalIgnorePatterns, ignoreErr := files.GlobalGitIgnorePatterns(fs)
	if ignoreErr != nil {
		fmt.Printf("WARN %s", ignoreErr)
		globalIgnorePatterns = nil
	}
	return globalIgnorePatterns
}
