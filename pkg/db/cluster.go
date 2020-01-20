package db

import()

type Cluster struct {
	BaseObject
	Name       string `json:"name"`
	NameSpaces string `json:"namespaces"`
	Version    string `json:"version"`
	PartyId    string `json:"party_id"`
	Chart      HelmChart   `json:"chart"`
}

type HelmChart struct {
	Name     string `json:"name"` 
	Value    string `json:"value"` 
	Template string `json:"template"` 
}

func NewCluster(name string, nameSpaces string, version string, partyId string, chart HelmChart, bo BaseObject) *Cluster {
	c := &Cluster{
		BaseObject: bo,
		Name: name,
		NameSpaces: nameSpaces,
		Version: version,
		PartyId: partyId,
		Chart: chart,
	}

	return c
}

func NewHelmChart(name string, value string, template string) *HelmChart {
	h := &HelmChart{
		Name: name,
		Value: value,
		Template: template,
	}

	return h
}