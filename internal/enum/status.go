package enum

// Type represents the status of a task, which can be "IN_PROGRESS", "DONE", or "UNKNOWN".
const (
	StatusInProgress Type = "IN_PROGRESS"
	StatusDone       Type = "DONE"
	StatusUnknown    Type = "UNKNOWN"
)

// StatusOf takes a string input and returns the corresponding Type for task status.
func StatusOf(target string) Type {
	parsed := Of(target)
	switch parsed {
	case StatusInProgress:
		return StatusInProgress
	case StatusDone:
		return StatusDone
	default:
		return StatusUnknown
	}
}
