package service

import (
	"github.com/iniakunhuda/logistik-tani/users/request"
	"github.com/iniakunhuda/logistik-tani/users/response"
)

type UserService interface {
	Create(user request.CreateUserRequest) error
	Update(userId int, user request.UpdateUserRequest) error
	Delete(userId int) error
	FindById(userId int) (response.UserResponse, error)
	FindAll() ([]response.UserResponse, error)
	FindByRole(role string) ([]response.UserResponse, error)
}
