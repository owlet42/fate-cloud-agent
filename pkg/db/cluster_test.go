package db

import (
	"testing"
)

var clusterJustAddedUuid string

func TestNewCluster(t *testing.T) {
	InitConfigForTest()
	party := NewParty("9999", "192.168.0.1", "normal")
	backend := NewComputingBackend("egg", "1")
	fate := NewCluster("fate-cluster1", "fate-nameSpaces", "v1.2.0", *backend, *party)
	clusterUuid, error := Save(fate)
	if error == nil {
		t.Log("uuid: ", clusterUuid)
		clusterJustAddedUuid = clusterUuid
	}
}

func TestFindCluster(t *testing.T) {
	InitConfigForTest()
	fate := &Cluster{}
	results, error := Find(fate)
	if error == nil {
		t.Log(ToJson(results))
	}
}

func TestFindClusterByUuid(t *testing.T) {
	InitConfigForTest()
	clusterJustAddedUuid = "f3a366f5-bf97-4be2-b49a-2137fe84a38b"
	t.Log("Find cluster just add: " + clusterJustAddedUuid)
	fate := &Cluster{}
	result, error := FindByUUID(fate, clusterJustAddedUuid)
	if error == nil {
		t.Log(ToJson(result))
		t.Log(result.(Cluster).Name)
	}
}

func TestUpdateCluster(t *testing.T) {
	InitConfigForTest()
	t.Log("Update: " + clusterJustAddedUuid)
	fate := &Cluster{}
	result, error := FindByUUID(fate, clusterJustAddedUuid)
	if error == nil {
		fate2Update := result.(Cluster)
		fate2Update.Name = "fate-cluster2"
		fate2Update.NameSpaces = "fate-nameSpaces"

		party := NewParty("10000", "192.168.0.1", "normal")
		backend := NewComputingBackend("egg", "1")
		fate2Update.Backend = *backend
		fate2Update.BootstrapParties = *party
		UpdateByUUID(&fate2Update, clusterJustAddedUuid)
	}

	result, error = FindByUUID(fate, clusterJustAddedUuid)
	if error == nil {
		t.Log(ToJson(result))
	}
}

func TestDeleteByUUID(t *testing.T) {
	InitConfigForTest()
	fate := &Cluster{}
	DeleteByUUID(fate, clusterJustAddedUuid)
}

func TestReturnMethods(t *testing.T) {
	InitConfigForTest()
	fate := &Cluster{}
	results, error := Find(fate)
	if error == nil {
		for _, v := range results {
			oneFate := v.(Cluster)
			t.Log(oneFate.GetUuid())
			t.Log(oneFate.Name)
		}
	}
}
