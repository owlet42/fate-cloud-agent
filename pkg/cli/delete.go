package cli

import (
	"fate-cloud-agent/pkg/service"
	"fmt"
	"github.com/urfave/cli/v2"
)

func deleteCommand() *cli.Command {
	return &cli.Command{
		Name:   "delete",
		Usage:  "delete",
		Action: delete,
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
	}
}

func delete(c *cli.Context) error {
	res, err := service.Delete(c.String("namespace"), c.String("name"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if res != nil && res.Info != "" {
		fmt.Println(res.Info)
	}

	fmt.Printf("release \"%s\" uninstalled\n", "")
	return nil
}
