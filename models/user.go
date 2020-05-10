package models

type User struct {
	ID   int64  `json:id`
	Name string `json:name`
	Sex  string `json:sex`
	Age  string `json:age`
}
