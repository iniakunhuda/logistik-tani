package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/users/model"
	"github.com/iniakunhuda/logistik-tani/users/repository"
	"github.com/iniakunhuda/logistik-tani/users/request"
	"github.com/iniakunhuda/logistik-tani/users/response"
)

type UserLandServiceImpl struct {
	UserLandRepository repository.UserLandRepository
	Validate           *validator.Validate
}

func NewUserLandServiceImpl(userRepository repository.UserLandRepository, validate *validator.Validate) UserLandService {
	return &UserLandServiceImpl{
		UserLandRepository: userRepository,
		Validate:           validate,
	}
}

func (t *UserLandServiceImpl) Create(land request.CreateUserLandRequest) error {
	userModel := model.UserLand{
		IDUser:      land.IDUser,
		LandName:    land.LandName,
		LandAddress: land.LandAddress,
		LandArea:    land.LandArea,
		TotalObat:   land.TotalObat,
		TotalPupuk:  land.TotalPupuk,
		TotalBibit:  land.TotalBibit,
		TotalTebu:   land.TotalTebu,
	}
	err := t.UserLandRepository.Save(userModel)

	if err != nil {
		return err
	}

	return nil
}

func (t *UserLandServiceImpl) Delete(landId int) error {
	err := t.UserLandRepository.Delete(landId)
	if err != nil {
		return err
	}
	return nil
}

func (t *UserLandServiceImpl) FindAll(land *model.UserLand) ([]response.UserLandResponse, error) {
	result, err := t.UserLandRepository.GetAllByQuery(*land)

	if err != nil {
		return nil, err
	}

	var lands []response.UserLandResponse
	for _, value := range result {
		user := response.UserLandResponse{
			UserLand: value,
		}
		lands = append(lands, user)
	}

	return lands, nil
}

func (t *UserLandServiceImpl) FindById(landId int) (response.UserLandResponse, error) {
	landData, err := t.UserLandRepository.FindById(landId)
	if err != nil {
		return response.UserLandResponse{}, err
	}

	formatResponse := response.UserLandResponse{
		UserLand: *landData,
	}
	return formatResponse, nil
}

func (t *UserLandServiceImpl) Update(landId int, user request.UpdateUserLandRequest) error {
	landData, err := t.UserLandRepository.FindById(landId)
	if err != nil {
		return err
	}

	// update all field
	landData.IDUser = user.IDUser

	// update only if not empty
	if user.LandName != "" {
		landData.LandName = user.LandName
	}
	if user.LandAddress != "" {
		landData.LandAddress = user.LandAddress
	}
	if user.LandArea != 0 {
		landData.LandArea = user.LandArea
	}
	if user.TotalObat != 0 {
		landData.TotalObat = user.TotalObat
	}
	if user.TotalPupuk != 0 {
		landData.TotalPupuk = user.TotalPupuk
	}
	if user.TotalBibit != 0 {
		landData.TotalBibit = user.TotalBibit
	}
	if user.TotalTebu != 0 {
		landData.TotalTebu = user.TotalTebu
	}

	t.UserLandRepository.Update(*landData)

	return nil
}
