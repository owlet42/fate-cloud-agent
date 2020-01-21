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
		t.Log(results)
	}
}

func TestFindFateClusterByUuid(t *testing.T) {
	fate := NewBaseFateCluster()
	results, error := FindByUUID(fate, "63452daa-8a25-46aa-8b34-c5fb2d57288b")
	if error == nil {
		t.Log(results)
	}
}
