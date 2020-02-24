package cli

import (
	"github.com/urfave/cli/v2"
)

func UserCommand() *cli.Command {
	return &cli.Command{
		Name: "user",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Value: "",
				Usage: "k8s Namespace",
			},
		},
		Subcommands: []*cli.Command{
			UserListCommand(),
			UserInfoCommand(),
		},
		Usage: "add a task to the list",
	}
}

func UserListCommand() *cli.Command {
	return &cli.Command{
		Name: "user",
		Flags: []cli.Flag{
		},
		Usage: "show job list",
		Action: func(c *cli.Context) error {
			User := new(User)
			return getItemList(User)
		},
	}
}

func UserInfoCommand() *cli.Command {
	return &cli.Command{
		Name: "describe",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "uuid",
				Value: "",
				Usage: "uuid",
			},
		},
		Usage: "show User info",
		Action: func(c *cli.Context) error {
			// todo uuid get
			uuid := c.Args().Get(0)
			User := new(User)
			return getItem(User, uuid)
		},
	}
}
