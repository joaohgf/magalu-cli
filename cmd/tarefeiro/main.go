package main

import (
	"context"

	"github.com/joaohgf/magalu-cli/internal/cli"
)

// main is the entry point of the application.
// It initializes the context and executes the CLI commands defined in the cli package.
func main() {
	ctx := context.Background()
	cli.Execute(ctx)
}
