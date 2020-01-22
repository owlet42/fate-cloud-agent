package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/satori/go.uuid"
)

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"` 
	Password string `json:"password"` 
}

func NewUser (username string, password string) *User {
	u := &User{
		Uuid: uuid.NewV4().String(),
		Username: username,
		Password: password,
	}

	return u
}

func (user *User) getCollection() string {
	return "user"
}

func (user *User) GetUuid() string {
	return user.Uuid
}

func (user *User) FromBson(m *bson.M) interface{}{
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, user)

	return *user
}