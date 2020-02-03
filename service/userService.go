package service

import (
	"go-web-learn/model"
	"go-web-learn/response"
)

func Login(user model.User) response.LoginResponse {

	return response.LoginResponse{user.Name + user.Password, 1001}

}
