package db

import (
	"testing"
)

var userUuid string

func TestAddUser(t *testing.T) {
	u := NewUser("Layne", "test")
	userUuid, error := Save(u)
	if error == nil {
		t.Log(userUuid)
	}
}

func TestFindUsers(t *testing.T) {
	user := &User{}
	results, _ := Find(user)
	t.Log(ToJson(results))
}

// func TestFindByUUID(t *testing.T) {
// 	user := &User{}
// 	results, _ := FindByUUID(user, "7dc85dd9-ef29-4854-a2d6-0a3f8f5d1ab4")
// 	t.Log(ToJson(results))
// }

// func TestDeleteByUUID(t *testing.T) {
// 	user := &User{}
// 	DeleteByUUID(user, "7dc85dd9-ef29-4854-a2d6-0a3f8f5d1ab4")
// }