package dao

import (
	database "github.com/DestroyerAlpha/COMSW4156-AuthDemo/db"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/pkg/errors"
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/user/model"
)

func CreateUser(user *model.User) error {
	db := database.GetDatabase()
	db.Users = append(db.Users, user)
	return nil
}

func AddFriend(userId, friendId string) error {
	db := database.GetDatabase()
	for _, u := range db.Users {
		if u.Id == userId {
			u.Friends = append(u.Friends, friendId)
			return nil
		}
	}
	return errors.ErrRecordNotFound
}
