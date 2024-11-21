package dao

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHashedPassword(password string) string {
	hash := md5.Sum([]byte(password))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
