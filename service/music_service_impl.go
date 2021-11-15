package service

import (
	"ananhartanto/lion-backend/exception"
	"ananhartanto/lion-backend/helper"
	"ananhartanto/lion-backend/model/web"
	"ananhartanto/lion-backend/repository"
	"context"
	"database/sql"
)

type MusicServiceImpl struct {
	MusicRepository repository.MusicRepository
	DB *sql.DB
}

func NewMusicService(musicRepository repository.MusicRepository, DB *sql.DB) MusicService {
	return &MusicServiceImpl{
		MusicRepository: musicRepository,
		DB: DB,
	}
}

func (service *MusicServiceImpl) FindById(ctx context.Context, musicId int) web.MusicResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	defer helper.CommitOrCallBack(tx)
	music, err := service.MusicRepository.FindById(ctx, tx, musicId)
	helper.Error(err)
	return helper.ToMusicResponse(music)
}

func (service *MusicServiceImpl) FindAll(ctx context.Context) []web.MusicResponse {
	tx, err := service.DB.Begin()
	helper.Error(err)
	defer helper.CommitOrCallBack(tx)
	allMusic := service.MusicRepository.FindAll(ctx, tx)
	return helper.ToMusicResponses(allMusic)
}