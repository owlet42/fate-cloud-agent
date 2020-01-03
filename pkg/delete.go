package pkg

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"os"
)

func Delete(args []string) error {

	cfg := new(action.Configuration)
	client := action.NewUninstall(cfg)
	out := os.Stdout

	if err := cfg.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), debug); err != nil {
		debug("%+v", err)
		os.Exit(1)
	}
	for i := 0; i < len(args); i++ {

		res, err := client.Run(args[i])
		if err != nil {
			return err
		}
		if res != nil && res.Info != "" {
			fmt.Fprintln(out, res.Info)
		}

		fmt.Fprintf(out, "release \"%s\" uninstalled\n", args[i])
	}
    return nil
}
