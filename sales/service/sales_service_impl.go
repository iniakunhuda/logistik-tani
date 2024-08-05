package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/iniakunhuda/logistik-tani/sales/model"
	"github.com/iniakunhuda/logistik-tani/sales/repository"
	"github.com/iniakunhuda/logistik-tani/sales/request"
	"github.com/iniakunhuda/logistik-tani/sales/response"
)

type InventoryServiceImpl struct {
	SalesRepository       repository.SalesRepository
	SalesRepositoryDetail repository.SalesDetailRepository
	Validate              *validator.Validate
}

func NewInventoryServiceImpl(salesRepository repository.SalesRepository, validate *validator.Validate) SalesService {
	return &InventoryServiceImpl{
		SalesRepository: salesRepository,
		Validate:        validate,
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

	// Validate the request
	// Check is product exist

	

	noInv, err := t.GenerateNoInvoice()
	if err != nil {
		return err
	}
	totalHargaSales := 0
	salesModel := model.Sales{
		NoInvoice:        noInv,
		IDPembeli:        sales.IDPembeli,
		IDPenjual:        sales.IDPenjual,
		Tanggal:          sales.Tanggal,
		Status:           "open",
		TotalHarga:       0,
		IsPurchasedByIGM: false,
		InvPurchasedIGM:  nil,
	}

	salesDetailModel := []model.SalesDetail{}
	for _, value := range sales.Produk {
		salesDetailModel = append(salesDetailModel, model.SalesDetail{
			IDProduk:   value.IDProduk,
			Jenis:      value.Jenis,
			Harga:      int(value.Harga),
			Qty:        int(value.Qty),
			TotalHarga: int(value.Harga) * int(value.Qty),
			Tanggal:    value.Tanggal,
		})

		totalHargaSales += int(value.Harga) * int(value.Qty)
	}

	salesModel.TotalHarga = totalHargaSales
	err = t.SalesRepository.Save(salesModel, salesDetailModel)
	if err != nil {
		return err
	}


	// Trigger stok update

	return nil
}

func (t *InventoryServiceImpl) Delete(salesId int) error {
	err := t.SalesRepository.Delete(salesId)
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
		newSalesDetail := response.SalesResponse{
			Sales: value,
		}
		sales = append(sales, newSalesDetail)
	}

	return sales, nil
}

func (t *InventoryServiceImpl) FindById(salesId int) (response.SalesResponse, error) {
	salesData, err := t.SalesRepository.FindById(salesId)
	if err != nil {
		return response.SalesResponse{}, err
	}

	formatResponse := response.SalesResponse{
		Sales: *salesData,
	}
	return formatResponse, nil
}
