package repository

import "github.com/iniakunhuda/logistik-tani/users/model"

type UserLandRepository interface {
	Save(land model.UserLand) error
	Update(land model.UserLand) error
	Delete(landId int) error
	FindById(landId int) (*model.UserLand, error)
	FindAll() (lands []model.UserLand, err error)
	GetAllByQuery(land model.UserLand) (lands []model.UserLand, err error)
	GetOneByQuery(land model.UserLand) (userData model.UserLand, err error)
}
