package patterns

import (
	"regexp"
	"strings"
)

// Precompiled regular expressions.
var (
	// Matches characters that are not valid in snake_case.
	invalidSnakeCaseCharacters = regexp.MustCompile(`[^a-z0-9._]+`)

	// Matches a valid snake_case string.
	snakeCase = regexp.MustCompile(`^[a-z0-9._]*$`)
)

// IsSnakeCase determines if a string is valid snake_case.
func IsSnakeCase(s string) bool {
	return snakeCase.MatchString(s)
}

// ToSnakeCase attempts to convert a string to valid snake_case.
func ToSnakeCase(s string) string {
	s = separators.ReplaceAllLiteralString(s, "_")
	s = uppers.ReplaceAllStringFunc(s, strings.ToLower)

	return invalidSnakeCaseCharacters.ReplaceAllLiteralString(s, "")
}
