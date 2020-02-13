package db

import (
	"testing"
)

var clusterJustAddedUuid string

func TestNewFateCluster(t *testing.T) {
	InitConfigForTest()
	helm := NewHelmChart("name", "value", "template")
	party := NewParty("9999", "192.168.0.1", "normal")
	backend := NewComputingBackend("egg", "1")
	fate := NewFateCluster("fate-cluster1", "fate-nameSpaces", "v1.2.0", *helm, *backend, *party)
	clusterUuid, error := Save(fate)
	if error == nil {
		t.Log("uuid: ", clusterUuid)
		clusterJustAddedUuid = clusterUuid
	}
}

func TestFindFateCluster(t *testing.T) {
	InitConfigForTest()
	fate := &FateCluster{}
	results, error := Find(fate)
	if error == nil {
		t.Log(ToJson(results))
	}
}

func TestFindFateClusterByUuid(t *testing.T) {
	InitConfigForTest()
	t.Log("Find cluster just add: " + clusterJustAddedUuid)
	fate := &FateCluster{}
	result, error := FindByUUID(fate, clusterJustAddedUuid)
	if error == nil {
		t.Log(ToJson(result))
		t.Log(result.(FateCluster).Name)
	}
}

func TestUpdateCluster(t *testing.T) {
	InitConfigForTest()
	t.Log("Update: " + clusterJustAddedUuid)
	fate := &FateCluster{}
	result, error := FindByUUID(fate, clusterJustAddedUuid)
	if error == nil {
		fate2Update := result.(FateCluster)
		fate2Update.Name = "fate-cluster2"
		fate2Update.NameSpaces = "fate-nameSpaces"

		helm := NewHelmChart("name2", "value2", "template2")
		party := NewParty("10000", "192.168.0.1", "normal")
		backend := NewComputingBackend("egg", "1")
		fate2Update.Chart = *helm
		fate2Update.Backend = *backend
		fate2Update.BootstrapParties = *party
		UpdateByUUID(&fate2Update, clusterJustAddedUuid)
	}

	result, error = FindByUUID(fate, clusterJustAddedUuid)
	if error == nil {
		t.Log(ToJson(result))
	}
}

func TestDeleteClusterByUUID(t *testing.T) {
	InitConfigForTest()
	fate := &FateCluster{}
	DeleteByUUID(fate, clusterJustAddedUuid)
}

func TestReturnMethods(t *testing.T) {
	InitConfigForTest()
	fate := &FateCluster{}
	results, error := Find(fate)
	if error == nil {
		for _, v := range results {
			oneFate := v.(FateCluster)
			t.Log(oneFate.GetUuid())
			t.Log(oneFate.Name)
		}
	}
}
