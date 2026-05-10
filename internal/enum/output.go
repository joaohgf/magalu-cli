package enum

// Type represents the output format for rendering data.
const (
	OutputDefaultKey Type = "OUTPUT"
	OutputJSON       Type = "JSON"
	OutputYAML       Type = "YAML"
	OutputTable      Type = "TABLE"
	OutputUnknown    Type = "UNKNOWN"
)

// OutputOf takes a string input and returns the corresponding Type for output format.
func OutputOf(target string) Type {
	parsed := Of(target)
	switch parsed {
	case OutputJSON:
		return OutputJSON
	case OutputYAML:
		return OutputYAML
	case OutputTable:
		return OutputTable
	default:
		return OutputUnknown
	}
}
