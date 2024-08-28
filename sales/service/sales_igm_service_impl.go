package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/sales/model"
	"github.com/iniakunhuda/logistik-tani/sales/repository"
	"github.com/iniakunhuda/logistik-tani/sales/repository/remote"
	"github.com/iniakunhuda/logistik-tani/sales/request"
	"github.com/iniakunhuda/logistik-tani/sales/response"
)

type SalesIgmServiceImpl struct {
	TokenAuth                 string
	SalesIgmRepository        repository.SalesIgmRepository
	UserRemoteRepository      remote.UserRemoteRepository
	InventoryRemoteRepository remote.InventoryRemoteRepository
	Validate                  *validator.Validate
}

func NewSalesIgmServiceImpl(salesRepository repository.SalesIgmRepository, validate *validator.Validate) SalesIgmService {
	return &SalesIgmServiceImpl{
		SalesIgmRepository:        salesRepository,
		UserRemoteRepository:      remote.NewUserRemoteRepositoryImpl(),
		InventoryRemoteRepository: remote.NewInventoryRemoteRepositoryImpl(""), // TODO: set bearer token
		Validate:                  validate,
	}
}

func (t *SalesIgmServiceImpl) GenerateNoInvoice() (string, error) {
	sales, err := t.SalesIgmRepository.FindLastRow()

	lastInv := 0
	if sales != nil {
		parts := strings.Split(sales.NoInvoice, "-")
		if len(parts) == 2 {
			lastInv, err = strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
		}
	}

	noInv := ""
	if lastInv+1 < 10 {
		noInv = fmt.Sprintf("SALES-000%d", lastInv+1)
	} else if lastInv+1 < 100 {
		noInv = fmt.Sprintf("SALES-00%d", lastInv+1)
	} else if lastInv+1 < 1000 {
		noInv = fmt.Sprintf("SALES-0%d", lastInv+1)
	}

	return noInv, nil
}

func (t *SalesIgmServiceImpl) Create(sales request.CreateSalesIgmRequest) error {

	noInv, err := t.GenerateNoInvoice()
	if err != nil {
		return err
	}
	salesModel := model.SalesIgm{
		NoInvoice:      noInv,
		IDSeller:       int(sales.IDSeller),
		SalesDate:      sales.SalesDate,
		Status:         "open",
		TotalPrice:     float64(sales.Product.Price) * float64(sales.Product.Qty),
		IDProductOwner: sales.Product.IDProduct,
		Qty:            sales.Product.Qty,
		Price:          float64(sales.Product.Price),
		Note:           sales.Product.Catatan,
		BuyerName:      sales.BuyerName,
		BuyerTelp:      sales.BuyerTelp,
		BuyerAddress:   sales.BuyerAddress,
	}

	// Inventory Service: Trigger stok update
	err = t.InventoryRemoteRepository.UpdateReduceStok(strconv.Itoa(int(sales.Product.IDProduct)), strconv.Itoa(int(sales.Product.Qty)))
	if err != nil {
		return err
	}

	err = t.SalesIgmRepository.Save(salesModel)
	if err != nil {
		return err
	}

	return nil
}

func (t *SalesIgmServiceImpl) Update(saleId int, userId int, sales request.UpdateSalesRequest) error {
	// Validate the request
	salesData, err := t.SalesIgmRepository.GetOneByQuery(model.SalesIgm{IDSeller: userId, ID: uint(saleId)})
	if err != nil {
		return err
	}

	if salesData.Status == "done" {
		return errors.New("Sales already closed")
	}

	statusArr := []string{"open", "pending", "done", "cancel"}
	isValid := false
	for _, value := range statusArr {
		if value == sales.Status {
			isValid = true
			break
		}
	}

	if !isValid {
		return errors.New("Invalid status. Available status: (open, pending, done, cancel)")
	}

	// Update the sales
	err = t.SalesIgmRepository.Update(model.SalesIgm{ID: uint(saleId), Status: sales.Status})
	if err != nil {
		return err
	}

	return nil
}

func (t *SalesIgmServiceImpl) Delete(salesId int) error {
	err := t.SalesIgmRepository.Delete(salesId)
	if err != nil {
		return err
	}
	return nil
}

func (t *SalesIgmServiceImpl) FindAll(sale *model.SalesIgm) ([]response.SalesIgmResponse, error) {
	result, err := t.SalesIgmRepository.GetAllByQuery(*sale)

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

	// get list of products
	products, err := t.InventoryRemoteRepository.GetAll()
	if err != nil {
		return nil, err
	}
	listProducts := map[uint]response.ProductResponse{}
	for _, value := range products {
		listProducts[value.ID] = value
	}

	var sales []response.SalesIgmResponse
	for _, value := range result {

		newSalesDetail := response.SalesIgmResponse{
			SalesIgm:     value,
			SellerDetail: listUser[uint(value.IDSeller)],
		}
		sales = append(sales, newSalesDetail)
	}

	return sales, nil
}

func (t *SalesIgmServiceImpl) FindById(salesId int, userId int) (response.SalesIgmResponse, error) {
	salesData, err := t.SalesIgmRepository.GetOneByQuery(model.SalesIgm{IDSeller: userId, ID: uint(salesId)})
	if err != nil {
		return response.SalesIgmResponse{}, err
	}

	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return response.SalesIgmResponse{}, err
	}
	listUser := map[uint]response.UserResponse{}
	for _, value := range users {
		listUser[value.ID] = value
	}

	// get list of products
	products, err := t.InventoryRemoteRepository.GetAll()
	if err != nil {
		return response.SalesIgmResponse{}, err
	}
	listProducts := map[uint]response.ProductResponse{}
	for _, value := range products {
		listProducts[value.ID] = value
	}

	formatResponse := response.SalesIgmResponse{
		SalesIgm:     salesData,
		SellerDetail: listUser[uint(salesData.IDSeller)],
	}
	return formatResponse, nil
}
