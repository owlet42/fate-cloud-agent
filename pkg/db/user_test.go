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
	t.Log(results)
}

func TestFindByUUID(t *testing.T) {
	user := &User{}
	results, _ := FindByUUID(user, "a1d14f92-6594-4a1d-8dec-4800e440e9b5")
	t.Log(results)
}

func TestDeleteByUUID(t *testing.T) {
	user := &User{}
	DeleteByUUID(user, "a1d14f92-6594-4a1d-8dec-4800e440e9b5")
}