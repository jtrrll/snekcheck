package main

// Common errors.
var (
	invalidFileNamesErr      = Error{Code: 1, Message: "invalid filenames found"}
	noPathsProvidedErr       = Error{Code: 2, Message: "no paths provided"}
	failedToBuildFileTreeErr = Error{Code: 3, Message: "failed to build file tree"}
)

// A command-line error.
type Error struct {
	Code    uint8
	Message string
}

// Returns the error's status code.
func (err Error) ExitCode() uint8 {
	return err.Code
}

// Returns the error's message.
func (err Error) Error() string {
	return err.Message
}
