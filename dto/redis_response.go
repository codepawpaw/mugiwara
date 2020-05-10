package dto

type RedisResponse struct {
	Type   string `json:"type"`
	Data   string `json:"data"`
	Status int    `json:"status"`
}
