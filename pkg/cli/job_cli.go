package cli

import (
	"github.com/urfave/cli/v2"
)

func JobCommand() *cli.Command {
	return &cli.Command{
		Name: "job",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "namespace",
				Value: "",
				Usage: "k8s Namespace",
			},
		},
		Subcommands: []*cli.Command{
			JobListCommand(),
			JobInfoCommand(),
			JobDeleteCommand(),
		},
		Usage: "add a task to the list",
	}
}

func JobListCommand() *cli.Command {
	return &cli.Command{
		Name: "list",
		Flags: []cli.Flag{
		},
		Usage: "show job list",
		Action: func(c *cli.Context) error {
			cluster := new(Job)
			return getItemList(cluster)
		},
	}
}

func JobDeleteCommand() *cli.Command {
	return &cli.Command{
		Name: "delete",
		Flags: []cli.Flag{
		},
		Usage: "show job list",
		Action: func(c *cli.Context) error {
			cluster := new(Job)
			uuid := c.String("uuid")
			return deleteItem(cluster, uuid)
		},
	}
}

func JobInfoCommand() *cli.Command {
	return &cli.Command{
		Name: "describe",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "uuid",
				Value: "",
				Usage: "uuid",
			},
		},
		Usage: "show job info",
		Action: func(c *cli.Context) error {
			// todo uuid get
			uuid := c.String("uuid")
			Job := new(Job)
			return getItem(Job, uuid)
		},
	}
}
