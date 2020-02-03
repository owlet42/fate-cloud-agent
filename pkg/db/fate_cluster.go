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
	Chart            Helm             `json:"chart"`
	Status           ClusterStatus    `json:"status"`
	Backend          ComputingBackend `json:"backend"`
	BootstrapParties Parties          `json:"bootstrap_parties"`
}

type Parties struct {
	PartyId   string `json:"party_id"`
	Endpoint  string `json:"endpoint"`
	PartyType string `json:"party_type"`
}

type ComputingBackend struct {
	BackendType string `json:"backend_type"`
	BackendInfo string `json:"backend_info"`
}

type ClusterStatus int

const (
	Creating ClusterStatus = iota
	Deleting
	Updating
	Running
	Unavailable
)

type Helm struct {
	Name     string `json:"name"` 
	Value    string `json:"value"` 
	Template string `json:"template"` 
}

func NewFateCluster(name string, nameSpaces string, version string, chart Helm, status ClusterStatus, backend ComputingBackend, party Parties) *FateCluster {
	fateCluster := &FateCluster{
		Uuid: uuid.NewV4().String(),
		Name: name,
		NameSpaces: nameSpaces,
		Version: version,
		Chart: chart,
		Status: status,
		Backend: backend,
		BootstrapParties: party,
	}

	return fateCluster
}

func NewParties(partyId string, endpoint string, partyType string) *Parties {
	party := &Parties{
		PartyId: partyId,
		Endpoint: endpoint,
		PartyType: partyType,
	}

	return party
}

func NewComputingBackend(BackendType string, BackendInfo string) *ComputingBackend {
	backend := &ComputingBackend{
		BackendType: BackendType,
		BackendInfo: BackendInfo,
	}

	return backend
}

func NewHelm(name string, value string, template string) *Helm {
	helm := &Helm{
		Name: name,
		Value: value,
		Template: template,
	}

	return helm
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
