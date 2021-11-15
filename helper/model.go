package helper

import (
	"ananhartanto/lion-backend/model/domain"
	"ananhartanto/lion-backend/model/web"
)

func ToRegisterResponse(user domain.User) web.UserRegisterResponse {
	return web.UserRegisterResponse{
		Id:   user.Id,
		FullName: user.FullName,
		Username: user.Username,
		Gender: user.Gender,
	}
}

func ToMusicResponse(music domain.Music) web.MusicResponse {
	return web.MusicResponse{
		Id:        music.Id,
		Title:     music.Title,
		Duration:  music.Duration,
		Singer:   music.Singer,
		CreatedOn: music.CreatedOn,
	}
}

func ToMusicResponses(musics []domain.Music) []web.MusicResponse {
	var musicResponses []web.MusicResponse
	for _, music := range musics {
		musicResponses = append(musicResponses, ToMusicResponse(music))
	}
	return musicResponses
}
