package auth

import (
	"github.com/DestroyerAlpha/COMSW4156-AuthDemo/auth/dao"
)

func CreateAuthCredentials(userName string, password string) error {
	return dao.CreateEntry(userName, password)
}
