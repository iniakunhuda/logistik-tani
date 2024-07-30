package service

import (
	"github.com/iniakunhuda/logistik-tani/sales/request"
	"github.com/iniakunhuda/logistik-tani/sales/response"
)

type InventoryService interface {
	Create(sale request.CreateProdukRequest) error
	Update(saleId int, sale request.UpdateUserRequest) error
	Delete(saleId int) error
	FindById(saleId int) (response.SalesResponse, error)
	FindAll() ([]response.SalesResponse, error)
}
