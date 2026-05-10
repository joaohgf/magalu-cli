package port

import "github.com/spf13/cobra"

// Runner defines the interface for executing a command with specific logic based on the provided command-line arguments.
type Runner[T any] interface {
	Run(cmd *cobra.Command, _ []string) error
}
