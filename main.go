package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		// TODO: this name is horrible
		Name:     "confluencer",
		Usage:    "...",
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
