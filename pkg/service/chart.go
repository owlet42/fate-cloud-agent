package service

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fate-cloud-agent/pkg/db"
	"fmt"
	"github.com/spf13/viper"

	"github.com/rs/zerolog/log"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/cli"
	"sigs.k8s.io/yaml"
)

type Chart interface {
	save(Chart) error
	read(version string) (Chart, error)
	load(version string) (Chart, error)
}

type FateChart struct {
	*db.HelmChart
}

func (fc *FateChart) save() error {
	helmUUID, err := db.Save(fc.HelmChart)
	if err != nil {
		return err
	}
	log.Debug().Str("helmUUID", helmUUID).Msg("helm chart save uuid")
	return nil
}

func (fc *FateChart) read(version string) (*FateChart, error) {

	helmChart, err := db.FindHelmByVersion(version)
	if err != nil {
		return nil, err
	}
	if helmChart == nil {
		return nil, errors.New("find chart error")
	}

	log.Debug().Interface("helmChart version", helmChart.Version).Msg("find chart from db success")

	return &FateChart{
		HelmChart: helmChart,
	}, nil
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

	helmChart, err := SaveChartFromPath(cp, "fate")
	if err != nil {
		return nil, err
	}

	return &FateChart{
		HelmChart: helmChart,
	}, nil
}

func (fc *FateChart) GetChartValuesTemplates() (string, error) {
	if fc.ValuesTemplate == "" {
		return "", errors.New("FateChart ValuesTemplate not exist")
	}
	return fc.ValuesTemplate, nil
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
	log.Debug().Interface("read error", err).Msg("read version FateChart err")

	fc, err = fc.load(version)
	if err != nil {
		return nil, err
	}
	err = fc.save()
	if err != nil {
		return nil, err
	}
	return fc, nil
}

func (fc *FateChart) ToHelmChart() (*chart.Chart, error) {
	if fc == nil || fc.HelmChart == nil {
		return nil, errors.New("FateChart not exist")
	}
	return ConvertToChart(fc.HelmChart)
}

func GetChart(version string) db.HelmChart {
	fateChart, err := GetFateChart(version)
	if err != nil {
		log.Err(err).Msg("get chart error")
	}
	return *fateChart.HelmChart
}

func GetChartPath(version string) string {
	ChartPath := viper.GetString("chart.path")
	ChartPath = fmt.Sprintf("%sfate/%s/", ChartPath, version)
	log.Debug().Str("ChartPath", ChartPath).Msg("ChartPath")
	return ChartPath
}

type Value struct {
	Val []byte
	T   string // type json yaml yml
}



func (v *Value) Unmarshal() (map[string]interface{}, error) {
	si := make(map[string]interface{})
	switch v.T {
	case "yaml":
		err := yaml.Unmarshal(v.Val, &si)
		return si, err
	case "json":
		err := json.Unmarshal(v.Val, &si)
		return si, err
	case "xml":
		err := xml.Unmarshal(v.Val, &si)
		return si, err
	}
	return nil, errors.New("unrecognized type")
}
