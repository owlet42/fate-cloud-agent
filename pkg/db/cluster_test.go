package db

import (
	"reflect"
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

func TestFindClusterFindByUUID(t *testing.T) {
	InitConfigForTest()
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		args    args
		want    *Cluster
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test",
			args:    args{
				uuid: "0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "test",
			args:    args{
				uuid: "a42d9679-7f44-47a6-a42a-89e3bedacd1f",
			},
			want:    &Cluster{
				Uuid:       "2f41aabe-1610-4e4a-bc1c-9b24e9f8ec11",
				Name:       "fate-8888",
				NameSpaces: "fate-8888",
				Version:    "v1.2.0",
				Metadata:   map[string]interface{}{},
				Status:     Creating_c,
				Backend: ComputingBackend{
					BackendType: "",
					BackendInfo: "",
				},
				BootstrapParties: Party{
					PartyId:   "",
					Endpoint:  "",
					PartyType: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ClusterFindByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClusterFindByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClusterFindByUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
