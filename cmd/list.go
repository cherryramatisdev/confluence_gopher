package cmd

import (
	"fmt"

	"github.com/cherryramatisdev/confluence_gopher/internal/config"
	"github.com/urfave/cli"
)

var ListArticles = cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "list all the possible articles",
	Action: func(c *cli.Context) error {
		config, err := config.GetConfig()

		if err != nil {
			return err
		}

		fmt.Println(config.Auth.UserMail)
		return nil
	},
}
