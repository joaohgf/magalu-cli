package enum

const (
	OutputDefaultKey Type = "OUTPUT"
	OutputJSON       Type = "JSON"
	OutputYAML       Type = "YAML"
	OutputTable      Type = "TABLE"
	OutputUnknown    Type = "UNKNOWN"
)

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
