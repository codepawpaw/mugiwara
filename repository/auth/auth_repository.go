package auth

import (
	"context"
	"database/sql"
	"fmt"

	dto "../../dto"
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

func (o *AuthRepository) Create(ctx context.Context, user *models.User, account *models.Account) (dto.Auth, error) {
	userQuery := "Insert users SET name=?, age=?, sex=?"
	accountQuery := "Insert accounts SET username=?, password=?, user_id=?"

	tx, _ := o.Connection.Begin()

	emptyAuthResponse := dto.Auth{}

	// ===========

	userStatement, err := tx.PrepareContext(ctx, userQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	userResponse, err := userStatement.ExecContext(ctx, user.Name, user.Age, user.Sex)
	defer userStatement.Close()
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	// ===========

	accountStatement, err := tx.PrepareContext(ctx, accountQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	userID, _ := userResponse.LastInsertId()
	user.ID = userID
	accountResponse, err := accountStatement.ExecContext(ctx, account.Username, account.Password, userID)
	defer accountStatement.Close()

	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	accountID, _ := accountResponse.LastInsertId()
	account.ID = accountID

	tx.Commit()

	authResponse := dto.Auth{
		Account: account,
		User:    user,
	}

	return authResponse, err
}

func (o *AuthRepository) Update(ctx context.Context, user *models.User, account *models.Account) (dto.Auth, error) {
	userQuery := "Update users SET name=?, age=?, sex=? Where id=?"
	accountQuery := "Update accounts SET username=?, password=? Where id=?"

	tx, _ := o.Connection.Begin()

	emptyAuthResponse := dto.Auth{}

	// ===========

	userStatement, err := tx.PrepareContext(ctx, userQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	userResponse, err := userStatement.ExecContext(ctx, user.Name, user.Age, user.Sex, user.ID)
	defer userStatement.Close()
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	// ===========

	accountStatement, err := tx.PrepareContext(ctx, accountQuery)
	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	fmt.Println(userResponse)

	accountResponse, err := accountStatement.ExecContext(ctx, account.Username, account.Password, account.ID)
	defer accountStatement.Close()

	if err != nil {
		tx.Rollback()

		return emptyAuthResponse, err
	}

	fmt.Println(accountResponse)

	tx.Commit()

	authResponse := dto.Auth{
		Account: account,
		User:    user,
	}

	return authResponse, err
}
