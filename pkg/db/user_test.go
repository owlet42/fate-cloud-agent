package db

import (
	"fate-cloud-agent/pkg/utility/config"
	"testing"

	"github.com/spf13/viper"
)

var userJustAddedUuid string

func InitConfigForTest() {
	config.InitViper()
	viper.AddConfigPath("../../")
	viper.ReadInConfig()
}

func TestAddUser(t *testing.T) {
	InitConfigForTest()
	u := NewUser("Layne", "test", "email@vmware.com", Deprecate)
	userUuid, error := Save(u)
	if error == nil {
		t.Log(userUuid)
		userJustAddedUuid = userUuid
	}
}

func TestFindUsers(t *testing.T) {
	InitConfigForTest()
	user := &User{}
	results, _ := Find(user)
	t.Log(ToJson(results))
}

func TestFindByUUID(t *testing.T) {
	InitConfigForTest()
	user := &User{}
	results, _ := FindByUUID(user, userJustAddedUuid)
	t.Log(ToJson(results))
}

func TestDeleteByUUID(t *testing.T) {
	InitConfigForTest()
	user := &User{}
	DeleteByUUID(user, userJustAddedUuid)
}
