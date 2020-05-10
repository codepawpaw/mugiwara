package models

import "time"

type Incident struct {
	ID          int64     `json:id`
	CityName    string    `json:cityName`
	Province    string    `json:province`
	Nation      string    `json:"nation"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Lat         string    `json:"lat"`
	Lang        string    `json:"lang"`
	UserId      int64     `json:"userID"`
}
