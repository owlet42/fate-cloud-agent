package db

import (
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"helm.sh/helm/v3/pkg/chart"
)

// HelmChart helm chart model
type HelmChart struct {
	Uuid      string        `json:"uuid"`
	Name      string        `json:"name"`
	Chart     string        `json:"chart"`
	Values    string        `json:"values"`
	Templates []*chart.File `json:"templates"`
	Version   string        `json:"version"`
}

// NewHelmChart create a new helm chart
func NewHelmChart(name string, chart string, values string, templates []*chart.File, version string) *HelmChart {
	helm := &HelmChart{

		Uuid:      uuid.NewV4().String(),
		Name:      name,
		Chart:     chart,
		Values:    values,
		Templates: templates,
		Version:   version,
	}

	return helm
}

func (helm *HelmChart) getCollection() string {
	return "helm"
}

// GetUuid get helm uuid
func (helm *HelmChart) GetUuid() string {
	return helm.Uuid
}

// FromBson convert bson to helm
func (helm *HelmChart) FromBson(m *bson.M) interface{} {
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, helm)
	return *helm
}

// FindHelmByNameAndVersion find helm chart via name and version
func (helm *HelmChart) FindHelmByNameAndVersion(name string, version string) *HelmChart {
	filter := bson.M{"name": name, "version": version}
	helms, err := FindByFilter(helm, filter)
	if err == nil && len(helms) != 0 {
		helm0 := helms[0]
		helm0o := helm0.(HelmChart)
		return &helm0o
		// return helms[0].(*HelmChart)
	}
	return nil
}
