// Package files is a collection of utiliies for file operations.
package files

import (
	"os"
	"path/filepath"
	"strings"
)

// An OS-specific path separator.
const pathSeparator = string(os.PathSeparator)

// Path is a separated file path.
type Path []string

// NewPath constructs a new Path by splitting the elements with an OS-specific separator.
func NewPath(path string) Path {
	return strings.Split(path, pathSeparator)
}

// Base returns the last element of the path. Will panic if the path is empty.
func (p Path) Base() string {
	return p[len(p)-1]
}

// Parent returns the every element of the path except the last. Will panic if the path is empty.
func (p Path) Parent() Path {
	return p[:len(p)-1]
}

// Converts a Path to a string by joining the elements with an OS-specific separator.
func (p Path) String() string {
	return filepath.Join(p...)
}
