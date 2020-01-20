package db

import (
)

type User struct {
	BaseObject
	Username string `json:"username"` 
	Password string `json:"password"` 
}

func NewUser (username string, password string, bo BaseObject) *User {
	u := &User{
		BaseObject: bo,
		Username: username,
		Password: password,
	}

	return u
}