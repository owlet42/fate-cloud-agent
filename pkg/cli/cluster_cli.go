package cli

import (
	"encoding/json"
	"github.com/urfave/cli/v2"
)

func ClusterCommand() *cli.Command {
	return &cli.Command{
		Name: "cluster",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "install",
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
		Subcommands: []*cli.Command{
			ClusterListCommand(),
			ClusterInfoCommand(),
			ClusterDeleteCommand(),
			ClusterInstallCommand(),
			ClusterUpgradeCommand(),
		},
		Usage: "add a task to the list",
	}
}

func ClusterListCommand() *cli.Command {
	return &cli.Command{
		Name: "list",
		Flags: []cli.Flag{
		},
		Usage: "show cluster list",
		Action: func(c *cli.Context) error {
			cluster := new(Cluster)
			return getItemList(cluster)
		},
	}
}

func ClusterInfoCommand() *cli.Command {
	return &cli.Command{
		Name: "describe",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "uuid",
				Value: "",
				Usage: "uuid",
			},
		},
		Usage: "show cluster info",
		Action: func(c *cli.Context) error {
			// todo uuid get
			uuid := c.String("uuid")
			cluster := new(Cluster)
			return getItem(cluster, uuid)
		},
	}
}

func ClusterDeleteCommand() *cli.Command {
	return &cli.Command{
		Name: "delete",
		Flags: []cli.Flag{
		},
		Usage: "cluster delete",
		Action: func(c *cli.Context) error {
			cluster := new(Cluster)
			uuid := c.String("uuid")
			return deleteItem(cluster, uuid)
		},
	}
}


func ClusterInstallCommand() *cli.Command {
	return &cli.Command{
		Name: "install",
		Flags: []cli.Flag{
		},
		Usage: "cluster delete",
		Action: func(c *cli.Context) error {
			cluster := new(Cluster)
			args := struct {
				Name      string
				Namespace string
				Version   string
				Data      []byte
			}{
				Namespace: "fate-10000",
				Name:      "fate-10000",
				Version:   "v1.2.0",
				Data:      []byte(`{"partyId":10000,"endpoint": {"ip":"10.184.111.187","port":30000}}`),
			}
			body,err := json.Marshal(args)
			if err!=nil{
				return err
			}
			return postItem(cluster, body)
		},
	}
}

func ClusterUpgradeCommand() *cli.Command {
	return &cli.Command{
		Name: "upgrade",
		Flags: []cli.Flag{
		},
		Usage: "cluster delete",
		Action: func(c *cli.Context) error {
			cluster := new(Cluster)
			args := struct {
				Name      string
				Namespace string
				Version   string
				Data      []byte
			}{
				Namespace: "fate-10000",
				Name:      "fate-10000",
				Version:   "v1.2.0",
				Data:      []byte(`{"partyId":10000,"endpoint": {"ip":"10.184.111.187","port":30000}}`),
			}
			body,err := json.Marshal(args)
			if err!=nil{
				return err
			}
			return postItem(cluster, body)
		},
	}
}