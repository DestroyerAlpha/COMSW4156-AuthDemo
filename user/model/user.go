package model

type User struct {
	Id      string   `json:"id"`
	Name    *Name    `json:"name"`
	Friends []string `json:"friends"`
}

type Name struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	Honorific  string `json:"honorific"`
}
