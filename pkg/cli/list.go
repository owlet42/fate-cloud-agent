package cli

import (
	"fate-cloud-agent/pkg/service"
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/urfave/cli/v2"
	"helm.sh/helm/v3/pkg/cli/output"
	"os"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name: "list",
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
			}, &cli.StringFlag{
				Name:  "chart",
				Value: "",
				Usage: "chart path",
			},
		},
		Usage:  "list",
		Action: list,
	}
}

func list(c *cli.Context) error {
	release, err := service.List(c.String("name"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	table := uitable.New()
	table.AddRow("NAME", "NAMESPACE", "REVISION", "UPDATED", "STATUS", "CHART", "APP VERSION")
	for _, r := range release.Releases {
		table.AddRow(r.Name, r.Namespace, r.Revision, r.Updated, r.Status, r.Chart, r.AppVersion)
	}
	return output.EncodeTable(os.Stdout, table)
}
