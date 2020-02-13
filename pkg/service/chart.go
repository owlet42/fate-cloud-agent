package service

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fate-cloud-agent/pkg/db"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"sigs.k8s.io/yaml"
)

type Chart interface {
	save(Chart) error
	read(version string) (Chart, error)
	load(version string) (Chart, error)
}

type FateChart struct {
	version string
	*chart.Chart
	*db.HelmChart
}

func (fc *FateChart) save() error {
	return nil
}

func (fc *FateChart) read(version string) (*FateChart, error) {

	return &FateChart{}, errors.New("not achieved")
}
func (fc *FateChart) load(version string) (*FateChart, error) {
	chartPath := GetChartPath(version)
	settings := cli.New()
	cfg := new(action.Configuration)
	client := action.NewInstall(cfg)
	cp, err := client.ChartPathOptions.LocateChart(chartPath, settings)
	if err != nil {
		return nil, err
	}

	log.Debug().Str("FateChart chartPath:", cp).Msg("chartPath:")

	// Check chartPath dependencies to make sure all are present in /charts
	chartRequested, err := loader.Load(cp)
	if err != nil {
		return nil, err
	}

	return &FateChart{
		version: version,
		Chart:   chartRequested,
	}, nil
}

func (fc *FateChart) GetChartValuesTemplates() (string, error) {
	chartPath := GetChartPath(fc.version)

	log.Debug().Str("path", chartPath+"values-templates.yaml").Msg("values-templates.yaml path,")

	values, err := ioutil.ReadFile(chartPath + "values-templates.yaml")
	if err != nil {
		log.Error().Msg("readFile values-templates.yaml error :" + err.Error())
		return "", err
	}
	return string(values), nil
}

func (fc *FateChart) GetChartValues(v map[string]interface{}) (map[string]interface{}, error) {
	// template to values
	template, err := fc.GetChartValuesTemplates()
	if err != nil {
		log.Err(err).Msg("GetChartValuesTemplates error")
	}
	values, err := MapToConfig(v, template)

	// values to map
	vals := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(values), &vals)
	if err != nil {
		log.Err(err).Msg("values yaml Unmarshal error")
	}
	return vals, nil
}

// todo  get chart by version from repository
func GetFateChart(version string) (*FateChart, error) {
	fc := new(FateChart)
	fc, err := fc.read(version)
	if err == nil {
		return fc, nil
	}

	fc, err = fc.load(version)
	err = fc.save()
	return fc, err
}

func (fc *FateChart) ToHelmChart() (*chart.Chart, error) {
	ch := new(chart.Chart)
	ch = fc.Chart
	return ch, nil
}

func GetChart(version string) db.HelmChart {

	return db.HelmChart{}
}

func GetChartValuesTemplates(chartPath string) string {
	values, err := ioutil.ReadFile(chartPath + "values-templates.yaml")
	if err != nil {
		log.Error().Msg("readFile values-templates.yaml error :" + err.Error())
	}
	return string(values)
}

func GetChartPath(version string) string {
	return "../../fate/"
}

type Value struct {
	Val string
	T   string // type json yaml yml
}

func (v *Value) Unmarshal() (map[string]interface{}, error) {
	si := make(map[string]interface{})
	switch v.T {
	case "yaml":
		err := yaml.Unmarshal([]byte(v.Val), &si)
		return si, err
	case "json":
		err := json.Unmarshal([]byte(v.Val), &si)
		return si, err
	case "xml":
		err := xml.Unmarshal([]byte(v.Val), &si)
		return si, err
	}
	return nil, errors.New("unrecognized type")
}
