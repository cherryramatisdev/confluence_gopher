package main

import (
	"log"
	"os"

	"github.com/cherryramatisdev/confluence_gopher/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		// TODO: this name is horrible
		Name:     "confluencer",
		Usage:    "...",
		Commands: []cli.Command{cmd.ListArticles},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
