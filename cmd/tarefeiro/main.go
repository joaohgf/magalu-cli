package main

import (
	"log"
	"os"

	"github.com/joaohgf/magalu-cli/internal/cli"
	"github.com/nanobox-io/scribble"
)

func main() {
	// todo add by env variable
	db, err := scribble.New("./tarefeiro", nil)
	if err != nil {
		logger := log.Default()
		logger.SetOutput(os.Stderr)
		logger.Printf("Failed to initialize database: %v", err)
		os.Exit(1)
	}
	cli.Execute(db)
}
