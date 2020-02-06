package service

import "fate-cloud-agent/pkg/db"

func GetChart(version string) db.HelmChart {

	return db.HelmChart{}
}

func GetChartPath(version string) string {

	return "fate/fate"
}
