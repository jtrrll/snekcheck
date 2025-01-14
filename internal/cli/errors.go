// A collection of utilities for building command-line interfaces.
package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// A command-line error with a status code.
type Error interface {
	error
	Code() uint8
}

// Terminates the current program.
// Also prints the given error to stderr if present.
func Exit(err Error) {
	if err == nil {
		os.Exit(0)
	}
	code := err.Code()
	if code == 0 {
		os.Exit(0)
	}
	fmt.Fprint(os.Stderr, color.RedString("error: %s\n", err.Error()))
	os.Exit(int(code))
}
