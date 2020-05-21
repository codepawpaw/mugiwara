package repsitory

import (
	"context"

	models "../models"
)

type AuthRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	GetByEmail(ctx context.Context, email string) ([]*models.User, error)
}

type IncidentRepository interface {
	Create(ctx context.Context, incident *models.Incident) (*models.Incident, error)
	GetByCityName(ctx context.Context, cityName string) ([]*models.Incident, error)
	Update(ctx context.Context, p *models.Incident) (*models.Incident, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
