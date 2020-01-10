package cli

import (
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"sort"
)

func initCommandLine() *cli.App {
	app := &cli.App{

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "fate name",
			},
			&cli.StringFlag{
				Name:  "config",
				Value: "",
				Usage: "kube config",
			},&cli.StringFlag{
				Name:  "chart",
				Value: "",
				Usage: "chart path",
			},
		},
		Commands: []*cli.Command{
			serviceCommand(),
			configCommand(),
			installCommand(),
			listCommand(),
			getCommand(),
			deleteCommand(),
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	return app
}

func Run(Args []string) {
	app := initCommandLine()
	err := app.Run(Args)
	if err != nil {
		log.Fatal(err)
	}
}
