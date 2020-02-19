package db

import (
	"bytes"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
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

func (cluster *Cluster) getCollection() string {
	return "cluster"
}

// GetUuid get cluster uuid
func (cluster *Cluster) GetUuid() string {
	return cluster.Uuid
}

// FromBson convert bson to cluster
func (cluster *Cluster) FromBson(m *bson.M) (interface{}, error) {
	bsonBytes, err := bson.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(bsonBytes, cluster)
	if err != nil {
		return nil, err
	}
	return *cluster, nil
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

// ClusterFindByUUID get cluster from via uuid
func ClusterFindByUUID(uuid string) (*Cluster, error) {
	result, err := FindOneByUUID(new(Cluster), uuid)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("cluster no find")
	}
	Cluster, ok := result.(Cluster)
	if !ok {
		return nil, errors.New("assertion type error")
	}
	log.Debug().Interface("Cluster", Cluster).Msg("find Cluster success")
	return &Cluster, nil
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

func ClusterDeleteByUUID(uuid string) error {

	err := DeleteOneByUUID(new(Cluster), uuid)
	if err != nil {
		return err
	}

	log.Debug().Interface("ClusterUuid", uuid).Msg("delete Cluster success")
	return nil
}
