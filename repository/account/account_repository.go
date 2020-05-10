package account

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitAccountRepository(Connection *sql.DB) pRepo.AccountRepository {
	return &AccountRepository{
		Connection: Connection,
	}
}

type AccountRepository struct {
	Connection *sql.DB
}

func (o *AccountRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Account, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Account, 0)
	for rows.Next() {
		data := new(models.Account)

		err := rows.Scan(
			&data.ID,
			&data.Username,
			&data.Password,
			&data.UserId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *AccountRepository) GetByID(ctx context.Context, id int64) (*models.Account, error) {
	query := "Select id, username, password, user_id From accounts where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Account{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *AccountRepository) GetByUsernameAndPassword(ctx context.Context, username string, password string) (*models.Account, error) {
	query := "Select id, username, password, user_id From accounts where username=? and password=?"

	rows, err := m.fetch(ctx, query, username, password)
	if err != nil {
		return nil, err
	}

	payload := &models.Account{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}
