package repository

import "github.com/iniakunhuda/logistik-tani/users/model"

type UserLandRepository interface {
	Save(user model.UserLand) error
	Update(user model.UserLand) error
	Delete(userId int) error
	FindById(userId int) (*model.UserLand, error)
	FindAll() (users []model.UserLand, err error)
	GetAllByQuery(user model.UserLand) (users []model.UserLand, err error)
	GetOneByQuery(user model.UserLand) (userData model.UserLand, err error)
}
