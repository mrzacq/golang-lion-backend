package controller

import (
	"ananhartanto/lion-backend/helper"
	"ananhartanto/lion-backend/model/web"
	"ananhartanto/lion-backend/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRegisterRequest := web.UserRegisterRequest{}
	helper.ReadFromRequestBody(request, &userRegisterRequest)

	userRegisterResponse := controller.UserService.Register(request.Context(), userRegisterRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   userRegisterResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
