package main

import (
	"context"
	"log/slog"

	"github.com/joaohgf/magalu-cli/internal/cli"
)

func main() {
	ctx := context.Background()
	slog.Log(ctx, slog.LevelInfo, "starting tarefeiro CLI application")
	cli.Execute(ctx)
}
