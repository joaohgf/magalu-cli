package port

import (
	"github.com/nanobox-io/scribble"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type (
	Option  func(command *cobra.Command) error
	Command func(db *scribble.Driver, tableWriter *tablewriter.Table) (*cobra.Command, error)
)
