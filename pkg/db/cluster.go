package db

import (
	"bytes"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Cluster struct {
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	NameSpaces string `json:"namespaces"`
	// Cluster version
	Version string `json:"version"`
	// Helm chart version, example: fate v1.2.0
	ChartVersion string `json:"chart_version"`
	// The value of this cluster for installing helm chart
	Values           string                 `json:"values"`
	ChartName        string                 `json:"chart_name"`
	Type             string                 `json:"cluster_type"`
	Metadata         map[string]interface{} `json:"metadata"`
	Status           ClusterStatus          `json:"status"`
	Backend          ComputingBackend       `json:"backend"`
	BootstrapParties Party                  `json:"bootstrap_parties"`
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

// MarshalJSON convert cluster status to string
func (s ClusterStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// NewCluster create cluster object with basic argument
func NewCluster(name string, nameSpaces string, version string, backend ComputingBackend, party Party) *Cluster {
	cluster := &Cluster{
		Uuid:             uuid.NewV4().String(),
		Name:             name,
		NameSpaces:       nameSpaces,
		Version:          version,
		Status:           Creating_c,
		Backend:          backend,
		BootstrapParties: party,
	}

	return cluster
}

// FindClusterFindByUUID get cluster from via uuid
func FindClusterFindByUUID(uuid string) (*Cluster, error) {
	result, err := FindByUUID(new(Cluster), uuid)
	fc := result.(Cluster)
	return &fc, err
}

// FindClusterList get all cluster list
func FindClusterList(args string) ([]*Cluster, error) {

	cluster := &Cluster{}
	result, err := Find(cluster)
	if err != nil {
		return nil, err
	}

	clusterList := make([]*Cluster, 0)
	for _, r := range result {
		cluster := r.(Cluster)
		clusterList = append(clusterList, &cluster)
	}
	return clusterList, nil
}

func (cluster *Cluster) getCollection() string {
	return "cluster"
}

// GetUuid get cluster uuid
func (cluster *Cluster) GetUuid() string {
	return cluster.Uuid
}

// FromBson convert bson to cluster
func (cluster *Cluster) FromBson(m *bson.M) interface{} {
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, cluster)
	return *cluster
}
