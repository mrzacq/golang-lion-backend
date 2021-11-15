package repository

import (
	"ananhartanto/lion-backend/model/domain"
	"context"
	"database/sql"
)


type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}