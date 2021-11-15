package repository

import (
	"ananhartanto/lion-backend/model/domain"
	"context"
	"database/sql"
)


type MusicRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Music
	FindById(ctx context.Context, tx *sql.Tx, musicId int) (domain.Music, error)
}