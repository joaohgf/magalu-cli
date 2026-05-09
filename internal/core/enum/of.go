package enum

import "strings"

// Of is a helper function that transforms a string into a format suitable for enum parsing.
// It trims whitespace, converts to uppercase, and replaces spaces with underscores.
func Of(target string) string {
	target = strings.TrimSpace(target)
	target = strings.ToUpper(target)
	target = strings.ReplaceAll(target, " ", "_")
	return target
}
