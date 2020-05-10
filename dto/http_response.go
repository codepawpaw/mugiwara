package dto

type HttpResponse struct {
	Data         string `json:"data"`
	ErrorMessage string `json:"errorMessage"`
	Status       int    `json:"status"`
}
