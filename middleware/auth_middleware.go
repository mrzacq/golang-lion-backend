package middleware

import (
	"ananhartanto/lion-backend/helper"
	"ananhartanto/lion-backend/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var token = "token_rahasia"
	if token == request.Header.Get("ACCESS_TOKEN") {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
	 
}