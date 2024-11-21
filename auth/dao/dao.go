package dao

import (
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/auth/model"
	database "github.com/DestroyerAlpha/COMSW4156-AuthDemo/db"
)

func IsValidUsernameAndPassword(username string, password string) bool {
	db := database.GetDatabase()

	for _, cred := range db.Credentials {
		if cred.Username == username && GetHashedPassword(password) == cred.HashedPassword {
			return true
		}
	}
	return false
}

func CreateEntry(userName, password string) error {
	cred := &model.Credentials{
		Username:       userName,
		HashedPassword: GetHashedPassword(password),
	}
	db := database.GetDatabase()
	db.Credentials = append(db.Credentials, cred)
	return nil
}
