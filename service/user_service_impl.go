package service

import (
	"ananhartanto/lion-backend/helper"
	"ananhartanto/lion-backend/model/domain"
	"ananhartanto/lion-backend/model/web"
	"ananhartanto/lion-backend/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserRegisterRequest) web.UserRegisterResponse {
	err := service.Validate.Struct(request)
	helper.Error(err)

	tx, err := service.DB.Begin()
	helper.Error(err)
	defer helper.CommitOrCallBack(tx)

	user := domain.User{
		FullName: request.FullName,
		Username: request.Username,
		Password: request.Password,
		Gender: request.Gender,
	}

	user = service.UserRepository.Register(ctx, tx, user)

	return helper.ToRegisterResponse(user)
}
