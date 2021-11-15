package repository

import (
	"ananhartanto/lion-backend/helper"
	"ananhartanto/lion-backend/model/domain"
	"context"
	"database/sql"
)

type UserRepositoryImpl struct {

}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users (full_name, username, password, gender) VALUES ($1, $2, $3, $4) RETURNING id, full_name, username, password, gender, created_on"

	result, err := tx.ExecContext(ctx, SQL, user.FullName, user.Username, user.Password, user.Gender)
	helper.Error(err)

	id, err := result.LastInsertId()

	helper.Error(err)

	user.Id = int(id)
	return user
}
