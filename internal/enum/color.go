package enum

// ANSI Colors
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
	colorBold   = "\033[1m"
)

func ColorByPriority(priority Type) string {
	switch priority {
	case PriorityHigh:
		return colorRed + "high" + colorReset
	case PriorityMedium:
		return colorYellow + "medium" + colorReset
	case PriorityLow:
		return colorGreen + "low" + colorReset
	default:
		return colorGray + "?" + colorReset
	}
}

func ColorByStatus(status Type) string {
	switch status {
	case StatusDone:
		return colorGreen + "done" + colorReset
	case StatusInProgress:
		return colorYellow + "in_progress" + colorReset
	default:
		return colorGray + "?" + colorReset
	}
}
