package main

import (
	"context"

	"github.com/joaohgf/magalu-cli/internal/cli"
)

func main() {
	ctx := context.Background()
	cli.Execute(ctx)
}
