package controller

import (
	"ananhartanto/lion-backend/helper"
	"ananhartanto/lion-backend/model/web"
	"ananhartanto/lion-backend/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MusicControllerImpl struct {
	MusicService service.MusicService
}

func NewMusicController(musicService service.MusicService) MusicController {
	return &MusicControllerImpl{
		MusicService: musicService,
	}
}

func (controller *MusicControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	musicId := params.ByName("musicId")
	id, err := strconv.Atoi(musicId)
	helper.Error(err)
	musicResponse := controller.MusicService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: musicResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MusicControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	musicResponses := controller.MusicService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: musicResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}