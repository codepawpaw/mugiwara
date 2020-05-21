package auth

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitAuthRepository(Connection *sql.DB) pRepo.AuthRepository {
	return &AuthRepository{
		Connection: Connection,
	}
}

type AuthRepository struct {
	Connection *sql.DB
}

func (o *AuthRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
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
			&data.DisplayName,
			&data.Email,
			&data.IdToken,
			&data.PhotoUrl,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *AuthRepository) GetByEmail(ctx context.Context, email string) ([]*models.User, error) {
	query := "Select * From users where email=?"

	return m.fetch(ctx, query, email)
}

func (o *AuthRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	userQuery := "Insert users SET name=?, displayName=?, email=?, idToken=?, photoUrl=?"

	tx, _ := o.Connection.Begin()

	// ===========

	userStatement, err := tx.PrepareContext(ctx, userQuery)
	if err != nil {
		tx.Rollback()

		return &models.User{}, err
	}

	userResponse, err := userStatement.ExecContext(ctx, user.Name, user.DisplayName, user.Email, user.IdToken, user.PhotoUrl)
	defer userStatement.Close()
	if err != nil {
		tx.Rollback()

		return &models.User{}, err
	}

	userID, _ := userResponse.LastInsertId()
	user.ID = userID

	tx.Commit()

	return user, err
}

func (o *AuthRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	userQuery := "Update users SET name=?, displayName=?, email=?, idToken=?, photoUrl=? Where email=?"

	tx, _ := o.Connection.Begin()

	// ===========

	userStatement, err := tx.PrepareContext(ctx, userQuery)
	if err != nil {
		tx.Rollback()

		return &models.User{}, err
	}

	_, err = userStatement.ExecContext(ctx, user.Name, user.DisplayName, user.Email, user.IdToken, user.PhotoUrl)
	defer userStatement.Close()
	if err != nil {
		tx.Rollback()

		return &models.User{}, err
	}

	tx.Commit()

	return user, err
}
