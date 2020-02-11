package db

import (
	"bytes"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Uuid     string     `json:"uuid"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Email    string     `json:"email"`
	Status   UserStatus `json:"userStatus"`
}

type UserStatus int

const (
	Deprecate_u UserStatus = iota
	Available_u
)

func (s UserStatus) String() string {
	names := []string{
		"Deprecate",
		"Available",
	}

	return names[s]
}

func (s UserStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func NewUser(username string, password string, email string) *User {
	u := &User{
		Uuid:     uuid.NewV4().String(),
		Username: username,
		Password: password,
		Email:    email,
		Status:   Deprecate_u,
	}

	return u
}

func (user *User) getCollection() string {
	return "user"
}

func (user *User) GetUuid() string {
	return user.Uuid
}

func (user *User) FromBson(m *bson.M) interface{} {
	bsonBytes, _ := bson.Marshal(m)
	bson.Unmarshal(bsonBytes, user)

	return *user
}

func (user *User) IsValid() bool {
	filter := bson.M{"username": user.Username, "password": user.Password}
	users, err := FindByFilter(user, filter)
	if err != nil || len(users) == 0 {
		return false
	}
	return true
}
