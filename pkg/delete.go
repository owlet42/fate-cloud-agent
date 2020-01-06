package pkg

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
	"os"
)

func Delete(args []string) (*release.UninstallReleaseResponse, error) {

	cfg := new(action.Configuration)
	client := action.NewUninstall(cfg)
	out := os.Stdout
	namespace := args[1]
	if err := cfg.Init(settings.RESTClientGetter(), Namespace(namespace), os.Getenv("HELM_DRIVER"), debug); err != nil {
		return nil, err
	}
	name := args[0]

	res, err := client.Run(name)
	if err != nil {
		return nil, err
	}
	if res != nil && res.Info != "" {
		fmt.Fprintln(out, res.Info)
	}

	fmt.Fprintf(out, "release \"%s\" uninstalled\n", args[0])

	return res, nil
}
