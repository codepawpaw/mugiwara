package models

type Account struct {
	ID       int64  `json:id`
	Username string `json:username`
	Password string `json:password`
	UserId   int64  `json:userId`
}
