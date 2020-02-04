package db
import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/satori/go.uuid"
)

type FateCluster struct {
	Uuid             string           `json:"uuid"`
	Name             string           `json:"name"`
	NameSpaces       string           `json:"namespaces"`
	Version          string           `json:"version"`
	Chart            HelmChart        `json:"chart"`
	Status           ClusterStatus    `json:"status"`
	Backend          ComputingBackend `json:"backend"`
	BootstrapParties Party            `json:"bootstrap_parties"`
}

type ClusterStatus int

const (
	Creating ClusterStatus = iota
	Deleting
	Updating
	Running
	Unavailable
)

func NewFateCluster(name string, nameSpaces string, version string, chart HelmChart, backend ComputingBackend, party Party) *FateCluster {
	fateCluster := &FateCluster{
		Uuid: uuid.NewV4().String(),
		Name: name,
		NameSpaces: nameSpaces,
		Version: version,
		Chart: chart,
		Status: Creating,
		Backend: backend,
		BootstrapParties: party,
	}

	return fateCluster
}

func (fate *FateCluster) getCollection() string {
	return "fate"
}

func (fate *FateCluster) GetUuid() string {
	return fate.Uuid
}

func (fate *FateCluster) FromBson(m *bson.M) interface{}{
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, fate)
	return *fate
}
