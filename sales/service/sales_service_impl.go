package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/sales/model"
	"github.com/iniakunhuda/logistik-tani/sales/repository"
	"github.com/iniakunhuda/logistik-tani/sales/request"
	"github.com/iniakunhuda/logistik-tani/sales/response"
)

type InventoryServiceImpl struct {
	SalesRepository repository.SalesRepository
	Validate        *validator.Validate
}

func NewInventoryServiceImpl(salesRepository repository.SalesRepository, validate *validator.Validate) InventoryService {
	return &InventoryServiceImpl{
		SalesRepository: salesRepository,
		Validate:        validate,
	}
}

func (t *InventoryServiceImpl) Create(produk request.CreateProdukRequest) error {

	// TODO: fix sales
	produkModel := model.Sales{}
	err := t.SalesRepository.Save(produkModel)

	if err != nil {
		return err
	}

	return nil
}

func (t *InventoryServiceImpl) Delete(produkId int) error {
	err := t.SalesRepository.Delete(produkId)
	if err != nil {
		return err
	}
	return nil
}

func (t *InventoryServiceImpl) FindAll() ([]response.SalesResponse, error) {
	result, err := t.SalesRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var sales []response.SalesResponse
	for _, value := range result {
		produk := response.SalesResponse{
			Sales: value,
		}
		sales = append(sales, produk)
	}

	return sales, nil
}

func (t *InventoryServiceImpl) FindById(produkId int) (response.SalesResponse, error) {
	salesData, err := t.SalesRepository.FindById(produkId)
	if err != nil {
		return response.SalesResponse{}, err
	}

	formatResponse := response.SalesResponse{
		Sales: *salesData,
	}
	return formatResponse, nil
}

func (t *InventoryServiceImpl) Update(produkId int, produk request.UpdateUserRequest) error {
	salesData, err := t.SalesRepository.FindById(produkId)
	if err != nil {
		return err
	}

	// update all field
	// TODO: update here

	t.SalesRepository.Update(*salesData)

	return nil
}
