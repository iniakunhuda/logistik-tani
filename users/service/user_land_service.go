package service

import (
	"github.com/iniakunhuda/logistik-tani/users/model"
	"github.com/iniakunhuda/logistik-tani/users/request"
	"github.com/iniakunhuda/logistik-tani/users/response"
)

type UserLandService interface {
	Create(land request.CreateUserLandRequest) error
	Update(landId int, land request.UpdateUserLandRequest) error
	Delete(landId int) error
	FindById(landId int) (response.UserLandResponse, error)
	FindAll(land *model.UserLand) ([]response.UserLandResponse, error)
}
