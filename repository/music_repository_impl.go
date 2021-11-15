package repository

import (
	"ananhartanto/lion-backend/helper"
	"ananhartanto/lion-backend/model/domain"
	"context"
	"database/sql"
	"errors"
)

type MusicRepositoryImpl struct {}

func NewMusicRepository() MusicRepository {
	return &MusicRepositoryImpl{}
}

func (repository *MusicRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, musicId int) (domain.Music, error) {
	SQL := "SELECT * FROM music WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, musicId)
	helper.Error(err)
	music := domain.Music{}
	if rows.Next() {
		err := rows.Scan(&music.Id, &music.Title, &music.Duration, &music.Singer, &music.CreatedOn)
		helper.Error(err)
		return music, nil
	} else {
		return music, errors.New("music not found")
	}
}

func (repository *MusicRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Music {
	SQL := "SELECT * FROM music"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.Error(err)
	var allMusic []domain.Music
	for rows.Next() {
		music := domain.Music{}
		err := rows.Scan(&music.Id, &music.Title, &music.Duration, &music.Singer, &music.CreatedOn)
		helper.Error(err)
		allMusic = append(allMusic, music)
	}
	return allMusic
}