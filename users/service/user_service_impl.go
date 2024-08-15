package service

import (
	"errors"
	"os"
	"strconv"
	"time"

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
		Name:         user.Name,
		Username:     user.Email,
		Email:        user.Email,
		Password:     hashPassword,
		Alamat:       user.Alamat,
		Telp:         user.Telp,
		Role:         user.Role,
		Saldo:        user.Saldo,
		// LastLogin:    "",
		AlamatKebun:  user.AlamatKebun,
		TotalObat:    user.TotalObat,
		TotalPupuk:   user.TotalPupuk,
		TotalBibit:   user.TotalBibit,
		TotalTebu:    user.TotalTebu,
		LuasLahan:    user.LuasLahan,
		Token:        nil,
		TokenExpired: nil,
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

	formatResponse := response.UserResponse{
		User: *userData,
	}
	return formatResponse, nil
}

func (t *UserServiceImpl) Update(userId int, user request.UpdateUserRequest) error {
	userData, err := t.UserRepository.FindById(userId)
	if err != nil {
		return err
	}

	// update all field

	if user.Name != "" {
		userData.Name = user.Name
	}
	if user.Email != "" {
		userData.Email = user.Email
	}
	if user.Alamat != "" {
		userData.Alamat = user.Alamat
	}
	if user.Telp != "" {
		userData.Telp = user.Telp
	}
	if user.Role != "" {
		userData.Role = user.Role
	}
	if user.Saldo != 0 {
		userData.Saldo = user.Saldo
	}

	// if user.LastLogin != "" {
	// 	userData.LastLogin = user.LastLogin
	// }
	if user.AlamatKebun != "" {
		userData.AlamatKebun = user.AlamatKebun
	}
	if user.TotalObat != 0 {
		userData.TotalObat = user.TotalObat
	}
	if user.TotalPupuk != 0 {
		userData.TotalPupuk = user.TotalPupuk
	}
	if user.TotalBibit != 0 {
		userData.TotalBibit = user.TotalBibit
	}
	if user.TotalTebu != 0 {
		userData.TotalTebu = user.TotalTebu
	}
	if user.LuasLahan != 0 {
		userData.LuasLahan = user.LuasLahan
	}

	if user.Token != nil {
		userData.Token = user.Token
	}

	if user.TokenExpired != nil {
		tokenExpired, err := time.Parse(time.RFC3339, *user.TokenExpired)
		if err != nil {
			return err
		}
		userData.TokenExpired = &tokenExpired
	}

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

func (t *UserServiceImpl) Login(email string, password string) (response.UserResponse, error) {
	userData, err := t.UserRepository.GetOneByQuery(model.User{Email: email})
	if err != nil {
		return response.UserResponse{}, err
	}

	checkPassword := util.VerifyPassword(password, userData.Password)
	if !checkPassword {
		return response.UserResponse{}, errors.New("password not match")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return response.UserResponse{}, errors.New("JWT_SECRET is not set")
	}

	jwt := util.NewJWT(secret)
	tokenExpired := time.Duration(720) * time.Hour

	token, err := jwt.CreateToken(strconv.Itoa(int(userData.ID)), userData.Email, tokenExpired)
	if err != nil {
		return response.UserResponse{}, err
	}

	tokenExpiredTime := time.Now().AddDate(0, 0, 30)
	userData.TokenExpired = &tokenExpiredTime
	userData.Token = &token

	// userData.LastLogin = util.GetTimeNow()

	err = t.UserRepository.Update(userData)
	if err != nil {
		return response.UserResponse{}, err
	}

	formatResponse := response.UserResponse{
		User: userData,
	}
	return formatResponse, nil
}

func (t *UserServiceImpl) Profile(token string) (*response.UserResponse, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET is not set")
	}

	jwt := util.NewJWT(secret)
	userID, err := jwt.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	userIDToInt, err := strconv.Atoi(userID)
	result, err := t.UserRepository.GetOneByQuery(model.User{ID: uint(userIDToInt)})
	if err != nil {
		return nil, err
	}

	formatResponse := &response.UserResponse{
		User: result,
	}

	return formatResponse, nil
}
