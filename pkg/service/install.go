package service

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
	"os"
	"strconv"
)

// install is create a cluster
// value is a json ,
func Install(namespace, name, version, value string) (*releaseElement, error) {

	ENV_CS.Lock()
	err := os.Setenv("HELM_NAMESPACE", namespace)
	if err != nil {
		panic(err)
	}
	settings := cli.New()
	ENV_CS.Unlock()

	cfg := new(action.Configuration)
	client := action.NewInstall(cfg)

	if err := cfg.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), debug); err != nil {
		return nil, err
	}

	chartPath := GetChartPath(version)

	template := GetChartValuesTemplates(chartPath)

	// template to values
	v := make(map[string]interface{})
	err = json.Unmarshal([]byte(value), &v)
	values, err := MapToConfig(v, template)

	// values to map
	vals := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(values), &vals)



	cp, err := client.ChartPathOptions.LocateChart(chartPath, settings)
	if err != nil {
		return nil, err
	}

	debug("CHART PATH: %s\n", cp)

	// Check chartPath dependencies to make sure all are present in /charts
	chartRequested, err := loader.Load(cp)
	if err != nil {
		return nil, err
	}


	rel, err := runInstall(name, chartRequested, client, vals, settings)
	if err != nil {
		return nil, err
	}

	return newReleaseWriter(rel), nil
}
func newReleaseWriter(releases *release.Release) *releaseElement {
	// Initialize the array so no results returns an empty array instead of null

	r := releases
	element := &releaseElement{
		Name:       r.Name,
		Namespace:  r.Namespace,
		Revision:   strconv.Itoa(r.Version),
		Status:     r.Info.Status.String(),
		Chart:      fmt.Sprintf("%s-%s", r.Chart.Metadata.Name, r.Chart.Metadata.Version),
		AppVersion: r.Chart.Metadata.AppVersion,
	}
	t := "-"
	if tspb := r.Info.LastDeployed; !tspb.IsZero() {
		t = tspb.String()
	}
	element.Updated = t
	return element
}
func runInstall(name string, chartRequested *chart.Chart, client *action.Install, vals map[string]interface{}, settings *cli.EnvSettings) (*release.Release, error) {
	debug("Original chartPath version: %q", client.Version)
	if client.Version == "" && client.Devel {
		debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}

	client.ReleaseName = name



	validInstallableChart, err := isChartInstallable(chartRequested)
	if !validInstallableChart {
		return nil, err
	}

	if chartRequested.Metadata.Deprecated {
		_, _ = fmt.Println( "WARNING: This chartPath is deprecated")
	}

	client.Namespace = settings.Namespace()

	return client.Run(chartRequested, vals)
}

func isChartInstallable(ch *chart.Chart) (bool, error) {
	switch ch.Metadata.Type {
	case "", "application":
		return true, nil
	}
	return false, errors.Errorf("%s charts are not installable", ch.Metadata.Type)
}
