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
	FindAllExclude(exclude string) ([]response.UserResponse, error)

	AddSaldoUser(userId int, saldo int) error

	Login(email string, password string) (response.UserResponse, error)
	Profile(token string) (*response.UserResponse, error)
}
