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

type InventoryServiceImpl struct {
	TokenAuth                 string
	SalesRepository           repository.SalesRepository
	UserRemoteRepository      remote.UserRemoteRepository
	InventoryRemoteRepository remote.InventoryRemoteRepository
	Validate                  *validator.Validate
}

func NewInventoryServiceImpl(salesRepository repository.SalesRepository, validate *validator.Validate) SalesService {
	return &InventoryServiceImpl{
		SalesRepository:           salesRepository,
		UserRemoteRepository:      remote.NewUserRemoteRepositoryImpl(),
		InventoryRemoteRepository: remote.NewInventoryRemoteRepositoryImpl(""), // TODO: set bearer token
		Validate:                  validate,
	}
}

func (t *InventoryServiceImpl) GenerateNoInvoice() (string, error) {
	sales, err := t.SalesRepository.FindLastRow()
	if err != nil {
		return "", err
	}

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

func (t *InventoryServiceImpl) Create(sales request.CreateSalesRequest) error {

	noInv, err := t.GenerateNoInvoice()
	if err != nil {
		return err
	}
	totalHargaSales := 0
	salesModel := model.Sales{
		NoInvoice:  noInv,
		IDBuyer:    int(sales.IDBuyer),
		IDSeller:   int(sales.IDSeller),
		SalesDate:  sales.SalesDate,
		Status:     "open",
		TotalPrice: 0,
	}

	pembeliDetail, err := t.UserRemoteRepository.Find(strconv.Itoa(int(sales.IDBuyer)))
	if err != nil {
		return err
	}

	// Check stok produk
	for _, value := range sales.Products {
		// Inventory Service: Check product stok
		inventoryDetail, err := t.InventoryRemoteRepository.GetDetail(strconv.Itoa(int(value.IDProduct)))
		if err != nil {
			return err
		}
		if inventoryDetail.Stock < uint(value.Qty) {
			return fmt.Errorf("stok produk %s tidak mencukupi. Stok tersedia: %d", inventoryDetail.Name, inventoryDetail.Stock)
		}
	}

	salesDetailModel := []model.SalesDetail{}
	for _, value := range sales.Products {

		// Inventory Service: Check product stok
		inventoryDetail, err := t.InventoryRemoteRepository.GetDetail(strconv.Itoa(int(value.IDProduct)))
		if err != nil {
			return err
		}

		salesDetailModel = append(salesDetailModel, model.SalesDetail{
			IDProductOwner: int(inventoryDetail.ID),
			Price:          float64(value.Price),
			Qty:            int(value.Qty),
			Subtotal:       float64(value.Price) * float64(value.Qty),
		})

		// Inventory Service: Trigger stok update
		err = t.InventoryRemoteRepository.UpdateReduceStok(strconv.Itoa(int(value.IDProduct)), strconv.Itoa(int(value.Qty)))
		if err != nil {
			return err
		}

		// Inventory Service: If pembeli is petani, then create inventory_petani
		if pembeliDetail.Role == "petani" {
			err = t.InventoryRemoteRepository.AutoCreateProdukPetani(inventoryDetail, uint(value.Qty), uint(sales.IDBuyer))
			if err != nil {
				return errors.New("AutoCreateProdukPetani: " + err.Error())
			}
		}

		totalHargaSales += int(value.Price) * int(value.Qty)
	}

	salesModel.TotalPrice = float64(totalHargaSales)
	err = t.SalesRepository.Save(salesModel, salesDetailModel)
	if err != nil {
		return err
	}

	return nil
}

func (t *InventoryServiceImpl) Update(saleId int, userId int, sales request.UpdateSalesRequest) error {
	// Validate the request
	salesData, err := t.SalesRepository.GetOneByQuery(model.Sales{IDSeller: userId, ID: uint(saleId)})
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
	err = t.SalesRepository.Update(model.Sales{ID: uint(saleId), Status: sales.Status})
	if err != nil {
		return err
	}

	return nil
}

func (t *InventoryServiceImpl) Delete(salesId int) error {
	err := t.SalesRepository.Delete(salesId)
	if err != nil {
		return err
	}
	return nil
}

func (t *InventoryServiceImpl) FindAll(sale *model.Sales) ([]response.SalesResponse, error) {
	result, err := t.SalesRepository.GetAllByQuery(*sale)

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

	var sales []response.SalesResponse
	for _, value := range result {

		// get product detail
		var salesDetail = value.SalesDetail
		for i, salesDetailItem := range salesDetail {
			productDetail := listProducts[uint(salesDetailItem.IDProductOwner)]
			salesDetail[i].Name = productDetail.Name
			salesDetail[i].Description = productDetail.Description
		}
		value.SalesDetail = salesDetail

		newSalesDetail := response.SalesResponse{
			Sales:        value,
			SellerDetail: listUser[uint(value.IDSeller)],
			BuyerDetail:  listUser[uint(value.IDBuyer)],
		}
		sales = append(sales, newSalesDetail)
	}

	return sales, nil
}

func (t *InventoryServiceImpl) FindById(salesId int, userId int) (response.SalesResponse, error) {
	salesData, err := t.SalesRepository.GetOneByQuery(model.Sales{IDSeller: userId, ID: uint(salesId)})
	if err != nil {
		return response.SalesResponse{}, err
	}

	// get user service
	users, err := t.UserRemoteRepository.GetAll()
	if err != nil {
		return response.SalesResponse{}, err
	}
	listUser := map[uint]response.UserResponse{}
	for _, value := range users {
		listUser[value.ID] = value
	}

	// get list of products
	products, err := t.InventoryRemoteRepository.GetAll()
	if err != nil {
		return response.SalesResponse{}, err
	}
	listProducts := map[uint]response.ProductResponse{}
	for _, value := range products {
		listProducts[value.ID] = value
	}

	// get product detail
	var salesDetail = salesData.SalesDetail
	for i, salesDetailItem := range salesDetail {
		productDetail := listProducts[uint(salesDetailItem.IDProductOwner)]
		salesDetail[i].Name = productDetail.Name
		salesDetail[i].Description = productDetail.Description
	}
	salesData.SalesDetail = salesDetail

	formatResponse := response.SalesResponse{
		Sales:        salesData,
		SellerDetail: listUser[uint(salesData.IDSeller)],
		BuyerDetail:  listUser[uint(salesData.IDBuyer)],
	}
	return formatResponse, nil
}
