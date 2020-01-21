package db
import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/satori/go.uuid"
)

type FateCluster struct {
	Uuid       string `json:"uuid"` 
	Name       string `json:"name"`
	NameSpaces string `json:"namespaces"`
	Version    string `json:"version"`
	PartyId    string `json:"party_id"`
	Chart      Helm   `json:"chart"`
}

type Helm struct {
	Name     string `json:"name"` 
	Value    string `json:"value"` 
	Template string `json:"template"` 
}

func NewFateCluster(name string, nameSpaces string, version string, partyId string, chart Helm) *FateCluster {
	fateCluster := &FateCluster{
		Uuid: uuid.NewV4().String(),
		Name: name,
		NameSpaces: nameSpaces,
		Version: version,
		PartyId: partyId,
		Chart: chart,
	}

	return fateCluster
}

func NewBaseFateCluster() *FateCluster {
	return new(FateCluster)
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

func (fate *FateCluster) FromBson(m *bson.M){
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, fate)
}
