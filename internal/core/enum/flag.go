package enum

import "strings"

func FlagOf(flag string) []string {
	if flag == "" {
		return nil
	}
	flags := strings.Split(flag, ",")
	for i := range flags {
		flags[i] = strings.TrimSpace(flags[i])
	}
	return flags
}
