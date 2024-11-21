package db

import (
	authModel "github.com/DestroyerAlpha/COMSW4156-AuthDemo/auth/model"
	userModel "github.com/DestroyerAlpha/COMSW4156-AuthDemo/user/model"
)

var localDB *InMemoryDatabase

type InMemoryDatabase struct {
	Credentials []*authModel.Credentials
	Users       []*userModel.User
}

func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{}
}

func GetDatabase() *InMemoryDatabase {
	if localDB == nil {
		localDB = NewInMemoryDatabase()
	}
	return localDB
}
