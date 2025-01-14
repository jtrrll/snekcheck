package main

// Common errors.
var (
	invalidFileNamesErr      = Error{code: 1, message: "invalid filenames found"}
	noPathsProvidedErr       = Error{code: 2, message: "no paths provided"}
	failedToBuildFileTreeErr = Error{code: 3, message: "failed to build file tree"}
)

// A command-line error.
type Error struct {
	code    uint8
	message string
}

// Returns the error's status code.
func (err Error) Code() uint8 {
	return err.code
}

// Returns the error's message.
func (err Error) Error() string {
	return err.message
}
