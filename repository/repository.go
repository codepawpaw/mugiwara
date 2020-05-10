package repsitory

import (
	"context"

	dto "../dto"
	models "../models"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
}

type AccountRepository interface {
	GetByID(ctx context.Context, id int64) (*models.Account, error)
	GetByUsernameAndPassword(ctx context.Context, username string, password string) (*models.Account, error)
}

type AuthRepository interface {
	Create(ctx context.Context, user *models.User, account *models.Account) (dto.Auth, error)
	Update(ctx context.Context, user *models.User, account *models.Account) (dto.Auth, error)
}

type IncidentRepository interface {
	Create(ctx context.Context, incident *models.Incident) (*models.Incident, error)
	GetByCityName(ctx context.Context, cityName string) ([]*models.Incident, error)
	Update(ctx context.Context, p *models.Incident) (*models.Incident, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
