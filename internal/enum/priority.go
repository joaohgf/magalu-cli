package enum

const (
	PriorityLow     Type = "LOW"
	PriorityMedium  Type = "MEDIUM"
	PriorityHigh    Type = "HIGH"
	PriorityUnknown Type = "UNKNOWN"
)

func PriorityOf(target string) Type {
	parsed := Of(target)
	switch parsed {
	case PriorityLow:
		return PriorityLow
	case PriorityMedium:
		return PriorityMedium
	case PriorityHigh:
		return PriorityHigh
	default:
		return PriorityUnknown
	}
}
