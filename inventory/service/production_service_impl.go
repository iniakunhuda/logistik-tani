package service

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/repository"
	"github.com/iniakunhuda/logistik-tani/inventory/repository/remote"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type ProductionServiceImpl struct {
	TokenAuth            string
	ProductionRepository repository.ProductionRepository
	InventoryRepository  repository.ProductOwnerRepository
	UserRemoteRepository remote.UserRemoteRepository
	Validate             *validator.Validate
}

func NewProductionServiceImpl(productionRepo repository.ProductionRepository, inventoryRepo repository.ProductOwnerRepository, validate *validator.Validate) ProductionService {
	return &ProductionServiceImpl{
		ProductionRepository: productionRepo,
		InventoryRepository:  inventoryRepo,
		UserRemoteRepository: remote.NewUserRemoteRepositoryImpl(),
		Validate:             validate,
	}
}

func (t *ProductionServiceImpl) Create(production request.CreateProductionRequest) error {
	// check land owner
	landOwner, err := t.UserRemoteRepository.GetLandByUserId(strconv.Itoa(production.IDUser), strconv.Itoa(production.IDLand))
	if err != nil {
		return err
	}

	if landOwner.Data == nil {
		return errors.New("User tidak memiliki lahan")
	}

	productionModel := model.Production{
		IDUser:       production.IDUser,
		IDUserLand:   production.IDLand,
		Title:        production.Title,
		DateStart:    production.DateStart,
		TotalPanenKg: production.TotalPanenKg,
		Status:       "pending",
	}

	// get date month from date start
	productionModel.DateMonth = int(production.DateStart.Month())
	productionModel.DateYear = production.DateStart.Year()

	if !production.DateEnd.IsZero() {
		productionModel.DateEnd = &production.DateEnd
	}

	err = t.ProductionRepository.Save(productionModel)
	if err != nil {
		return err
	}

	return nil
}

func (t *ProductionServiceImpl) Update(saleId int, production request.UpdateProductionRequest) error {
	// Validate the request
	productionDb, err := t.ProductionRepository.GetOneByQuery(model.Production{ID: uint(saleId)})
	if err != nil {
		return err
	}

	if productionDb.Status == "done" {
		return errors.New("Panen sudah selesai")
	}

	statusArr := []string{"pending", "done"}
	isValid := false
	for _, value := range statusArr {
		if value == production.Status {
			isValid = true
			break
		}
	}

	if !isValid {
		return errors.New("Status tidak ditemukan. Status yang tersedia: (pending, done)")
	}

	if !production.DateStart.IsZero() {
		productionDb.DateStart = production.DateStart
	}

	if !production.DateEnd.IsZero() {
		productionDb.DateEnd = &production.DateEnd
	}

	if production.TotalPanenKg != 0 {
		productionDb.TotalPanenKg = int(production.TotalPanenKg)
	}

	if production.Status != "" {
		productionDb.Status = production.Status
	}

	// Update the production
	err = t.ProductionRepository.Update(productionDb)
	if err != nil {
		return err
	}

	return nil
}

func (t *ProductionServiceImpl) Delete(productionId int) error {
	err := t.ProductionRepository.Delete(productionId)
	if err != nil {
		return err
	}
	return nil
}

func (t *ProductionServiceImpl) FindAll(sale *model.Production) ([]response.ProductionResponse, error) {
	result, err := t.ProductionRepository.GetAllByQuery(*sale)

	if err != nil {
		return nil, err
	}

	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return nil, err
	}
	listUser := map[uint]response.UserResponse{}
	for _, value := range users {
		listUser[value.ID] = value
	}

	var production []response.ProductionResponse
	for _, value := range result {
		newproductionDetail := response.ProductionResponse{
			Production: value,
			UserDetail: listUser[uint(value.IDUser)],
		}
		production = append(production, newproductionDetail)
	}

	return production, nil
}

func (t *ProductionServiceImpl) FindById(productionId int) (response.ProductionResponse, error) {
	productionDb, err := t.ProductionRepository.GetOneByQuery(model.Production{ID: uint(productionId)})
	if err != nil {
		return response.ProductionResponse{}, err
	}

	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return response.ProductionResponse{}, err
	}
	listUser := map[uint]response.UserResponse{}
	for _, value := range users {
		listUser[value.ID] = value
	}

	formatResponse := response.ProductionResponse{
		Production: productionDb,
		UserDetail: listUser[uint(productionDb.IDUser)],
	}
	return formatResponse, nil
}
