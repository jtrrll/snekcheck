package main

// Common errors.
var (
	errInvalidFileNames = Error{Code: 1, Message: "invalid filenames found"}
	errNoPathsProvided  = Error{Code: 2, Message: "no paths provided"}
)

// Error is a command-line error.
type Error struct {
	Message string
	Code    uint8
}

// ExitCode returns the error's status code.
func (err Error) ExitCode() uint8 {
	return err.Code
}

// Error returns the error's message.
func (err Error) Error() string {
	return err.Message
}
