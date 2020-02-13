package service

import (
	"fate-cloud-agent/pkg/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"helm.sh/helm/v3/pkg/chart"
	"io/ioutil"
)

type Chart interface {
	save(Chart) error
	read(version string) (Chart, error)
	load(version string) (Chart, error)
}

type FateChart struct {
}

func (fc *FateChart) save() error {
	return nil
}

func (fc *FateChart) read(version string) (*FateChart, error) {
	return &FateChart{}, nil
}
func (fc *FateChart) load(version string) (*FateChart, error) {
	return &FateChart{}, nil
}

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

func FateChartToHelmChart(fc *FateChart) (*chart.Chart, error) {
	ch := new(chart.Chart)


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
	chartRepository := viper.GetString("chart.repository")
	var fileName string
	if ok := checkVersion(version); ok {
		fileName = "fate-" + version + ".tgz"
	} else {
		fileName = "fate-latest.tgz"
	}

	return chartRepository + fileName
}

func checkVersion(version string) bool {
	if version != "1.2.0" {
		return false
	}
	return true
}
