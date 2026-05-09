package enum

type Output string

const (
	OutputJSON    Output = "JSON"
	OutputYAML    Output = "YAML"
	OutputTable   Output = "TABLE"
	OutputUnknown Output = "UNKNOWN"
)

func (s Output) String() string {
	return string(s)
}

func OutputOf(target string) Output {
	target = Of(target)
	switch target {
	case OutputJSON.String():
		return OutputJSON
	case OutputYAML.String():
		return OutputYAML
	case OutputTable.String():
		return OutputTable
	default:
		return OutputUnknown
	}
}
