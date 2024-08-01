package repository

import "github.com/iniakunhuda/logistik-tani/users/model"

type UserRepository interface {
	Save(user model.User) error
	Update(user model.User) error
	Delete(userId int) error
	FindById(userId int) (*model.User, error)
	FindAll() (users []model.User, err error)
	GetAllByQuery(user model.User) (users []model.User, err error)
	GetOneByQuery(user model.User) (userData model.User, err error)
}
