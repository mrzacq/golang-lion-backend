package service

import (
	"ananhartanto/lion-backend/model/web"
	"context"
)

type MusicService interface {
	FindById(ctx context.Context, musicId int) web.MusicResponse
	FindAll(ctx context.Context) []web.MusicResponse
}