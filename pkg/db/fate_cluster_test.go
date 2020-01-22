package db

import (
	"testing"
)

func TestNewFateCluster(t *testing.T) {
	helm := NewHelm("name","value","template")
	fate := NewFateCluster("fate-cluster1","fate-nameSpaces","v1.2.0","party-1111",*helm)
	uuid, error := Save(fate)
	if error ==nil {
		t.Log("uuid: ", uuid)
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
	result, error := FindByUUID(fate, "fff51a40-bc90-4124-827c-7a88d4cdd970")
	if error == nil {
		t.Log(ToJson(result))
	}
}
