package main

// Common errors.
var (
	errInvalidFileNames      = Error{Code: 1, Message: "invalid filenames found"}
	errNoPathsProvided       = Error{Code: 2, Message: "no paths provided"}
	errFailedToBuildFileTree = Error{Code: 3, Message: "failed to build file tree"}
)

// A command-line error.
type Error struct {
	Message string
	Code    uint8
}

// Returns the error's status code.
func (err Error) ExitCode() uint8 {
	return err.Code
}

// Returns the error's message.
func (err Error) Error() string {
	return err.Message
}
