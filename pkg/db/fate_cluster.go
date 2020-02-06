package db

import (
	"bytes"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
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
	Creating_c ClusterStatus = iota
	Deleting_c
	Updating_c
	Available_c
	Unavailable_c
)

func (s ClusterStatus) String() string {
	names := []string{
		"Creating",
		"Deleting",
		"Updating",
		"Running",
		"Unavailable",
	}

	return names[s]
}

func (s ClusterStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func NewFateCluster(name string, nameSpaces string, version string, chart HelmChart, backend ComputingBackend, party Party) *FateCluster {
	fateCluster := &FateCluster{
		Uuid:             uuid.NewV4().String(),
		Name:             name,
		NameSpaces:       nameSpaces,
		Version:          version,
		Chart:            chart,
		Status:           Creating_c,
		Backend:          backend,
		BootstrapParties: party,
	}

	return fateCluster
}
func FindFateClusterFindByUUID(uuid string) (*FateCluster, error) {
	result, err := FindByUUID(new(FateCluster), uuid)
	fc := result.(FateCluster)
	return &fc, err
}
func FindFateClusterList(args string) ([]*FateCluster, error) {

	fate := &FateCluster{}
	result, err := Find(fate)
	if err != nil {
		return nil, err
	}

	clusterList := make([]*FateCluster, 0)
	for _, r := range result {
		cluster := r.(FateCluster)
		clusterList = append(clusterList, &cluster)
	}
	return clusterList, nil
}

func (fate *FateCluster) getCollection() string {
	return "fate"
}

func (fate *FateCluster) GetUuid() string {
	return fate.Uuid
}

func (fate *FateCluster) FromBson(m *bson.M) interface{} {
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, fate)
	return *fate
}
