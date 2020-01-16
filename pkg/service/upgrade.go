package service

import (
	"fmt"
	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
	"os"
)

func Upgrade(namespace, name, chartPath string) error {

	ENV_CS.Lock()
	err := os.Setenv("HELM_NAMESPACE", namespace)
	if err != nil {
		panic(err)
	}
	settings := cli.New()
	ENV_CS.Unlock()

	cfg := new(action.Configuration)
	client := action.NewUpgrade(cfg)
	valueOpts := &values.Options{}

	if err := cfg.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), debug); err != nil {
		return err
	}

	client.Namespace = settings.Namespace()

	if client.Version == "" && client.Devel {
		debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}

	vals, err := valueOpts.MergeValues(getter.All(settings))
	if err != nil {
		return err
	}

	// Check chart dependencies to make sure all are present in /charts
	ch, err := loader.Load(chartPath)
	if err != nil {
		return err
	}
	if req := ch.Metadata.Dependencies; req != nil {
		if err := action.CheckDependencies(ch, req); err != nil {
			return err
		}
	}

	if ch.Metadata.Deprecated {
		fmt.Println("WARNING: This chart is deprecated")
	}

	_, err = client.Run(name, ch, vals)
	if err != nil {
		return errors.Wrap(err, "UPGRADE FAILED")
	}

	return nil

}