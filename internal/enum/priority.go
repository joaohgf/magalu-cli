package enum

// Type represents the priority level of a task.
const (
	PriorityLow     Type = "LOW"
	PriorityMedium  Type = "MEDIUM"
	PriorityHigh    Type = "HIGH"
	PriorityUnknown Type = "UNKNOWN"
)

// PriorityOf takes a string input and returns the corresponding Type for priority level.
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
