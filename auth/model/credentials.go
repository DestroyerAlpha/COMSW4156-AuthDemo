package model

type Credentials struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashedpassword"`
}
