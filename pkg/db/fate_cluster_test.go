package db

import (
	"testing"
)

var clusterUuid string

func TestNewFateCluster(t *testing.T) {
	helm := NewHelm("name","value","template")
	fate := NewFateCluster("fate-cluster1","fate-nameSpaces","v1.2.0","party-1111",*helm)
	clusterUuid, error := Save(fate)
	if error ==nil {
		t.Log("uuid: ", clusterUuid)
	}
}

func TestFindFateCluster(t *testing.T) {
	fate := NewBaseFateCluster()
	results, error := Find(fate)
	if error == nil {
		t.Log(ToJson(results))
	}
}

func TestFindFateClusterByUuid(t *testing.T) {
	fate := NewBaseFateCluster()
	result, error := FindByUUID(fate, clusterUuid)
	if error == nil {
		t.Log(ToJson(result))
	}
}
