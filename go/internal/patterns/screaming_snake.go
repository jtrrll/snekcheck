package patterns

import (
	"regexp"
	"strings"
)

// Precompiled regular expressions.
var (
	// Matches characters that are not valid in SCREAMING_SNAKE_CASE.
	invalidScreamingSnakeCaseCharacters = regexp.MustCompile(`[^A-Z0-9._]+`)

	// Matches a valid SCREAMING_SNAKE_CASE string.
	screamingSnakeCase = regexp.MustCompile(`^[A-Z0-9._]*$`)
)

// IsScreamingSnakeCase determines if a string is valid SCREAMING_SNAKE_CASE.
func IsScreamingSnakeCase(s string) bool {
	return screamingSnakeCase.MatchString(s)
}

// ToScreamingSnakeCase attempts to convert a string to valid SCREAMING_SNAKE_CASE.
func ToScreamingSnakeCase(s string) string {
	s = separators.ReplaceAllLiteralString(s, "_")
	s = lowers.ReplaceAllStringFunc(s, strings.ToUpper)

	return invalidScreamingSnakeCaseCharacters.ReplaceAllLiteralString(s, "")
}
