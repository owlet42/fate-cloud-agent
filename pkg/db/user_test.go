package db

// import (
// 	"testing"
// )


// func TestAddUser(t *testing.T) {
// 	u := NewUser("Layne", "test", *NewBaseObject())
// 	result, error := u.Save("user", u)
// 	if error == nil {
// 		t.Log(result)
// 	}
// }

// func TestFindUsers(t *testing.T) {
// 	b := NewBaseObject()
// 	user := &User{}
// 	results, _ := b.Find("user", *user)
// 	t.Log(results)
// }

// func TestFindByUUID(t *testing.T) {
// 	b := NewBaseObject()
// 	user := &User{}
// 	results, _ := b.FindByUUID("user", "a1d14f92-6594-4a1d-8dec-4800e440e9b5", *user)
// 	t.Log(results)
// }

// func TestDeleteByUUID(t *testing.T) {
// 	b := NewBaseObject()
// 	b.DeleteByUUID("user", "a1d14f92-6594-4a1d-8dec-4800e440e9b5")
// }