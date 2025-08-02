package files

import (
	"bufio"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

// GitIgnore is a list of gitignore patterns in order of increasing priority.
type GitIgnore []gitignore.Pattern

// Match determines if a file path is ignored by Git.
func (gi GitIgnore) Match(path Path, isDir bool) bool {
	return gitignore.NewMatcher(gi).Match(path, isDir)
}

// GlobalGitIgnorePatterns parses the list of global gitignore patterns.
func GlobalGitIgnorePatterns(fs billy.Filesystem) ([]gitignore.Pattern, error) {
	basePatterns := []gitignore.Pattern{gitignore.ParsePattern(".git/", nil)}
	systemPatterns, systemErr := gitignore.LoadSystemPatterns(fs)
	userPatterns, userErr := gitignore.LoadGlobalPatterns(fs)

	allErr := errors.Join(systemErr, userErr)
	if allErr != nil {
		return nil, fmt.Errorf("failed to load gitignore patterns: %w", allErr)
	}

	return slices.Concat(basePatterns, systemPatterns, userPatterns), nil
}

// ParseGitIgnore parses the .git/info/exclude patterns in a directory.
func ParseGitIgnore(fs billy.Filesystem, path Path) ([]gitignore.Pattern, error) {
	ignorePatterns, parseErr := parseGitIgnoreFile(fs, append(path, ".gitignore"))
	if parseErr != nil {
		ignorePatterns = nil
	}

	excludePatterns, parseErr := parseGitIgnoreFile(fs, append(path, ".git/info/exclude"))
	if parseErr != nil {
		excludePatterns = nil
	}

	return append(ignorePatterns, excludePatterns...), nil
}

// Parses the patterns from a given gitignore file.
func parseGitIgnoreFile(fs billy.Filesystem, path Path) ([]gitignore.Pattern, error) {
	f, openErr := fs.Open(path.String())
	if openErr != nil {
		return nil, openErr
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	var patterns []gitignore.Pattern

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if !strings.HasPrefix(s, "#") && len(strings.TrimSpace(s)) > 0 {
			patterns = append(patterns, gitignore.ParsePattern(s, path.Parent()))
		}
	}

	return patterns, nil
}
