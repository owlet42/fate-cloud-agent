package db

import (
	"testing"
)

var userJustAddedUuid string

func TestAddUser(t *testing.T) {
	u := NewUser("Layne", "test", "email@vmware.com", Deprecate)
	userUuid, error := Save(u)
	if error == nil {
		t.Log(userUuid)
		userJustAddedUuid = userUuid
	}
}

func TestFindUsers(t *testing.T) {
	user := &User{}
	results, _ := Find(user)
	t.Log(ToJson(results))
}

func TestFindByUUID(t *testing.T) {
	user := &User{}
	results, _ := FindByUUID(user, userJustAddedUuid)
	t.Log(ToJson(results))
}

func TestDeleteByUUID(t *testing.T) {
	user := &User{}
	DeleteByUUID(user, userJustAddedUuid)
}