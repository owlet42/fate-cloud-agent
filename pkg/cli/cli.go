package cli

import (
	"fate-cloud-agent/pkg/service"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"sort"
)

func CommandLine() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "api",
				Value: "api",
				Usage: "run api api",
			},
			&cli.StringFlag{
				Name:  "install",
				Usage: "install fate",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					_, _ = service.Install([]string{"fate-10000", "E:\\machenlong\\AI\\github\\owlet42\\KubeFATE\\k8s-deploy\\fate-10000"})
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list",
				Action: func(c *cli.Context) error {
					service.List("JSON")
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
