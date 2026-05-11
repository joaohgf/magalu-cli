package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/joaohgf/magalu-cli/internal/cli"
)

// main is the entry point of the application.
// It initializes the context and executes the CLI commands defined in the cli package.
func main() {
	ctx := context.Background()
	err := cli.Execute(ctx)
	if err != nil {
		slog.Log(ctx, slog.LevelError, "failed to execute CLI", "error", err)
		os.Exit(1)
		return
	}
}
