package cli

import (
	"fate-cloud-agent/pkg/service"
	"fmt"
	"github.com/urfave/cli/v2"
	"time"
)

func getCommand() *cli.Command {
	return &cli.Command{
		Name:   "get",
		Usage:  "get",
		Action: get,
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

func get(c *cli.Context) error {
	release, err := service.Get(c.String("namespace"), c.String("name"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	_, _ = fmt.Printf("NAME: %s\n", release.Name)
	if !release.Info.LastDeployed.IsZero() {
		fmt.Printf("LAST DEPLOYED: %s\n", release.Info.LastDeployed.Format(time.ANSIC))
	}
	fmt.Printf("NAMESPACE: %s\n", release.Namespace)
	fmt.Printf("STATUS: %s\n", release.Info.Status.String())
	fmt.Printf("REVISION: %d\n", release.Version)
	return nil
}
