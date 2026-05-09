package enum

const (
	StatusInProgress Status = "IN_PROGRESS"
	StatusDone       Status = "DONE"
	StatusUnknown    Status = "UNKNOWN"
)

type Status string

func (s Status) String() string {
	return string(s)
}

func StatusOf(target string) Status {
	target = Of(target)
	switch target {
	case StatusInProgress.String():
		return StatusInProgress
	case StatusDone.String():
		return StatusDone
	default:
		return StatusUnknown
	}
}
