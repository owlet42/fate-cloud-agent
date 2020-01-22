package db

import (
	"testing"
)


func TestAddUser(t *testing.T) {
	u := NewUser("Layne", "test")
	result, error := Save(u)
	if error == nil {
		t.Log(result)
	}
}

func TestFindUsers(t *testing.T) {
	user := &User{}
	results, _ := Find(user)
	t.Log(ToJson(results))
}

func TestFindByUUID(t *testing.T) {
	user := &User{}
	results, _ := FindByUUID(user, "14cdfbae-0013-4f3c-b389-a04395eb50f4")
	t.Log(ToJson(results))
}

func TestDeleteByUUID(t *testing.T) {
	user := &User{}
	DeleteByUUID(user, "14cdfbae-0013-4f3c-b389-a04395eb50f4")
}