package enum

const (
	StatusInProgress Type = "IN_PROGRESS"
	StatusDone       Type = "DONE"
	StatusUnknown    Type = "UNKNOWN"
)

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
