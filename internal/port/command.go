package port

import "github.com/spf13/cobra"

type (
	Option func(command *cobra.Command)
)
