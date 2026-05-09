package enum

type Output string

const (
	OutputJSON  Output = "json"
	OutputYaml  Output = "yaml"
	OutputTable Output = "table"
)

func (s Output) String() string {
	return string(s)
}

func OutputOf(s string) Output {
	switch s {
	case OutputJSON.String():
		return OutputJSON
	case OutputYaml.String():
		return OutputYaml
	case OutputTable.String():
		return OutputTable
	default:
		return OutputYaml
	}
}
