package enum

// ANSI Colors
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorGray   = "\033[90m"
)

// ColorByPriority returns a colored string representation of the task priority based on the provided priority type.
func ColorByPriority(priority Type) string {
	switch priority {
	case PriorityHigh:
		return colorRed + priority.String() + colorReset
	case PriorityMedium:
		return colorYellow + priority.String() + colorReset
	case PriorityLow:
		return colorGreen + priority.String() + colorReset
	default:
		return colorGray + priority.String() + colorReset
	}
}

// ColorByStatus returns a colored string representation of the task status based on the provided status type.
func ColorByStatus(status Type) string {
	switch status {
	case StatusDone:
		return colorGreen + status.String() + colorReset
	case StatusInProgress:
		return colorYellow + status.String() + colorReset
	default:
		return colorGray + status.String() + colorReset
	}
}
