package models

type Auth struct {
	User    User    `json:user`
	Account Account `json:account`
}
