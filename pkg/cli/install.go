package cli

import (
	"fate-cloud-agent/pkg/service"
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/urfave/cli/v2"
)

func installCommand() *cli.Command {
	return &cli.Command{
		Name:    "install",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "fate name",
			},
			&cli.StringFlag{
				Name:  "namespace",
				Value: "",
				Usage: "k8s Namespace",
			},&cli.StringFlag{
				Name:  "chart",
				Value: "",
				Usage: "chart path",
			},
		},
		Usage:   "add a task to the list",
		Action:  install,
	}
}

func install(c *cli.Context) error {
	fmt.Println(c.String("name"))
	r, err := service.Install(c.String("namespace"), c.String("name"), c.String("chart"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	table := uitable.New()
	table.AddRow("NAME", "NAMESPACE", "REVISION", "UPDATED", "STATUS", "CHART", "APP VERSION")
	table.AddRow(r.Name, r.Namespace, r.Revision, r.Updated, r.Status, r.Chart, r.AppVersion)
	fmt.Println(table)
	return nil
}
