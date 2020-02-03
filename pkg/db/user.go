package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"github.com/satori/go.uuid"
)

type User struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"` 
	Password string `json:"password"` 
	Email    string `json:"email`
	Status   UserStatus `json:"userStatus"`
}

type UserStatus int

const (
	Deprecate UserStatus = iota
	Available
)

func NewUser (username string, password string, email string, userStatus UserStatus) *User {
	u := &User{
		Uuid: uuid.NewV4().String(),
		Username: username,
		Password: password,
		Email: email,
		Status: userStatus,
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