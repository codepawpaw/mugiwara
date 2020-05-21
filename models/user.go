package models

type User struct {
	ID          int64  `json:id`
	Name        string `json:name`
	DisplayName string `json:displayName`
	Email       string `json:email`
	IdToken     string `json:idToken`
	PhotoUrl    string `json:photoUrl`
	JWTToken    string `json:jwtToken`
}
