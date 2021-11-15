package service

import (
	"ananhartanto/lion-backend/model/web"
	"context"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) web.UserRegisterResponse
}