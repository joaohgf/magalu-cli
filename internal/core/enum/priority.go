package enum

const (
	PriorityLow     Priority = "LOW"
	PriorityMedium  Priority = "MEDIUM"
	PriorityHigh    Priority = "HIGH"
	PriorityUnknown Priority = "UNKNOWN"
)

type Priority string

func (p Priority) String() string {
	return string(p)
}

func PriorityOf(target string) Priority {
	target = Of(target)
	switch target {
	case string(PriorityLow):
		return PriorityLow
	case string(PriorityMedium):
		return PriorityMedium
	case string(PriorityHigh):
		return PriorityHigh
	default:
		return PriorityUnknown
	}
}
