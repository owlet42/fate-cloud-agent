package service

import (
	"fate-cloud-agent/pkg/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io/ioutil"
)

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
	return "E:/machenlong/AI/gitlab/fate-cloud-agent/fate/"
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
