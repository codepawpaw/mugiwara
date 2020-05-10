package dto

import (
	models "../models"
)

type Auth struct {
	Account *models.Account `json:"account"`
	User    *models.User    `json:"user"`
	Token   string          `json:"token"`
}
