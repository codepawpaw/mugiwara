package user

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitUserRepository(Connection *sql.DB) pRepo.UserRepository {
	return &UserRepository{
		Connection: Connection,
	}
}

type UserRepository struct {
	Connection *sql.DB
}

func (o *UserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Sex,
			&data.Age,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *UserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := "Select * From users where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.User{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}
