package main

import (
	"fmt"
	"github.com/richardjaytea/teago/cmd/teago/app"
	"os"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	s := app.Server{}

	err := s.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	if err != nil {
		return err
	}

	s.Run(":8010")

	return nil
}
