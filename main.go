package main

import (
	"ananhartanto/lion-backend/app"
	"ananhartanto/lion-backend/controller"
	"ananhartanto/lion-backend/exception"
	"ananhartanto/lion-backend/helper"
	"ananhartanto/lion-backend/repository"
	"ananhartanto/lion-backend/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	
	musicRepository := repository.NewMusicRepository()
	musicService := service.NewMusicService(musicRepository, db)
	musicController := controller.NewMusicController(musicService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()


	router.POST("/user/register", userController.Register)

	router.GET("/musics", musicController.FindAll)
	router.GET("/musics/:musicId", musicController.FindById)
	
	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.Error(err)
}