package enum

import "strings"

// Type is a generic type that can be used to represent any string-based enumeration.
type Type string

func (t Type) String() string {
	return string(t)
}

func (t Type) ToLower() string {
	return strings.ToLower(t.String())
}

// Of is a helper function that transforms a string into a format suitable for enum parsing.
// It trims whitespace, converts to uppercase, and replaces spaces with underscores.
func Of(target string) Type {
	target = strings.TrimSpace(target)
	target = strings.ToUpper(target)
	target = strings.ReplaceAll(target, " ", "_")
	return Type(target)
}
