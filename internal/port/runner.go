package port

import "github.com/spf13/cobra"

type Runner[T any] interface {
	Run(cmd *cobra.Command, _ []string) error
}
