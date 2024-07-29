package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/users/model"
	"github.com/iniakunhuda/logistik-tani/users/repository"
	"github.com/iniakunhuda/logistik-tani/users/request"
	"github.com/iniakunhuda/logistik-tani/users/response"
	"github.com/iniakunhuda/logistik-tani/users/util"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (t *UserServiceImpl) Create(user request.CreateUserRequest) error {

	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	userModel := model.User{
		Name:        user.Name,
		Username:    user.Email,
		Email:       user.Email,
		Password:    hashPassword,
		Alamat:      user.Alamat,
		Telp:        user.Telp,
		Role:        user.Role,
		Saldo:       user.Saldo,
		LastLogin:   "",
		AlamatKebun: user.AlamatKebun,
		TotalObat:   user.TotalObat,
		TotalPupuk:  user.TotalPupuk,
		TotalBibit:  user.TotalBibit,
		TotalTebu:   user.TotalTebu,
		LuasLahan:   user.LuasLahan,
	}
	err = t.UserRepository.Save(userModel)

	if err != nil {
		return err
	}

	return nil
}

func (t *UserServiceImpl) Delete(userId int) error {
	err := t.UserRepository.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}

func (t *UserServiceImpl) FindAll() ([]response.UserResponse, error) {
	result, err := t.UserRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			User: value,
		}
		users = append(users, user)
	}

	return users, nil
}

func (t *UserServiceImpl) FindById(userId int) (response.UserResponse, error) {
	userData, err := t.UserRepository.FindById(userId)
	if err != nil {
		return response.UserResponse{}, err
	}

	tagResponse := response.UserResponse{
		User: *userData,
	}
	return tagResponse, nil
}

func (t *UserServiceImpl) Update(userId int, user request.UpdateUserRequest) error {
	userData, err := t.UserRepository.FindById(userId)
	if err != nil {
		return err
	}
	userData.Name = user.Name
	t.UserRepository.Update(*userData)

	return nil
}

func (t *UserServiceImpl) FindByRole(role string) ([]response.UserResponse, error) {
	result, err := t.UserRepository.GetAllByQuery(model.User{Role: role})
	if err != nil {
		return nil, err
	}

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			User: value,
		}
		users = append(users, user)
	}

	return users, nil
}
